package main

import (
	"log"
	"qflux/game"
	"qflux/pkg/config"
	_ "qflux/scenes/gameover"
	_ "qflux/scenes/gameplay"
	_ "qflux/scenes/pause"
	_ "qflux/scenes/settings"
	_ "qflux/scenes/splash"
	_ "qflux/scenes/title"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(config.ScreenWidth, config.ScreenHeight)
	ebiten.SetWindowTitle("Quantum Flux")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
