package menu

import (
	"image/color"
	"log"
	"rpg-tutorial/assets/audio/sfx"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Menu struct {
	Items []MenuItem
	Index int
}

type MenuItem struct {
	Label  string
	Action func()
}

const (
	NewGame = iota
	Settings
	HighScores
	Difficulty
	Controls
	AudioVideo
	Accessibility
)

// Update updates the menu and handles user input.
// If Enter is pressed, the MenuItem's action function is called.
func (m *Menu) Update() error {
	// handle down arrow
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		sfx.Sounds["Menu Select"].Player.Rewind()
		sfx.Sounds["Menu Select"].Player.Play()
		m.Index++
		if m.Index >= len(m.Items) {
			m.Index = 0
		}
		log.Printf("Highlighted menu %d", m.Index)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		sfx.Sounds["Menu Select"].Player.Rewind()
		sfx.Sounds["Menu Select"].Player.Play()
		m.Index--
		if m.Index < 0 {
			m.Index = len(m.Items) - 1
		}
		log.Printf("Highlighted menu %d", m.Index)
	}

	// run the caller's function upon enter
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		if err := sfx.Sounds["Menu Confirm"].Player.Rewind(); err != nil {
			log.Printf("Failed to rewind %v", err)
		}
		sfx.Sounds["Menu Confirm"].Player.Play()
		log.Printf("Selected menu %d", m.Index)
		m.Items[m.Index].Action()
	}
	return nil
}

func (m *Menu) Draw(screen *ebiten.Image, x, y float64, fontFace text.Face, clr color.Color) {
	lineHeight := 40.0
	padding := 20.0
	maxWidth := 0.0

	// set the max width for consistent sizes with varying lengths of text
	for _, item := range m.Items {
		w, _ := text.Measure(item.Label, fontFace, 0)
		if w > maxWidth {
			maxWidth = w
		}
	}

	// check the height as well for consistent rectangle highlight shapes
	for i, item := range m.Items {
		_, h := text.Measure(item.Label, fontFace, 0)

		// if the index matches, draw the rectangle
		if i == m.Index {
			boxX := x - maxWidth/2 - padding/2
			boxY := y - float64(h)/2
			vector.DrawFilledRect(
				screen,
				float32(boxX),
				float32(boxY),
				float32(maxWidth+padding),
				float32(h),
				clr,
				false,
			)
		}

		// draw the text last so it is on top
		opts := &text.DrawOptions{}
		opts.PrimaryAlign = text.AlignCenter
		opts.SecondaryAlign = text.AlignCenter
		opts.GeoM.Translate(x, y)
		text.Draw(screen, item.Label, fontFace, opts)
		// increment the lineHeight to create the vertical stack of menu items
		y += lineHeight
	}
}
