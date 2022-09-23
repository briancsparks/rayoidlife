package raygun

type Rules struct {
  Attraction  float32
  Radius      float32
}

type GlobalRules struct {
  GravPerAttr   float32
}
var TheGlobalRules *GlobalRules = &GlobalRules{
  GravPerAttr: 1.0 / -100.0,
}

