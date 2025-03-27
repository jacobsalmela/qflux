package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (s *PauseScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = GameSceneId
		return nil
	}

	s.next = PauseSceneId
	return nil
}
