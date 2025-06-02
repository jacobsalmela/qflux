package main

import (
	"log"
	"qflux/game"
	"qflux/pkg/config"

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
