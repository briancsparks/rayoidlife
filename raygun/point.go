package raygun

import rl "github.com/gen2brain/raylib-go/raylib"

// -------------------------------------------------------------------------------------------------------------------

type Point struct {
  pos, vel    rl.Vector2
  r           float32
  Mass        float32

  Species *Species
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointAtV(pos, vel rl.Vector2) (*Point, error) {
  pt := Point{
    pos:  pos,
    vel:  vel,
    r:    6,
    Mass: 1,
  }

  return &pt, nil
}

// -------------------------------------------------------------------------------------------------------------------

// Update can update the Point, and returns true if it does, false otherwise.
func (pt *Point) Update(st *ComputeStats) bool {
  return false
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Draw() {
  if pt.Species.QuasiType == "center" {
    rl.DrawCircleV(pt.pos, pt.r + 3, rl.Black)
  }
  rl.DrawCircleV(pt.pos, pt.r, pt.Species.Color)
}
