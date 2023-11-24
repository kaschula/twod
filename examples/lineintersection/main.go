package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/kaschula/twod/color"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/inputs"
	"github.com/kaschula/twod/physics"
	"log"
)

const (
	screenW, screenH = 1000, 600
)

func main() {
	// make a game with 6 otherLines that cross and print out where the touches occur

	//test physics.GetIntersection()
	game := NewLineSegmentsGame(
		physics.NewLineSegment(physics.New(100, 100), physics.New(100, 500)),
		physics.NewLineSegment(physics.New(50, 250), physics.New(500, 250)),
		physics.NewLineSegment(physics.New(100, 350), physics.New(500, 350)),
		physics.NewLineSegment(physics.New(250, 50), physics.New(250, 450)),
	)

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Self Driving Car")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type LineSegmentsGame struct {
	user       *physics.LineSegment
	otherLines []*physics.LineSegment
	touches    []physics.Touch
}

func (l *LineSegmentsGame) Update() error {
	userLine := l.user
	if inputs.KeyPressAndHold(ebiten.KeyUp) {
		userLine = userLine.MoveEnd(physics.New(0, -10))
	}

	if inputs.KeyPressAndHold(ebiten.KeyDown) {
		userLine = userLine.MoveEnd(physics.New(0, 10))
	}

	if inputs.KeyPressAndHold(ebiten.KeyLeft) {
		userLine = userLine.MoveEnd(physics.New(-10, 0))
	}

	if inputs.KeyPressAndHold(ebiten.KeyRight) {
		userLine = userLine.MoveEnd(physics.New(10, 0))
	}
	l.user = userLine

	l.touches = []physics.Touch{}
	for _, line := range l.otherLines {
		touch := physics.GetLineIntersection(l.user, line)
		if !touch.Empty() {
			l.touches = append(l.touches, touch)
		}
	}

	return nil
}

func (l *LineSegmentsGame) Draw(screen *ebiten.Image) {
	draw.Line(screen, l.user.Start(), l.user.End(), color.Red)

	for _, line := range l.otherLines {
		draw.Line(screen, line.Start(), line.End(), color.White)
	}

	// loop through touches and draw them

	if len(l.touches) == 0 {
		return
	}

	for _, touch := range l.touches {
		draw.Vector(screen, touch.Vector(), color.Green, 5)
	}

	// show first touch as example

	touch := l.touches[0]

	// DrawUnit the line based on the offset, get t1 and t2 from touch, will need to add line to touch

	//touchLine := physics.NewLineSegment(l.user.Start(), touch.Vector()).Transpose(physics.New(500, 0))

	originalLine := touch.LineA.Transpose(physics.New(500, 0))
	t1 := touch.LineAStartToOffSet().Transpose(physics.New(520, 0))
	t2 := touch.LineAEndToOffSet().Transpose(physics.New(540, 0))

	draw.Line(screen, t1.Start(), t1.End(), color.Blue)
	ebitenutil.DebugPrintAt(screen, "T1", t1.Start().XInt(), t1.Start().YInt())
	draw.Line(screen, t2.Start(), t2.End(), color.Blue)
	t2Debug := t2.Start().Add(physics.New(4, 0))
	ebitenutil.DebugPrintAt(screen, "T2", t2Debug.XInt(), t2Debug.YInt())
	draw.Line(screen, originalLine.Start(), originalLine.End(), color.Blue)
	original := originalLine.Start().Add(physics.New(-4, -16))
	ebitenutil.DebugPrintAt(screen, "Original Line", original.XInt(), original.YInt())

}

func (l *LineSegmentsGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewLineSegmentsGame(lines ...*physics.LineSegment) *LineSegmentsGame {
	return &LineSegmentsGame{
		user:       lines[0],
		otherLines: lines[1:],
	}
}
