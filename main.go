package main

import (
	"golang.org/x/image/font/basicfont"
	"image"
	"io"
	"flag"
	"os"
	"bufio"
	"strings"
	"github.com/jroimartin/gocui"
)

var (
	text       = flag.String("t", "", "Specify custom text")
	inputFile  = flag.String("i", "", "Input file")
	pipedInput = flag.Bool("p", false, "Read from stdin instead")
)

var glyphs []uint8 = basicfont.Face7x13.Mask.(*image.Alpha).Pix

const (
	glyphW = 6
	glyphH = 13
)

func main() {
	flag.Parse()

	switch {
	case *pipedInput:
		in := bufio.NewReader(os.Stdin)
		text, _ := in.ReadString(0x3)
		generateASCIIArt(os.Stdout, text)

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
		generateASCIIArt(os.Stdout, text)

	case *text != "":
		generateASCIIArt(os.Stdout, *text)

	default:
		gui, err := gocui.NewGui(gocui.OutputNormal)
		if err != nil {
			panic(err)
		}
		gui.Close()
	}

}

func generateASCIIArt(w io.Writer, text string) {
	buffer := bufio.NewWriter(w)
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
						buffer.WriteRune('\u2588')
					}
				}
				buffer.WriteRune(' ')
			}
			buffer.WriteRune('\n')
		}
	}
	buffer.Flush()
}
