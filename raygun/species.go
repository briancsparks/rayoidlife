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

  QuasiType string

  Rules   map[string]*Rules
}

// -------------------------------------------------------------------------------------------------------------------

var allSpecies map[string]*Species

func init() {
  allSpecies = map[string]*Species{}
}

// -------------------------------------------------------------------------------------------------------------------

func NewSpecies(name string, color color.RGBA) (*Species, error) {
  s := &Species{
    Name:  name,
    Color: color,
    Rules: map[string]*Rules{},
  }

  allSpecies[name] = s

  // Defaults for species interactions
  for _, species := range allSpecies {
    if species == s {
      // Self
      s.InteractWith(s, &Rules{
        Attraction: TheGlobalRules.SelfAttractionDef,
        Radius:     TheGlobalRules.SelfRadiusDef,
      })
      continue
    }

    // Ignore each other
    species.InteractWith(s, Ignore)
    s.InteractWith(species, Ignore)
  }


  return s, nil
}

// -------------------------------------------------------------------------------------------------------------------

func NewQuasiSpecies(name string) (*Species, error) {
  s, err := NewSpecies(name, rl.DarkPurple)
  s.QuasiType = name

  return s, err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointAt(x, y float32) (*Point, error) {
  pt, err := NewPointAt(x, y, 0.0, 0.0)
  return s.integrate(pt), err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointGoing(dx, dy float32) (*Point, error) {
  pt, err := NewPointGoing(dx, dy)
  return s.integrate(pt), err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointAtGoing(x, y, dx, dy float32) (*Point, error) {
  pt, err := NewPointAt(x, y, dx, dy)
  return s.integrate(pt), err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointsAt(x, y float32, n int) {
  for i := 0; i < n; i++ {
    pt, _ := NewPointAt(x, y, 0.0, 0.0)
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointsGoing(dx, dy float32, n int) {
  for i := 0; i < n; i++ {
    pt, _ := NewPointGoing(dx, dy)
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointsAtGoing(x, y, dx, dy float32, n int) {
  for i := 0; i < n; i++ {
    pt, _ := NewPointAt(x, y, dx, dy)
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePoints(n int) {
  for i := 0; i < n; i++ {
    pos := rl.Vector2{X: randUpToN(CurrentScreenWidth), Y: randUpToN(CurrentScreenHeight)}
    pt, _ := NewPointAtV(pos, rl.Vector2{})
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakeBigPoints(n int, size float32) {
  for i := 0; i < n; i++ {
    pos := rl.Vector2{X: randUpToN(CurrentScreenWidth), Y: randUpToN(CurrentScreenHeight)}
    pt, _ := NewPointAtV(pos, rl.Vector2{})
    pt.Mass = size
    pt.r *= float32(math.Log10(float64(size * 10)))
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakeBigPointsAt(n int, size float32, pos rl.Vector2 /*, x, y float32*/) {
  for i := 0; i < n; i++ {
    pt, _ := NewPointAtV(pos, rl.Vector2{})
    pt.Mass = size
    pt.r *= float32(math.Log10(float64(size * 10)))
    s.integrate(pt)
  }
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

func UpdateAllSpecies() {
  for _, species := range allSpecies {
    species.Update()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func DrawAllSpecies() {
  for _, species := range allSpecies {
    species.Draw()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Update() {
  for _, point := range s.Points {
    //point.Update()

    fx, fy := float32(0.0), float32(0.0)

    for otherColor, rules := range s.Rules {
      other := allSpecies[otherColor]
      grav := rules.Attraction * TheGlobalRules.GravPerAttr
      rulesDistSq := rules.Radius * rules.Radius

      // -------------- Loop over other group
      for _, otherPt := range other.Points {
        if point == otherPt {
          continue
        }

        dist := rl.Vector2Subtract(point.pos, otherPt.pos)
        pairDistSq2 := rl.Vector2LenSqr(dist)
        if pairDistSq2 > rulesDistSq {
         continue
        }

        if pairDistSq2 == 0.0 {
         continue
        }

        pairDist2 := float32(math.Sqrt(float64(pairDistSq2)))
        fx += otherPt.Mass * dist.X / pairDist2
        fy += otherPt.Mass * dist.Y / pairDist2
      }

      point.vel = rl.Vector2Add(point.vel, rl.Vector2{X: fx * grav, Y: fy * grav})
    }

    // ---------- Finalize computations ----------

    clampV2(&point.vel, TheGlobalRules.MaxVelocity)

    // Update position
    point.pos = rl.Vector2Add(point.pos, point.vel)

    // Bounce off edges
    if clamped(&point.pos.X, 0, float32(CurrentScreenWidth)) {
     point.vel.X *= -1
    }
    if clamped(&point.pos.Y, 0, float32(CurrentScreenHeight)) {
     point.vel.Y *= -1
    }

  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) UpdateOne(pt *Point) {
  pt.pos = rl.Vector2Add(pt.pos, pt.vel)
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Draw() {
  for _, point := range s.Points {
    point.Draw()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) DrawOne(pt *Point) {
  if pt.Species.QuasiType == "center" {
   rl.DrawCircleV(pt.pos, pt.r + 3, rl.Black)
  }
  rl.DrawCircleV(pt.pos, pt.r, pt.Species.Color)
}

