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

  blues, _ := NewSpecies("blues", rl.Blue)
  pinks, _ := NewSpecies("pinks", rl.Pink)
  robots, _ := NewSpecies("robots", rl.Black)

  adam, _ := blues.MakePointGoing(1, 1)
  eve, _  := pinks.MakePointGoing(-1, -1)
  robot, _ := robots.MakePointGoing(1, -1)

  _,_,_ = adam,eve,robot

  robot.Color = rl.Black

  rl.InitWindow(screenWidth, screenHeight, "Two, what did you expect?")

  //camera := rl.Camera2D{}
  rl.SetTargetFPS(60)


  for !rl.WindowShouldClose() {

    // --------------------------------- Update -------------------------------------
    blues.Update()
    pinks.Update()
    robots.Update()


    // --------------------------------- Draw ---------------------------------------
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)

    //rl.BeginMode2D(camera)

    blues.Draw()
    pinks.Draw()
    robots.Draw()


    //rl.EndMode2D()

    rl.EndDrawing()
  }

  rl.CloseWindow()
}


