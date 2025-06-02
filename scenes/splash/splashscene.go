package splash

import (
	"bytes"
	"qflux/assets/images/thirdparty"
	"qflux/scenes"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SplashScene struct {
	scenes.Base
	images     []*ebiten.Image
	elapsed    float64
	fadeInTime time.Duration
}

var _ scenes.Scene = (*SplashScene)(nil)

func NewScene() scenes.Scene {
	return &SplashScene{
		Base: scenes.Base{
			Loaded: false,
			Id:     scenes.SplashId,
			Next:   scenes.SplashId,
			Name:   "Splash",
		},
		fadeInTime: 4 * time.Second,
	}
}

func init() {
	scenes.Register(scenes.SplashId, NewScene)
}

func (s *SplashScene) Slug() string {
	return s.Name
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
	s.Loaded = true
	return nil
}

func (s *SplashScene) IsLoaded() bool {
	return s.Loaded
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
		s.Next = scenes.TitleId
		return nil
	}

	s.Next = scenes.SplashId
	return nil
}

func (s *SplashScene) GetID() scenes.Id {
	return s.Id
}

func (s *SplashScene) GetNext() scenes.Id {
	return s.Next
}

func (s *SplashScene) SetFreezeFrame(screen *ebiten.Image) {
	s.FreezeFrame = screen
}
