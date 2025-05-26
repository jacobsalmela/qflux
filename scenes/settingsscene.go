package scenes

import (
	"rpg-tutorial/assets/audio/sfx"
	"rpg-tutorial/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type SettingsScene struct {
	scene
	elapsed float64
}

var _ Scene = (*SettingsScene)(nil)

func NewSettingsScene() *SettingsScene {
	s := &SettingsScene{
		scene: scene{
			loaded: false,
			id:     SettingsId,
			next:   SettingsId,
		},
	}

	s.menu = &menu.Menu{
		Items: []menu.MenuItem{
			// FIXME: add scenes for these, or should there be multiple menus instead?
			{Label: "Difficulty", Action: func() { s.next = SettingsId }},
			{Label: "Controls", Action: func() { s.next = SettingsId }},
			{Label: "Audio Video", Action: func() { s.next = SettingsId }},
			{Label: "Accessibility", Action: func() { s.next = SettingsId }},
		},
		Index: menu.Settings,
	}

	return s
}

func (s *SettingsScene) Init() error {
	// TODO: loop this (and all) scenes sounds to initialize them
	if err := sfx.Sounds["Menu Select"].Init(); err != nil {
		panic("failed to initialize sound")
	}
	if err := sfx.Sounds["Menu Confirm"].Init(); err != nil {
		panic("failed to initialize sound")
	}
	s.loaded = true
	return nil
}

func (s *SettingsScene) IsLoaded() bool {
	return s.loaded
}

func (s *SettingsScene) OnEnter() error {
	return nil
}

func (s *SettingsScene) OnExit() error {
	return nil
}

func (s *SettingsScene) ID() SceneId {
	return s.id
}

func (s *SettingsScene) Next() SceneId {
	return s.next
}

func (s *SettingsScene) SetFreezeFrame(screen *ebiten.Image) {
	s.freezeFrame = screen
}
