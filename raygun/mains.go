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

  // ---------- SpeciesCohort ----------

  // Quasi-species
  center, _ := NewQuasiSpecies("center")
  center.MakeBigPointsAt(1, 100, CurrentScreenCenter)

  // Colors SpeciesCohort
  redSpecies, _     := NewSpecies("red", rl.Red)
  greenSpecies, _   := NewSpecies("green", rl.Green)
  blueSpecies, _    := NewSpecies("blue", rl.Blue)
  whiteSpecies, _   := NewSpecies("white", rl.White)
  robotSpecies, _   := NewSpecies("robot", rl.Black)

  _ = robotSpecies

  // ---------- Populations ----------
  reds := redSpecies.MakePoints(100)
  //redSpecies.MakePoints(100)
  greens := greenSpecies.MakePoints(1)
  blues := blueSpecies.MakePoints(100)
  whites := whiteSpecies.MakePoints(100)
  robots := robotSpecies.MakeBigPoints(10, 10)
  _,_,_,_,_ = reds, greens, blues, whites, robots

  //_,_,_ = adam,eve,robot

  // ---------- Interaction Rules ----------

  redSpecies.InteractWith(blueSpecies, NewRules(-300.0, 200.0))
  redSpecies.InteractWith(whiteSpecies, NewRules(100.0, 288 /*float32(CurrentScreenRadius) / 12*/)) // 288

  blueSpecies.InteractWith(redSpecies, NewRules(-10.0, 175.0))
  blueSpecies.InteractWith(whiteSpecies, Likes(400))

  whiteSpecies.InteractWith(robotSpecies, Friendly(float32(CurrentScreenRadius) / 12))

  robotSpecies.InteractWith(robotSpecies, Ignore)
  robotSpecies.InteractWith(center, NewRules(1.0, CurrentScreenRadius))


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


