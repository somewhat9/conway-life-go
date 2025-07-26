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
		0, 0, config.ScreenWidth/3, config.GridOffsetY,
		"Play", false, false,
		color.RGBA{0, 255, 0, 255}, color.RGBA{0, 200, 0, 255}, color.White, 
		g.Font,
	)
	g.PauseBtn = game.NewButton(
		config.ScreenWidth/3, 0, config.ScreenWidth/3, config.GridOffsetY,
		"Pause", false, false,
		color.Gray{215}, color.Gray{115}, color.White, 
		g.Font,
	)

	g.SpeedBtn = &game.DynamicButton {
		States: []*game.Button{
			game.NewButton(
				config.ScreenWidth*2/3, 0, config.ScreenWidth/3, config.GridOffsetY,
				"Slow", false, false,
				config.SpeedColors[0], config.SpeedHovColors[0], color.White, 
				g.Font,
			),
			game.NewButton(
				config.ScreenWidth*2/3, 0, config.ScreenWidth/3, config.GridOffsetY,
				"Normal", false, false,
				config.SpeedColors[1], config.SpeedHovColors[1], color.White, 
				g.Font,
			),
			game.NewButton(
				config.ScreenWidth*2/3, 0, config.ScreenWidth/3, config.GridOffsetY,
				"Quick", false, false,
				config.SpeedColors[2], config.SpeedHovColors[2], color.White, 
				g.Font,
			),
			game.NewButton(
				config.ScreenWidth*2/3, 0, config.ScreenWidth/3, config.GridOffsetY,
				"Fast", false, false,
				config.SpeedColors[3], config.SpeedHovColors[3], color.White, 
				g.Font,
			),
			game.NewButton(
				config.ScreenWidth*2/3, 0, config.ScreenWidth/3, config.GridOffsetY,
				"Very Fast", false, false,
				config.SpeedColors[4], config.SpeedHovColors[4], color.White, 
				g.Font,
			),
		},
	}
	
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
