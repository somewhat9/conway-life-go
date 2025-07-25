package game

import "github.com/somewhat9/conway-life-go/internal/config"

type Game struct{
	board [config.ScreenHeight][config.ScreenWidth]uint8
}
