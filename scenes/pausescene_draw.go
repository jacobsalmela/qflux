package scenes

import (
	"image/color"
	"rpg-tutorial/assets/fonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *PauseScene) Draw(screen *ebiten.Image) {
	if s.freezeFrame != nil {
		screen.DrawImage(s.freezeFrame, nil)
	}
	s.drawPausedText(screen)
}

func (s *PauseScene) drawPausedText(screen *ebiten.Image) {
	// draw a translucent overlay
	overlay := ebiten.NewImage(320, 240)
	overlay.Fill(color.RGBA{0, 0, 0, 100})
	screen.DrawImage(overlay, nil)

	// draw paused text
	label := "-- PAUSED --"
	top := &text.DrawOptions{}
	top.PrimaryAlign = text.AlignCenter
	top.SecondaryAlign = text.AlignCenter
	top.GeoM.Reset()
	top.GeoM.Translate(320/2, 240/2)
	text.Draw(screen, label, fonts.CommonFontFace, top)
}
