package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type TitleScene struct {
	scene
	elapsed float64
}

var _ Scene = (*TitleScene)(nil)

func NewTitleScene() *TitleScene {
	return &TitleScene{
		scene: scene{
			loaded: false,
			id:     TitleSceneId,
			next:   TitleSceneId,
		},
	}
}

func (s *TitleScene) Init() error {
	s.loaded = true
	return nil
}

func (s *TitleScene) IsLoaded() bool {
	return s.loaded
}

func (s *TitleScene) OnEnter() error {
	return nil
}

func (s *TitleScene) OnExit() error {
	return nil
}

func (s *TitleScene) Update() error {
	s.elapsed += 0.1
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = GameSceneId
		return nil
	}

	s.next = TitleSceneId
	return nil
}

func (s *TitleScene) ID() SceneId {
	return s.id
}

func (s *TitleScene) Next() SceneId {
	return s.next
}

func (s *TitleScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
