package raygun

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  "image/color"
  "math"
)

// -------------------------------------------------------------------------------------------------------------------

type Species struct {
  Name   string
  Points []*Point
  Color   color.RGBA

  Rules   map[string]*Rules
}

// -------------------------------------------------------------------------------------------------------------------

var allSpecies map[string]*Species

func init() {
  allSpecies = map[string]*Species{}
}

// -------------------------------------------------------------------------------------------------------------------

func NewSpecies(name string, color color.RGBA) (*Species, error) {
  s := Species{
    Name:  name,
    Color: color,
    Rules: map[string]*Rules{},
  }

  allSpecies[name] = &s

  return &s, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointGoing(dx, dy float32) (*Point, error) {
  pt, err := NewPointGoing(dx, dy)
  return s.integrate(pt), err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointAt(x, y, dx, dy float32) (*Point, error) {
  pt, err := NewPointAt(x, y, dx, dy)
  return s.integrate(pt), err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) integrate(pt *Point) *Point {

  pt.Species = s

  s.Points = append(s.Points, pt)
  return pt
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) InteractWith(other *Species, rules *Rules) {
  s.Rules[other.Name] = rules
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Update() {
  for _, point := range s.Points {
    //point.Update()

    fx, fy := float32(0.0), float32(0.0)

    for otherColor, rules := range s.Rules {
      other := allSpecies[otherColor]
      grav := rules.Attraction * TheGlobalRules.GravPerAttr
      rsq := rules.Radius * rules.Radius

      // -------------- Loop over other group
      for _, otherPt := range other.Points {
        dx := point.X - otherPt.X
        dy := point.Y - otherPt.Y
        pairSq := dx*dx + dy*dy
        if pairSq > rsq {
          continue
        }

        if pairSq != 0.0 {
          r := float32(math.Sqrt(float64(pairSq)))
          fx += dx / r
          fy += dy / r
        }

      }

      point.Dx += fx * grav
      point.Dy += fy * grav
    }

    // ---------- Finalize computations ----------

    // TODO: clamp velocity

    // Update position
    point.X += point.Dx
    point.Y += point.Dy

    // TODO: Keep within bounds
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) UpdateOne(pt *Point) {
  pt.X += pt.Dx
  pt.Y += pt.Dy
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Draw() {
  for _, point := range s.Points {
    point.Draw()
  }
}

func (s *Species) DrawOne(pt *Point) {
  rl.DrawCircle(int32(pt.X), int32(pt.Y), 10, pt.Species.Color)
}

