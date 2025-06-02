package game

import (
	"fmt"
	"qflux/assets/audio"
	"qflux/assets/audio/music"
	"qflux/assets/audio/sfx"
	"qflux/assets/fonts"
)

func init() {
	fmt.Println("Loading font assets...")
	fonts.Init()
	fmt.Println("Loading audio assets...")
	audio.Init()
	fmt.Println("Loading sfx assets...")
	sfx.Init()
	fmt.Println("Loading music assets...")
	music.Init()
}
