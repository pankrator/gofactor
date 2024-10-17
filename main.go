package main

import (
	"fmt"
	"gogame/anim"
	"gogame/game"
	"gogame/util"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var fontFace = text.NewGoXFace(bitmapfont.Face)

type Game struct {
	keys        []ebiten.Key
	leftPressed bool

	fillerImg *ebiten.Image
	beltImg   *ebiten.Image

	objects []*game.Object

	moveable   *game.Moveable
	animetable *anim.Animate

	ghost any

	mouseX int
	mouseY int
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

	f, err := os.Open("resources/belts2.png")
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(img.Bounds())

	g.beltImg = ebiten.NewImageFromImage(img)

	g.animetable = anim.NewAnimate(g.beltImg, 3, 0, 0, 845, 460)

	g.ghost = game.NewObject(
		ebiten.GeoM{},
		game.WithImage(g.beltImg.SubImage(image.Rect(0, 0, 845, 460)).(*ebiten.Image)),
	)
}

func (g *Game) Update() error {
	g.handleInput()
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	x, y := ebiten.CursorPosition()
	g.mouseX = x
	g.mouseY = y

	if g.leftPressed {
		geom := ebiten.GeoM{}
		geom.Scale(32, 32)
		geom.Translate(float64(g.mouseX), float64(g.mouseY))

		game.NewObject(
			geom,
			game.WithAnimateable(anim.NewAnimate(g.beltImg, 3, 0, 0, 845, 460)),
		)
		g.objects = append(g.objects, game.NewObject(geom))
	}

	g.moveable.Update()
	g.animetable.Update()

	if g.ghost != nil {
		g.ghost.(*game.Object).Geom.Reset()
		g.ghost.(*game.Object).Geom.Scale(util.SizeTo(g.ghost.(*game.Object).Img.Bounds().Size(), image.Pt(64, 64)))
		boxX := g.mouseX / 64
		boxY := g.mouseY / 64
		g.ghost.(*game.Object).Geom.Translate(float64(boxX*64), float64(boxY*64))
	}

	return nil
}

func (g *Game) handleInput() {
	g.leftPressed = inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
}

func (g *Game) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Scale(5, 5)
	textOp.LineSpacing = fontFace.Metrics().HLineGap + fontFace.Metrics().HAscent + fontFace.Metrics().HDescent

	text.Draw(screen, fmt.Sprintf("%d %d", g.mouseX/64, g.mouseY/64), fontFace, textOp)

	// geom := ebiten.GeoM{}
	// geom.Scale(64, 64)
	// geom.Translate(32, 32)

	g.animetable.Draw(screen)

	for _, o := range g.objects {
		screen.DrawImage(g.fillerImg.SubImage(image.Rect(0, 0, 1, 1)).(*ebiten.Image),
			&ebiten.DrawImageOptions{
				GeoM:  o.Geom,
				Blend: ebiten.BlendCopy,
			})
	}

	screen.DrawImage(g.fillerImg.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawImageOptions{
		GeoM:  g.moveable.GetObject().Geom,
		Blend: ebiten.BlendCopy,
	})

	if g.ghost != nil {
		colorScale := &ebiten.ColorScale{}
		colorScale.ScaleAlpha(0.4)
		g.ghost.(*game.Object).Draw(screen, colorScale)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1024, 768
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetFullscreen(false)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetVsyncEnabled(true)

	gameWorld := &Game{}
	gameWorld.Init()

	if err := ebiten.RunGame(gameWorld); err != nil {
		log.Fatal(err)
	}
}
