package game

import "github.com/hajimehoshi/ebiten/v2"

type Transform struct {
	X, Y           float64
	ScaleX, ScaleY float64
	geom           ebiten.GeoM
}

type TestCamera struct {
	*Transform
	ZoomFactor float64
}

func (tc TestCamera) ToWorldCoords(t *Transform) Transform {
	return Transform{
		X: (t.X + tc.X),
		Y: (t.Y + tc.Y),
	}
}

func (t *Transform) ToGeom(camera *TestCamera) ebiten.GeoM {
	t.geom.Reset()

	if camera != nil {
		t.geom.Scale(t.ScaleX*float64(camera.ZoomFactor), t.ScaleY*float64(camera.ZoomFactor))
		t.geom.Translate((t.X-camera.X)*camera.ZoomFactor, (t.Y-camera.Y)*camera.ZoomFactor)
	} else {
		t.geom.Scale(t.ScaleX, t.ScaleY)
		t.geom.Translate(t.X, t.Y)
	}

	/*
		camera.x === 0
		camera.y === 0

		camera.x = 30
		camera.y = 30

		camera.zoom = 5

		item.world.x = 5
		item.world.y = 5
	*/
	return t.geom
}
