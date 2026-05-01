package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Empty = iota
	ColorRed
	ColorBlue
	ColorGreen
	ColorYellow
	ColorOrange
)

func GetColor(id int) rl.Color {
	switch id {
	case ColorRed:
		return rl.Red
	case ColorBlue:
		return rl.Blue
	case ColorGreen:
		return rl.Lime
	case ColorYellow:
		return rl.Yellow
	case ColorOrange:
		return rl.Orange
	default:
		return rl.White
	}
}
