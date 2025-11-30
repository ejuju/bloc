package main

import (
	"image"
	"image/png"
	"os"

	"github.com/ejuju/bloc/pkg/bloc"
)

func main() {
	var lines []string
	if len(os.Args) >= 2 {
		lines = []string{os.Args[1]}
	} else {
		lines = []string{
			"Pack my box with five dozen liquor jugs.",
			"Waltz, bad nymph, for quick jigs vex.",
			"0123456789",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			"abcdefghijklmnopqrstuvwxyz",
			"!\"#$%&'()*+,-./",
			":;<=>?@",
			"[\\]^_`",
			"{|}~",
		}
	}

	maxLineNumChars := 0
	for _, line := range lines {
		maxLineNumChars = max(maxLineNumChars, len(line))
	}

	f, err := os.OpenFile("tmp/bloc.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img := image.NewGray(image.Rect(0, 0, (bloc.Width+2)*maxLineNumChars, (bloc.Height+2)*len(lines)))
	for y, line := range lines {
		for x, c := range line {
			copyImg(bloc.CharFromRune(c).Image(), img, x*(bloc.Width+2), y*(bloc.Height+2))
		}
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func copyImg(src, dst *image.Gray, offsetX, offsetY int) {
	for y := range bloc.Height {
		for x := range bloc.Width {
			dst.SetGray(offsetX+x, offsetY+y, src.GrayAt(x, y))
		}
	}
}
