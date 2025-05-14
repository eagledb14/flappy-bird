package main

import (

	"github.com/hajimehoshi/ebiten/v2"
)

type Bird struct {
	img *ebiten.Image
	pos Point
	velocity float64
	pressedTimer int
	alive bool
}

func newBird(h int) *Bird {
	img := loadImage("./bird.png")
	return &Bird {
		img: img,
		pos: Point{20 - float64(img.Bounds().Max.X / 2), float64(h) / 2 - 25 - float64(img.Bounds().Max.Y / 2)},
		pressedTimer: 0,
		alive: true,
	}
}

func (b *Bird) GetBox() *Hitbox {
	return &Hitbox{b.pos.X, b.pos.Y, float64(b.img.Bounds().Max.X), float64(b.img.Bounds().Max.Y)}
}

func (b *Bird) Update() {
	if b.pressedTimer <= 0  && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft){
		b.pressedTimer = 25
		b.velocity = -7
	} else {
		_, h := ebiten.WindowSize()
		b.pressedTimer -= 1
		b.pos.Y += b.velocity
		b.pos.Y = clamp(0,float64(h) - 100, b.pos.Y)
		b.velocity = clamp(-7, 5, b.velocity + 0.4)
	}
}

func (b *Bird) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.pos.X, b.pos.Y)
	screen.DrawImage(b.img, op)
}
