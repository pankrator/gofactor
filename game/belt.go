package game

import (
	"gogame/anim"

	"github.com/hajimehoshi/ebiten/v2"
)

type Item struct {
	geom *ebiten.GeoM
}

type Belt struct {
	animate *anim.Animate
	geom    *ebiten.GeoM

	prev *Belt
	next *Belt

	items []*Item
}

func NewBelt(animate *anim.Animate, geom *ebiten.GeoM) *Belt {
	return &Belt{
		animate: animate,
		geom:    geom,
	}
}

func (b *Belt) Draw(screen *ebiten.Image) {
	b.animate.Draw(screen)
}

func (b *Belt) Update() {
	b.animate.Update()

	// TODO: Update items to move them along the belt
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
