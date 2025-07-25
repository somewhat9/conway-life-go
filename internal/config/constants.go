package config

const (
	ScreenWidth = 600
	ScreenHeight = 600+(SquareSize*2)

	NumSquares = 30
	SquareSize = ScreenWidth / NumSquares

	GridOffsetX = 0
	GridOffsetY = SquareSize*2

	FontSize = 36
	Title = "Conway's Game of Life Go"
)

