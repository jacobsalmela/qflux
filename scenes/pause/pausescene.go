package pause

import (
	"qflux/assets/audio/music"
	"qflux/assets/audio/sfx"
	"qflux/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type PauseScene struct {
	scenes.Base
}

var _ scenes.Scene = (*PauseScene)(nil)

func NewScene() scenes.Scene {
	return &PauseScene{
		Base: scenes.Base{
			Loaded: false,
			Id:     scenes.PauseId,
			Next:   scenes.PauseId,
			Name:   "Pause",
		},
	}
}

func init() {
	scenes.Register(scenes.PauseId, NewScene)
}

func (s *PauseScene) Slug() string {
	return s.Name
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
	s.Loaded = true
	return nil
}

func (s *PauseScene) IsLoaded() bool {
	return s.Loaded
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

func (s *PauseScene) GetID() scenes.Id {
	return s.Id
}

func (s *PauseScene) GetNext() scenes.Id {
	return s.Next
}

func (s *PauseScene) SetFreezeFrame(screen *ebiten.Image) {
	s.FreezeFrame = screen
}
