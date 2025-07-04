package gameplay

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"qflux/assets/images/billboards"
	"qflux/assets/images/vehicles"
	"qflux/camera"
	"qflux/constants"
	"qflux/entities"
	"qflux/scenes"
)

type GameScene struct {
	scenes.Base
	Player     entities.Player
	Road       []entities.Entity
	Billboards []entities.Entity
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
		Player: entities.Player{ // TODO: needs constructor function
			Camera: camera.NewCamera(0, 0, 0, 68, 4.8),
			Entity: &entities.Entity{
				X: 0,
				Y: 0,
				Z: 0,
			},
			Speed:        0.0, // start game in stopped state
			LateralSpeed: 350.0,
		},
		Road: initRoadSegments(1000),
		Billboards: []entities.Entity{
			{
				X: 250, // off to the right a bit
				Y: 20,  // slightly off the ground
				Z: 100,
			},
		},
	}
}

const (
	// fixed timestep: ebitengine's Update() always runs at 60 TPS
	dt = 1.0 / 60.0

	// forward motion (units in m/s2)
	forwardAccel    = 90.0    // acceleration when "w" is pressed
	forwardDecel    = 50.0    // natural deceleration (friction) when no key is pressed
	brakeDecel      = 90.0    // braking deceleration when "s" is pressed while moving
	maxForwardSpeed = 56.8224 // m/s 96.5606 km/h ~= 26.8224 m/s

	// reverse motion
	reverseAccel    = 90.0 // reverse acceleration
	maxReverseSpeed = 90.0 // max reverse speed FIXME

)

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
	var err error
	// FIXME: loop all assets
	playerImageImage, _, err := image.Decode(bytes.NewReader(vehicles.Gtr))
	if err != nil {
		return fmt.Errorf("failed to decode player image: %v", err)
	}
	s.Player.Img = ebiten.NewImageFromImage(playerImageImage)
	billboardImageImage, _, err := image.Decode(bytes.NewReader(billboards.Ebitengine))
	if err != nil {
		return fmt.Errorf("failed to decode billboard image: %v", err)
	}
	s.Billboards[0].Img = ebiten.NewImageFromImage(billboardImageImage)
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

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		// accelerate forward
		s.Player.Speed += forwardAccel * dt
		if s.Player.Speed > maxForwardSpeed {
			s.Player.Speed = maxForwardSpeed
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		// accelerate in reverse if speed is zero or negative
		s.Player.Speed -= reverseAccel * dt
		if s.Player.Speed < -maxReverseSpeed {
			s.Player.Speed = -maxReverseSpeed
		}
	} else {
		// no forward or reverse input; apply natural deceleration toward 0
		if s.Player.Speed > 0 {
			s.Player.Speed -= forwardDecel * dt
			if s.Player.Speed < 0 {
				s.Player.Speed = 0
			}
		} else if s.Player.Speed < 0 {
			// if reversing, decelerate toward 0 (reduce reverse speed)
			s.Player.Speed += forwardDecel * dt
			if s.Player.Speed > 0 {
				s.Player.Speed = 0
			}
		}
	}

	// move left or right
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		s.Player.Camera.X -= s.Player.LateralSpeed * dt
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		s.Player.Camera.X += s.Player.LateralSpeed * dt
	}

	// update player position
	s.Player.Camera.Z += s.Player.Speed * dt

	log.Printf("Speed: %.3f m/s, Position: %.3f", s.Player.Speed, s.Player.Camera.Z)
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
