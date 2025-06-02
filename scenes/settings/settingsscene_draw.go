package settings

import (
	"image/color"
	"qflux/assets/fonts"
	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *SettingsScene) Draw(screen *ebiten.Image) {
	s.drawMenu(screen)
}

func (s *SettingsScene) drawMenu(screen *ebiten.Image) {
	x := float64(config.ScreenWidth / 2)
	y := float64(25) //FIXME: hardcode
	s.Menu.Draw(screen, x, y, fonts.RobotoMediumFontFace, color.RGBA{R: 0, G: 0, B: 125, A: 255})
}
