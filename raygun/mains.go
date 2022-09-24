package raygun

import rl "github.com/gen2brain/raylib-go/raylib"

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

const res4kW = 3840
const res4kH = 2160

var (
  InitialScreenWidth  int32 = 0.90 * res4kW
  InitialScreenHeight int32 = 0.90 * res4kH
  CurrentScreenWidth  int32
  CurrentScreenHeight int32
  CurrentScreenRadius int32
  CurrentScreenMidX   int32
  CurrentScreenMidY   int32
)

func MainTwo() {

  // --------------------------------- Initialize ---------------------------------
  var screenWidth, screenHeight int32 = InitialScreenWidth, InitialScreenHeight
  CurrentScreenWidth, CurrentScreenHeight = InitialScreenWidth, InitialScreenHeight
  CurrentScreenMidX, CurrentScreenMidY = CurrentScreenWidth/2, CurrentScreenHeight/2

  CurrentScreenRadius = maxInt(CurrentScreenWidth, CurrentScreenHeight)

  // ---------- Species ----------

  // Quasi-species
  center, _ := NewQuasiSpecies("center")
  center.MakeBigPointsAt(1, 100, float32(CurrentScreenMidX), float32(CurrentScreenMidY))

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
  reds.InteractWith(whites, NewRules(100.0, float32(CurrentScreenRadius) / 12))   // 288

  blues.InteractWith(reds, UnFriendly(175.0))
  blues.InteractWith(whites, Likes(400))

  whites.InteractWith(robots, Friendly(float32(CurrentScreenRadius) / 12))

  robots.InteractWith(robots, Ignore)
  robots.InteractWith(center, NewRules(1.0, float32(CurrentScreenRadius)))


  rl.InitWindow(screenWidth, screenHeight, "Two, what did you expect?")

  //camera := rl.Camera2D{}
  rl.SetTargetFPS(60)


  for !rl.WindowShouldClose() {

    // --------------------------------- Update -------------------------------------
    UpdateAllSpecies()

    // --------------------------------- Draw ---------------------------------------
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)
    //rl.BeginMode2D(camera)

    DrawAllSpecies()

    //rl.EndMode2D()
    rl.EndDrawing()
  }

  rl.CloseWindow()
}


