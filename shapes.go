package main

type ShapeTemplate struct {
	Matrix  [4][4]int
	ColorID int
}

var ShapeLibrary = []ShapeTemplate{
	{
		Matrix: [4][4]int{
			{0, 0, 0, 0},
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
		},
		ColorID: ColorBlue,
	},
	{
		Matrix: [4][4]int{
			{1, 1, 1, 0},
			{0, 1, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		ColorID: ColorRed,
	},
	{
		Matrix: [4][4]int{
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 0, 0, 0},
		},
		ColorID: ColorGreen,
	}, {
		Matrix: [4][4]int{
			{1, 0, 0, 0},
			{1, 0, 0, 0},
			{1, 1, 0, 0},
			{0, 0, 0, 0},
		},
		ColorID: ColorOrange,
	}, {
		Matrix: [4][4]int{
			{1, 1, 0, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		ColorID: ColorYellow,
	},
}

func RotateMatrix(matrix [4][4]int) [4][4]int {
	var result [4][4]int
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			// This math performs a 90-degree clockwise rotation
			result[x][3-y] = matrix[y][x]
		}
	}
	return result
}
