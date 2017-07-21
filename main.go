package main

import (
	"golang.org/x/image/font/basicfont"
	"image"
	"io"
	"flag"
	"os"
	"bufio"
	"fmt"
)

var (
	text = flag.String("t", "Hello, World!", "Specify custom text")
	pipedInput = flag.Bool("p", false, "Read from stdin instead")
)

var (
	glyphs []uint8 = basicfont.Face7x13.Mask.(*image.Alpha).Pix
	glyphW         = 6
	glyphH         = 13
)

func main() {
	flag.Parse()

	if *pipedInput {
		in := bufio.NewReader(os.Stdin)
		text, _ := in.ReadString(0x3)
		fmt.Println(text)
		generateASCIIArt(os.Stdout, text)

	} else {
		generateASCIIArt(os.Stdout, *text)
	}
}

func generateASCIIArt(w io.Writer, text string) {
	for row := 0; row < glyphH; row++ {
		for _, r := range text {
			if int(r) > 0x7f {
				r = 0x7f
			}
			glyphIndex := int(r) - 0x20
			for col := 0; col < glyphW; col++ {
				dot := glyphs[glyphW*glyphH*glyphIndex+(row*glyphW+col)]
				if dot == 0x00 {
					w.Write([]byte(" "))
				} else if dot == 0xff {
					w.Write([]byte("\u2588"))
				}
			}
			w.Write([]byte(" "))
		}
		w.Write([]byte("\n"))
	}
}
