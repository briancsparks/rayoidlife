package raygun

type SingleRule struct {
  Factor              float32
  Radius              float32
  RadiusSq            float32
}

// -------------------------------------------------------------------------------------------------------------------

type Rules struct {
  Attraction          float32
  Radius              float32
  RadiusSq            float32

  SepFactor           float32
  SepRadius           float32
  SepRadiusSq         float32

  ThemRules           map[string]*SingleRule
}

// -------------------------------------------------------------------------------------------------------------------

type GlobalRules struct {
  GravPerAttr         float32
  MaxVelocity         float32
  MaxForce            float32
  InitialVelocity     float32

  SelfAttractionDef   float32
  SelfRadiusDef       float32
  SelfSepFactorDef    float32
  SelfSepRadiusDef    float32

  SkipAttractionRule  bool
  SkipSeparationRule  bool
  SkipNewSepRules     bool

  QuadTreeCap         int

  QuadTreeStrat       int
  QuadTreeArea        float32

  QuadTreeCmp         bool
}
var TheGlobalRules *GlobalRules = &GlobalRules{
  GravPerAttr: 1.0 / -500,
  MaxVelocity: 10,
  MaxForce:    2.5,
  InitialVelocity: -1,

  SelfAttractionDef: 4.0,
  SelfRadiusDef: 180.0,
  //SelfSepFactorDef: -0.3,
  //SelfSepRadiusDef:  30,
  //SelfSepFactorDef: -30.0,           /* TODO: Seems too big */
  //SelfSepRadiusDef:  8,
  SelfSepFactorDef: -4.0,
  SelfSepRadiusDef:  8,

  SkipAttractionRule: false,
  SkipSeparationRule: false,
  SkipNewSepRules:    false,

  QuadTreeCap: 100,

  //QuadTreeStrat: 0,         /* 0==preferAppend, 1==smallArea(preferSplit)*/
  //QuadTreeArea: -1.0,
  QuadTreeStrat: 1,
  QuadTreeArea: 10.0 * 10.0,

  QuadTreeCmp:        true,
}

func MaxInitialVelocity() float32 {
  if TheGlobalRules.InitialVelocity < 0 {
    return TheGlobalRules.MaxVelocity
  }
  return TheGlobalRules.InitialVelocity
}

// -------------------------------------------------------------------------------------------------------------------

func NewRules(a, r float32) *Rules {
  rules := &Rules{
    Attraction:  a,
    Radius:      r,
    RadiusSq:    r * r,
    SepFactor:   1,
    ThemRules:   map[string]*SingleRule{},
  }

  rules.ThemRules["attraction"] = &SingleRule{
    Factor:   a,
    Radius:   r,
    RadiusSq: r * r,
  }

  rules.ThemRules["separation"] = &SingleRule{
    Factor:   1,
  }

  return rules
}

// -------------------------------------------------------------------------------------------------------------------

func NewRulesWithSep(a, r, s, sr float32) *Rules {
  rules := &Rules{
    Attraction: a,
    Radius:     r,
    RadiusSq:   r*r,
    SepFactor:  s,
    SepRadius:  sr,
    SepRadiusSq: sr*sr,
    ThemRules:   map[string]*SingleRule{},
  }

  rules.ThemRules["attraction"] = &SingleRule{
    Factor:   a,
    Radius:   r,
    RadiusSq: r * r,
  }

  rules.ThemRules["separation"] = &SingleRule{
    Factor:   s,
    Radius:   sr,
    RadiusSq: sr * sr,
  }

  return rules
}

// -------------------------------------------------------------------------------------------------------------------

var Ignore *Rules = NewRules(0, 0)

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

