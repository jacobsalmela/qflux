package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type StartScene struct {
	scene
}

var _ Scene = (*StartScene)(nil)

func NewStartScene() *StartScene {
	return &StartScene{
		scene: scene{
			loaded: false,
			id:     StartSceneId,
			next:   StartSceneId,
		},
	}
}

func (s *StartScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{255, 0, 0, 255})
	ebitenutil.DebugPrint(screen, "Press enter to start.")
}

func (s *StartScene) Init() error {
	s.loaded = true
	return nil
}

func (s *StartScene) IsLoaded() bool {
	return s.loaded
}

func (s *StartScene) OnEnter() error {
	return nil
}

func (s *StartScene) OnExit() error {
	return nil
}

func (s *StartScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = GameSceneId
		return nil
	}

	s.next = StartSceneId
	return nil
}

func (s *StartScene) ID() SceneId {
	return s.id
}

func (s *StartScene) Next() SceneId {
	return s.next
}

func (s *StartScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
