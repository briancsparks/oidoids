package pg

import "math"

type IVec2d struct {
  X int32
  Y int32
}

func (v1 IVec2d) Add(v2 IVec2d) IVec2d {
  return IVec2d{v1.X + v2.X, v1.Y + v2.Y}
}

func (v1 IVec2d) Subtract(v2 IVec2d) IVec2d {
  return IVec2d{v1.X - v2.X, v1.Y - v2.Y}
}

func (v1 IVec2d) Multiply(v2 IVec2d) IVec2d {
  return IVec2d{v1.X * v2.X, v1.Y * v2.Y}
}


//func (v1 IVec2d) AddV(d int32) IVec2d {
//  return IVec2d{v1.x + d, v1.y + d}
//}

func (v1 IVec2d) ScaleV(d int32) IVec2d {
  return IVec2d{v1.X * d, v1.Y * d}
}

func (v1 IVec2d) DivV(d int32) IVec2d {
  return IVec2d{v1.X / d, v1.Y / d}
}

//func (v1 IVec2d) Limit(lower, upper int32) IVec2d {
//  return IVec2d{math.Min(math.Max(v1.X, lower), upper),
//    math.Min(math.Max(v1.Y, lower), upper)}
//}

func (v1 *IVec2d) ClampX(low, high int32) IVec2d {
  return IVec2d{iclamp(v1.X, low, high), v1.Y}
}

func (v1 *IVec2d) ClampY(low, high int32) IVec2d {
  return IVec2d{iclamp(v1.X, low, high), v1.Y}
}

func (v1 IVec2d) Distance(v2 IVec2d) float64 {
  return math.Sqrt(math.Pow(float64(v1.X-v2.X), 2) + math.Pow(float64(v1.Y-v2.Y), 2))
}

func (v1 IVec2d) IsWithinSq(v2 IVec2d, d2 int32) bool {
  dx,dy := v1.Subtract(v2).Xy()

  return dx*dx+dy*dy < d2
}

func (v1 IVec2d) IsWithin(v2 IVec2d, d int32) bool {
  return v1.IsWithinSq(v2, d*d)
}

func (v1 IVec2d) Xy() (int32, int32) {
  return v1.X, v1.Y
}

