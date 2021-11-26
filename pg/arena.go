package pg

import (
  "github.com/hajimehoshi/ebiten/v2"
  "image/color"
)

var (
 red     = color.RGBA{R: 255, G: 10, B: 10, A: 255}
 //bgImage = ebiten.NewImage(arenaWidth, arenaHeight)
)

//func init() {
//  bgImage.Fill(red)
//}

type Arena struct {
  dim IVec2d    /* dimension of the arena */
  bgImage *ebiten.Image
}

func NewArena(dim IVec2d) *Arena {
  w, h := dim.Xy()

  a := Arena{
    dim:     dim,
    bgImage: ebiten.NewImage(int(w), int(h)),
  }

  a.bgImage.Fill(red)

  return &a
}


// The draw function draws the critter on an ebiten.Image.
func (arena *Arena) draw(screen *ebiten.Image, g *Game) error {

  screen.DrawImage(arena.bgImage, nil)

  return nil
}


