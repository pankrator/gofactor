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
	width  int
	height int

	keys        []ebiten.Key
	leftPressed bool

	fillerImg *ebiten.Image
	beltImg   *ebiten.Image

	items    []*game.Object
	objects  []*game.Belt
	beltGrid [][]*game.Belt

	moveable   *game.Moveable
	animetable *anim.Animate

	ghost any

	mouseX int
	mouseY int
}

func (g *Game) Init() {
	g.beltGrid = make([][]*game.Belt, 3000/64)
	for i := 0; i < 3000/64; i++ {
		g.beltGrid[i] = make([]*game.Belt, 3000/64)
	}

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

	moveableTransform := game.Transform{
		X:      100,
		Y:      100,
		ScaleX: 32,
		ScaleY: 32,
	}

	// g.items = append(g.items, game.NewObject(
	// 	&game.Transform{
	// 		X:      100,
	// 		Y:      100,
	// 		ScaleX: 100,
	// 		ScaleY: 100,
	// 	},
	// 	game.WithImage(g.fillerImg.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)),
	// ))

	g.moveable = game.NewMoveable(
		game.NewObject(
			&moveableTransform,
			game.WithImage(g.fillerImg.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)),
		),
	)

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
		&game.Transform{},
		game.WithImage(g.beltImg.SubImage(image.Rect(0, 0, 845, 460)).(*ebiten.Image)),
	)
}

func (g *Game) Update() error {
	g.handleInput()
	g.placeItem()

	g.keys = inpututil.AppendPressedKeys(g.keys[:0])

	x, y := ebiten.CursorPosition()
	g.mouseX = x
	g.mouseY = y

	if g.leftPressed {
		g.handleLeftClick()
	}

	for _, o := range g.objects {
		o.Update()
	}

	g.moveable.Update()
	g.animetable.Update()

	if g.ghost != nil {
		boxX := g.mouseX / 64
		boxY := g.mouseY / 64
		scaleX, scaleY := util.SizeTo(g.ghost.(*game.Object).Img.Bounds().Size(), image.Pt(64, 64))
		g.ghost.(*game.Object).Transform.ScaleX = scaleX
		g.ghost.(*game.Object).Transform.ScaleY = scaleY
		g.ghost.(*game.Object).Transform.X = float64(boxX * 64)
		g.ghost.(*game.Object).Transform.Y = float64(boxY * 64)
	}

	return nil
}

func (g *Game) placeItem() {
	if ebiten.IsKeyPressed(ebiten.KeyG) {
		boxX := g.mouseX / 64
		boxY := g.mouseY / 64

		belt := g.beltGrid[boxX][boxY]
		if belt == nil {
			return
		}

		item := game.NewObject(
			&game.Transform{
				ScaleX: 32,
				ScaleY: 32,
			},
			game.WithImage(g.fillerImg.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)),
		)

		g.items = append(g.items, item)

		belt.AddItem(&game.Item{
			Transform: item.Transform,
		})
	}
}

func (g *Game) handleLeftClick() {
	boxX := g.mouseX / 64
	boxY := g.mouseY / 64

	if obj := g.beltGrid[boxX][boxY]; obj != nil {
		return
	}

	tr := game.Transform{
		X:      float64(boxX * 64),
		Y:      float64(boxY * 64),
		ScaleX: 64,
		ScaleY: 64,
	}

	belt := game.NewBelt(anim.NewAnimate(g.beltImg, 3, 0, 0, 845, 460), tr)

	g.beltGrid[boxX][boxY] = belt
	g.objects = append(g.objects, belt)

	if boxX-1 >= 0 {
		left := g.beltGrid[boxX-1][boxY]
		if left != nil {
			left.ConnectAfter(belt)
			belt.ConnectBefore(left)
		}
	}
	if boxX+1 < len(g.beltGrid) {
		right := g.beltGrid[boxX+1][boxY]
		if right != nil {
			belt.ConnectAfter(right)
			right.ConnectBefore(belt)
		}
	}
}

func (g *Game) handleInput() {
	g.leftPressed = inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0)
}

func (g *Game) Draw(screen *ebiten.Image) {
	textOp := &text.DrawOptions{}
	textOp.GeoM.Scale(3, 3)
	textOp.LineSpacing = fontFace.Metrics().HLineGap + fontFace.Metrics().HAscent + fontFace.Metrics().HDescent

	text.Draw(screen, fmt.Sprintf("%d %d - %d", g.mouseX/64, g.mouseY/64, ebiten.TPS()), fontFace, textOp)

	geom := ebiten.GeoM{}

	geom.Scale(util.SizeTo(image.Pt(845, 460), image.Pt(64, 64)))
	geom.Translate(130, 60)

	g.animetable.Draw(screen, &ebiten.DrawImageOptions{
		GeoM: geom,
	})

	for _, o := range g.objects {
		o.Draw(screen, &ebiten.DrawImageOptions{})
	}

	for _, item := range g.items {
		item.Draw(screen, nil, nil)
	}

	g.moveable.Draw(screen)
	// screen.DrawImage(g.fillerImg.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image), &ebiten.DrawImageOptions{
	// 	GeoM:  g.moveable.GetObject().Geom,
	// 	Blend: ebiten.BlendCopy,
	// })

	if g.ghost != nil {
		colorScale := &ebiten.ColorScale{}
		colorScale.ScaleAlpha(0.4)
		g.ghost.(*game.Object).Draw(screen, colorScale, nil)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return width, height
	g.width = outsideWidth
	g.height = outsideHeight
	return outsideWidth, outsideHeight
}

func main() {
	// ebiten.SetWindowSize(2300, 1400)
	// ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetVsyncEnabled(true)

	gameWorld := &Game{}
	gameWorld.Init()

	if err := ebiten.RunGame(gameWorld); err != nil {
		log.Fatal(err)
	}
}
