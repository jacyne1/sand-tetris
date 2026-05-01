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
}
