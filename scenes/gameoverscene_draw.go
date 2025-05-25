package scenes

import (
	"rpg-tutorial/assets/fonts"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *GameOverScene) Draw(screen *ebiten.Image) {
	s.drawGameOverText(screen)
}

func (s *GameOverScene) drawGameOverText(screen *ebiten.Image) {
	label := "Game Over"
	top := &text.DrawOptions{}
	top.PrimaryAlign = text.AlignCenter
	top.SecondaryAlign = text.AlignCenter
	top.GeoM.Reset()
	top.GeoM.Translate(320/2, 240/2)
	text.Draw(screen, label, fonts.CommonFontFace, top)
}
