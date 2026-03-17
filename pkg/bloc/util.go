package bloc

import (
	"image"
	"image/color"
	"image/draw"
	"strings"
)

func Image(txt string, fg, bg color.Color) (img *image.RGBA) {
	lines := strings.Split(txt, "\n")
	maxLineSize := 0
	for _, line := range lines {
		maxLineSize = max(maxLineSize, len(line))
	}
	w, h := (Width+2)*maxLineSize, (Height+2)*len(lines)
	img = image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, image.Rect(0, 0, w, h), image.NewUniform(bg), image.Point{}, draw.Over)
	for y, line := range lines {
		for x, c := range line {
			copyChar(CharFromRune(c).Image(fg, bg), img, x*(Width+2), y*(Height+2))
		}
	}
	return img
}

func copyChar(src, dst *image.RGBA, offsetX, offsetY int) {
	for y := range Height {
		for x := range Width {
			dst.Set(offsetX+x, offsetY+y, src.At(x, y))
		}
	}
}
