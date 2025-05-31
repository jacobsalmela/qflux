package entities

import "github.com/hajimehoshi/ebiten/v2"

// Entity is any in-world object that has a position and possibly a sprite
type Entity struct {
	X      float64 // in-world x coordinate
	Y      float64 // in-world y coordinate
	Z      float64 // in-world z coordinate
	Dx, Dy float64
	Img    *ebiten.Image // the sprite
}
