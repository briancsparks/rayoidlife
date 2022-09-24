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

func NewPoint() (*Point, error) {
  return NewPointAt(float32(CurrentScreenMidX), float32(CurrentScreenMidY), 1, 1)
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointGoing(xvel, yvel float32) (*Point, error) {
  return NewPointAt(float32(CurrentScreenMidX), float32(CurrentScreenMidY), xvel, yvel)
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointAt(x, y, xvel, yvel float32) (*Point, error) {
  pt := Point{
    pos:  rl.Vector2{X: x, Y: y},
    vel:  rl.Vector2{X: xvel, Y: yvel},

    r:    10,
    Mass: 1,
  }

  return &pt, nil
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
  //pt.X += pt.xvel
  //pt.Y += pt.yvel
  pt.Species.UpdateOne(pt)
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Draw() {
  //rl.DrawCircle(pt.X, pt.Y, 10, pt.Species.Color)
  pt.Species.DrawOne(pt)
}
