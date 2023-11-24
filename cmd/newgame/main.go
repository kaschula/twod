package main

import (
	"fmt"
	"log"
	"os"
)

var MainFileTemplat = `package main

import (
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenW, screenH = 600, 600
)

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetWindowTitle("New Game")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	
}

func (g *Game) Update() error {
	return errors.New("implement update")
}

func (g *Game) DrawUnit(screen *ebiten.Image) {
	panic("implement draw")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}

func NewGame() *Game {
	return &Game{}
}
`

// go run cmd/newgame/main.go examples/nameofgame
func main() {
	// create directory and template file directory

	args := os.Args

	if len(args) != 2 {
		log.Fatalf("required 2 args got %v", len(args)-1)
	}

	directory := args[1]

	fmt.Println(directory)

	err := os.Mkdir(directory, os.ModePerm)
	if err != nil {
		log.Fatalf("mkdir error %s", err.Error())
	}

	f, err := os.Create(directory + "/main.go")
	if err != nil {
		log.Fatalf("create main.go %s", err.Error())
	}

	_, err = f.Write([]byte(MainFileTemplat))
	if err != nil {
		log.Fatalf("write to main.go %s", err.Error())
	}

}
