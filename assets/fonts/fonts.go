package fonts

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/opentype"
)

var (
	CommonFontFace       text.Face
	RobotoMediumFontFace text.Face
	RobotoBoldFontFace   text.Face
	// uncomment with your own font or keep the ebiten font below
	////go:embed "HopeGold/HopeGold.ttf"
	// commonFontBytes []byte

	commonFontBytes = fonts.MPlus1pRegular_ttf

	//go:embed "Roboto/static/Roboto-Medium.ttf"
	roboto_Medium []byte
	//go:embed "Roboto/static/Roboto-Bold.ttf"
	roboto_Bold []byte
)

func Init() {
	var err error
	CommonFontFace, err = loadFont(commonFontBytes, 32)
	if err != nil {
		log.Fatal(err)
	}
	RobotoMediumFontFace, err = loadOpenFont(roboto_Medium, 24)
	if err != nil {
		os.Exit(1)
	}
	RobotoBoldFontFace, err = loadOpenFont(roboto_Bold, 36)
	if err != nil {
		os.Exit(1)
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

func loadOpenFont(fontData []byte, size float64) (text.Face, error) {
	// Parse the TTF font
	ttfFont, err := opentype.Parse(fontData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse TTF font: %v", err)
	}

	// Create a font.Face from the TTF font
	golangFontFace, err := opentype.NewFace(ttfFont, &opentype.FaceOptions{
		Size: size,
		DPI:  72,
	})
	if err != nil {
		return nil, err
	}
	// wrap the golang font face into a text/v2.Face
	return text.NewGoXFace(golangFontFace), nil
}
