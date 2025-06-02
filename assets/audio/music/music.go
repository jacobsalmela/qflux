package music

import (
	"bytes"
	_ "embed"
	"log"
	"time"

	internalaudio "qflux/assets/audio"

	"github.com/hajimehoshi/ebiten/v2/audio"
	// "github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	exampleaudio "github.com/hajimehoshi/ebiten/v2/examples/resources/audio"
)

var (
	_20_Pause = exampleaudio.Ragtime_ogg
	// uncomment to replace the above and add your own file
	////go:embed "20. Pause.wav"
	// _20_Pause []byte
	Musics map[string]*Music
)

type Music struct {
	Data                []byte // raw data from file
	Player              *audio.Player
	name                string
	sampleRate          int64
	bufferSize          time.Duration
	bitsPerSecond       int64
	introLengthInSecond int64
	loopLengthInSecond  int64
}

func Init() {
	Musics = make(map[string]*Music)
	assets := map[string][]byte{
		"Paused": _20_Pause,
	}

	for name, data := range assets {
		sound := &Music{
			Data:                data,
			sampleRate:          44100,                 // FIXME: hardcode
			bufferSize:          50 * time.Millisecond, // FIXME: hardcode
			name:                name,
			bitsPerSecond:       2304,
			introLengthInSecond: 0,
			loopLengthInSecond:  7,
		}

		// add the sound to the map
		Musics[name] = sound
	}
}

func (s *Music) Init() error {
	var player *audio.Player
	var err error

	// create the reader
	r := bytes.NewReader(s.Data)
	// uncomment if using a wav or use another audio package from ebitengine
	// d, err := wav.DecodeF32(r)
	d, err := vorbis.DecodeF32(r)
	if err != nil {
		log.Fatalf("failed to decode sound: %v", err)
	}

	m := audio.NewInfiniteLoopWithIntroF32(
		d,
		s.introLengthInSecond*s.bitsPerSecond*s.sampleRate,
		s.loopLengthInSecond*s.bitsPerSecond*s.sampleRate,
	)

	player, err = internalaudio.AudioContext44100.NewPlayerF32(m)
	if err != nil {
		log.Fatalf("failed to create player: %v", err)
	}
	s.Player = player
	s.Player.SetBufferSize(s.bufferSize)
	return nil
}
