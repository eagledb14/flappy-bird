package main

import (
	"fmt"
	"image"
	"image/png"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	activeScene IScene
	background *ebiten.Image
	bgop *ebiten.DrawImageOptions
}

func (g *Game) Update() error {
	msg, err := g.activeScene.Update()
	switch msg.msg {
	case "start":
		g.activeScene = NewGameScene(g.Layout(0,0))
	case "end":
		w, h := g.Layout(0,0)

		g.activeScene = NewMenuScene(w, h, msg.data.(int))
	}
	return err
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 360, 643
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, g.bgop)
	g.activeScene.Draw(screen)
}

func newRect(w,h int, color color.Color) image.Image {
	rect := image.NewRGBA(image.Rect(0,0,w,h))

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			rect.Set(x, y, color)
		}
	}
	return rect
}

func drawBorder(img *ebiten.Image) {
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	color := color.RGBA{255,0,0,255}

	// top and bottom
	for y := 0; y < w; y++ {
		img.Set(0, y, color)
		img.Set(h - 1, y, color)
	}

	// left and right
	for x := 0; x < h; x++ {
		img.Set(x, 0, color)
		img.Set(x, w - 1, color)
	}
}

func loadImage(path string) *ebiten.Image {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return ebiten.NewImageFromImage(newRect(100, 100, color.RGBA{255,192,203,255}))
	}

	img, err := png.Decode(file)
	if err != nil {
		return ebiten.NewImageFromImage(newRect(100, 100, color.RGBA{255,192,203,255}))
	}

	return ebiten.NewImageFromImage(img)
}

func main() {
	ebiten.SetWindowSize(360, 643)
	ebiten.SetWindowTitle("Flappy")
	game := &Game{}

	game.background = loadImage("./background.png")
	game.bgop = &ebiten.DrawImageOptions{}
	game.bgop.GeoM.Scale(3, 3)

	w, h := game.Layout(0,0)
	game.activeScene = NewMenuScene(w, h,0)

	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
	}
}
