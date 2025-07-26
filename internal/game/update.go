package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Update() error {
	g.PlayBtn.Update()
	g.PauseBtn.Update()
	g.SpeedBtn.Update()
	g.StatusUpdate()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.GridUpdate(true)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.GridUpdate(false)
	}

	if g.running {
		now := time.Now()
		if now.Sub(g.lastTick) >= time.Second/time.Duration(g.Speed) {
			g.GridTick()
			g.lastTick = now
		}
	}

	return nil
}
