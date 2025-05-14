package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type IScene interface {
	Update() (Msg, error)
	Draw(screen *ebiten.Image)
	Layout() (int, int)
}


type GameScene struct {
	w, h int
	pipes *Pipes
	bird *Bird
	score *Score
}

func NewGameScene(w, h int) *GameScene {
	return &GameScene{
		w: w,
		h: h,
		pipes: newPipes(),
		bird: newBird(h),
		score: newScore(),
	}
}

func (g *GameScene) Layout() (int, int) {
	return g.w, g.h
}

func (g *GameScene) Update() (Msg, error) {

	g.pipes.Update(g.Layout())
	g.bird.Update()
	for _, pb := range g.pipes.GetBoxes() {
		_ = pb
		if g.bird.GetBox().IsOverlap(&pb) {
			return Msg{"end", g.pipes.Score}, nil
		}
	}

	return Msg{}, nil
}

func (g *GameScene) Draw(screen *ebiten.Image) {
	g.pipes.Draw(screen)
	g.score.Draw(screen, g.pipes.Score)
	g.bird.Draw(screen)
}


type MenuScene struct {
	w, h int
	tick int
	op *ebiten.DrawImageOptions
	button *ebiten.Image
	buttonBox Hitbox
	title *ebiten.Image
	top *ebiten.DrawImageOptions
	score *Score
	highScore int
}

func NewMenuScene(w, h, score int) *MenuScene {
	title := loadImage("./title.png")
	top := &ebiten.DrawImageOptions{}
	x := (float64(w) / 2) - float64(title.Bounds().Max.X / 2)
	y := (float64(h) / 2)- float64(title.Bounds().Max.Y / 2) - 100
	top.GeoM.Translate(x, y)

	op := &ebiten.DrawImageOptions{}
	img := loadImage("./start.png")
	x = (float64(w) / 2) - float64(img.Bounds().Max.X / 2)
	y = (float64(h) / 2)- float64(img.Bounds().Max.Y / 2) + 100
	op.GeoM.Translate(x, y)

	return &MenuScene {
		w: w,
		h: h,
		op: op,
		button: img,
		buttonBox: Hitbox{
			X: x, 
			Y: y,
			Width: 100,
			Height: 100,
		},
		title: title,
		top: top,
		score: newScore(),
		highScore: score,
	}
}

func (m *MenuScene) Layout() (int, int) {
	return m.w, m.h
}

func (m *MenuScene) Update() (Msg, error) {

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		mouseBox := Hitbox{float64(mx), float64(my), 1, 1}

		if m.buttonBox.IsOverlap(&mouseBox) {
			return Msg{msg: "start"}, nil
		}
	}

	return Msg{}, nil
}

func (m *MenuScene) Draw(screen *ebiten.Image) {
	screen.DrawImage(m.title, m.top)
	screen.DrawImage(m.button, m.op)
	if m.highScore > 0 {
		m.score.Draw(screen, m.highScore)
	}
}



