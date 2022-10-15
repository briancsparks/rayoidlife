package raygun

type Rules struct {
  Attraction          float32
  Radius              float32

  SepFactor           float32
  SepRadius           float32
}

// -------------------------------------------------------------------------------------------------------------------

type GlobalRules struct {
  GravPerAttr         float32
  MaxVelocity         float32
  InitialVelocity     float32

  SelfAttractionDef   float32
  SelfRadiusDef       float32
  SelfSepFactorDef    float32
  SelfSepRadiusDef    float32

  SkipAttractionRule  bool
  SkipSeparationRule  bool

  QuadTreeCap         int

  QuadTreeStrat       int
  QuadTreeArea        float32
}
var TheGlobalRules *GlobalRules = &GlobalRules{
  GravPerAttr: 1.0 / -500.0,
  MaxVelocity: 10.0,
  InitialVelocity: -1,

  SelfAttractionDef: 40.0,
  SelfRadiusDef: 180.0,
  SelfSepFactorDef: -3,
  SelfSepRadiusDef:  30,

  SkipAttractionRule: false,
  SkipSeparationRule: false,

  QuadTreeCap: 100,

  //QuadTreeStrat: 0,
  //QuadTreeArea: -1.0,
  QuadTreeStrat: 1,
  QuadTreeArea: 10.0 * 10.0,
}

func MaxInitialVelocity() float32 {
  if TheGlobalRules.InitialVelocity < 0 {
    return TheGlobalRules.MaxVelocity
  }
  return TheGlobalRules.InitialVelocity
}

// -------------------------------------------------------------------------------------------------------------------

var Ignore *Rules = &Rules{Attraction: 0, Radius: 0}

// -------------------------------------------------------------------------------------------------------------------

func NewRules(a, r float32) *Rules {
  return &Rules{
    Attraction: a,
    Radius:     r,
    SepFactor:  1,
  }
}

// -------------------------------------------------------------------------------------------------------------------

func NewRulesWithSep(a, r, s, sr float32) *Rules {
  return &Rules{
    Attraction: a,
    Radius:     r,
    SepFactor:  s,
    SepRadius:  sr,
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

