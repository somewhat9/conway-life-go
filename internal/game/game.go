package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/somewhat9/conway-life-go/internal/config"
	"golang.org/x/image/font"
)

type Game struct{
	Font font.Face
	PlayBtn *Button
	PauseBtn *Button
	board [config.NumSquares][config.NumSquares]uint8
	running bool
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
}

func (g *Game) GridUpdate(val uint8) {
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
	var new_board [config.NumSquares][config.NumSquares]uint8 = g.board
	for y, col := range g.board {
		for x, value := range col {
			n := g.CountNeighbors(x, y)
			if value == 1 {
				//fmt.Printf("(x: %d, y: %d) %d\n", x, y, n)
				if n < 2 || n > 3 {
					new_board[y][x] = 0
				}
			} else if n == 3 {
				new_board[y][x] = 1
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
			if g.board[ny][nx] == 1 {
				count++
			}
		}
	}
	return count
}
