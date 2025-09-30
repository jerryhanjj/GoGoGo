package interfacego

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func ImagesDemo() {
	// m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// fmt.Println(m.Bounds())
	// fmt.Println(m.At(0, 0).RGBA())
	m := Image{}
	pic.ShowImage(m)
}
