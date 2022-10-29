package raygun

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  "image/color"
)


// -------------------------------------------------------------------------------------------------------------------

var capacity int = TheGlobalRules.QuadTreeCap

// -------------------------------------------------------------------------------------------------------------------

type HQuadTree struct {
  left, right  float32
  parent, aTree, bTree *VQuadTree
  isA bool

  aggregated  *Point

  Cohort  *SpeciesCohort
  Points []*Point
  Color   color.RGBA
}

// -------------------------------------------------------------------------------------------------------------------

type VQuadTree struct {
  top, bottom  float32
  parent, aTree, bTree *HQuadTree
  isA bool

  aggregated  *Point

  Cohort  *SpeciesCohort
  Points []*Point
  Color   color.RGBA
}

// -------------------------------------------------------------------------------------------------------------------

func NewQuadTree(l, t, r, b float32, cohort *SpeciesCohort, c color.RGBA) *HQuadTree {
  vtree := &VQuadTree{
    top:    t,
    bottom: b,
    isA:    true,
    Color:  c,
    Cohort: cohort,
  }
  //_=vtree

  htree := &HQuadTree{
    left:  l,
    right: r,
    parent: vtree,
    aTree: vtree,
    bTree: &VQuadTree{}, /* t,b are zero. has no area, will contain no points */
    isA:   true,
    Color: c,
    Cohort: cohort,
  }

  htree.aTree.parent = htree
  htree.bTree.parent = htree

  return htree
}

// ===================================================================================================================

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) addPoints(pts []*Point, st *ComputeStats) {
  for _, pt := range pts {
    t.addPoint(pt, st)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) addPoint(pt *Point, st *ComputeStats) bool {
  if !t.contains(pt) {
    return false
  }

  added := false

  // If we're already using subtrees, just pass along
  if t.aTree != nil {
    added = added || t.aTree.addPoint(pt, st)
    added = added || t.bTree.addPoint(pt, st)
    return added
  }

  preferAppend := true
  if TheGlobalRules.QuadTreeStrat == 1 {
    if t.area() > TheGlobalRules.QuadTreeArea {
      preferAppend = false
    }
  }

  if preferAppend {
    if len(t.Points) < capacity {
      t.Points = append(t.Points, pt)
      return true
    }
  }


  // Upon entering, there were no subtrees, but we are at capacity, so we need to create
  //   the subtrees here. Recursing is the simplest way to then add to myself.
  t.aTree, t.bTree = t.parent.split(t)
  return t.addPoint(pt, st)
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) contains(pt *Point) bool {
  // if we're 'A' (top), we know the pt is below our top - just check bottom
  if t.isA {
    return pt.pos.Y <= t.bottom
  }
  return pt.pos.Y > t.top
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) count() int {
  result := len(t.Points)

  if t.aTree != nil {
    result += t.aTree.count()
    result += t.bTree.count()
  }

  return result
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) split(h *HQuadTree) (*VQuadTree, *VQuadTree) {
  topHeight := (t.bottom - t.top) / 2

  top := &VQuadTree{
    top:    t.top,
    bottom: t.top + topHeight,
    parent: h,
    isA:    true,
    Color:  t.Color,
    Cohort: t.Cohort,
  }

  bottom := &VQuadTree{
    top:    t.top + topHeight,
    bottom: t.bottom,
    parent: h,
    Color:  t.Color,
    Cohort: t.Cohort,
  }

  return top, bottom
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) Draw() {
  left, right := t.parent.left, t.parent.right

  rl.DrawRectangleLines(int32(left), int32(t.top), int32(right - left), int32(t.bottom - t.top), t.Color)


  if t.aTree != nil {
    t.aTree.Draw()
    t.bTree.Draw()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) area() float32 {
  w := t.parent.right - t.parent.left
  h := t.bottom - t.top

  return w * h
}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) getPoints(point *Point, rules *Rules, pts *[]*Point) {

  all := false
  if t.area() <= TheGlobalRules.QuadTreeArea {
    all = true
  }

  // If we are 'A' (top), this whole tree is excluded if the point+r is less than our left
  if t.isA {
    if point.pos.Y + rules.Radius < t.top {
      return
    }

    all = all && (point.pos.Y + rules.Radius >= t.bottom)
  } else {
    if point.pos.Y - rules.Radius > t.bottom {
      return
    }

    all = all && (point.pos.Y - rules.Radius <= t.top)
  }

  if all {
    if t.aggregated == nil {
      t.aggregated, _ = NewPointAt(rl.Vector2{
        X: (t.parent.left + t.parent.right) / 2,
        Y: (t.bottom + t.top) / 2,
      },t.Cohort, 66664)
      //t.aggregated.Mass = maxFloat32(t.totalMass(), 1)
      t.aggregated.Mass = t.totalMass()
      t.aggregated.Count = t.count()
      //t.aggregated.Id = 66664
    }
    dist := rl.Vector2Distance(point.pos, t.aggregated.pos)
    all = all && (dist <= rules.Radius)
  }

  if all {
    *pts = append(*pts, t.aggregated)
    return
  }

  for _, p := range t.Points {
    *pts = append(*pts, p)
  }

  if t.aTree != nil {
    t.aTree.getPoints(point, rules, pts)
    t.bTree.getPoints(point, rules, pts)
  }

}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) getPoints2(point *Point, rule *SingleRule, rulesz *Rules, pts *[]*Point) {

  all := false
  if t.area() <= TheGlobalRules.QuadTreeArea {
    all = true
  }

  // If we are 'A' (top), this whole tree is excluded if the point+r is less than our left
  if t.isA {
    if point.pos.Y + rule.Radius < t.top {
      return
    }

    all = all && (point.pos.Y + rule.Radius >= t.bottom)
  } else {
    if point.pos.Y - rule.Radius > t.bottom {
      return
    }

    all = all && (point.pos.Y - rule.Radius <= t.top)
  }

  if all {
    if t.aggregated == nil {
      t.aggregated, _ = NewPointAt(rl.Vector2{
        X: (t.parent.left + t.parent.right) / 2,
        Y: (t.bottom + t.top) / 2,
      },t.Cohort, 66660)
      //t.aggregated.Mass = maxFloat32(t.totalMass(), 1)
      t.aggregated.Mass = t.totalMass()
      t.aggregated.Count = t.count()
      //t.aggregated.Id = 66660
    }
    dist := rl.Vector2Distance(point.pos, t.aggregated.pos)
    all = all && (dist <= rule.Radius)
  }

  if all {
    *pts = append(*pts, t.aggregated)
    return
  }

  for _, p := range t.Points {
    *pts = append(*pts, p)
  }

  if t.aTree != nil {
    t.aTree.getPoints2(point, rule, rulesz, pts)
    t.bTree.getPoints2(point, rule, rulesz, pts)
  }

}

