package pause

import (
	"qflux/scenes"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (s *PauseScene) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.Next = scenes.GameId
		return nil
	}

	s.Next = scenes.PauseId
	return nil
}
