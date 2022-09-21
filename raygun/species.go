package raygun

import "image/color"

type Species struct {
  name string
  Color color.RGBA
}

func NewSpecies(name string, color color.RGBA) (*Species, error) {
  s := Species{
    name:  name,
    Color: color,
  }

  return &s, nil
}

func (s *Species) MakePointGoing(dx, dy int32) (*Point, error) {
  pt, err := NewPointGoing(dx, dy)

  pt.Color = s.Color

  return pt, err
}

func (s *Species) MakePointAt(x, y, dx, dy int32) (*Point, error) {
  pt, err := NewPointAt(x, y, dx, dy)

  pt.Color = s.Color

  return pt, err
}

