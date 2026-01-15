package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	width, height int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.width, m.height)
}

func (m Image) At(x, y int) color.Color {
	c := uint8((x ^ y) % 256)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{
		width:  128,
		height: 128,
	}
	pic.ShowImage(m)
}

// Alternative approach using a configuration struct slow
// type Config struct {
// 	colorModel color.Model
// 	bounds     image.Rectangle
// 	at         func(x, y int) color.Color
// }

// func showImageWithConfig(cfg Config) {
// 	// process the image according to cfg
// }

// func runWithCustomImage() {
// 	cfg := Config{
// 		colorModel: color.RGBAModel,
// 		bounds:     image.Rect(0, 0, 128, 128),
// 		at: func(x, y int) color.Color {
// 			c := uint8((x ^ y) % 256)
// 			return color.RGBA{c, c, 255, 255}
// 		},
// 	}

// 	showImageWithConfig(cfg)
// }
