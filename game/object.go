package game

import "github.com/hajimehoshi/ebiten/v2"

type Object struct {
	Geom ebiten.GeoM
	// img  ebiten.Image
}

func NewObject(geom ebiten.GeoM) *Object {
	return &Object{
		Geom: geom,
	}
}
