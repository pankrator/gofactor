package game

import "github.com/hajimehoshi/ebiten/v2"

type Moveable struct {
	obj *Object
}

func NewMoveable(obj *Object) *Moveable {
	return &Moveable{
		obj: obj,
	}
}

func (m *Moveable) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m.obj.Geom.Translate(-2, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		m.obj.Geom.Translate(2, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		m.obj.Geom.Translate(0, -2)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		m.obj.Geom.Translate(0, 2)
	}
}

func (m *Moveable) GetObject() *Object {
	return m.obj
}
