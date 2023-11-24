package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type KeyEvent struct {
	key    ebiten.Key
	action Action
}

func NewKeyEvent(key ebiten.Key, action Action) KeyEvent {
	return KeyEvent{key: key, action: action}
}

func (e KeyEvent) Key() ebiten.Key {
	return e.key
}

func (e KeyEvent) Action() Action {
	return e.action
}

type KeyProducer struct {
	keys []ebiten.Key
}

func NewKeyProducer(keys ...ebiten.Key) *KeyProducer {
	return &KeyProducer{keys: keys}
}

func (producer *KeyProducer) Update() []InputEvent {
	var events []InputEvent
	for _, key := range producer.keys {
		if KeyPressAndHold(key) {
			events = append(events, NewInputKeyEvent(NewKeyEvent(key, KeyHeld)))
			continue
		}

		if KeyJustPressed(key) {
			events = append(events, NewInputKeyEvent(NewKeyEvent(key, KeyPressed)))
		}
	}

	return events
}
