package scenes

import (
	"rpg-tutorial/assets/audio/music"
	"rpg-tutorial/assets/audio/sfx"

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
	err := sfx.Sounds["Pause"].Init()
	if err != nil {
		return err
	}
	err = music.Musics["Paused"].Init()
	if err != nil {
		return err
	}
	s.loaded = true
	return nil
}

func (s *PauseScene) IsLoaded() bool {
	return s.loaded
}

func (s *PauseScene) OnEnter() error {
	err := sfx.Sounds["Pause"].Player.Rewind()
	if err != nil {
		return err
	}
	sfx.Sounds["Pause"].Player.Play()
	music.Musics["Paused"].Player.Play()
	return nil
}

func (s *PauseScene) OnExit() error {
	music.Musics["Paused"].Player.Pause()
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
