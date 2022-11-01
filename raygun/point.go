package raygun

import (
  "fmt"
  rl "github.com/gen2brain/raylib-go/raylib"
)

// -------------------------------------------------------------------------------------------------------------------

type Point struct {
  pos, vel    rl.Vector2
  r           float32
  Mass        float32

  Count       int

  Cohort *SpeciesCohort
  CoName string
  CoId   int
  Id     int
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointAtV(pos, vel rl.Vector2, sco *SpeciesCohort, id int) (*Point, error) {
  pt := Point{
    pos:    pos,
    vel:    vel,
    r:      6,
    Mass:   1,
    CoId:   id,
    Cohort: sco,
    Id:     (sco.Id * 10000) + id,
  }

  if sco != nil {
    pt.CoName = sco.CoName
  }

  return &pt, nil
}

// -------------------------------------------------------------------------------------------------------------------

func NewPointAt(pos rl.Vector2, sco *SpeciesCohort, id int) (*Point, error) {

  pt := Point{
    pos:    pos,
    r:      1,
    Mass:   1,
    CoId:   id,
    Cohort: sco,
    Id:     (sco.Id * 10000) + id,
  }

  if sco != nil {
    pt.CoName = sco.CoName
  }

  return &pt, nil
}

// -------------------------------------------------------------------------------------------------------------------

func NewAggPointAt(mass float32, pos rl.Vector2, sco *SpeciesCohort, id, count, area int) (*Point, error) {

 pt := Point{
   pos:    pos,
   r:      1,
   Mass:   mass,
   CoId:   id,
   Cohort: sco,
   Id:     (sco.Id * 10000) + id,
 }

 if sco != nil {
   pt.CoName = sco.CoName
 }
 pt.CoName += fmt.Sprintf("-from%03d-a%04d", count, area)

 return &pt, nil
}

func (pt *Point) clamp(x, y float32) bool {
  isxClampped := true
  if pt.pos.X > x {
    pt.pos.X = x
  } else if pt.pos.X < 0 {
    pt.pos.X = 0
  } else {
    isxClampped = false
  }

  isyClampped := true
  if pt.pos.Y > y {
    pt.pos.Y = y
  } else if pt.pos.Y < 0 {
    pt.pos.Y = 0
  } else {
    isyClampped = false
  }

  return isxClampped || isyClampped
}

// -------------------------------------------------------------------------------------------------------------------

// Update can update the Point, and returns true if it does, false otherwise.
func (pt *Point) Update(st *ComputeStats) bool {
  return false
}

// -------------------------------------------------------------------------------------------------------------------

func (pt *Point) Draw() {
  if pt.Cohort.Species.QuasiType == "center" {
    rl.DrawCircleV(pt.pos, pt.r + 3, rl.Black)
  }
  rl.DrawCircleV(pt.pos, pt.r, pt.Cohort.Species.Color)
}
