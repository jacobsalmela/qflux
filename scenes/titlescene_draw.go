package scenes

import (
	"rpg-tutorial/assets/fonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *TitleScene) Draw(screen *ebiten.Image) {
	s.drawTitle(screen)
}

func (s *TitleScene) drawTitle(screen *ebiten.Image) {
	label := "Quantum Flux"
	top := &text.DrawOptions{}
	top.PrimaryAlign = text.AlignCenter
	top.SecondaryAlign = text.AlignCenter
	top.GeoM.Reset()
	top.GeoM.Translate(320/2, 240/2)
	text.Draw(screen, label, fonts.CommonFontFace, top)
}
