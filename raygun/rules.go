package raygun

type Rules struct {
  Attraction  float32
  Radius      float32
}

type GlobalRules struct {
  GravPerAttr         float32
  MaxVelocity         float32
  SelfAttractionDef   float32
  SelfRadiusDef       float32
}
var TheGlobalRules *GlobalRules = &GlobalRules{
  GravPerAttr: 1.0 / -500.0,
  MaxVelocity: 10.0,
  SelfAttractionDef: 40.0,
  SelfRadiusDef: 180.0,
}

var Ignore *Rules = &Rules{Attraction: 0, Radius: 0}
//var Friendly *Rules = &Rules{Attraction: 10.0, Radius: 100.0}
//var Unfriendly *Rules = &Rules{Attraction: -10.0, Radius: 100.0}
//var Afraid *Rules = &Rules{Attraction: -10.0, Radius: 100.0}

func NewRules(a, r float32) *Rules {
  return &Rules{
    Attraction: a,
    Radius:     r,
  }
}

// -------------------------------------------------------------------------------------------------------------------

func Friendly(r float32) *Rules {
  return NewRules(10.0, r)
}

// -------------------------------------------------------------------------------------------------------------------

func Likes(r float32) *Rules {
  return NewRules(40.0, r)
}

// -------------------------------------------------------------------------------------------------------------------

func UnFriendly(r float32) *Rules {
  return NewRules(-10.0, r)
}

// -------------------------------------------------------------------------------------------------------------------

func Afraid(r float32) *Rules {
  return NewRules(-50.0, r)
}

// -------------------------------------------------------------------------------------------------------------------

func Obsessed(r float32) *Rules {
  return NewRules(100.0, r)
}

