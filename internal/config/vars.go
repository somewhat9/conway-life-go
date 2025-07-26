package config

import (
	"image/color"
)

var Cells = map[bool]color.Color {
	false: color.Black,
	true: color.White,
}

var Speeds = [...]uint8{1, 2, 5, 10, 30} 

var SpeedColors = [...]color.RGBA{
	{55, 40, 250, 255},
	{40, 225, 250, 255},
	{40, 250, 90, 255},
	{240, 250, 40, 255},
	{250, 40, 40, 255},
}

var SpeedHovColors = [...]color.RGBA{
	{35, 20, 230, 255},
	{20, 205, 230, 255},
	{20, 230, 70, 255},
	{220, 230, 20, 255},
	{230, 20, 20, 255},
}
