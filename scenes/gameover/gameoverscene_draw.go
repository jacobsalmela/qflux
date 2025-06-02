package gameover

import (
	"image"
	"qflux/assets/fonts"
	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	s.drawGameOverText(screen)
	meltImage(screen, s.elapsed)
}

func (s *GameOverScene) drawGameOverText(screen *ebiten.Image) {
	label := "Game Over"
	top := &text.DrawOptions{}
	top.PrimaryAlign = text.AlignCenter
	top.SecondaryAlign = text.AlignCenter
	top.GeoM.Reset()
	top.GeoM.Translate(config.ScreenWidth/2, config.ScreenHeight/2)
	text.Draw(screen, label, fonts.CommonFontFace, top)
}

// meltImage creates a melting effect on the screen by slicing the image into vertical strips
// and offsetting each strip vertically based on position and amount
func meltImage(screen *ebiten.Image, amount float64) {
	const sliceWidth = 4 // adjust as desired
	w, h := screen.Bounds().Dx(), screen.Bounds().Dy()

	// copy the screen content as the source for the melting effect
	overlay := ebiten.NewImage(w, h)
	opts := &ebiten.DrawImageOptions{}
	overlay.DrawImage(screen, opts)
	screen.Clear()

	// iterate over the screen width in increments of sliceWidth
	for x := 0; x < w; x += sliceWidth {
		//calculate vertical offset for this slice
		// offset increases as the slice moves further right
		offset := amount * float64(x) / float64(w)

		// define a rect for the current slice
		rect := image.Rect(x, 0, x+sliceWidth, h)
		slice := overlay.SubImage(rect).(*ebiten.Image)

		// translate the slice to its original x position withing the calculated vertical offset
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(x), offset)

		screen.DrawImage(slice, opts)
	}
}
