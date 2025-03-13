package game

import (
	"log"
	"rpg-tutorial/scenes"
)

type Game struct {
	sceneMap map[scenes.SceneId]scenes.Scene
	current  scenes.Scene
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.GameSceneId:  scenes.NewGameScene(),
		scenes.StartSceneId: scenes.NewStartScene(),
		scenes.PauseSceneId: scenes.NewPauseScene(),
	}
	activeSceneId := scenes.StartSceneId
	for _, s := range sceneMap {
		if err := sceneMap[s.ID()].Init(); err != nil {
			log.Fatal(err)
		}
	}
	return &Game{
		sceneMap: sceneMap,
		current:  sceneMap[activeSceneId],
	}
}
