package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/kaschula/twod/draw"
	"github.com/kaschula/twod/physics"
	"time"
)

var (
	spaceShipColour = draw.RGBA(115, 169, 173, 255)
	obsticleColour  = draw.RGBA(144, 200, 172, 255)
)

// SpaceShipGame create a poly shape and move it around by increasing / decreasing the rigid body acceleration
type SpaceShipGame struct {
	spaceShip physics.RigidPoly
	// used a map for easy deleting
	projectiles    map[int]physics.RigidCircle
	count          int
	nextProjectile int64
}

func NewSpaceShipGame() *SpaceShipGame {
	spaceShip := physics.NewPoly(
		physics.New(600, 600), 3, 25,
		physics.WithMaxAccelerate(3),
		physics.WithSpeed(3.0),
	)

	return &SpaceShipGame{
		spaceShip:   spaceShip,
		count:       0,
		projectiles: map[int]physics.RigidCircle{},
	}
}

func (s *SpaceShipGame) Update() error {
	if inpututil.KeyPressDuration(ebiten.KeyArrowUp) > 0 {
		s.spaceShip.Accelerate(0.2)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		s.spaceShip.Rotate(15)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		s.spaceShip.Rotate(-15)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.fireProjectile()
	}

	s.spaceShip.Accelerate(-0.03)
	s.spaceShip.Update()

	for i, p := range s.projectiles {
		p.Update()

		projectileLocation := p.Location()
		if projectileLocation.OutOfBounds(screenW, screenH) {
			delete(s.projectiles, i)
		}
	}

	return nil
}

func (s *SpaceShipGame) fireProjectile() {
	now := time.Now()
	if now.Unix() < s.nextProjectile {
		return
	}

	s.nextProjectile = now.Add(800 * time.Millisecond).Unix()

	bullet := physics.NewCircle(s.spaceShip.RotatedEdgePoint(), 3, 3, physics.WithAngle(s.spaceShip.Direction()), physics.WithAcceleration(1), physics.WithSpeed(30))
	s.projectiles[s.count] = bullet
	s.count++
}

func (s *SpaceShipGame) Draw(screen *ebiten.Image) {
	// to do next
	// add poly rotation and movement for triangle and mo
	// add projectile firing

	// triangle
	spaceShip := s.spaceShip
	draw.Polygon(screen, spaceShip, spaceShipColour)

	for _, p := range s.projectiles {
		draw.Circle(screen, p, colourRed())
	}
}

func (s SpaceShipGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenW, screenH
}
