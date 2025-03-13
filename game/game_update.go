package game

import (
	"fmt"
	"os"
	"rpg-tutorial/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		fmt.Println("Game ended via regular termination.")
		os.Exit(0)
	}

	// if the current scene is not the same as the next, a switch is happening
	if g.current.ID() != g.current.Next() {
		// pause logic
		g.pauseCheck()

		// Exit the existing scene
		if err := g.sceneMap[g.current.ID()].OnExit(); err != nil {
			return err
		}

		// Switch scenes
		g.current = g.sceneMap[g.current.Next()]

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
	if g.current.ID() != scenes.PauseSceneId {
		// create a new image
		ff := ebiten.NewImage(1280, 800) // FIXME hardcode
		// draw the current scene to the image
		g.current.Draw(ff)
		// set the pause screen freeze frame
		g.sceneMap[scenes.PauseSceneId].SetFreezeFrame(ff)
	}
}
