package game

import "github.com/somewhat9/conway-life-go/internal/config"

func (g *Game) Layout(outsideWidth, oustideHeight int) (screenWidth, screenHeight int) {
	return config.ScreenWidth, config.ScreenHeight
}
