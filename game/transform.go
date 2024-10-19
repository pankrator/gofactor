package game

import "github.com/hajimehoshi/ebiten/v2"

type Transform struct {
	X, Y           float64
	ScaleX, ScaleY float64
	geom           ebiten.GeoM
}

func (t *Transform) ToGeom() ebiten.GeoM {
	t.geom.Reset()
	t.geom.Scale(t.ScaleX, t.ScaleY)
	t.geom.Translate(t.X, t.Y)

	return t.geom
}
