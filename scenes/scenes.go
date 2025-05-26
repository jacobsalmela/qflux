package scenes

import (
	"rpg-tutorial/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneId uint

const (
	SplashSceneId SceneId = iota
	TitleSceneId
	GameSceneId
	StartSceneId
	PauseSceneId
	ExitSceneId
	GameOverSceneId
	SettingsId
	HighScoresId
	DifficultyId
	ControlsId
	AudioVideoId
	AccessibilityId
)

type scene struct {
	loaded      bool
	id          SceneId
	next        SceneId
	freezeFrame *ebiten.Image
	menu        *menu.Menu
}

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Init() error
	OnEnter() error
	OnExit() error
	IsLoaded() bool
	ID() SceneId
	Next() SceneId
	SetFreezeFrame(screen *ebiten.Image)
}
