package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	)

type MyImage struct{
	W int
	H int
	col uint8
}

func (m MyImage) Bounds() (image.Rectangle) {
	return image.Rect(0, 0, m.W, m.H)
}

func (m MyImage) ColorModel() (color.Model){
	return color.RGBAModel
}

func (m MyImage) At(x, y int) (color.Color){
	return color.RGBA{uint8(x*y), uint8(y), uint8(y), uint8(x)}
}

func main() {
	m := MyImage{300, 300, 100}
	pic.ShowImage(m)
}
