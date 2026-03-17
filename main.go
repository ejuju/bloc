package main

import (
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/ejuju/bloc/pkg/bloc"
)

func main() {
	txt := ""
	if len(os.Args) >= 2 {
		txt = strings.Join(os.Args[1:], "\n")
	} else {
		txt = "" +
			"Pack my box with five dozen liquor jugs.\n" +
			"Waltz, bad nymph, for quick jigs vex.\n" +
			"0123456789\n" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ\n" +
			"abcdefghijklmnopqrstuvwxyz\n" +
			"!\"#$%&'()*+,-./\n" +
			":;<=>?@\n" +
			"[\\]^_`\n" +
			"{|}~"
	}

	f, err := os.OpenFile("tmp/bloc.png", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img := bloc.Image(txt, image.White, image.Black)

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
