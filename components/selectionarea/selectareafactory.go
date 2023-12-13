package selectionarea

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/containers"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/inputs"
	"github.com/kaschula/twod/physics"
	"image/color"
)

type SelectionAreaFactory struct {
	topLeft         physics.V
	bottomRight     physics.V
	latestSelection physics.Rectangle
}

func NewSelectionAreaFactory() *SelectionAreaFactory {

	return &SelectionAreaFactory{
		topLeft:         physics.ZeroVector,
		bottomRight:     physics.ZeroVector,
		latestSelection: nil,
	}
}

func (s *SelectionAreaFactory) Update(mouseEvents []inputs.InputEvent) containers.Maybe[physics.Rectangle] {
	if len(mouseEvents) == 0 {
		// selection has finished
		latestSelection := s.latestSelection

		// reset selection
		s.topLeft = physics.ZeroVector
		s.bottomRight = physics.ZeroVector
		s.latestSelection = nil

		if latestSelection == nil {
			return containers.Nothing[physics.Rectangle]()
		}

		return containers.Just(latestSelection)
	}

	event := mouseEvents[0]
	if !event.IsMouse() {
		return containers.Nothing[physics.Rectangle]()
	}

	mEvent := event.GetMouse()
	if mEvent.Button() == ebiten.MouseButtonLeft && mEvent.Action() == inputs.MouseButtonPressed {
		if s.topLeft.Equal(physics.ZeroVector) {
			s.topLeft = event.GetMouse().At()
			s.bottomRight = event.GetMouse().At()
		} else {
			s.bottomRight = event.GetMouse().At()
		}
	}

	if !s.topLeft.Equal(physics.ZeroVector) || !s.bottomRight.Equal(physics.ZeroVector) {
		r := physics.NewRectangleFromCorners(s.topLeft, s.bottomRight, 0)
		if r != nil {
			s.latestSelection = physics.NewRectangleFromCorners(s.topLeft, s.bottomRight, 0)
		}

	}

	return containers.Nothing[physics.Rectangle]()
}

func (s *SelectionAreaFactory) Draw(screen *ebiten.Image, colour color.Color) {
	if s.latestSelection != nil {
		draw.Polygon(screen, s.latestSelection, colour)
	}
}
