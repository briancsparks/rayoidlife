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

var (
  InitialScreenWidth  int32 = 1200
  InitialScreenHeight int32 = 800
  CurrentScreenWidth  int32
  CurrentScreenHeight int32
  CurrentScreenMidX   int32
  CurrentScreenMidY   int32
)

func MainTwo() {

  // --------------------------------- Initialize ---------------------------------
  var screenWidth, screenHeight int32 = InitialScreenWidth, InitialScreenHeight
  CurrentScreenWidth, CurrentScreenHeight = InitialScreenWidth, InitialScreenHeight
  CurrentScreenMidX, CurrentScreenMidY = CurrentScreenWidth/2, CurrentScreenHeight/2

  adam, _ := NewPointGoing(1, 1)
  eve, _  := NewPointGoing(-1, -1)

  adam.Color = rl.Blue
  eve.Color  = rl.Pink

  rl.InitWindow(screenWidth, screenHeight, "Two, what did you expect?")

  //camera := rl.Camera2D{}
  rl.SetTargetFPS(60)


  for !rl.WindowShouldClose() {

    // --------------------------------- Update -------------------------------------
    adam.Update()
    eve.Update()


    // --------------------------------- Draw ---------------------------------------
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)

    //rl.BeginMode2D(camera)

    // Adam and Eve
    adam.Draw()
    eve.Draw()


    //rl.EndMode2D()

    rl.EndDrawing()
  }

  rl.CloseWindow()
}


