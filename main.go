package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(screenWidth, screenHeight, "raylib [core] example - basic window")
	rl.SetTargetFPS(60)

	state := GameState{
		GameOver: false,
	}

	state.SpawnNewPiece()

	for !rl.WindowShouldClose() {

		if !state.SquareActive && !state.GameOver {
			state.SpawnNewPiece()
		}

		state.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Violet)

		state.Draw()

		rl.EndDrawing()
	}
}
