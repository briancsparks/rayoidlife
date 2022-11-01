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

  theCenterPt := center.Cohorts["center-00"].Points[0]
  _=theCenterPt

  // Colors SpeciesCohort
  redSpecies, _     := NewSpecies("red", rl.Red)
  greenSpecies, _   := NewSpecies("green", rl.Green)
  blueSpecies, _    := NewSpecies("blue", rl.Blue)
  whiteSpecies, _   := NewSpecies("white", rl.White)
  robotSpecies, _   := NewSpecies("robot", rl.Black)

  _ = robotSpecies

  //// ---------- Populations ----------
  //xdel := rl.Vector2{
  //  X: 250,
  //  Y: 0,
  //}
  //ydel := rl.Vector2{
  //  X: 0,
  //  Y: 250,
  //}
  //
  //curr := CurrentScreenCenter
  //_,_,_=xdel,ydel,curr
  //reds := redSpecies.MakePoints(1)
  //for i, pt := range reds.Points {
  //  _=i
  //  pt.pos = curr
  //  curr = rl.Vector2Add(curr, rl.Vector2Scale(xdel, 2))
  //  pt.clamp(CurrentScreenWidth, CurrentScreenHeight)
  //}
  //
  ////redSpecies.MakePoints(100)
  //greens := greenSpecies.MakePoints(1)
  //
  //blues := blueSpecies.MakePoints(20)
  //curr = CurrentScreenCenter
  //for i, pt := range blues.Points {
  //  _=i
  //  pt.pos = curr
  //  //rl.Vector2Scale(xdel, 2)
  //  curr = rl.Vector2Add(curr, rl.Vector2Scale(xdel, .25))
  //  //curr = rl.Vector2Add(curr, xdel)
  //  pt.clamp(CurrentScreenWidth, CurrentScreenHeight)
  //}
  //
  //whites := whiteSpecies.MakePoints(4)
  //robots := robotSpecies.MakeBigPoints(10, 10)
  //_,_,_,_,_ = reds, greens, blues, whites, robots

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

  redSpecies.InteractWith(blueSpecies, NewRules(-17.5, 200))
  redSpecies.InteractWith(whiteSpecies, NewRules(10.0, 500 /*float32(CurrentScreenRadius) / 12*/)) // 288
  //redSpecies.InteractWith2(redSpecies,
  //  TheGlobalRules.SelfAttractionDef,
  //  TheGlobalRules.SelfRadiusDef,
  //  0 /*TheGlobalRules.SelfSepFactorDef*/,
  //  0 /*TheGlobalRules.SelfSepRadiusDef*/,
  //)

  blueSpecies.InteractWith(redSpecies, NewRules(-1.0, 175))
  blueSpecies.InteractWith(whiteSpecies, NewRules(10.0, 400))
  blueSpecies.InteractWith2(blueSpecies,
    TheGlobalRules.SelfAttractionDef,
    TheGlobalRules.SelfRadiusDef,
    TheGlobalRules.SelfSepFactorDef / 2,
    TheGlobalRules.SelfSepRadiusDef,
  )

  whiteSpecies.InteractWith(robotSpecies, NewRules(10.0, float32(CurrentScreenRadius) / 4))
  whiteSpecies.InteractWith2(whiteSpecies,
    TheGlobalRules.SelfAttractionDef,
    TheGlobalRules.SelfRadiusDef,
    TheGlobalRules.SelfSepFactorDef,
    TheGlobalRules.SelfSepRadiusDef,
    //0,
    //0,
  )

  //robotSpecies.InteractWith(robotSpecies, Ignore)
  robotSpecies.InteractWith2(robotSpecies,
    TheGlobalRules.SelfAttractionDef,
    TheGlobalRules.SelfRadiusDef,
    0 /*TheGlobalRules.SelfSepFactorDef*/,
    0 /*TheGlobalRules.SelfSepRadiusDef*/,
  )
  robotSpecies.InteractWith(center, NewRules(10.0, CurrentScreenRadius))


  rl.InitWindow(int32(screenWidth), int32(screenHeight), "Two, what did you expect?")

  //camera := rl.Camera2D{}
  rl.SetTargetFPS(60)


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
    heavyProcPercent := float32(st.PointsProcHeavy) / float32(st.Points)
    rl.SetWindowTitle(fmt.Sprintf("FPS: %f, sqrts: %08d, cmps: %08d #Pts: %08d/%08d (%f) #HPts: %08d/%08d (%f)",
      rl.GetFPS(), st.Sqrts, st.Cmps,
      st.PointsProc, st.Points, procPercent,
      st.PointsProcHeavy, st.Points, heavyProcPercent))

    //rl.EndMode2D()
    rl.EndDrawing()
  }

  rl.CloseWindow()
}


