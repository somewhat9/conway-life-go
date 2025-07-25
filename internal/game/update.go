package game

func (g *Game) Update() error {
	g.PlayBtn.Update()
	g.PauseBtn.Update()
	return nil
}
