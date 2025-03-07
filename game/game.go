package game

import "rpg-tutorial/scenes"

type Game struct {
	sceneMap      map[scenes.SceneId]scenes.Scene
	activeSceneId scenes.SceneId
}

func NewGame() *Game {
	sceneMap := map[scenes.SceneId]scenes.Scene{
		scenes.GameSceneId:  scenes.NewGameScene(),
		scenes.StartSceneId: scenes.NewStartScene(),
		scenes.PauseSceneId: scenes.NewPauseScene(),
	}
	activeSceneId := scenes.StartSceneId
	sceneMap[activeSceneId].FirstLoad()
	return &Game{
		sceneMap,
		activeSceneId,
	}
}
