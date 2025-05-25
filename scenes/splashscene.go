package scenes

import (
	"bytes"
	"rpg-tutorial/assets/images/thirdparty"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SplashScene struct {
	scene
	images     []*ebiten.Image
	elapsed    float64
	fadeInTime time.Duration
}

var _ Scene = (*SplashScene)(nil)

func NewSplashScene() *SplashScene {
	return &SplashScene{
		scene: scene{
			loaded: false,
			id:     SplashSceneId,
			next:   SplashSceneId,
		},
		fadeInTime: 4 * time.Second,
	}
}

func (s *SplashScene) Init() error {
	// loop through desired byte images for the splash screen
	for _, sponsor := range [][]byte{
		thirdparty.Splash_144x144_trasparent,
	} {
		// create a new ebiten image
		splash, _, err := ebitenutil.NewImageFromReader(
			bytes.NewReader(sponsor),
		)
		if err != nil {
			return err
		}
		// append it to the list of images to display
		s.images = append(s.images, splash)
	}
	s.loaded = true
	return nil
}

func (s *SplashScene) IsLoaded() bool {
	return s.loaded
}

func (s *SplashScene) OnEnter() error {
	return nil
}

func (s *SplashScene) OnExit() error {
	return nil
}

func (s *SplashScene) Update() error {
	s.elapsed += 0.1
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		s.next = TitleSceneId
		return nil
	}

	s.next = SplashSceneId
	return nil
}

func (s *SplashScene) ID() SceneId {
	return s.id
}

func (s *SplashScene) Next() SceneId {
	return s.next
}

func (s *SplashScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
