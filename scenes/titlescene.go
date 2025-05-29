package scenes

import (
	"image/color"
	"rpg-tutorial/assets/audio/sfx"
	"rpg-tutorial/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type TitleScene struct {
	scene
	elapsed float64
	stars   []*Star
}

type Star struct {
	X, Y   float64
	Dx, Dy float64
	Shape  string // plus, square
	Clr    color.Color
}

var _ Scene = (*TitleScene)(nil)

func NewTitleScene() *TitleScene {
	s := &TitleScene{
		scene: scene{
			loaded: false,
			id:     TitleSceneId,
			next:   TitleSceneId,
		},
	}

	s.menu = &menu.Menu{
		Items: []menu.MenuItem{
			{Label: "New Game", Action: func() { s.next = GameSceneId }},
			{Label: "Settings", Action: func() { s.next = SettingsId }},
			{Label: "High Scores", Action: func() { s.next = HighScoresId }},
		},
		Index: menu.NewGame,
	}

	s.stars = initStars(50, 320, 240) // 50 stars within screen dimensions

	return s
}

func (s *TitleScene) Init() error {
	// TODO: Loop and init all sfx
	if err := sfx.Sounds["Menu Select"].Init(); err != nil {
		return err
	}
	if err := sfx.Sounds["Menu Confirm"].Init(); err != nil {
		return err
	}
	s.loaded = true
	return nil
}

func (s *TitleScene) IsLoaded() bool {
	return s.loaded
}

func (s *TitleScene) OnEnter() error {
	return nil
}

func (s *TitleScene) OnExit() error {
	return nil
}

func (s *TitleScene) ID() SceneId {
	return s.id
}

func (s *TitleScene) Next() SceneId {
	return s.next
}

func (s *TitleScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
