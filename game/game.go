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
		scenes.SplashSceneId:   scenes.NewSplashScene(),
		scenes.TitleSceneId:    scenes.NewTitleScene(),
		scenes.SettingsId:      scenes.NewSettingsScene(),
		scenes.HighScoresId:    scenes.NewSettingsScene(), // FIXME: link to actual scene
		scenes.DifficultyId:    scenes.NewSettingsScene(), // FIXME: link to actual scene
		scenes.ControlsId:      scenes.NewSettingsScene(), // FIXME: link to actual scene
		scenes.AudioVideoId:    scenes.NewSettingsScene(), // FIXME: link to actual scene
		scenes.AccessibilityId: scenes.NewSettingsScene(), // FIXME: link to actual scene
		scenes.GameSceneId:     scenes.NewGameScene(),
		scenes.StartSceneId:    scenes.NewStartScene(),
		scenes.PauseSceneId:    scenes.NewPauseScene(),
		scenes.GameOverSceneId: scenes.NewGameOverScene(),
	}
	activeSceneId := scenes.SplashSceneId
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
