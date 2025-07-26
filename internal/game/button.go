package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	x, y float64
	w, h int
	label string
	hov bool
	click bool
	clr color.Color
	hovClr color.Color
	fntClr color.Color
	face font.Face
}

func NewButton(x, y float64, w, h int, label string, hov, click bool, clr, hovClr, fntClr color.Color, face font.Face) *Button {
	return &Button {x, y, w, h, label, hov, click, clr, hovClr, fntClr, face}
}

func (b *Button) Update() error {
	int_x, int_y := ebiten.CursorPosition()
	x := float64(int_x)
	y := float64(int_y)
	b.hov = x >= b.x && 
			x <= b.x+float64(b.w) &&
			y >= b.y &&
			y <= b.y+float64(b.h)
	
	if b.hov && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		b.click = true
	} else {
		b.click = false
	}
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	rect := ebiten.NewImage(b.w, b.h)

	if b.hov {
		rect.Fill(b.clr)
	} else {
		rect.Fill(b.hovClr)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.x), float64(b.y))
	screen.DrawImage(rect, op)

	tBounds := text.BoundString(b.face, b.label)
	textX := int(b.x) + (b.w-tBounds.Dx())/2
	textY := int(b.y) + (b.h+tBounds.Dy())/2

	text.Draw(screen, b.label, b.face, textX, textY, b.fntClr)
}
