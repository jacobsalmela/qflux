package game

import (
	"fmt"
	"rpg-tutorial/assets/audio"
	"rpg-tutorial/assets/audio/music"
	"rpg-tutorial/assets/audio/sfx"
	"rpg-tutorial/assets/fonts"
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
