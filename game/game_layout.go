package game

import "qflux/pkg/config"

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// always render at 320×240, even when the user drags to resize the window,
	// so Ebiten will handle scaling for you:
	return config.ScreenWidth, config.ScreenHeight
}
