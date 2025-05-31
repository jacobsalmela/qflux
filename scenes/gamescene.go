package scenes

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"rpg-tutorial/constants"
	"rpg-tutorial/entities"
)

type GameScene struct {
	scene
	Player entities.Entity
	Road   []entities.Entity
}

var _ Scene = (*GameScene)(nil)

func NewGameScene() *GameScene {
	return &GameScene{
		scene: scene{
			loaded: false,
			id:     GameSceneId,
			next:   GameSceneId,
		},
		Player: entities.Entity{ // TODO: needs constructor function
			X: 0,
			Y: 0,
			Z: 0,
		},
		Road: initRoadSegments(1000),
	}
}

func (s *GameScene) IsLoaded() bool {
	return s.loaded
}

func (s *GameScene) Init() error {

	s.loaded = true
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
		s.next = GameOverSceneId
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		s.next = ExitSceneId
		return nil
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = PauseSceneId
		return nil
	}

	s.next = GameSceneId
	return nil
}

func (s *GameScene) ID() SceneId {
	return s.id
}

func (s *GameScene) Next() SceneId {
	return s.next
}

func (s *GameScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
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
