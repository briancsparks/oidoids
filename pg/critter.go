package pg

import (
  "github.com/hajimehoshi/ebiten/v2"
  "time"
)

type  Critter struct {
  loc    Vec2d /* location */
  id     int
  health float64
}

// NewCritter creates a new Critter.
func NewCritter() *Critter {
  c := Critter{
    loc: Vec2d{X: 30, Y: 30},
    id: 0,
    health: 0.0,
  }
  return &c
}

func (c *Critter) run() {
  go func() {

    for {
      c.loc = c.loc.Add(Vec2d{
        X: 1,
        Y: 1,
      })

      //time.Sleep(250 * time.Millisecond)
      time.Sleep(1000 * time.Millisecond)
    }

  }()
}

// The draw function draws the critter on an ebiten.Image.
func (c *Critter) draw(screen *ebiten.Image, g *Game) error {

  if zoommode == FourSquare {
    screen.Set(int(c.loc.X),   int(c.loc.Y),   critterColor)
    screen.Set(int(c.loc.X+1), int(c.loc.Y),   critterColor)
    screen.Set(int(c.loc.X),   int(c.loc.Y+1), critterColor)
    screen.Set(int(c.loc.X+1), int(c.loc.Y+1), critterColor)

  } else if zoommode == BigFour {
    screen.Set(int(c.loc.X-1), int(c.loc.Y-1), critterColor)
    screen.Set(int(c.loc.X+1), int(c.loc.Y-1), critterColor)
    screen.Set(int(c.loc.X-1), int(c.loc.Y+1), critterColor)
    screen.Set(int(c.loc.X+1), int(c.loc.Y+1), critterColor)

    screen.Set(int(c.loc.X-2), int(c.loc.Y-2), critterColor)
    screen.Set(int(c.loc.X+2), int(c.loc.Y-2), critterColor)
    screen.Set(int(c.loc.X-2), int(c.loc.Y+2), critterColor)
    screen.Set(int(c.loc.X+2), int(c.loc.Y+2), critterColor)

  } else {
    screen.Set(int(c.loc.X),   int(c.loc.Y),   critterColor)
    screen.Set(int(c.loc.X+1), int(c.loc.Y),   critterColor)
    screen.Set(int(c.loc.X),   int(c.loc.Y+1), critterColor)
    screen.Set(int(c.loc.X+1), int(c.loc.Y+1), critterColor)

  }

  return nil
}


