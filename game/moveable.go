package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Moveable struct {
	obj *Object
}

// TODO: Set the geometry that needs moving and embed this into other structs that need moving
func NewMoveable(obj *Object) *Moveable {
	return &Moveable{
		obj: obj,
	}
}

func (m *Moveable) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		m.obj.Transform.X -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		m.obj.Transform.X += 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		m.obj.Transform.Y -= 2
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		m.obj.Transform.Y += 2
	}
}

func (m *Moveable) Draw(screen *ebiten.Image) {
	m.obj.Draw(screen, &ebiten.ColorScale{}, &ebiten.BlendCopy)
}

func (m *Moveable) GetObject() *Object {
	return m.obj
}
