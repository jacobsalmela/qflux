package billboards

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

var (
	Ebitengine = images.Gophers_jpg
	// use your own images here
	////go:embed "MyBillboard.png"
	// Ebitengine []byte
)
