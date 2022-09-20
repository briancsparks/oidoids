package pg

import (
  "github.com/hajimehoshi/ebiten/v2"
  "image/color"
  "log"
  "math/rand"
  "time"
)


type Drawinger interface {
  draw(screen *ebiten.Image, g *Game) error
}

type ZoomMode int

const (
  OneToOne    ZoomMode = 1
  FourSquare  ZoomMode = 4
  BigFour     ZoomMode = 5
  Zoom        ZoomMode = iota + 5
)

func (z ZoomMode) String() string {
  return [...]string{"OneToOne", "FourSquare", "Zoom"}[z]
}


const (
  arenaWidth, arenaHeight  = 640, 480
  screenWidth, screenHeight = 640, 480
  zoom = 2
)

var(
  green   = color.RGBA{R: 10, G: 255, B: 50, A: 255}
  blue    = color.RGBA{B: 255, A: 255}
  cyan    = color.RGBA{G: 255, B: 255, A: 255}
  yellow  = color.RGBA{R: 255, G: 255, A: 255}
  red     = color.RGBA{R: 255, G: 10, B: 10, A: 255}

  arena *Arena
  oid Drawinger

  //
  zoommode ZoomMode = BigFour
  critterColor = blue
  backgroundColor = cyan
)

type Game struct{
  oids []*Critter
}

func newGame() *Game {
  g := Game{
    oids: make([]*Critter, 2),
  }

  for i := 0; i < len(g.oids); i++ {
    g.oids[i] = &Critter{
      loc: Vec2d{X: float64(rand.Intn(arenaWidth)), Y: float64(rand.Intn(arenaHeight))},
      id: i,
      health: 100,
    }

    g.oids[i].run()
  }

  return &g
}

func (g *Game) Update() error {
  return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
  //screen.Fill(color.RGBA{R: 0xff, G: 0, B: 0, A: 0xff})
  //ebitenutil.DebugPrint(screen, "Hello, World!")

  _ = arena.draw(screen, g)
  //_ = g.drawCritter(screen, oid)

  for _, critter := range g.oids {
    critter.draw(screen, g)
  }

}

func (g *Game) drawCritter(screen *ebiten.Image, dinger Drawinger) error {
  return dinger.draw(screen, g)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (w, h int) {
  return screenWidth, screenHeight
}

func EbMain() {
  rand.Seed(time.Now().UnixNano())
  game := newGame()

  var dinger Drawinger = &Critter{
    loc:    Vec2d{X: 90, Y: 90},
    //vel:    Vec2d{X: 0, Y: 0},
    id:     0,
    health: 0.0,
  }

  oid = dinger

  arena = NewArena(IVec2d{
    X: arenaWidth,
    Y: arenaHeight,
  })

  //ebiten.SetWindowSize(screenWidth * zoom, screenHeight * zoom)
  ebiten.SetWindowSize(screenWidth, screenHeight)
  ebiten.SetWindowTitle("Hello, World!")
  if err := ebiten.RunGame(game); err != nil {
    log.Fatal(err)
  }

}

