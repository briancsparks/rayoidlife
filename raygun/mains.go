package raygun

import (
  "fmt"
  rl "github.com/gen2brain/raylib-go/raylib"
)

// -------------------------------------------------------------------------------------------------------------------

func MainOne() {
  rl.InitWindow(800, 450, "raylib [core] example - basic window")
  rl.SetTargetFPS(60)

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()

    rl.ClearBackground(rl.RayWhite)
    rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

    rl.EndDrawing()
  }

  rl.CloseWindow()
}

// -------------------------------------------------------------------------------------------------------------------

const res4kW = float32(3840)
const res4kH = float32(2160)

var (
  InitialScreenWidth  float32 = 0.90 * res4kW
  InitialScreenHeight float32 = 0.90 * res4kH
  CurrentScreenWidth  float32
  CurrentScreenHeight float32
  CurrentScreenRadius float32
  CurrentScreenMidX   float32
  CurrentScreenMidY   float32

  CurrentScreenCenter rl.Vector2
)

func MainTwo() {

  // --------------------------------- Initialize ---------------------------------
  var screenWidth, screenHeight float32 = InitialScreenWidth, InitialScreenHeight
  CurrentScreenWidth, CurrentScreenHeight = InitialScreenWidth, InitialScreenHeight
  CurrentScreenMidX, CurrentScreenMidY = CurrentScreenWidth/2, CurrentScreenHeight/2
  CurrentScreenCenter = rl.Vector2{X: CurrentScreenMidX, Y: CurrentScreenMidY}

  CurrentScreenRadius = maxFloat32(CurrentScreenWidth, CurrentScreenHeight)

  // ---------- Species ----------

  // Quasi-species
  center, _ := NewQuasiSpecies("center")
  center.MakeBigPointsAt(1, 100, CurrentScreenCenter)

  // Colors Species
  reds, _     := NewSpecies("red", rl.Red)
  greens, _   := NewSpecies("green", rl.Green)
  blues, _    := NewSpecies("blue", rl.Blue)
  whites, _   := NewSpecies("white", rl.White)
  robots, _   := NewSpecies("robot", rl.Black)

  _ = robots

  // ---------- Populations ----------
  reds.MakePoints(100)
  greens.MakePoints(1)
  blues.MakePoints(100)
  whites.MakePoints(100)
  robots.MakeBigPoints(10, 10)

  //_,_,_ = adam,eve,robot

  // ---------- Interaction Rules ----------

  reds.InteractWith(blues, NewRules(-300.0, 200.0))
  reds.InteractWith(whites, NewRules(100.0, 288/*float32(CurrentScreenRadius) / 12*/))   // 288

  blues.InteractWith(reds, NewRules(-10.0, 175.0))
  blues.InteractWith(whites, Likes(400))

  whites.InteractWith(robots, Friendly(float32(CurrentScreenRadius) / 12))

  robots.InteractWith(robots, Ignore)
  robots.InteractWith(center, NewRules(1.0, CurrentScreenRadius))


  rl.InitWindow(int32(screenWidth), int32(screenHeight), "Two, what did you expect?")

  //camera := rl.Camera2D{}
  //rl.SetTargetFPS(60)


  stats := StartComputeStatsAgent()
  for !rl.WindowShouldClose() {
    stats.Reset()

    // --------------------------------- Update -------------------------------------
    UpdateAllSpecies(stats)

    // --------------------------------- Draw ---------------------------------------
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)
    //rl.BeginMode2D(camera)

    DrawAllSpecies(stats)

    st := stats.GetData()
    procPercent := float32(st.PointsProc) / float32(st.Points)
    rl.SetWindowTitle(fmt.Sprintf("FPS: %f, sqrts: %08d, cmps: %08d #Pts: %08d/%08d (%f)",
      rl.GetFPS(), st.Sqrts, st.Cmps, st.PointsProc, st.Points, procPercent))

    //rl.EndMode2D()
    rl.EndDrawing()
  }

  rl.CloseWindow()
}


