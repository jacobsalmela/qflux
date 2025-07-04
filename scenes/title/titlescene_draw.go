package title

import (
	"image/color"
	"qflux/assets/fonts"
	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

func (s *TitleScene) Draw(screen *ebiten.Image) {
	screen.Fill(s.calculateBgColor())
	for _, star := range s.stars {
		star.Draw(screen)
	}
	s.drawTitle(screen)
	s.drawMenu(screen)
}

func (s *TitleScene) drawTitle(screen *ebiten.Image) {
	label := "Quantum Flux"
	top := &text.DrawOptions{}
	top.PrimaryAlign = text.AlignCenter
	top.SecondaryAlign = text.AlignCenter
	top.GeoM.Reset()
	top.GeoM.Translate(config.ScreenWidth/2, 20) //FIXME: hardcode
	text.Draw(screen, label, fonts.CommonFontFace, top)
}

func (s *TitleScene) drawMenu(screen *ebiten.Image) {
	x := float64(config.ScreenWidth / 2)
	y := float64(75) //FIXME: hardcode
	s.Menu.Draw(screen, x, y, fonts.RobotoMediumFontFace, color.RGBA{R: 0, G: 0, B: 125, A: 255})
}
