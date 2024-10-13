package anim

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Animate struct {
	img *ebiten.Image

	frames    int
	maxFrames int

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
		frames:    0,
		img:       img,
		maxFrames: maxFrames,
		beginX:    beginX,
		beginY:    beginY,
		offsetX:   offsetX,
		offsetY:   offsetY,
	}
}

// TODO: Add geom data somehow. Maybe calculate in Update?? Or comes from other obj
func (a *Animate) Draw(screen *ebiten.Image) {
	geom := ebiten.GeoM{}
	geom.Translate(10, 10)

	startX := a.beginX + (a.offsetX * a.frames)
	startY := a.beginY
	screen.DrawImage(a.img.SubImage(image.Rect(startX, startY, startX+a.offsetX, startY)).(*ebiten.Image), &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (a *Animate) Update() {
	a.frames++
	if a.frames > 5 {
		a.frames = 0
	}
}
