package raygun

import (
  rl "github.com/gen2brain/raylib-go/raylib"
  "math"
  "math/rand"
  "time"
)

// -------------------------------------------------------------------------------------------------------------------

var rnd *rand.Rand

func init() {
  rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// -------------------------------------------------------------------------------------------------------------------

func randBetween(min, max float32) float32 {
  return (rnd.Float32() * (max - min)) + min
}

// -------------------------------------------------------------------------------------------------------------------

func randUpTo(max float32) float32 {
  return randBetween(0.0, max)
}

// -------------------------------------------------------------------------------------------------------------------

func randBetweenN(min, max int32) float32 {
  return randBetween(float32(min), float32(max))
}

// -------------------------------------------------------------------------------------------------------------------

func randUpToN(max int32) float32 {
  return randBetween(0.0, float32(max))
}

// -------------------------------------------------------------------------------------------------------------------

func maxInt(a, b int32) int32 {
  if a > b {
    return a
  }
  return b
}

// -------------------------------------------------------------------------------------------------------------------

func maxFloat32(a, b float32) float32 {
  if a > b {
    return a
  }
  return b
}

// -------------------------------------------------------------------------------------------------------------------

func clamped(x *float32, a, b float32) bool {
  if *x < a {
    *x = a
    return true
  }

  if *x > b {
    *x = b
    return true
  }

  return false
}

// -------------------------------------------------------------------------------------------------------------------

func clampxy(x, y *float32, maxLen float32) {
  maxLenSq := maxLen * maxLen
  vSq := *x * *x + *y * *y
  if vSq > maxLenSq {
    vLen := float32(math.Sqrt(float64(vSq)))

    *x /= vLen
    *y /= vLen
  }
}

// -------------------------------------------------------------------------------------------------------------------

func clampV2(v *rl.Vector2, maxLen float32) {
  clampxy(&v.X, &v.Y, maxLen)
}

