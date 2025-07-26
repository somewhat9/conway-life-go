package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type DynamicButton struct {
	States []*Button
	State uint8
	lastClick time.Time
}

func (dyB *DynamicButton) Update() error {
	now := time.Now()
	if now.Sub(dyB.lastClick) >= time.Second/6 {
		dyB.lastClick = now
		return dyB.States[dyB.State].Update()
	}
	return nil
}

func (dyB *DynamicButton) Draw(screen *ebiten.Image) {
	dyB.States[dyB.State].Draw(screen)
}
