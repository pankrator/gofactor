package game

import (
	"gogame/anim"
	"gogame/util"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Item struct {
	Transform *Transform
	init      bool
}

type Belt struct {
	transform Transform
	animate   *anim.Animate

	prev *Belt
	next *Belt

	items  []*Item
	camera *TestCamera
}

func NewBelt(animate *anim.Animate, transform Transform, camera *TestCamera) *Belt {
	return &Belt{
		animate:   animate,
		transform: transform,
		camera:    camera,
	}
}

func (b *Belt) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	scaleX, scaleY := util.SizeTo(image.Pt(845, 460), image.Pt(int(b.transform.ScaleX), int(b.transform.ScaleY)))

	tempTransform := b.transform
	tempTransform.ScaleX = scaleX
	tempTransform.ScaleY = scaleY

	geom := tempTransform.ToGeom(b.camera)

	opts = &ebiten.DrawImageOptions{
		GeoM: geom,
	}

	if b.next != nil || b.prev != nil {
		opts.ColorScale.SetB(0.1)
	}

	b.animate.Draw(screen, opts)
}

func (b *Belt) Update() {
	b.animate.Update()

	for i := len(b.items) - 1; i >= 0; i-- {
		item := b.items[i]

		if !item.init {
			item.Transform.X = b.transform.X
			item.Transform.Y = b.transform.Y
			item.init = true
		}

		item.Transform.X += 0.5
		if item.Transform.X+item.Transform.ScaleX >= b.transform.X+b.transform.ScaleX {
			item.Transform.X = b.transform.X + b.transform.ScaleX - item.Transform.ScaleX
			if b.next != nil {
				b.next.AddItem(item)
				b.items = append(b.items[:i], b.items[i+1:]...)
			}
		}
	}
}

func (b *Belt) ConnectBefore(belt *Belt) {
	b.prev = belt
}

func (b *Belt) ConnectAfter(belt *Belt) {
	b.next = belt
}

func (b *Belt) AddItem(item *Item) {
	b.items = append(b.items, item)
}
