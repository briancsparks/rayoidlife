package raygun

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  "image/color"
  "log"
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
var speciesQuadTrees map[string]*HQuadTree

func init() {
  allSpecies = map[string]*Species{}
  speciesQuadTrees = map[string]*HQuadTree{}
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
      s.InteractWith(s, NewRulesWithSep(
        TheGlobalRules.SelfAttractionDef,
        TheGlobalRules.SelfRadiusDef,
        TheGlobalRules.SelfSepFactorDef,
        TheGlobalRules.SelfSepRadiusDef,
      ))
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

func (s *Species) MakePoints(n int) {
  for i := 0; i < n; i++ {
    pos := rl.Vector2{X: randUpTo(CurrentScreenWidth), Y: randUpTo(CurrentScreenHeight)}
    vel := rl.Vector2{}
    if s.QuasiType != "center" {
      vel = randVector2(MaxInitialVelocity())
    }
    pt, _ := NewPointAtV(pos, vel)
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakeBigPoints(n int, size float32) {
  for i := 0; i < n; i++ {
    pos := rl.Vector2{X: randUpTo(CurrentScreenWidth), Y: randUpTo(CurrentScreenHeight)}
    vel := rl.Vector2{}
    if s.QuasiType != "center" {
      vel = randVector2(MaxInitialVelocity())
    }
    pt, _ := NewPointAtV(pos, vel)
    pt.Mass = size
    pt.r *= float32(math.Log10(float64(size * 10)))
    s.integrate(pt)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakeBigPointsAt(n int, size float32, pos rl.Vector2 /*, x, y float32*/) {
  for i := 0; i < n; i++ {
    vel := rl.Vector2{}
    if s.QuasiType != "center" {
      vel = randVector2(MaxInitialVelocity())
    }
    pt, _ := NewPointAtV(pos, vel)
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
  for name, species := range allSpecies {
    quadTree := NewQuadTree(0, 0, CurrentScreenWidth, CurrentScreenHeight, species.Color)
    quadTree.addPoints(species.Points)

    quadCount := quadTree.count()
    specCount := len(species.Points)
    if quadCount != specCount {
      log.Panic("Wrong counts")
    }

    speciesQuadTrees[name] = quadTree
  }

  for _, species := range allSpecies {
    species.Update()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func DrawAllSpecies() {
  for _, species := range allSpecies {
    species.Draw()
  }

  for _, tree := range speciesQuadTrees {
    tree.Draw()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Update() {
  for _, point := range s.Points {
    // Give the point a chance to do its own update
    if point.Update() {
      continue
    }

    for otherColor, rules := range s.Rules {
      other := allSpecies[otherColor]
      grav := rules.Attraction * TheGlobalRules.GravPerAttr
      rulesDistSq := rules.Radius * rules.Radius
      rulesSepDistSq := rules.SepRadius * rules.SepRadius

      // TODO: Make a Vector2
      fx, fy := float32(0.0), float32(0.0)

      // -------------- Loop over other group
      for _, otherPt := range other.Points {
        if point == otherPt {
          continue
        }

        fxOther, fyOther := float32(0.0), float32(0.0)

        dist := rl.Vector2Subtract(point.pos, otherPt.pos)
        pairDistSq := rl.Vector2LenSqr(dist)

        // Attraction
        if !TheGlobalRules.SkipAttractionRule {

          if pairDistSq <= rulesDistSq && pairDistSq != 0.0 {
            pairDist := float32(math.Sqrt(float64(pairDistSq)))
            fxOther += otherPt.Mass * dist.X / pairDist
            fyOther += otherPt.Mass * dist.Y / pairDist
          }
        }

        // Separation
        if !TheGlobalRules.SkipSeparationRule {

          if pairDistSq <= rulesSepDistSq && rules.SepFactor != 1 {
            // We are too close
            fxOther *= rules.SepFactor
            fyOther *= rules.SepFactor
          }
        }

        fx += fxOther
        fy += fyOther
      }

      point.vel = rl.Vector2Add(point.vel, rl.Vector2{X: fx * grav, Y: fy * grav})
    }

    // ---------- Finalize computations ----------

    clampV2(&point.vel, TheGlobalRules.MaxVelocity)

    // Update position
    point.pos = rl.Vector2Add(point.pos, point.vel)

    // Bounce off edges
    if clamped(&point.pos.X, 0, CurrentScreenWidth) {
     point.vel.X *= -1
    }
    if clamped(&point.pos.Y, 0, CurrentScreenHeight) {
     point.vel.Y *= -1
    }

  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Draw() {
  for _, point := range s.Points {
    point.Draw()
  }
}

