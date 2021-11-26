package pg

type Critter interface {

}

type  Critter0 struct {
  loc    Vec2d /* location */
  vel    Vec2d /* velocity */
  id     int
  health float64
}

func NewCritter() *Critter0 {
  c := Critter0{
    loc: Vec2d{X: 30, Y: 30},
    vel: Vec2d{X: 0, Y: 0},
    id: 0,
    health: 0.0,
  }
  return &c
}
