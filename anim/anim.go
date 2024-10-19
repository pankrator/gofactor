package anim

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animate struct {
	img *ebiten.Image

	frames    int
	maxFrames int
	counter   int

	beginX  int
	beginY  int
	offsetX int
	offsetY int
}

func NewAnimate(
	img *ebiten.Image,
	maxFrames int,
	beginX int,
	beginY int,
	offsetX int,
	offsetY int,
) *Animate {
	return &Animate{
		img:       img,
		maxFrames: maxFrames,
		beginX:    beginX,
		beginY:    beginY,
		offsetX:   offsetX,
		offsetY:   offsetY,
	}
}

func (a *Animate) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	startX := a.beginX + (a.offsetX * a.frames)
	startY := a.beginY

	f := a.img.SubImage(image.Rect(startX, startY, startX+a.offsetX, startY+a.offsetY)).(*ebiten.Image)

	screen.DrawImage(f, opts)
}

func (a *Animate) Update() {
	a.counter++

	if a.counter > 10 {
		a.frames++
		a.counter = 0
	}

	if a.frames >= a.maxFrames {
		a.frames = 0
	}
}
