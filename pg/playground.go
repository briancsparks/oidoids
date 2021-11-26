package pg

import (
  "github.com/hajimehoshi/ebiten/v2"
  "image/color"
  "log"
)


type Drawinger interface {
  draw(screen *ebiten.Image, g *Game) error
}


const (
  screenWidth, screenHeight = 640, 480
  zoom = 2
)

var(
  green   = color.RGBA{10, 255, 50, 255}

  oid Drawinger
)

type Game struct{}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  //screen.Fill(color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
  //ebitenutil.DebugPrint(screen, "Hello, World!")

  _ = g.drawCritter(screen, oid)

}

func (g *Game) drawCritter(screen *ebiten.Image, dinger Drawinger) error {
  return dinger.draw(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
  return screenWidth, screenHeight
}

func EbMain() {
  var dinger Drawinger = &Critter{
    loc:    Vec2d{X: 90, Y: 90},
    vel:    Vec2d{X: 0, Y: 0},
    id:     0,
    health: 0.0,
  }

  oid = dinger

  ebiten.SetWindowSize(screenWidth * zoom, screenHeight * zoom)
  ebiten.SetWindowTitle("Hello, World!")
  if err := ebiten.RunGame(&Game{}); err != nil {
    log.Fatal(err)
  }

}

