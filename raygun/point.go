package raygun

// -------------------------------------------------------------------------------------------------------------------

type Point struct {
  X, Y float32
  Dx, Dy float32

  Species *Species
}

// -------------------------------------------------------------------------------------------------------------------

func NewPoint() (*Point, error) {
  return NewPointAt(float32(CurrentScreenMidX), float32(CurrentScreenMidY), 1, 1)
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointGoing(dx, dy float32) (*Point, error) {
  return NewPointAt(float32(CurrentScreenMidX), float32(CurrentScreenMidY), dx, dy)
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointAt(x, y, dx, dy float32) (*Point, error) {
  pt := Point{
    X:     x,
    Y:     y,
    Dx:    dx,
    Dy:    dy,
    //Color: rl.White,
  }

  return &pt, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Update() {
  //pt.X += pt.Dx
  //pt.Y += pt.Dy
  pt.Species.UpdateOne(pt)
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Draw() {
  //rl.DrawCircle(pt.X, pt.Y, 10, pt.Species.Color)
  pt.Species.DrawOne(pt)
}
