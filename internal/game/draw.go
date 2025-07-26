package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/somewhat9/conway-life-go/internal/config"
)

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Gray{15})
	for y, col := range g.board {
		for x, value := range col {
			vector.DrawFilledRect(screen, float32(x*config.SquareSize)+1+config.GridOffsetX, float32(y*config.SquareSize)+1+config.GridOffsetY, config.SquareSize-2, config.SquareSize-2, config.Colors[value], false)
		}
	}
	
	g.PlayBtn.Draw(screen)
	g.PauseBtn.Draw(screen)
}
