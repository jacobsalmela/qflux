package pause

import (
	"image/color"
	"qflux/assets/fonts"
	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *PauseScene) Draw(screen *ebiten.Image) {
	if s.FreezeFrame != nil {
		screen.DrawImage(s.FreezeFrame, nil)
	}
	s.drawPausedText(screen)
}

func (s *PauseScene) drawPausedText(screen *ebiten.Image) {
	// draw a translucent overlay
	overlay := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
	overlay.Fill(color.RGBA{0, 0, 0, 100})
	screen.DrawImage(overlay, nil)

	// draw paused text
	label := "-- PAUSED --"
	top := &text.DrawOptions{}
	top.PrimaryAlign = text.AlignCenter
	top.SecondaryAlign = text.AlignCenter
	top.GeoM.Reset()
	top.GeoM.Translate(config.ScreenWidth/2, config.ScreenHeight/2)
	text.Draw(screen, label, fonts.CommonFontFace, top)
}
