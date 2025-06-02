package scenes

import (
	"errors"
	"image/color"
	"qflux/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type Id uint

const (
	SplashId        Id = iota // 0
	TitleId                   // 1
	GameId                    // 2
	StartId                   // 3
	PauseId                   // 4
	ExitId                    // 5
	GameOverId                // 6
	SettingsId                // 7
	HighScoresId              // 8
	DifficultyId              // 9
	ControlsId                // 10
	AudioVideoId              // 11
	AccessibilityId           // 12
)

type Base struct {
	Loaded      bool
	Id          Id
	Next        Id
	FreezeFrame *ebiten.Image
	Menu        *menu.Menu
	Name        string
}

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Init() error
	OnEnter() error
	OnExit() error
	IsLoaded() bool
	GetID() Id
	GetNext() Id
	SetFreezeFrame(screen *ebiten.Image)
	Slug() string
}

// registry of scene constructors
var (
	registry = map[Id]func() Scene{}
)

// called by each subpackage in its init()
func Register(id Id, factory func() Scene) {
	registry[id] = factory
}

// build all scenes
func NewSceneMap() (map[Id]Scene, Scene, error) {
	sceneMap := make(map[Id]Scene, len(registry))
	for id, factory := range registry {
		s := factory()
		if err := s.Init(); err != nil {
			return nil, nil, err
		}
		sceneMap[id] = s
	}
	start, ok := sceneMap[SplashId]
	if !ok {
		return nil, nil, errors.New("no SplashScreen registered")
	}
	return sceneMap, start, nil
}

// lerp performs linear interpolation between two values
// a. starting value
// b. ending value
// t. interpolation factor (0.0 to 1.0)
func Lerp(a, b uint8, t float64) uint8 {
	return uint8(float64(a) + (float64(b)-float64(a))*t)
}

// lerpColor performs linear interpolation between two colors
// c1. starting color
// c2. ending color
// t. interpolation factor (0.0 to 1.0)
func LerpColor(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: Lerp(c1.R, c2.R, t),
		G: Lerp(c1.G, c2.G, t),
		B: Lerp(c1.B, c2.B, t),
		A: Lerp(c1.A, c2.A, t),
	}
}
