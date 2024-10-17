package game

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Draw(screen *ebiten.Image, opts ebiten.DrawImageOptions)
}

type Animateable interface {
	Drawable
	Update()
}

type Object struct {
	Geom        ebiten.GeoM
	Img         *ebiten.Image
	animateable Animateable
}

type ObjOption func(*Object)

func WithImage(img *ebiten.Image) ObjOption {
	return func(o *Object) {
		o.Img = img
	}
}

func WithAnimateable(drawable Animateable) ObjOption {
	return func(o *Object) {
		o.animateable = drawable
	}
}

func NewObject(geom ebiten.GeoM, opts ...ObjOption) *Object {
	obj := &Object{
		Geom: geom,
	}

	for _, opt := range opts {
		opt(obj)
	}

	return obj
}

func (o Object) Update() {
	if o.animateable != nil {
		o.animateable.Update()
	}
}

func (o *Object) Draw(screen *ebiten.Image, colorScale *ebiten.ColorScale) {
	if o.Img != nil {
		screen.DrawImage(o.Img, &ebiten.DrawImageOptions{
			GeoM:       o.Geom,
			ColorScale: *colorScale,
		})
	}

	if o.animateable != nil {
		o.animateable.Draw(screen, ebiten.DrawImageOptions{
			GeoM: o.Geom,
		})
	}
}
