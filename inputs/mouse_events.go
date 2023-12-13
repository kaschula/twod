package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/containers"
	"github.com/kaschula/twod/physics"
	"time"
)

type MouseEvent struct {
	button ebiten.MouseButton
	action Action
	at     physics.V
}

func (e MouseEvent) Button() ebiten.MouseButton {
	return e.button
}

func (e MouseEvent) Action() Action {
	return e.action
}

func (e MouseEvent) At() physics.V {
	return e.at
}

func NewMouseEvent(button ebiten.MouseButton, action Action, at physics.V) MouseEvent {
	return MouseEvent{button: button, action: action, at: at}
}

func NewMouseInputProducer(currentMouseLocation func() physics.V, mouseLeftPressed, mouseRightPressed MousePressedFN, leftThrottle, rightThrottle time.Duration) *MouseInputProducer {
	leftThrottler := mouseLeftPressed
	if leftThrottle > 0 {
		leftThrottler = NewMousePressedThrottler(leftThrottle, mouseLeftPressed).Pressed
	}

	rightThrottler := mouseRightPressed
	if rightThrottle > 0 {
		rightThrottler = NewMousePressedThrottler(rightThrottle, mouseRightPressed).Pressed
	}

	return &MouseInputProducer{
		currentMouseLocation:    currentMouseLocation,
		leftPressed:             leftThrottler,
		rightPressed:            rightThrottler,
		wasLeftPressedLastFrame: false,
	}
}

type MouseInputProducer struct {
	leftPressed              MousePressedFN
	rightPressed             MousePressedFN
	currentMouseLocation     func() physics.V
	wasLeftPressedLastFrame  bool
	wasRightPressedLastFrame bool
}

func (producer *MouseInputProducer) Update() []InputEvent {
	return append(producer.left(), producer.right()...)
}

func (producer *MouseInputProducer) left() []InputEvent {
	var events []InputEvent
	if v, pressed := producer.leftPressed(); pressed {
		producer.wasLeftPressedLastFrame = true

		mouseEvent := NewMouseEvent(ebiten.MouseButtonLeft, MouseButtonPressed, v)
		return append(events, NewInputMouseEvent(mouseEvent))

	}

	if producer.wasLeftPressedLastFrame {
		mouseEvent := NewMouseEvent(ebiten.MouseButtonLeft, MouseButtonReleased, producer.currentMouseLocation())
		events = append(events, NewInputMouseEvent(mouseEvent))
		// reset
		producer.wasLeftPressedLastFrame = false

	}

	return events
}

func (producer *MouseInputProducer) right() []InputEvent {
	var events []InputEvent
	if v, pressed := producer.rightPressed(); pressed {
		producer.wasRightPressedLastFrame = true

		mouseEvent := NewMouseEvent(ebiten.MouseButtonRight, MouseButtonPressed, v)
		return append(events, NewInputMouseEvent(mouseEvent))

	}

	if producer.wasRightPressedLastFrame {
		mouseEvent := NewMouseEvent(ebiten.MouseButtonRight, MouseButtonReleased, producer.currentMouseLocation())
		events = append(events, NewInputMouseEvent(mouseEvent))
		// reset
		producer.wasRightPressedLastFrame = false

	}

	return events
}

func GetFirstMouseClickVector(button ebiten.MouseButton, events []InputEvent) containers.Maybe[physics.V] {
	if len(events) == 0 {
		return containers.Nothing[physics.V]()
	}

	for _, event := range events {
		if !event.IsMouse() {
			continue
		}

		mouseEvent := event.GetMouse()

		if mouseEvent.Button() != button {
			continue
		}

		return containers.Just(mouseEvent.At())
	}

	return containers.Nothing[physics.V]()
}
