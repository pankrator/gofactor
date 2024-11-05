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
	Transform   *Transform
	Img         *ebiten.Image
	animateable Animateable
	camera      *TestCamera
}

type ObjOption func(*Object)

func WithImage(img *ebiten.Image) ObjOption {
	return func(o *Object) {
		o.Img = img
	}
}

func WithCamera(camera *TestCamera) ObjOption {
	return func(o *Object) {
		o.camera = camera
	}
}

func WithAnimateable(drawable Animateable) ObjOption {
	return func(o *Object) {
		o.animateable = drawable
	}
}

func NewObject(transform *Transform, opts ...ObjOption) *Object {
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
		tempTransform := &Transform{
			X:      o.Transform.X,
			Y:      o.Transform.Y,
			ScaleX: o.Transform.ScaleX,
			ScaleY: o.Transform.ScaleY,
		}

		if o.camera != nil {
			// tempTransform.X *= o.camera.ZoomFactor
			// tempTransform.Y *= o.camera.ZoomFactor
			// tempTransform.ScaleX *= o.camera.ZoomFactor
			// tempTransform.ScaleY *= o.camera.ZoomFactor
		}

		opts := &ebiten.DrawImageOptions{
			GeoM: tempTransform.ToGeom(
				o.camera,
			),
		}

		if colorScale != nil {
			opts.ColorScale = *colorScale
		}
		if blend != nil {
			opts.Blend = *blend
		}
		screen.DrawImage(o.Img, opts)
	}

	if o.animateable != nil {
		o.animateable.Draw(screen, &ebiten.DrawImageOptions{
			GeoM: o.Transform.ToGeom(o.camera),
		})
	}
}
