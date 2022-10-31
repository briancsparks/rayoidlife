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
  Species     *Species
  CoName      string
  Points      []*Point
  QuadTree    *HQuadTree
}

// -------------------------------------------------------------------------------------------------------------------

type Species struct {
  Name      string
  Color     color.RGBA
  Cohorts   map[string]*SpeciesCohort

  QuasiType string

  Rules     map[string]*Rules
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
    pt, _ := NewPointAtV(pos, vel, &sco, i)
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
    pt, _ := NewPointAtV(pos, vel, &sco, i)
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
    pt, _ := NewPointAtV(pos, vel, &sco, i)
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

  pt.Cohort = sco

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

func (sco *SpeciesCohort) getPoints2(pts []*Point, point *Point, rule *SingleRule, rules *Rules) []*Point {
  //pts := make([]*Point, 0, len(sco.Points))

  qtree := sco.QuadTree
  qtree.getPoints2(point, rule, rules, &pts)

  return pts
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
      quadTree := NewQuadTree(0, 0, CurrentScreenWidth, CurrentScreenHeight, sco, sco.Species.Color)
      quadTree.addPoints(sco.Points, st)

      quadCount := quadTree.count()
      specCount := len(sco.Points)
      if quadCount != specCount {
        log.Panic(fmt.Sprintf("Wrong counts: qt: %d, sp: %d", quadCount, specCount))
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

func (sco *SpeciesCohort) getCohorts(callback func(rules *Rules, species *Species, cohort *SpeciesCohort, color string)) {
  for rulesColor, rules := range sco.Species.Rules {
    for otherColor, species := range allSpecies {

      // Only doing cohorts from species of the color named by the rule
      if otherColor != rulesColor {
        continue
      }

      for _, otherCohort := range species.Cohorts {
        callback(rules, species, otherCohort, rulesColor)
      }
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

    origVel := point.vel
    _=origVel
    fearPoints := 0

    if TheGlobalRules.QuadTreeCmp {
      // First, evaluate the separation rules. If there is separation, do not do attraction
      if !TheGlobalRules.SkipNewSepRules {

        // Do all separation rules - if we are influenced by ANY other point, skip attraction rules
        sco.getCohorts(func(rules *Rules, species *Species, otherCohort *SpeciesCohort, color string) {
          stats.Points += len(otherCohort.Points)

          rule := rules.ThemRules["separation"]
          if rule.Radius <= 0 /*|| rule.Factor <= 0*/ {
            return
          }

          pts := make([]*Point, 0, len(otherCohort.Points))
          otherCohortSepPoints := otherCohort.getPoints2(pts, point, rule, rules)
          grav := rule.Factor * TheGlobalRules.GravPerAttr

          // TODO: Make a Vector2
          fx, fy := float32(0.0), float32(0.0)

          for iCohort, otherCohortPoint := range otherCohortSepPoints {
            _=iCohort

            if point == otherCohortPoint {
              wewe := 10
              _=wewe
              continue
            }
            stats.PointsProc += 1

            // TODO: put back
            //if otherCohortPoint.Mass <= 0 {
            //  continue
            //}
            stats.PointsProcHeavy += 1

            // TODO: Make a Vector2
            fxOther, fyOther := float32(0.0), float32(0.0)

            dist := rl.Vector2Subtract(point.pos, otherCohortPoint.pos)
            pairDistSq := rl.Vector2LenSqr(dist)

            // Separation
            stats.Cmps += 1
            if pairDistSq != 0 && pairDistSq <= rule.RadiusSq /*&& rule.Factor != 1*/ {
              // We are too close

              pairDist := float32(math.Sqrt(float64(pairDistSq)))
              stats.Sqrts += 1
              if pairDist != 0 {
                fearPoints += 1

                //// Kill all normal (attractive) momentum when we have to separate
                //if fearPoints == 1 {
                //  point.vel = rl.Vector2{}
                //}

                fxOther += (otherCohortPoint.Mass * dist.X / pairDist) * rule.Factor
                fyOther += (otherCohortPoint.Mass * dist.Y / pairDist) * rule.Factor
              }
            }

            fx += fxOther
            fy += fyOther
          }

          point.vel = rl.Vector2Add(point.vel, rl.Vector2{X: fx * grav, Y: fy * grav})
        })

        // Do attraction rules only if we are not separating
        if fearPoints <= 0 {
          sco.getCohorts(func(rules *Rules, species *Species, otherCohort *SpeciesCohort, color string) {
            stats.Points += len(otherCohort.Points)

            rule := rules.ThemRules["attraction"]
            if rule.Radius <= 0 /*|| rule.Factor <= 0*/ {
              return
            }

            pts := make([]*Point, 0, len(otherCohort.Points))
            otherCohortAttrPoints := otherCohort.getPoints2(pts, point, rule, rules)
            grav := rule.Factor * TheGlobalRules.GravPerAttr

            // TODO: Make a Vector2
            fx, fy := float32(0.0), float32(0.0)

            for iCohort, otherCohortPoint := range otherCohortAttrPoints {
              _=iCohort

              if point == otherCohortPoint {
                wewe := 10
                _=wewe
                continue
              }
              stats.PointsProc += 1

              // TODO: put back
              //if otherCohortPoint.Mass <= 0 {
              //  continue
              //}
              stats.PointsProcHeavy += 1

              // TODO: Make a Vector2
              fxOther, fyOther := float32(0.0), float32(0.0)

              dist := rl.Vector2Subtract(point.pos, otherCohortPoint.pos)
              pairDistSq := rl.Vector2LenSqr(dist)

              // Separation
              stats.Cmps += 1
              if pairDistSq != 0 && pairDistSq <= rule.RadiusSq /*&& rule.Factor != 1*/ {
                // We are too close

                pairDist := float32(math.Sqrt(float64(pairDistSq)))
                stats.Sqrts += 1
                if pairDist != 0 {

                  fxOther += (otherCohortPoint.Mass * dist.X / pairDist) * rule.Factor
                  fyOther += (otherCohortPoint.Mass * dist.Y / pairDist) * rule.Factor
                }
              }

              fx += fxOther
              fy += fyOther
            }

            point.vel = rl.Vector2Add(point.vel, rl.Vector2{X: fx * grav, Y: fy * grav})
          })
        }
      }

      if TheGlobalRules.SkipNewSepRules {
        if fearPoints <= 0 {
          for rulesColor, rules := range sco.Species.Rules {
            _=rulesColor
            grav := rules.Attraction * TheGlobalRules.GravPerAttr
            // TODO: put back
            //if rules.Attraction == 0 {
            //  continue
            //}

            for otherColor, species := range allSpecies {
              _=otherColor

              // Only doing cohorts from species of the color named by the rule
              if otherColor != rulesColor {
                continue
              }

              for _, otherCohort := range species.Cohorts {

                stats.Points += len(otherCohort.Points)

                // TODO: Make a Vector2
                fx, fy := float32(0.0), float32(0.0)

                // -------------- Loop over otherCohort group
                for _, otherCohortPoint := range otherCohort.getPoints(point, rules) {
                  stats.PointsProc += 1
                  if point == otherCohortPoint {
                    continue
                  }

                  // TODO: Make a Vector2
                  fxOther, fyOther := float32(0.0), float32(0.0)

                  dist := rl.Vector2Subtract(point.pos, otherCohortPoint.pos)
                  pairDistSq := rl.Vector2LenSqr(dist)

                  // Attraction
                  if !TheGlobalRules.SkipAttractionRule {

                    stats.Cmps += 1
                    if pairDistSq <= rules.RadiusSq && pairDistSq != 0.0 {
                      pairDist := float32(math.Sqrt(float64(pairDistSq)))
                      stats.Sqrts += 1
                      fxOther += otherCohortPoint.Mass * dist.X / pairDist
                      fyOther += otherCohortPoint.Mass * dist.Y / pairDist
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
      }

      iii := 10
      _=iii

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

    //// Restore original vel?
    //if fearPoints > 0 {
    //  point.vel = origVel
    //}

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

