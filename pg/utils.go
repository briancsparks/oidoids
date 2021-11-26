package pg


func clamp(x, low, high float64) float64  {
  if x < low {
    return low
  }

  if x > high {
    return high
  }

  return x
}

func iclamp(x, low, high int32) int32  {
  if x < low {
    return low
  }

  if x > high {
    return high
  }

  return x
}

