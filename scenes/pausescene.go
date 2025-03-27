package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type PauseScene struct {
	scene
}

var _ Scene = (*PauseScene)(nil)

func NewPauseScene() *PauseScene {
	return &PauseScene{
		scene: scene{
			loaded: false,
			id:     PauseSceneId,
			next:   PauseSceneId,
		},
	}
}

func (s *PauseScene) Init() error {
	s.loaded = true
	return nil
}

func (s *PauseScene) IsLoaded() bool {
	return s.loaded
}

func (s *PauseScene) OnEnter() error {
	return nil
}

func (s *PauseScene) OnExit() error {
	return nil
}

func (s *PauseScene) ID() SceneId {
	return s.id
}

func (s *PauseScene) Next() SceneId {
	return s.next
}

func (s *PauseScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
