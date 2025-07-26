package assets

import (
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed Bitcount-Regular.ttf
var fontBytes []byte

func LoadFont(size float64) text.Face {
	tt, err := opentype.Parse(fontBytes); 
	if err != nil {
		log.Fatal(err)
	}
	fntFace, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size: size,
		DPI: 72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return text.NewGoXFace(fntFace)
}
