package scenes

import (
	"errors"
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
