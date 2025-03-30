package audio

import "github.com/hajimehoshi/ebiten/v2/audio"

var (
	AudioContext44100 *audio.Context
)

func Init() {
	// only one audio context can exist
	AudioContext44100 = audio.NewContext(44100)
}
