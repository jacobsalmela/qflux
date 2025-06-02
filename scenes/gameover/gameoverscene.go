package gameover

import (
	"qflux/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameOverScene struct {
	scenes.Base
	elapsed float64
}

var _ scenes.Scene = (*GameOverScene)(nil)

func NewScene() scenes.Scene {
	return &GameOverScene{
		Base: scenes.Base{
			Loaded: false,
			Id:     scenes.GameOverId,
			Next:   scenes.GameOverId,
			Name:   "Game Over",
		},
	}
}

func init() {
	scenes.Register(scenes.GameOverId, NewScene)
}

func (s *GameOverScene) Slug() string {
	return s.Name
}

func (s *GameOverScene) Init() error {
	s.Loaded = true
	return nil
}

func (s *GameOverScene) IsLoaded() bool {
	return s.Loaded
}

func (s *GameOverScene) OnEnter() error {
	return nil
}

func (s *GameOverScene) OnExit() error {
	return nil
}

func (s *GameOverScene) Update() error {
	// restart to title screen with "R"
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		s.Next = scenes.TitleId
		return nil
	}

	s.elapsed += 0.6
	s.Next = scenes.GameOverId
	return nil
}

func (s *GameOverScene) GetID() scenes.Id {
	return s.Id
}

func (s *GameOverScene) GetNext() scenes.Id {
	return s.Next
}

func (s *GameOverScene) SetFreezeFrame(screen *ebiten.Image) {
	s.FreezeFrame = screen
}
