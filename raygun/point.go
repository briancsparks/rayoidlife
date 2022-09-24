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
    r:    10,
    Mass: 1,
  }

  return &pt, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Update() {
  pt.Species.UpdateOne(pt)
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Draw() {
  pt.Species.DrawOne(pt)
}
