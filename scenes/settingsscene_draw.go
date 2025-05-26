package scenes

import (
	"image/color"
	"rpg-tutorial/assets/fonts"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *SettingsScene) Draw(screen *ebiten.Image) {
	s.drawMenu(screen)
}

func (s *SettingsScene) drawMenu(screen *ebiten.Image) {
	x := float64(320 / 2)
	y := float64(25) //FIXME: hardcode
	s.menu.Draw(screen, x, y, fonts.RobotoMediumFontFace, color.RGBA{R: 0, G: 0, B: 125, A: 255})
}
