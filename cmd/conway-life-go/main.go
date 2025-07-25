package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/somewhat9/conway-life-go/internal/config"
	"github.com/somewhat9/conway-life-go/internal/game"
)

func main() {
	g := &game.Game{}
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle(config.Title)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
