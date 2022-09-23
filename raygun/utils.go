package raygun

import (
  "math/rand"
  "time"
)

// -------------------------------------------------------------------------------------------------------------------

var rnd *rand.Rand

func init() {
  rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// -------------------------------------------------------------------------------------------------------------------

func numBetween(min, max float32) float32 {
  return (rnd.Float32() * (max - min)) + min
}

// -------------------------------------------------------------------------------------------------------------------

func numUpTo(max float32) float32 {
  return numBetween(0.0, max)
}

// -------------------------------------------------------------------------------------------------------------------

func numBetweenN(min, max int32) float32 {
  return numBetween(float32(min), float32(max))
}

// -------------------------------------------------------------------------------------------------------------------

func numUpToN(max int32) float32 {
  return numBetween(0.0, float32(max))
}

