package sfx

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	internalaudio "rpg-tutorial/assets/audio"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	exampleaudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
)

var (
	Menu_Confirm = exampleaudio.Jab8_wav
	// uncomment to replace the above and add your own file
	////go:embed "Menu Confirm.wav"
	// Menu_Confirm []byte
	Sounds map[string]*Sound
)

type Sound struct {
	Data       []byte // raw data from file
	Player     *audio.Player
	name       string
	sampleRate int
	bufferSize time.Duration
}

func Init() {
	Sounds = make(map[string]*Sound)
	assets := map[string][]byte{
		"Pause": Menu_Confirm,
	}

	for name, data := range assets {
		sound := &Sound{
			Data:       data,
			sampleRate: 44100,                 // FIXME: hardcode
			bufferSize: 50 * time.Millisecond, // FIXME: hardcode
			name:       name,
		}

		// add the sound to the map
		Sounds[name] = sound
	}
}

func (s *Sound) Init() error {
	var player *audio.Player
	var err error

	// create the reader
	r := bytes.NewReader(s.Data)
	d, err := wav.DecodeF32(r)
	if err != nil {
		log.Fatalf("failed to decode sound: %v", err)
	}

	player, err = internalaudio.AudioContext44100.NewPlayerF32(d)
	if err != nil {
		log.Fatalf("failed to create player: %v", err)
	}
	s.Player = player
	s.Player.SetBufferSize(s.bufferSize)
	return nil
}
