package game

import (
	"github.com/somewhat9/conway-life-go/internal/config"
	"golang.org/x/image/font"
)

type Game struct{
	Font font.Face
	PlayBtn *Button
	PauseBtn *Button
	board [config.ScreenHeight][config.ScreenWidth]uint8
}
