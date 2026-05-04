package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (g *GameState) Draw() {
	// Draw the Sand Grid
	for y := 0; y < GridHeight; y++ {
		for x := 0; x < GridWidth; x++ {
			if g.Grid[y][x] != Empty {
				rl.DrawRectangle(int32(x*CellSize), int32(y*CellSize), CellSize, CellSize, GetColor(g.Grid[y][x]))
			}
		}
	}

	// Draw the active Falling Square
	if g.SquareActive {
		//rl.DrawRectangle(int32(g.SquareX), int32(g.SquareY), 20, 20, rl.Orange)
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				if g.ActiveShape[y][x] == 1 {
					renderX := int32(g.SquareX) + int32(x*BlockSize)
					renderY := int32(g.SquareY) + int32(y*BlockSize)

					rl.DrawRectangle(renderX, renderY, BlockSize, BlockSize, GetColor(g.CurrentColor))
				}
			}
		}
	}

	if g.GameOver && !g.Win {
		rl.DrawRectangle(0, 0, screenWidth, screenHeight, rl.Fade(rl.Black, 0.8))

		title := "YOU LOOSE"
		fontSize := int32(40)

		// Calculate centered X position
		textWidth := rl.MeasureText(title, fontSize)
		textX := (int32(rl.GetScreenWidth()) / 2) - (textWidth / 2)
		textY := (int32(rl.GetScreenHeight()) / 2) - 20

		// Draw the main text
		rl.DrawText(title, textX, textY, fontSize, rl.Red)

		// Optional: Subtitle for restarting
		subTitle := "Press R to Restart"
		subFontSize := int32(20)
		subWidth := rl.MeasureText(subTitle, subFontSize)
		rl.DrawText(subTitle, (int32(rl.GetScreenWidth())/2)-(subWidth/2), textY+60, subFontSize, rl.White)
	}
}
