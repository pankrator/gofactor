package anim

import (
	"gogame/util"
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

// TODO: Add geom data somehow. Maybe calculate in Update?? Or comes from other obj
func (a *Animate) Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions) {
	startX := a.beginX + (a.offsetX * a.frames)
	startY := a.beginY

	f := a.img.SubImage(image.Rect(startX, startY, startX+a.offsetX, startY+a.offsetY)).(*ebiten.Image)
	size := f.Bounds().Size()

	geom := ebiten.GeoM{}
	geom.Scale(util.SizeTo(size, image.Pt(64, 64)))
	geom.Translate(130, 60)

	// fmt.Println(startX, startY, startX+a.offsetX, startY+a.offsetY)

	// 64 = size.X * x
	screen.DrawImage(f, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	// screen.DrawImage(a.img, &ebiten.DrawImageOptions{
	// 	GeoM: geom,
	// })
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
