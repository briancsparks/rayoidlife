package raylib

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

func MainTwo() {

  // --------------------------------- Initialize ---------------------------------
  var screenWidth, screenHeight int32 = 1200, 800
  midPtX, midPtY := screenWidth/2, screenHeight/2
  adamX, adamY, adamDx, adamDy := midPtX, midPtY, int32(1), int32(1)
  eveX,  eveY,  eveDx,  eveDy  := midPtX, midPtY, int32(-1), int32(-1)

  rl.InitWindow(screenWidth, screenHeight, "Two, what did you expect?")

  //camera := rl.Camera2D{}
  rl.SetTargetFPS(60)

  
  for !rl.WindowShouldClose() {

    // --------------------------------- Update -------------------------------------
    adamX += adamDx
    adamY += adamDy

    eveX  += eveDx
    eveY  += eveDy


    // --------------------------------- Draw ---------------------------------------
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)

    //rl.BeginMode2D(camera)

    // Adam and Eve
    rl.DrawCircle(adamX, adamY, 10, rl.Blue)
    rl.DrawCircle(eveX,  eveY,  10, rl.Pink)


    //rl.EndMode2D()

    rl.EndDrawing()
  }

  rl.CloseWindow()
}


