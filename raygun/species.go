package raygun

import (
  "image/color"
)

// -------------------------------------------------------------------------------------------------------------------

type Species struct {
  name string
  Points []*Point
  Color color.RGBA
}

// -------------------------------------------------------------------------------------------------------------------

func NewSpecies(name string, color color.RGBA) (*Species, error) {
  s := Species{
    name:  name,
    Color: color,
  }

  return &s, nil
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointGoing(dx, dy int32) (*Point, error) {
  pt, err := NewPointGoing(dx, dy)

  pt.Color = s.Color

  s.Points = append(s.Points, pt)
  return pt, err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) MakePointAt(x, y, dx, dy int32) (*Point, error) {
  pt, err := NewPointAt(x, y, dx, dy)

  pt.Color = s.Color

  s.Points = append(s.Points, pt)
  return pt, err
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Update() {
  for _, point := range s.Points {
    point.Update()
  }
}

// -------------------------------------------------------------------------------------------------------------------

func (s *Species) Draw() {
  for _, point := range s.Points {
    point.Draw()
  }
}


