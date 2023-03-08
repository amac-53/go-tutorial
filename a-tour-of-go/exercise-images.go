package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	w, h int
}

func (im Image) Bounds() image.Rectangle {
	w, h := im.w, im.h 
	return image.Rect(0, 0, w, h)
}

func (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	v := (x + y) / 2
	return color.RGBA{byte(v), byte(v), 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
