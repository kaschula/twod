package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	color2 "github.com/kaschula/twod/color"
	"github.com/kaschula/twod/components/selectionarea"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/inputs"
	"github.com/kaschula/twod/physics"
	"log"
)

const (
	screenW, screenH = 600, 600
)

var (
	selectAreaColor = color2.RGBA(250, 250, 250, 60)
)

type SelectableEntityUnit struct {
	body       physics.Rectangle
	isSelected bool
}

func (s *SelectableEntityUnit) Center() physics.V {
	return s.body.Location()
}

func (s *SelectableEntityUnit) Body() physics.Rectangle {
	return s.body
}

func (s *SelectableEntityUnit) Select() {
	s.isSelected = true
}

func (s *SelectableEntityUnit) Unselect() {
	s.isSelected = false
}

type SelectableEntityStructure struct {
	body       physics.Rectangle
	isSelected bool
}

func (s *SelectableEntityStructure) Center() physics.V {
	return s.body.Location()
}

func (s *SelectableEntityStructure) Body() physics.Rectangle {
	return s.body
}

func (s *SelectableEntityStructure) Select() {
	s.isSelected = true
}

func (s *SelectableEntityStructure) Unselect() {
	s.isSelected = false
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("Area Select")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	mouseInputs                   *inputs.MouseInputProducer
	selectionAreaServiceUnits     selectionarea.Service[*SelectableEntityUnit]
	selectionAreaServiceBuildings selectionarea.Service[*SelectableEntityStructure]
	selectableUnits               []*SelectableEntityUnit
	selectableBuildings           []*SelectableEntityStructure
}

func (g *Game) Update() error {
	mouseInputs := g.mouseInputs.Update()

	// the service is bound to a single type this example is selecting multiple differnet types of things
	g.selectionAreaServiceUnits.Update(mouseInputs, g.selectableUnits)
	g.selectionAreaServiceBuildings.Update(mouseInputs, g.selectableBuildings)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// do next
	// print text on screen
	for _, entity := range g.selectableUnits {
		if entity.isSelected {
			draw.Polygon(screen, entity.body, color2.Green)
		} else {
			draw.Polygon(screen, entity.body, color2.Red)
		}
	}

	for _, entity := range g.selectableBuildings {
		if entity.isSelected {
			draw.Polygon(screen, entity.body, color2.Green)
		} else {
			draw.Polygon(screen, entity.body, color2.Blue)
		}
	}

	// Only draw one selection service
	g.selectionAreaServiceUnits.Draw(screen, selectAreaColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewGame() *Game {
	return &Game{
		mouseInputs:                   inputs.NewMouseInputProducer(inputs.MouseCursor, inputs.MouseLeftPressed, inputs.MouseRightPressed, 0, 0),
		selectionAreaServiceUnits:     selectionarea.NewSelectionAreaService[*SelectableEntityUnit](),
		selectionAreaServiceBuildings: selectionarea.NewSelectionAreaService[*SelectableEntityStructure](),
		selectableUnits: []*SelectableEntityUnit{
			{body: physics.NewRectangle(physics.New(50, 100), 16, 16, 100)},
			{body: physics.NewRectangle(physics.New(58, 125), 16, 16, 100)},
			{body: physics.NewRectangle(physics.New(90, 112), 16, 16, 100)},
			{body: physics.NewRectangle(physics.New(150, 190), 16, 16, 100)},
		},

		selectableBuildings: []*SelectableEntityStructure{
			{body: physics.NewRectangle(physics.New(140, 249), 32, 32, 100)},
			{body: physics.NewRectangle(physics.New(142, 108), 32, 32, 100)},
			{body: physics.NewRectangle(physics.New(200, 1400), 32, 32, 100)},
			{body: physics.NewRectangle(physics.New(210, 490), 32, 32, 100)},
		},
	}
}
