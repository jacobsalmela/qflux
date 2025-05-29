package scenes

import (
	"image/color"
	"rpg-tutorial/menu"

	"github.com/hajimehoshi/ebiten/v2"
)

type SceneId uint

const (
	SplashSceneId SceneId = iota
	TitleSceneId
	GameSceneId
	StartSceneId
	PauseSceneId
	ExitSceneId
	GameOverSceneId
	SettingsId
	HighScoresId
	DifficultyId
	ControlsId
	AudioVideoId
	AccessibilityId
)

type scene struct {
	loaded      bool
	id          SceneId
	next        SceneId
	freezeFrame *ebiten.Image
	menu        *menu.Menu
}

type Scene interface {
	Update() error
	Draw(screen *ebiten.Image)
	Init() error
	OnEnter() error
	OnExit() error
	IsLoaded() bool
	ID() SceneId
	Next() SceneId
	SetFreezeFrame(screen *ebiten.Image)
}

// lerp performs linear interpolation between two values
// a. starting value
// b. ending value
// t. interpolation factor (0.0 to 1.0)
func lerp(a, b uint8, t float64) uint8 {
	return uint8(float64(a) + (float64(b)-float64(a))*t)
}

// lerpColor performs linear interpolation between two colors
// c1. starting color
// c2. ending color
// t. interpolation factor (0.0 to 1.0)
func lerpColor(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: lerp(c1.R, c2.R, t),
		G: lerp(c1.G, c2.G, t),
		B: lerp(c1.B, c2.B, t),
		A: lerp(c1.A, c2.A, t),
	}
}
