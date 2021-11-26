package pg

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/ebitenutil"
  "image/color"
  "log"
)


const (
  screenWidth, screenHeight = 640, 480
  zoom = 2
)


var(
  green   = color.RGBA{10, 255, 50, 255}

  oid     = NewCritter()
  //oid       Critter0
)

type Game struct{}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  screen.Fill(color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
  ebitenutil.DebugPrint(screen, "Hello, World!")

  screen.Set(int(oid.loc.X), int(oid.loc.Y), green)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
  return screenWidth, screenHeight
}

func EbMain() {
  ebiten.SetWindowSize(screenWidth * zoom, screenHeight * zoom)
  ebiten.SetWindowTitle("Hello, World!")
  if err := ebiten.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }

}


