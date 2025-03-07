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
	nextSceneId := g.sceneMap[g.activeSceneId].Update()
	// switched scenes
	if nextSceneId == scenes.ExitSceneId {
		g.sceneMap[g.activeSceneId].OnExit()
		return ebiten.Termination
	}
	if nextSceneId != g.activeSceneId {
		nextScene := g.sceneMap[nextSceneId]
		// if not loaded? then load in
		if !nextScene.IsLoaded() {
			nextScene.FirstLoad()
		}
		nextScene.OnEnter()
		g.sceneMap[g.activeSceneId].OnExit()
	}
	g.activeSceneId = nextSceneId
	return nil
}
