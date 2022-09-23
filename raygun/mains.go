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

  reds, _     := NewSpecies("red", rl.Red)
  //greens, _ := NewSpecies("green", rl.Green)
  blues, _    := NewSpecies("blue", rl.Blue)
  //whites, _ := NewSpecies("white", rl.White)
  robots, _   := NewSpecies("robot", rl.Black)

  //reds.InteractWith(reds, &Rules{Attraction: 40.0, Radius: 180.0})
  //reds.InteractWith(blues, &Rules{Attraction: -5.0, Radius: 180.0})

  reds.InteractWith(reds, &Rules{Attraction: 40.0, Radius: float32(screenWidth)})
  reds.InteractWith(blues, &Rules{Attraction: -5.0, Radius: float32(screenWidth)})

  //blues.InteractWith(reds, &Rules{Attraction: -5.0, Radius: 180.0})
  //blues.InteractWith(blues, &Rules{Attraction: 10.0, Radius: 180.0})

  blues.InteractWith(reds, &Rules{Attraction: -5.0, Radius: float32(screenWidth)})
  blues.InteractWith(blues, &Rules{Attraction: 10.0, Radius: float32(screenWidth)})

  adam, _ := blues.MakePointAt(numUpToN(CurrentScreenWidth), numUpToN(CurrentScreenHeight), 0, 0)
  eve, _  := reds.MakePointAt(numUpToN(CurrentScreenWidth), numUpToN(CurrentScreenHeight), 0, 0)
  robot, _ := robots.MakePointAt(numUpToN(CurrentScreenWidth), numUpToN(CurrentScreenHeight), 0, 0)

  _,_,_ = adam,eve,robot

  rl.InitWindow(screenWidth, screenHeight, "Two, what did you expect?")

  //camera := rl.Camera2D{}
  rl.SetTargetFPS(60)


  for !rl.WindowShouldClose() {

    // --------------------------------- Update -------------------------------------
    blues.Update()
    reds.Update()
    robots.Update()


    // --------------------------------- Draw ---------------------------------------
    rl.BeginDrawing()
    rl.ClearBackground(rl.SkyBlue)

    //rl.BeginMode2D(camera)

    blues.Draw()
    reds.Draw()
    robots.Draw()


    //rl.EndMode2D()

    rl.EndDrawing()
  }

  rl.CloseWindow()
}


