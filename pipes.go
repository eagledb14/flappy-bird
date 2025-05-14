package main

import (
	// "image"
	// "image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Pipes struct {
	img *ebiten.Image
	op *ebiten.DrawImageOptions
	pos []Point
	Score int
}

type Point struct {
	X float64
	Y float64
}

func newPipes() *Pipes {
	w, _ := ebiten.WindowSize()
	img := loadImage("./pipes.png")

	pos := []Point{}

	// top
	for i := range 2 {
		pos = append(pos, Point{X: float64(w) + (float64(i) * 350) , Y:-300})
	}

	// bottom
	for i := range 2 {
		pos = append(pos, Point{X: float64(w) + (float64(i) * 350), Y:350})
	}
	
	return &Pipes{
		img: ebiten.NewImageFromImage(img),
		pos: pos,
		op: &ebiten.DrawImageOptions{},
	}
}

func (p *Pipes) GetBoxes() []Hitbox {
	boxes := []Hitbox{}
	for i := range p.pos {
		boxes = append(boxes, Hitbox{p.pos[i].X, p.pos[i].Y, float64(p.img.Bounds().Max.X), float64(p.img.Bounds().Max.Y)})
		_ = i
	}
	return boxes
}

func (p *Pipes) Update(w, h int) {
	for i := range p.pos {
		if p.pos[i].X < -50 {
			p.pos[i].X = float64(w) + 200
			num := p.pos[i].Y + (rand.Float64() * 100) - 50
			p.pos[i].Y = clamp(0 - float64(p.img.Bounds().Max.Y) + 10, float64(h) - 10, num)
			p.Score += 1
		} else {
			p.pos[i].X -= 4
		}
	}
}

func (p *Pipes) Draw(screen *ebiten.Image) {
	// for _, b := range p.GetBoxes() {
	// 	b.Draw(screen)
	// }
	for i := range p.pos {
		p.op.GeoM.Reset()
		p.op.GeoM.Translate(p.pos[i].X, p.pos[i].Y)

		screen.DrawImage(p.img, p.op)
	}
}

func clamp(min, max float64, num float64) float64 {
	if num <= min {
		return min
	} else if num >= max {
		return max
	} else {
		return num
	}
}
