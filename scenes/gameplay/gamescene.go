package gameplay

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"qflux/constants"
	"qflux/entities"
	"qflux/scenes"
)

type GameScene struct {
	scenes.Base
	Player entities.Entity
	Road   []entities.Entity
}

var _ scenes.Scene = (*GameScene)(nil)

func NewScene() scenes.Scene {
	return &GameScene{
		Base: scenes.Base{
			Loaded: false,
			Id:     scenes.GameId,
			Next:   scenes.GameId,
			Name:   "Gameplay",
		},
		Player: entities.Entity{ // TODO: needs constructor function
			X: 0,
			Y: 0,
			Z: 0,
		},
		Road: initRoadSegments(1000),
	}
}

func init() {
	scenes.Register(scenes.GameId, NewScene)
}

func (s *GameScene) Slug() string {
	return s.Name
}

func (s *GameScene) IsLoaded() bool {
	return s.Loaded
}

func (s *GameScene) Init() error {

	s.Loaded = true
	return nil
}

func (s *GameScene) OnEnter() error {
	return nil
}

func (s *GameScene) OnExit() error {
	return nil
}

func (s *GameScene) Update() error {
	// "T" will temporarily trigger a game over
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		s.Next = scenes.GameOverId
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		s.Next = scenes.ExitId
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.Next = scenes.PauseId
		return nil
	}

	s.Next = scenes.GameId
	return nil
}

func (s *GameScene) GetID() scenes.Id {
	return s.Id
}

func (s *GameScene) GetNext() scenes.Id {
	return s.Next
}

func (s *GameScene) SetFreezeFrame(screen *ebiten.Image) {
	s.FreezeFrame = screen
}

func CheckCollisionHorizontal(sprite *entities.Entity, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(
			image.Rect(
				int(sprite.X),
				int(sprite.Y),
				int(sprite.X)+constants.Tilesize,
				int(sprite.Y)+constants.Tilesize,
			),
		) {
			if sprite.Dx > 0.0 {
				sprite.X = float64(collider.Min.X) - constants.Tilesize
			} else if sprite.Dx < 0.0 {
				sprite.X = float64(collider.Max.X)
			}
		}
	}
}

func CheckCollisionVertical(sprite *entities.Entity, colliders []image.Rectangle) {
	for _, collider := range colliders {
		if collider.Overlaps(
			image.Rect(
				int(sprite.X),
				int(sprite.Y),
				int(sprite.X)+constants.Tilesize,
				int(sprite.Y)+constants.Tilesize,
			),
		) {
			if sprite.Dy > 0.0 {
				sprite.Y = float64(collider.Min.Y) - constants.Tilesize
			} else if sprite.Dy < 0.0 {
				sprite.Y = float64(collider.Max.X)
			}
		}
	}
}
