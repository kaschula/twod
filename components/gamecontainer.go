package components

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// GameContainer is object that can cycle through ebiten games, allowing to switch between different game object
// cycling is done with left and right bracket keys
// game state persisted between cycling
type GameContainer struct {
	current     int
	games       []ebiten.Game
	left, right ebiten.Key
}

func NewGameContainer(games []ebiten.Game, left, right ebiten.Key) *GameContainer {
	return &GameContainer{games: games, left: left, right: right}
}

func (g *GameContainer) Update() error {
	if inpututil.IsKeyJustPressed(g.left) {
		fmt.Println("left key", g.current)
		g.previous()
	}

	if inpututil.IsKeyJustPressed(g.right) {
		g.next()
	}

	return g.currentGame().Update()
}

func (g *GameContainer) Draw(screen *ebiten.Image) {
	g.currentGame().Draw(screen)
}

func (g *GameContainer) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.currentGame().Layout(outsideWidth, outsideHeight)
}

func (g *GameContainer) next() {
	g.current = g.cycle(g.current + 1)
}

func (g *GameContainer) previous() {
	g.current = g.cycle(g.current - 1)
}

func (g *GameContainer) cycle(n int) int {
	highestIndex := len(g.games) - 1
	if n > highestIndex {
		return 0
	}

	if n < 0 {
		return highestIndex
	}

	return n
}

func (g *GameContainer) currentGame() ebiten.Game {
	return g.games[g.current]
}
