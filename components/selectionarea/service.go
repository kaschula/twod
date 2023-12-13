package selectionarea

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/containers"
	"github.com/kaschula/twod/inputs"
	"github.com/kaschula/twod/physics"
	collision2 "github.com/kaschula/twod/physics/collision"
	"image/color"
)

type Selectable interface {
	Center() physics.V
	Body() physics.Rectangle
	Select()
	Unselect()
}

type Service[T Selectable] interface {
	// Update returns true if something was selected, and returns a Maybe selection area
	Update(mouseEvents []inputs.InputEvent, selectables []T) (bool, containers.Maybe[physics.Rectangle])
	Draw(screen *ebiten.Image, colour color.Color)
}

type AreaSelectService[T Selectable] struct {
	selectionAreaFactory *SelectionAreaFactory
}

func NewSelectionAreaService[T Selectable]() *AreaSelectService[T] {
	return &AreaSelectService[T]{
		selectionAreaFactory: NewSelectionAreaFactory(),
	}
}

func (s AreaSelectService[T]) Update(mouseInputs []inputs.InputEvent, selectables []T) (bool, containers.Maybe[physics.Rectangle]) {
	somethingHasBeenSelected := false
	mLatestSelectionArea := s.selectionAreaFactory.Update(mouseInputs)

	if !mLatestSelectionArea.Nothing() {
		selectArea := mLatestSelectionArea.Just()
		for _, entity := range selectables {
			entity.Unselect()
			collision := collision2.PointInRectangle(selectArea, entity.Center())
			if collision {
				somethingHasBeenSelected = true
				entity.Select()
			}
		}
	} else if mMouseVector := inputs.GetFirstMouseClickVector(ebiten.MouseButtonLeft, mouseInputs); !mMouseVector.Nothing() {
		v := mMouseVector.Just()
		for _, entity := range selectables {
			entity.Unselect()
			collision := collision2.PointInRectangle(entity.Body(), v)
			if collision {
				somethingHasBeenSelected = true
				entity.Select()
			}
		}
	}

	if inputs.KeyJustPressed(ebiten.KeyEscape) {
		for _, entity := range selectables {
			entity.Unselect()
			somethingHasBeenSelected = false
		}
	}

	return somethingHasBeenSelected, mLatestSelectionArea
}

func (s AreaSelectService[T]) Draw(screen *ebiten.Image, colour color.Color) {
	s.selectionAreaFactory.Draw(screen, colour)
}
