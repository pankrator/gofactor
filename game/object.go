package game

import "github.com/hajimehoshi/ebiten/v2"

type Drawable interface {
	Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions)
}

type Animateable interface {
	Drawable
	Update()
}

type Object struct {
	Transform   Transform
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

func NewObject(transform Transform, opts ...ObjOption) *Object {
	obj := &Object{
		Transform: transform,
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

func (o *Object) Draw(screen *ebiten.Image, colorScale *ebiten.ColorScale, blend *ebiten.Blend) {
	if o.Img != nil {
		opts := &ebiten.DrawImageOptions{
			GeoM:       o.Transform.ToGeom(),
			ColorScale: *colorScale,
		}
		if blend != nil {
			opts.Blend = *blend
		}
		screen.DrawImage(o.Img, opts)
	}

	if o.animateable != nil {
		o.animateable.Draw(screen, &ebiten.DrawImageOptions{
			GeoM: o.Transform.ToGeom(),
		})
	}
}
