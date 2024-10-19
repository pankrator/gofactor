package game

import (
	"gogame/anim"

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

	items []*Item
}

func NewBelt(animate *anim.Animate, transform Transform) *Belt {
	return &Belt{
		animate:   animate,
		transform: transform,
	}
}

func (b *Belt) Draw(screen *ebiten.Image, opts *ebiten.DrawImageOptions) {
	geom := ebiten.GeoM{}
	geom.Scale(b.transform.ScaleX, b.transform.ScaleY)
	geom.Translate(b.transform.X, b.transform.Y)

	b.animate.Draw(screen, &ebiten.DrawImageOptions{
		GeoM: geom,
	})
}

func (b *Belt) Update() {
	b.animate.Update()

	for _, item := range b.items {
		if !item.init {
			item.Transform.X = b.transform.X
			item.Transform.Y = b.transform.Y
			item.init = true
		}

		item.Transform.X += 0.1
		if item.Transform.X >= b.transform.X+b.transform.ScaleX {
			item.Transform.X = b.transform.X + b.transform.ScaleX
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
