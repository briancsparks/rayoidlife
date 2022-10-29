package raygun

import (
  "fmt"
  rl "github.com/gen2brain/raylib-go/raylib"
  "image/color"
  "log"
  "math"
)

// -------------------------------------------------------------------------------------------------------------------

type SpeciesCohort struct {
  Species *Species
  CoName  string
  Points  []*Point
  QuadTree  *HQuadTree
}

// -------------------------------------------------------------------------------------------------------------------

type Species struct {
  Name    string
  Color   color.RGBA
  Cohorts   map[string]*SpeciesCohort

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
    Name:      name,
    Color:     color,
    Cohorts:   map[string]*SpeciesCohort{},
    QuasiType: "",
    Rules:     map[string]*Rules{},
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

func (s *Species) MakePoints(n int) *SpeciesCohort {
  coNum := len(s.Cohorts)
  sco := SpeciesCohort{
    Species: s,
    CoName:  fmt.Sprintf("%s-%02d", s.Name, coNum),
    Points:  nil,
    QuadTree: nil,
  }

  // Generate points
  for i := 0; i < n; i++ {
    pos := rl.Vector2{X: randUpTo(CurrentScreenWidth), Y: randUpTo(CurrentScreenHeight)}
    vel := rl.Vector2{}
    if s.QuasiType != "center" {
      vel = randVector2(MaxInitialVelocity())
    }
    pt, _ := NewPointAtV(pos, vel)
    sco.integrate(pt)
  }

  // Add the cohort to the list of cohorts in Species
  s.Cohorts[sco.CoName] = &sco

  return &sco
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakeBigPoints(n int, size float32) *SpeciesCohort {
  coNum := len(s.Cohorts)
  sco := SpeciesCohort{
    Species: s,
    CoName:  fmt.Sprintf("%s-%02d", s.Name, coNum),
    Points:  nil,
    QuadTree: nil,
  }

  // Generate points
  for i := 0; i < n; i++ {
    pos := rl.Vector2{X: randUpTo(CurrentScreenWidth), Y: randUpTo(CurrentScreenHeight)}
    vel := rl.Vector2{}
    if s.QuasiType != "center" {
      vel = randVector2(MaxInitialVelocity())
    }
    pt, _ := NewPointAtV(pos, vel)
    pt.Mass = size
    pt.r *= float32(math.Log10(float64(size * 10)))
    sco.integrate(pt)
  }

  // Add the cohort to the list of cohorts in Species
  s.Cohorts[sco.CoName] = &sco

  return &sco
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakeBigPointsAt(n int, size float32, pos rl.Vector2 /*, x, y float32*/) *SpeciesCohort {
  coNum := len(s.Cohorts)
  sco := SpeciesCohort{
    Species: s,
    CoName:  fmt.Sprintf("%s-%02d", s.Name, coNum),
    Points:  nil,
    QuadTree: nil,
  }

  // Generate points
  for i := 0; i < n; i++ {
    vel := rl.Vector2{}
    if s.QuasiType != "center" {
      vel = randVector2(MaxInitialVelocity())
    }
    pt, _ := NewPointAtV(pos, vel)
    pt.Mass = size
    pt.r *= float32(math.Log10(float64(size * 10)))
    sco.integrate(pt)
  }

  // Add the cohort to the list of cohorts in Species
  s.Cohorts[sco.CoName] = &sco

  return &sco
}

// -------------------------------------------------------------------------------------------------------------------

func (sco *SpeciesCohort) integrate(pt *Point) *Point {

  pt.SpeciesCo = sco

  sco.Points = append(sco.Points, pt)
  return pt
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) InteractWith(other *Species, rules *Rules) {
  s.Rules[other.Name] = rules
}

// -------------------------------------------------------------------------------------------------------------------

func (sco *SpeciesCohort) getPoints(point *Point, rules *Rules) []*Point {
  //rulesDistSq := rules.Radius * rules.Radius
  //rulesSepDistSq := rules.SepRadius * rules.SepRadius

  // #4 fastest
  //var pts []*Point = nil
  //ch := s.getPointsChan()
  //for otherPt := range ch {
  //  pts = append(pts, otherPt)
  //}
  //return pts

  // #3 fastest, but carries the length with the slice
  pts := make([]*Point, 0, len(sco.Points))

  //qtree := speciesQuadTrees[sco.CoName]
  qtree := sco.QuadTree
  qtree.getPoints(point, rules, &pts)

  //for _, otherPt := range s.Points {
  // dist := rl.Vector2Subtract(point.pos, otherPt.pos)
  // pairDistSq := rl.Vector2LenSqr(dist)
  //
  // // Attraction
  // if !TheGlobalRules.SkipAttractionRule {
  //
  //   if pairDistSq <= rules.RadiusSq && pairDistSq != 0.0 {
  //     pts = append(pts, otherPt)
  //   }
  // }
  // // Separation
  // if !TheGlobalRules.SkipSeparationRule {
  //
  //   if pairDistSq <= rules.SepRadiusSq && rules.SepFactor != 1 {
  //     // We are too close
  //     pts = append(pts, otherPt)
  //   }
  // }
  //
  //}
  return pts

  // #2 fastest (very close to #3, tho)
  //pts := make([]*Point, len(s.Points))
  //for i, otherPt := range s.Points {
  // pts[i] = otherPt
  //}
  //return pts

  // #1 fastest
  //return s.Points
}

// -------------------------------------------------------------------------------------------------------------------

func (sco *SpeciesCohort) getPointsChan() chan *Point {
  ch := make(chan *Point, 500)

  go func() {
    for _, pt := range sco.Points {
      ch <- pt
    }
    close(ch)
  }()

  return ch
}

// -------------------------------------------------------------------------------------------------------------------

func UpdateAllSpecies(st *ComputeStats) {
  for _, species := range allSpecies {
    for _, sco := range species.Cohorts {
      quadTree := NewQuadTree(0, 0, CurrentScreenWidth, CurrentScreenHeight, sco.Species.Color)
      quadTree.addPoints(sco.Points, st)

      quadCount := quadTree.count()
      specCount := len(sco.Points)
      if quadCount != specCount {
        log.Panic("Wrong counts")
      }

      sco.QuadTree = quadTree
    }
  }

  for _, species := range allSpecies {
    for _, sco := range species.Cohorts {
      sco.Update(st)
    }
  }
}

// -------------------------------------------------------------------------------------------------------------------

func DrawAllSpecies(st *ComputeStats) {
  for _, species := range allSpecies {
    for _, sco := range species.Cohorts {
      sco.Draw(st)
    }
  }

  for _, species := range allSpecies {
    for _, sco := range species.Cohorts {
      sco.QuadTree.Draw()
    }
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (sco *SpeciesCohort) Update(st *ComputeStats) {
  stats := ComputeStatsData{}

  for _, point := range sco.Points {
    // Give the point a chance to do its own update
    if point.Update(st) {
      continue
    }

    if TheGlobalRules.QuadTreeCmp {
      for rulesColor, rules := range sco.Species.Rules {
        for otherColor, species := range allSpecies {

          // Only doing cohorts from species of the color named by the rule
          if otherColor != rulesColor {
            continue
          }

          for _, other := range species.Cohorts {
            grav := rules.Attraction * TheGlobalRules.GravPerAttr

            stats.Points += len(other.Points)

            // TODO: Make a Vector2
            fx, fy := float32(0.0), float32(0.0)

            // -------------- Loop over other group
            for _, otherPt := range other.getPoints(point, rules) {
              stats.PointsProc += 1
              if point == otherPt {
                continue
              }

              // TODO: Make a Vector2
              fxOther, fyOther := float32(0.0), float32(0.0)

              dist := rl.Vector2Subtract(point.pos, otherPt.pos)
              pairDistSq := rl.Vector2LenSqr(dist)

              // Attraction
              if !TheGlobalRules.SkipAttractionRule {

                stats.Cmps += 1
                if pairDistSq <= rules.RadiusSq && pairDistSq != 0.0 {
                  pairDist := float32(math.Sqrt(float64(pairDistSq)))
                  stats.Sqrts += 1
                  fxOther += otherPt.Mass * dist.X / pairDist
                  fyOther += otherPt.Mass * dist.Y / pairDist
                }
              }

              // Separation
              if !TheGlobalRules.SkipSeparationRule {

                stats.Cmps += 1
                if pairDistSq <= rules.SepRadiusSq && rules.SepFactor != 1 {
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
        }
      }

    } else {
      for rulesColor, rules := range sco.Species.Rules {
        for otherColor, species := range allSpecies {

          // Only doing cohorts from species of the color named by the rule
          if otherColor != rulesColor {
            continue
          }

          for _, other := range species.Cohorts {
            grav := rules.Attraction * TheGlobalRules.GravPerAttr

            // TODO: Make a Vector2
            fx, fy := float32(0.0), float32(0.0)

            // -------------- Loop over other group
            for _, otherPt := range other.Points {
              if point == otherPt {
                continue
              }

              // TODO: Make a Vector2
              fxOther, fyOther := float32(0.0), float32(0.0)

              dist := rl.Vector2Subtract(point.pos, otherPt.pos)
              pairDistSq := rl.Vector2LenSqr(dist)

              // Attraction
              if !TheGlobalRules.SkipAttractionRule {

                stats.Cmps += 1
                if pairDistSq <= rules.RadiusSq && pairDistSq != 0.0 {
                  pairDist := float32(math.Sqrt(float64(pairDistSq)))
                  stats.Sqrts += 1
                  fxOther += otherPt.Mass * dist.X / pairDist
                  fyOther += otherPt.Mass * dist.Y / pairDist
                }
              }

              // Separation
              if !TheGlobalRules.SkipSeparationRule {

                stats.Cmps += 1
                if pairDistSq <= rules.SepRadiusSq && rules.SepFactor != 1 {
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
        }
      }
    }


    // ---------- Finalize computations ----------

    // TODO -- this calls sqrt, update stats
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

  st.addStats(stats)
}

// -------------------------------------------------------------------------------------------------------------------

func (sco *SpeciesCohort) Draw(st *ComputeStats) {
  for _, point := range sco.Points {
    point.Draw()
  }
}

