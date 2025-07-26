package config

import "image/color"

var Colors = map[bool]color.Color {
	false: color.Black,
	true: color.White,
}
