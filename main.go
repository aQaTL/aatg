package main

import (
	"golang.org/x/image/font/basicfont"
	"image"
	"io"
	"flag"
	"os"
	"bufio"
	"strings"
	"unicode/utf8"
)

var (
	text       = flag.String("t", "", "Specify custom text")
	inputFile  = flag.String("i", "", "Input file")
	pipedInput = flag.Bool("p", false, "Read from stdin")
	glyph = flag.String("c", "\u2588", "Custom ascii character")
)

var glyphs []uint8 = basicfont.Face7x13.Mask.(*image.Alpha).Pix

const (
	glyphW = 6
	glyphH = 13
)

func main() {
	flag.Parse()
	glyph, _ := utf8.DecodeRuneInString(*glyph)

	switch {
	case *pipedInput:
		in := bufio.NewReader(os.Stdin)
		text, _ := in.ReadString(0x3)
		generateASCIIArt(os.Stdout, text, glyph)

	case *inputFile != "":
		file, err := os.Open(*inputFile)
		if err != nil {
			panic(err)
		}
		r := bufio.NewReader(file)
		text, err := r.ReadString(0x3)
		if err != io.EOF {
			panic(err)
		}
		generateASCIIArt(os.Stdout, text, glyph)

	case *text != "":
		generateASCIIArt(os.Stdout, *text, glyph)

	default:
		startGui()
	}

}

func generateASCIIArt(w io.Writer, text string, glyph rune) {
	buffer := bufio.NewWriter(w)
	if r, rSize := utf8.DecodeLastRuneInString(text); r == '\n' {
		text = text[:len(text)-rSize] //Prevents drawing empty line at the end
	}
	for _, text := range strings.Split(text, "\n") {
		for row := 0; row < glyphH; row++ {
			for _, r := range text {
				if int(r) > 0x7f || int(r) < 0x20 {
					r = 0x7f
				}
				glyphIndex := int(r) - 0x20
				for col := 0; col < glyphW; col++ {
					dot := glyphs[glyphW*glyphH*glyphIndex+(row*glyphW+col)]
					if dot == 0x00 {
						buffer.WriteRune(' ')
					} else if dot == 0xff {
						buffer.WriteRune(glyph)
					}
				}
				buffer.WriteRune(' ')
			}
			buffer.WriteRune('\n')
		}
	}
	buffer.Flush()
}
