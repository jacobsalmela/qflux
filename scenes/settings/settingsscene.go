package settings

import (
	"qflux/assets/audio/sfx"
	"qflux/menu"
	"qflux/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type SettingsScene struct {
	scenes.Base
	elapsed float64
}

var _ scenes.Scene = (*SettingsScene)(nil)

func NewScene() scenes.Scene {
	s := &SettingsScene{
		Base: scenes.Base{
			Loaded: false,
			Id:     scenes.SettingsId,
			Next:   scenes.SettingsId,
			Name:   "Settings",
		},
	}

	s.Menu = &menu.Menu{
		Items: []menu.MenuItem{
			// FIXME: add scenes for these, or should there be multiple menus instead?
			{Label: "Difficulty", Action: func() { s.Next = scenes.SettingsId }},
			{Label: "Controls", Action: func() { s.Next = scenes.SettingsId }},
			{Label: "Audio Video", Action: func() { s.Next = scenes.SettingsId }},
			{Label: "Accessibility", Action: func() { s.Next = scenes.SettingsId }},
		},
		Index: menu.Settings,
	}

	return s
}

func init() {
	scenes.Register(scenes.SettingsId, NewScene)
}

func (s *SettingsScene) Slug() string {
	return s.Name
}

func (s *SettingsScene) Init() error {
	// TODO: loop this (and all) scenes sounds to initialize them
	if err := sfx.Sounds["Menu Select"].Init(); err != nil {
		panic("failed to initialize sound")
	}
	if err := sfx.Sounds["Menu Confirm"].Init(); err != nil {
		panic("failed to initialize sound")
	}
	s.Loaded = true
	return nil
}

func (s *SettingsScene) IsLoaded() bool {
	return s.Loaded
}

func (s *SettingsScene) OnEnter() error {
	return nil
}

func (s *SettingsScene) OnExit() error {
	return nil
}

func (s *SettingsScene) GetID() scenes.Id {
	return s.Id
}

func (s *SettingsScene) GetNext() scenes.Id {
	return s.Next
}

func (s *SettingsScene) SetFreezeFrame(screen *ebiten.Image) {
	s.FreezeFrame = screen
}
