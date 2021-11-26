package pg

import (
  "math"
)

type Vec2d struct {
  X float64
  Y float64
}

func (v1 Vec2d) Add(v2 Vec2d) Vec2d {
  return Vec2d{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 Vec2d) Subtract(v2 Vec2d) Vec2d {
  return Vec2d{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 Vec2d) Multiply(v2 Vec2d) Vec2d {
  return Vec2d{v1.X * v2.X, v1.Y * v2.Y}
}


//func (v1 Vec2d) AddV(d float64) Vec2d {
//  return Vec2d{v1.x + d, v1.y + d}
//}

func (v1 Vec2d) ScaleV(d float64) Vec2d {
  return Vec2d{v1.X * d, v1.Y * d}
}

func (v1 Vec2d) DivV(d float64) Vec2d {
  return Vec2d{v1.X / d, v1.Y / d}
}

func (v1 Vec2d) Limit(lower, upper float64) Vec2d {
  return Vec2d{math.Min(math.Max(v1.X, lower), upper),
    math.Min(math.Max(v1.Y, lower), upper)}
}

func (v1 *Vec2d) ClampX(low, high float64) Vec2d {
  return Vec2d{clamp(v1.X, low, high), v1.Y}
}

func (v1 *Vec2d) ClampY(low, high float64) Vec2d {
  return Vec2d{clamp(v1.X, low, high), v1.Y}
}

func (v1 Vec2d) Distance(v2 Vec2d) float64 {
  return math.Sqrt(math.Pow(v1.X-v2.X, 2) + math.Pow(v1.Y-v2.Y, 2))
}

func (v1 Vec2d) IsWithinSq(v2 Vec2d, d2 float64) bool {
  dx,dy := v1.Subtract(v2).Xy()

  return dx*dx+dy*dy < d2
}

func (v1 Vec2d) IsWithin(v2 Vec2d, d float64) bool {
  return v1.IsWithinSq(v2, d*d)
}

func (v1 Vec2d) Xy() (float64, float64) {
  return v1.X, v1.Y
}
