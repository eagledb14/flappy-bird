package main

import (
	"image"
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Score struct {
	img *ebiten.Image
	imgs []*ebiten.Image
	op ebiten.DrawImageOptions
}

func newScore() *Score {
	op := ebiten.DrawImageOptions{}
	w, _ := ebiten.WindowSize()
	op.GeoM.Translate(float64(w) / 2  - 117, 10)
	op.GeoM.Scale(2, 2)

	img := loadImage("./numbers.png")
	imgs := []*ebiten.Image{}
	for val := range 10 {
		switch val {
		case 0:
		      imgs = append(imgs, getDigit(img, 0,0,30, 45))
		case 1:
		      imgs = append(imgs, getDigit(img, 31, 0, 48, 45))
		case 2:
		      imgs = append(imgs, getDigit(img, 50, 0, 80, 45))
		case 3:
		      imgs = append(imgs, getDigit(img, 80, 0, 110, 45))
		case 4:
		      imgs = append(imgs, getDigit(img, 110, 0, 140, 45))
		case 5:
		      imgs = append(imgs, getDigit(img, 0, 45, 30, 90))
		case 6:
		      imgs = append(imgs, getDigit(img, 30, 45, 60, 90))
		case 7:
		      imgs = append(imgs, getDigit(img, 60, 45, 90, 90))
		case 8:
		      imgs = append(imgs, getDigit(img, 90, 45, 120, 90))
		case 9:
			imgs = append(imgs, getDigit(img, 120, 45, 150, 90))
		default:
			imgs = append(imgs, ebiten.NewImageFromImage(newRect(100, 100, color.RGBA{255, 0, 0, 255})))
		}
	}

	return &Score {
		img: loadImage("./numbers.png"),
		imgs: imgs,
		op: op,
	}
}

func getDigit(img *ebiten.Image, x, y, w, h int) *ebiten.Image {
	r := image.Rectangle{}
	r.Min = image.Point{x, y}
	r.Max = image.Point{w, h}
	out := img.SubImage(r)
	return ebiten.NewImageFromImage(out)
}


func (s *Score) Draw(screen *ebiten.Image, score int) {
	score_split := strings.Split(strconv.Itoa(score / 2), "")
	if len(score_split) == 1 {
		score_split = append([]string{"0"}, score_split...)
	}


	op := s.op
	for _, img := range score_split {
		index, _ := strconv.Atoi(img)
		screen.DrawImage(s.imgs[index], &op)
		op.GeoM.Translate(65, 0)
	}
}
