package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/somewhat9/conway-life-go/internal/config"
)

type Game struct{
	Font text.Face
	PlayBtn *Button
	PauseBtn *Button
	SpeedBtn *DynamicButton
	board [config.NumSquares][config.NumSquares]bool
	running bool
	lastTick time.Time
	speed uint8
}

func (g *Game) StatusUpdate() {
	if g.PlayBtn.click {
		g.running = true
		g.PlayBtn.click = false
	}
	if g.PauseBtn.click {
		g.running = false
		g.PauseBtn.click = false
	}
	if g.SpeedBtn.States[g.SpeedBtn.State].click {
		g.SpeedBtn.States[g.SpeedBtn.State].click = false
		g.speed = config.Speeds[g.SpeedBtn.State]
		g.SpeedBtn.State++
		if g.SpeedBtn.State >= uint8(len(g.SpeedBtn.States)) {
			g.SpeedBtn.State = 0
		}
	}
}

func (g *Game) GridUpdate(val bool) {
	mx, my := ebiten.CursorPosition()
	for y, col := range g.board {
		for x := range col {
			x0 := float32(x*config.SquareSize)+1+config.GridOffsetX 
			y0 := float32(y*config.SquareSize)+1+config.GridOffsetY
			x1 := x0 + config.SquareSize-2
			y1 := y0 + config.SquareSize-2

			if float32(mx) >= x0 && float32(mx) <= x1 && float32(my) >= y0 && float32(my) <= y1 {
				g.board[y][x] = val
			}
		}
	}
}

func (g *Game) GridTick() {
	var new_board [config.NumSquares][config.NumSquares]bool = g.board
	for y, col := range g.board {
		for x, value := range col {
			n := g.CountNeighbors(x, y)
			if value { 
				new_board[y][x] = !(n < 2 || n > 3)
			} else {
				new_board[y][x] = (n == 3)
			}
		}
	}
	g.board = new_board
}

func (g *Game) CountNeighbors(x, y int) uint8 {
	var count uint8 = 0
	for dy := -1; dy <= 1; dy++ {
		ny := (y + dy + len(g.board)) % len(g.board)
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx := (x + dx + len(g.board[ny])) % len(g.board[ny])
			if g.board[ny][nx] { 
				count++
			}
		}
	} 
	return count
}
