package fonts

import (
	_ "embed"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	CommonFontFace text.Face
	// uncomment with your own font or keep the ebiten font below
	////go:embed "HopeGold/HopeGold.ttf"
	// commonFontBytes []byte
	commonFontBytes = fonts.MPlus1pRegular_ttf
)

func Init() {
	var err error
	CommonFontFace, err = loadFont(commonFontBytes, 32)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFont(fontData []byte, size float64) (text.Face, error) {
	// optionally, read from a local file
	// ttfFont, err := os.ReadFile(string(fontData))
	// if err != nil {
	// 	return nil, nil
	// }

	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		return nil, err
	}

	golangFontFace := truetype.NewFace(ttfFont, &truetype.Options{
		Size: size,
		DPI:  72,
	})

	// wrap the golang font face into a text/v2.Face
	return text.NewGoXFace(golangFontFace), nil
}
