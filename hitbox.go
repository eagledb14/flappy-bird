package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Hitbox struct {
	X float64
	Y float64
	Width  float64
	Height float64
}

func (h1 *Hitbox) IsOverlap(h2 *Hitbox) bool {
	xOverlap := h1.X < h2.X+h2.Width && h1.X+h1.Width > h2.X
	yOverlap := h1.Y < h2.Y+h2.Height && h1.Y+h1.Height > h2.Y

	return xOverlap && yOverlap
}

func (h1 *Hitbox) Draw(img *ebiten.Image) {
	h := int(h1.Width)
	w := int(h1.Height)
	x := int(h1.X)
	y := int(h1.Y)

	color := color.RGBA{255,0,0,255}

	// top and bottom
	for i := 0; i < w; i++ {
		img.Set(0 + x, i + y, color)
		img.Set(h - 1, i + y, color)
	}

	// left and right
	for i := 0; i < h; i++ {
		img.Set(i + x, 0 + y, color)
		img.Set(i + x, w - 0 + y, color)
	}
}
