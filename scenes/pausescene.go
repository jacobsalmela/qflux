package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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

func (s *PauseScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 255, 0, 255})
	ebitenutil.DebugPrint(screen, "Press enter to unpause.")
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

func (s *PauseScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = GameSceneId
		return nil
	}

	s.next = PauseSceneId
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
