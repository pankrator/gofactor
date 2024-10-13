package main

import (
	"gogame/game"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	keys        []ebiten.Key
	leftPressed bool

	fillerImg *ebiten.Image

	objects []*game.Object

	moveable *game.Moveable
}

func (g *Game) Init() {
	testImg := ebiten.NewImage(2, 2)
	testImg.Set(0, 0, color.RGBA{
		R: 255,
		G: 0,
		B: 0,
	})

	testImg.Set(1, 1, color.RGBA{
		R: 0,
		G: 255,
		B: 0,
	})

	g.fillerImg = testImg

	moveableGeom := ebiten.GeoM{}
	moveableGeom.Scale(32, 32)
	moveableGeom.Translate(100, 100)

	g.moveable = game.NewMoveable(game.NewObject(moveableGeom))
}

func (g *Game) Update() error {
	g.handleInput()
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	if g.leftPressed {
		x, y := ebiten.CursorPosition()
		geom := ebiten.GeoM{}
		geom.Scale(32, 32)
		geom.Translate(float64(x), float64(y))
		g.objects = append(g.objects, game.NewObject(geom))
	}

	g.moveable.Update()

	return nil
}

func (g *Game) handleInput() {
	g.leftPressed = inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, World!")

	// geom := ebiten.GeoM{}
	// geom.Scale(64, 64)
	// geom.Translate(32, 32)

	for _, o := range g.objects {
		screen.DrawImage(g.fillerImg.SubImage(image.Rect(0, 0, 1, 1)).(*ebiten.Image), &ebiten.DrawImageOptions{GeoM: o.Geom})
	}

	screen.DrawImage(g.fillerImg.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawImageOptions{
		GeoM: g.moveable.GetObject().Geom,
		// ColorScale: colorScale,
		// Blend:      ebiten.BlendDestination,
	})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowTitle("Hello, World!")

	gameWorld := &Game{}
	gameWorld.Init()

	if err := ebiten.RunGame(gameWorld); err != nil {
		log.Fatal(err)
	}
}
