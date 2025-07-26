package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Update() error {
	g.PlayBtn.Update()
	g.PauseBtn.Update()
	g.StatusUpdate()

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		g.GridUpdate(true)
	} else if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		g.GridUpdate(false)
	}

	if g.running {
		g.GridTick()
	}

	return nil
}
