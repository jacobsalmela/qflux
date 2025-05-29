package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameOverScene struct {
	scene
	elapsed float64
}

var _ Scene = (*GameOverScene)(nil)

func NewGameOverScene() *GameOverScene {
	return &GameOverScene{
		scene: scene{
			loaded: false,
			id:     GameOverSceneId,
			next:   GameOverSceneId,
		},
	}
}

func (s *GameOverScene) Init() error {
	s.loaded = true
	return nil
}

func (s *GameOverScene) IsLoaded() bool {
	return s.loaded
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
		s.next = TitleSceneId
		return nil
	}

	s.elapsed += 0.6
	s.next = GameOverSceneId
	return nil
}

func (s *GameOverScene) ID() SceneId {
	return s.id
}

func (s *GameOverScene) Next() SceneId {
	return s.next
}

func (s *GameOverScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
