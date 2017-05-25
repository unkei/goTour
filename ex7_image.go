package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{
	rect image.Rectangle
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return img.rect
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x*y) % 256)
	return color.RGBA{v, v, 0, 255}
}


func main() {
	m := Image{image.Rect(0, 0, 255, 255)}
	pic.ShowImage(m)
}

