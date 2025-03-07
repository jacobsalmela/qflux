package main

import (
	"log"
	"rpg-tutorial/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 800)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