// -------------------------------------------------------------------------------------------------------------------

func (t *VQuadTree) totalMass() float32 {
  var result float32 = 0
  for _, point := range t.Points {
    result += point.Mass
  }

  if t.aTree != nil {
    result += t.aTree.totalMass()
    result += t.bTree.totalMass()
  }

  return result
}


// ===================================================================================================================
// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) addPoints(pts []*Point, st *ComputeStats) {
  for _, pt := range pts {
    t.addPoint(pt, st)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) addPoint(pt *Point, st *ComputeStats) bool {
  if !t.contains(pt) {
    return false
  }

  added := false

  // If we're already using subtrees, just pass along
  if t.aTree != nil {
    added = added || t.aTree.addPoint(pt, st)
    added = added || t.bTree.addPoint(pt, st)
    return added
  }

  preferAppend := true
  if TheGlobalRules.QuadTreeStrat == 1 {
    if t.area() > TheGlobalRules.QuadTreeArea {
      preferAppend = false
    }
  }

  if preferAppend {
    if len(t.Points) < capacity {
      t.Points = append(t.Points, pt)
      return true
    }
  }

  // Upon entering, there were no subtrees, but we are at capacity, so we need to create
  //   the subtrees here. Recursing is the simplest way to then add to myself.
  t.aTree, t.bTree = t.parent.split(t)
  return t.addPoint(pt, st)
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) contains(pt *Point) bool {
  // if we're 'A' (left), we know the pt is less than our left - just check if the point is to our right
  if t.isA {
    return pt.pos.X <= t.right
  }
  return pt.pos.X > t.left
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) count() int {
  result := len(t.Points)

  if t.aTree != nil {
    result += t.aTree.count()
    result += t.bTree.count()
  }

  return result
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) split(v *VQuadTree) (*HQuadTree, *HQuadTree) {
  leftWidth := (t.right - t.left) / 2

  left := &HQuadTree{
    left:   t.left,
    right:  t.left + leftWidth,
    parent: v,
    isA:    true,
    Color:  t.Color,
    Cohort: t.Cohort,
  }

  right := &HQuadTree{
    left:   t.left + leftWidth,
    right:  t.right,
    parent: v,
    Color:  t.Color,
    Cohort: t.Cohort,
  }

  return left, right
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) Draw() {
  top, bottom := t.parent.top, t.parent.bottom

  rl.DrawRectangleLines(int32(t.left), int32(top), int32(t.right-t.left), int32(bottom-top), t.Color)

  if t.aTree != nil {
    t.aTree.Draw()
    t.bTree.Draw()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) area() float32 {
  w := t.right - t.left
  h := t.parent.bottom - t.parent.top

  return w * h
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) getPoints(point *Point, rules *Rules, pts *[]*Point) {

  all := false
  if t.area() <= TheGlobalRules.QuadTreeArea {
    all = true
  }

  // If we are 'A' (left), this whole tree is excluded if the point+r is less than our left
  if t.isA {
    if point.pos.X + rules.Radius < t.left {
      return
    }

    all = all && (point.pos.X + rules.Radius >= t.right)
  } else {
    if point.pos.X - rules.Radius > t.right {
      return
    }

    all = all && (point.pos.X - rules.Radius <= t.left)
  }

  if all {
    if t.aggregated == nil {
      t.aggregated, _ = NewPointAt(rl.Vector2{
        X: (t.left + t.right) / 2,
        Y: (t.parent.bottom + t.parent.top) / 2,
      },t.Cohort, 66665)
      //t.aggregated.Mass = maxFloat32(t.totalMass(), 1)
      t.aggregated.Mass = t.totalMass()
      t.aggregated.Count = t.count()
      //t.aggregated.Id = 66665
    }
    dist := rl.Vector2Distance(point.pos, t.aggregated.pos)
    all = all && (dist <= rules.Radius)
  }

  if all {
    *pts = append(*pts, t.aggregated)
    return
  }

  for _, p := range t.Points {
    *pts = append(*pts, p)
  }

  if t.aTree != nil {
    t.aTree.getPoints(point, rules, pts)
    t.bTree.getPoints(point, rules, pts)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) getPoints2(point *Point, rule *SingleRule, rulesz *Rules, pts *[]*Point) {

  all := false
  if t.area() <= TheGlobalRules.QuadTreeArea {
    all = true
  }

  // If we are 'A' (left), this whole tree is excluded if the point+r is less than our left
  if t.isA {
    if point.pos.X + rule.Radius < t.left {
      return
    }

    all = all && (point.pos.X + rule.Radius >= t.right)
  } else {
    if point.pos.X - rule.Radius > t.right {
      return
    }

    all = all && (point.pos.X - rule.Radius <= t.left)
  }

  if all {
    if t.aggregated == nil {
      t.aggregated, _ = NewPointAt(rl.Vector2{
        X: (t.left + t.right) / 2,
        Y: (t.parent.bottom + t.parent.top) / 2,
      },t.Cohort, 66661)
      //t.aggregated.Mass = maxFloat32(t.totalMass(), 1)
      t.aggregated.Mass = t.totalMass()
      t.aggregated.Count = t.count()
      //t.aggregated.Id = 66661
    }
    dist := rl.Vector2Distance(point.pos, t.aggregated.pos)
    all = all && (dist <= rule.Radius)
  }

  if all {
    *pts = append(*pts, t.aggregated)
    return
  }

  for _, p := range t.Points {
    *pts = append(*pts, p)
  }

  if t.aTree != nil {
    t.aTree.getPoints2(point, rule, rulesz, pts)
    t.bTree.getPoints2(point, rule, rulesz, pts)
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (t *HQuadTree) totalMass() float32 {
  var result float32 = 0
  for _, point := range t.Points {
    result += point.Mass
  }

  if t.aTree != nil {
    result += t.aTree.totalMass()
    result += t.bTree.totalMass()
  }

  return result
}



