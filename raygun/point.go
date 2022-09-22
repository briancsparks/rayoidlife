package raygun

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  "image/color"
)

// -------------------------------------------------------------------------------------------------------------------

type Point struct {
  X, Y int32
  Dx, Dy int32

  Color color.RGBA
}

// -------------------------------------------------------------------------------------------------------------------

func NewPoint() (*Point, error) {
  return NewPointAt(CurrentScreenMidX, CurrentScreenMidY, 1, 1)
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointGoing(dx, dy int32) (*Point, error) {
  return NewPointAt(CurrentScreenMidX, CurrentScreenMidY, dx, dy)
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointAt(x, y, dx, dy int32) (*Point, error) {
  pt := Point{
    X:     x,
    Y:     y,
    Dx:    dx,
    Dy:    dy,
    Color: rl.White,
  }

  return &pt, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Update() {
  pt.X += pt.Dx
  pt.Y += pt.Dy
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Draw() {
  rl.DrawCircle(pt.X, pt.Y, 10, pt.Color)
}
