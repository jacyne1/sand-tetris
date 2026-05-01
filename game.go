package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *GameState) Update() {
	if g.SquareActive {

		g.HandleInput()

		g.SquareY += 2 // Move down

		// Check for collision with bottom or sand
		//gridX := int(g.SquareX / CellSize)
		//gridY := int((g.SquareY + 20) / CellSize) // +20 is height of square

		if g.CheckCollision() {
			// snaps to cell so for better sand formation
			//g.SquareY = float32(int(g.SquareY/CellSize) * CellSize)
			//if gridY >= GridHeight || g.Grid[gridY][gridX] != 0 {
			//g.SquareActive = false
			//g.ConvertSquareToSand()

			g.SquareX = float32(int(g.SquareX/float32(BlockSize)) * BlockSize)
			g.SquareY = float32(int(g.SquareY/float32(BlockSize)) * BlockSize)

			g.SquareActive = false
			g.ConvertSquareToSand()
		}
	}
	g.UpdateSandPhysics()
}

func (g *GameState) CheckCollision() bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			// Only check blocks that are actually part of the shape
			if g.ActiveShape[y][x] != 0 {
				// Calculate the grid position of the BOTTOM of this specific cell
				// We use (y + 1) to look at the space immediately below this block
				nextGridY := int((g.SquareY + float32((y+1)*BlockSize)) / CellSize)

				// Calculate the grid position for the X edges
				gridXLeft := int((g.SquareX + float32(x*BlockSize)) / CellSize)
				gridXRight := int((g.SquareX + float32(x*BlockSize) + float32(BlockSize-1)) / CellSize)

				// 1. Check Floor
				if nextGridY >= GridHeight {
					return true
				}

				// 2. Check Sand below (Checking both left and right edges of the block)
				if nextGridY >= 0 && nextGridY < GridHeight {
					//if g.Grid[nextGridY][gridXLeft] != 0 || g.Grid[nextGridY][gridXRight] != 0 {
					//	return true
					//}
					// Safety check for Left Edge
					if gridXLeft >= 0 && gridXLeft < GridWidth {
						if g.Grid[nextGridY][gridXLeft] != 0 {
							return true
						}
					}
					// Safety check for Right Edge
					if gridXRight >= 0 && gridXRight < GridWidth {
						if g.Grid[nextGridY][gridXRight] != 0 {
							return true
						}
					}

				}
			}
		}
	}
	return false
}

func (g *GameState) ConvertSquareToSand() {
	baseX := int(g.SquareX / CellSize)
	baseY := int(g.SquareY / CellSize)

	// Turn a block of cells into sand
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if g.ActiveShape[y][x] != 0 {

				for sy := 0; sy < CellsInPiece; sy++ {
					for sx := 0; sx < CellsInPiece; sx++ {
						targetY := baseY + (y * CellsInPiece) + sy
						targetX := baseX + (x * CellsInPiece) + sx

						// Strict boundary check
						if targetY >= 0 && targetY < GridHeight &&
							targetX >= 0 && targetX < GridWidth {
							g.Grid[targetY][targetX] = g.CurrentColor // Use the INT ID
						}
					}
				}
			}
		}
	}
}

func (g *GameState) UpdateSandPhysics() {
	for y := GridHeight - 2; y >= 0; y-- {
		for x := 0; x < GridWidth; x++ {
			// 1. Get the value of the current cell
			currentID := g.Grid[y][x]

			// 2. If it's NOT empty (any color), try to move it
			if currentID != Empty {
				if g.Grid[y+1][x] == Empty { // Directly below
					g.Grid[y+1][x] = currentID // Move the ACTUAL color ID
					g.Grid[y][x] = Empty
				} else if x > 0 && g.Grid[y+1][x-1] == Empty { // Down-left
					g.Grid[y+1][x-1] = currentID
					g.Grid[y][x] = Empty
				} else if x < GridWidth-1 && g.Grid[y+1][x+1] == Empty { // Down-right
					g.Grid[y+1][x+1] = currentID
					g.Grid[y][x] = Empty
				}
			}
		}
	}
}

func (g *GameState) SpawnNewPiece() {

	randomIndex := rl.GetRandomValue(0, int32(len(ShapeLibrary)-1))
	template := ShapeLibrary[randomIndex]

	g.ActiveShape = template.Matrix
	g.CurrentColor = template.ColorID

	g.SquareX = float32(screenWidth / 2) // Middle of screen
	g.SquareY = 0                        // Top of screen
	g.SquareActive = true
}

func (g *GameState) CanMove(targetX, targetY float32) bool {
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if g.ActiveShape[y][x] != 0 {
				// Calculate where THIS specific block would be
				blockX := int((targetX + float32(x*BlockSize)) / CellSize)
				blockY := int((targetY + float32(y*BlockSize)) / CellSize)

				// Check Left, Right, and Floor boundaries
				if blockX < 0 || blockX >= GridWidth || blockY >= GridHeight {
					return false
				}

				// Check if there is already sand there
				if blockY >= 0 && g.Grid[blockY][blockX] != 0 {
					return false
				}
			}
		}
	}
	return true
}

func (g *GameState) HandleInput() {
	if !g.SquareActive {
		return
	}

	//moveSpeed := float32(BlockSize)

	if rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyDown(rl.KeyLeft) {
		//if g.SquareX-moveSpeed >= 0 {
		//	g.SquareX -= moveSpeed
		//}
		if g.CanMove(g.SquareX-float32(BlockSize), g.SquareY) {
			g.SquareX -= float32(BlockSize)
		}
	}

	if rl.IsKeyPressed(rl.KeyRight) || rl.IsKeyDown(rl.KeyRight) {
		//if g.SquareX+BlockSize+moveSpeed <= screenWidth {
		//	g.SquareX += moveSpeed
		//}
		if g.CanMove(g.SquareX+float32(BlockSize), g.SquareY) {
			g.SquareX += float32(BlockSize)
		}
	}

	if rl.IsKeyDown(rl.KeyDown) {
		g.SquareY += 5
	}

	if rl.IsKeyPressed(rl.KeyUp) {
		rotated := RotateMatrix(g.ActiveShape)
		oldShape := g.ActiveShape
		oldX := g.SquareX

		g.ActiveShape = rotated

		if g.CanMove(g.SquareX, g.SquareY) {
			return
		}

		kicks := []float32{-float32(BlockSize), float32(BlockSize), -float32(BlockSize * 2), float32(BlockSize * 2)}

		success := false
		for _, offset := range kicks {
			if g.CanMove(g.SquareX+offset, g.SquareY) {
				g.SquareX += offset
				success = true
				break
			}
		}

		// 3. If nothing worked, revert both shape and position
		if !success {
			g.ActiveShape = oldShape
			g.SquareX = oldX
		}
	}
}
