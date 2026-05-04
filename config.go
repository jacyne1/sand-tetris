package main

const (
	screenWidth  = 800
	screenHeight = 600

	CellSize         = 10
	CellsInPiece     = 4
	BlockSize        = CellSize * CellsInPiece
	GridWidth        = screenWidth / CellSize
	GridHeight       = screenHeight / CellSize
	DangerZoneHeight = 2
	//topMargin    = 150 // Room for Score & Turn Info
	//bottomMargin = 50  // Room for Settings/Buttons
	//padding      = 20
)

type GameState struct {
	// 0 = Empty, 1 = Sand (you can use your iota here)
	Grid [GridHeight][GridWidth]int

	ActiveShape [4][4]int

	// The active falling square
	SquareX      float32
	SquareY      float32
	SquareActive bool
	CurrentColor int
	GameOver     bool
	Win          bool
}
