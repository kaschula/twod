package gameobject

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/kaschula/twod/components/animation"
	"github.com/kaschula/twod/physics"
)

// todo for GameObject
// There is loads to finish about the state transitions but as proof of concpet it works
// one thing to do next is add the ability to have sub objects that are transposed at an off set vector, these will move as the centeral gameobjkect body moves around and rotates.
// An example of this is to have a small body that reproesents a spawning point for bullets for a top down player when the player has a gun

type GameObject struct {
	state      string
	states     map[string]interface{}
	animations map[string]animation.SpriteRenderer
	// This game object needs to be able to be Rectangle, Circle or Poly
	body     physics.RigidPoly
	updateFN func(gameObject *GameObject)
}

type OptionFN = func(gameObject *GameObject)

func WithRigidPoly(r physics.RigidPoly) OptionFN {
	return func(gameObject *GameObject) {
		gameObject.body = r
	}
}

// first state is initial state

func NewGameObject(states []string, animations map[string]animation.SpriteRenderer, updateFN func(gameObject *GameObject), options ...OptionFN) *GameObject {
	// args
	// withDefaultAnimation -> used if the state animation does not exist

	statesM := map[string]interface{}{}
	for _, state := range states {
		statesM[state] = nil
	}

	gameObject := &GameObject{state: states[0], states: statesM, animations: animations, updateFN: updateFN}

	// options
	for _, option := range options {
		option(gameObject)
	}
	// withRigidBody ->
	// withAnimations -> for drawing
	// withController -> ebiten controller | no op
	// add options

	return gameObject
}

func (gameObject *GameObject) SwitchState(state string) {
	if gameObject.state == state {
		return
	}

	_, ok := gameObject.states[state]
	if !ok {
		return
	}

	gameObject.state = state
}

func (gameObject *GameObject) Body() physics.RigidPoly {
	return gameObject.body
}

func (gameObject *GameObject) Update() {

	// todo eventually the animation could return if its first frame or last frame. You could have callabcks attached to the game object for certain animations.

	if gameObject.updateFN != nil {
		gameObject.updateFN(gameObject)
	}

	gameObject.animation().Update()
}

func (gameObject *GameObject) Draw(screen *ebiten.Image) {
	gameObject.animation().Draw(gameObject.body.Location(), screen)
}

func (gameObject *GameObject) animation() animation.SpriteRenderer {
	return gameObject.animations[gameObject.state]
}
