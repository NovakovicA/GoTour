package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{byte(x), byte(y), byte(x + y % 255), 255}
}


func main() {
	m := Image{}
	pic.ShowImage(m)
}
