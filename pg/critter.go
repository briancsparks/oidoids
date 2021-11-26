package pg

import "github.com/hajimehoshi/ebiten/v2"

type  Critter struct {
  loc    Vec2d /* location */
  vel    Vec2d /* velocity */
  id     int
  health float64
}

// NewCritter creates a new Critter.
func NewCritter() *Critter {
  c := Critter{
    loc: Vec2d{X: 30, Y: 30},
    vel: Vec2d{X: 0, Y: 0},
    id: 0,
    health: 0.0,
  }
  return &c
}

// The draw function draws the critter on an ebiten.Image.
func (critter *Critter) draw(screen *ebiten.Image, g *Game) error {
  screen.Set(int(critter.loc.X), int(critter.loc.Y), green)
  screen.Set(int(critter.loc.X+1), int(critter.loc.Y), green)
  screen.Set(int(critter.loc.X), int(critter.loc.Y+1), green)
  screen.Set(int(critter.loc.X+1), int(critter.loc.Y+1), green)

  return nil
}


