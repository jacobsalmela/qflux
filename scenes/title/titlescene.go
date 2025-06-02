package title

import (
	"image/color"
	"qflux/assets/audio/sfx"
	"qflux/menu"
	"qflux/pkg/config"
	"qflux/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type TitleScene struct {
	scenes.Base
	elapsed float64
	stars   []*Star
}

type Star struct {
	X, Y   float64
	Dx, Dy float64
	Shape  string // plus, square
	Clr    color.Color
}

var _ scenes.Scene = (*TitleScene)(nil)

func NewScene() scenes.Scene {
	s := &TitleScene{
		Base: scenes.Base{
			Loaded: false,
			Id:     scenes.TitleId,
			Next:   scenes.TitleId,
			Name:   "Title",
		},
	}

	s.Menu = &menu.Menu{
		Items: []menu.MenuItem{
			{Label: "New Game", Action: func() { s.Next = scenes.GameId }},
			{Label: "Settings", Action: func() { s.Next = scenes.SettingsId }},
			{Label: "High Scores", Action: func() { s.Next = scenes.HighScoresId }},
		},
		Index: menu.NewGame,
	}

	s.stars = initStars(50, config.ScreenWidth, config.ScreenHeight) // 50 stars within screen dimensions

	return s
}

func init() {
	scenes.Register(scenes.TitleId, NewScene)
}

func (s *TitleScene) Slug() string {
	return s.Name
}

func (s *TitleScene) Init() error {
	// TODO: Loop and init all sfx
	if err := sfx.Sounds["Menu Select"].Init(); err != nil {
		return err
	}
	if err := sfx.Sounds["Menu Confirm"].Init(); err != nil {
		return err
	}
	s.Loaded = true
	return nil
}

func (s *TitleScene) IsLoaded() bool {
	return s.Loaded
}

func (s *TitleScene) OnEnter() error {
	return nil
}

func (s *TitleScene) OnExit() error {
	return nil
}

func (s *TitleScene) GetID() scenes.Id {
	return s.Id
}

func (s *TitleScene) GetNext() scenes.Id {
	return s.Next
}

func (s *TitleScene) SetFreezeFrame(screen *ebiten.Image) {
	s.FreezeFrame = screen
}
