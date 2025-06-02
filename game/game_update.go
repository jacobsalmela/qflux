package game

import (
	"fmt"
	"os"
	"qflux/pkg/config"
	"qflux/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		fmt.Println("Game ended via regular termination.")
		os.Exit(0)
	}

	// if the current scene is not the same as the next, a switch is happening
	if g.current.GetID() != g.current.GetNext() {
		// pause logic
		g.pauseCheck()

		// Exit the existing scene
		if err := g.sceneMap[g.current.GetID()].OnExit(); err != nil {
			return err
		}

		// Switch scenes
		g.current = g.sceneMap[g.current.GetNext()]

		// Enter the next scene
		if err := g.current.OnEnter(); err != nil {
			return err
		}
	}

	// Let each scene handle its own logic/updates/inputs
	if err := g.current.Update(); err != nil {
		return err
	}
	return nil
}

func (g *Game) pauseCheck() {
	if g.current.GetNext() == scenes.PauseId {
		freezeFrame := ebiten.NewImage(config.ScreenWidth, config.ScreenHeight)
		g.current.Draw(freezeFrame)
		g.sceneMap[scenes.PauseId].SetFreezeFrame(freezeFrame)
	}
}
