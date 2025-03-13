package scenes

import "github.com/hajimehoshi/ebiten/v2"

type SceneId uint

const (
	GameSceneId SceneId = iota
	StartSceneId
	PauseSceneId
	ExitSceneId
)

type scene struct {
	loaded      bool
	id          SceneId
	next        SceneId
	freezeFrame *ebiten.Image
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
