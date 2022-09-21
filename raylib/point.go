package raylib

type Point struct {
  X, Y int32
  Dx, Dy int32
}

func NewPoint() (*Point, error) {
  return NewPointAt(CurrentScreenMidX, CurrentScreenMidY, 1, 1)
}

func NewPointGoing(dx, dy int32) (*Point, error) {
  return NewPointAt(CurrentScreenMidX, CurrentScreenMidY, dx, dy)
}

func NewPointAt(x, y, dx, dy int32) (*Point, error) {
  pt := Point{
    X:  x,
    Y:  y,
    Dx: dx,
    Dy: dy,
  }

  return &pt, nil
}
