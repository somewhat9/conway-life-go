package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/somewhat9/conway-life-go/internal/assets"
	"github.com/somewhat9/conway-life-go/internal/config"
	"github.com/somewhat9/conway-life-go/internal/game"
)

func main() {
	g := &game.Game{}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle(config.Title)

	g.Font = assets.LoadFont(config.FontSize)
	g.PlayBtn = game.NewButton(
		0, 0, config.ScreenWidth/2, config.GridOffsetY,
		"Play", false, false,
		color.RGBA{0, 255, 0, 255}, color.RGBA{0, 200, 0, 255}, color.White, 
		g.Font,
	)
	g.PauseBtn = game.NewButton(
		config.ScreenWidth/2, 0, config.ScreenWidth/2, config.GridOffsetY,
		"Pause", false, false,
		color.Gray{215}, color.Gray{115}, color.White, 
		g.Font,
	)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
