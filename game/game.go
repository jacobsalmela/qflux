package game

import (
	"fmt"
	"log"
	"qflux/scenes"
)

type Game struct {
	sceneMap map[scenes.Id]scenes.Scene
	current  scenes.Scene
}

func NewGame() *Game {
	sceneMap, firstScene, err := scenes.NewSceneMap()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range sceneMap {
		fmt.Printf("Initializing (%d) %v scene...\n", s.GetID(), s.Slug())
		if err := s.Init(); err != nil {
			log.Fatal(err)
		}
	}
	return &Game{
		sceneMap: sceneMap,
		current:  sceneMap[firstScene.GetID()],
	}
}
