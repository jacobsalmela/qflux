package title

import (
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func initStars(qty, screenWidth, screenHeight int) []*Star {
	stars := make([]*Star, 0)
	shapeTypes := []string{"plus", "square"}

	for i := 0; i < qty; i++ {
		s := &Star{
			X: rand.Float64() * float64(screenWidth),  // random X position within screen dimensions
			Y: rand.Float64() * float64(screenHeight), // random Y position within screen dimensions
			// Adjust velocity as desired
			Dx:    (rand.Float64()*2 - 1) * 0.9,           // random horizontal velocity between -3 and 3
			Dy:    (rand.Float64()*2 - 1) * 0.9,           // random horizontal velocity between -3 and 3
			Shape: shapeTypes[rand.Intn(len(shapeTypes))], // random shape
			Clr: color.RGBA{
				R: uint8(rand.Intn(256)),
				G: uint8(rand.Intn(256)),
				B: uint8(rand.Intn(256)),
				A: 255,
			},
		}
		stars = append(stars, s)
	}
	return stars
}

func (star *Star) Draw(screen *ebiten.Image) {
	size := 5.0
	switch star.Shape {
	case "plus":
		// draw vertical line
		vector.StrokeLine(screen, float32(star.X-size), float32(star.Y), float32(star.X+size), float32(star.Y), 1, star.Clr, false)
		// draw horizontal line
		vector.StrokeLine(screen, float32(star.X), float32(star.Y-size), float32(star.X), float32(star.Y+size), 1, star.Clr, false)
	case "square":
		vector.DrawFilledRect(screen, float32(star.X-size), float32(star.Y-size), float32(size*2), float32(size*2), star.Clr, false)
	}
}

func (star *Star) Update(screenWidth, screenHeight int) error {
	star.X += star.Dx // update x position based on horizontal velocity
	star.Y += star.Dy // update y position based on vertical velocity

	// wrap around screen horizontally
	if star.X < 0 {
		star.X = float64(screenWidth)
	}
	if star.X > float64(screenWidth) {
		star.X = 0
	}

	// wrap around screen vertically
	if star.Y < 0 {
		star.Y = float64(screenHeight)
	}
	if star.Y > float64(screenHeight) {
		star.Y = 0
	}
	return nil
}
