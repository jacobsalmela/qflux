package gameplay

import (
	"image/color"
	"qflux/entities"
	"qflux/pkg/pseudo"

	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cameraDepth  = 4.8 * 1.0 // looks good here, but leave adjustable
	cameraHeight = 68        // hardcoded starting value
	horizonY     = config.ScreenHeight / 2
)

func initRoadSegments(qty int) []entities.Entity {
	segments := make([]entities.Entity, 0)

	// segments can be placed anywhere from startZ to endZ
	startZ := 0.0
	endZ := 1000.0

	// loop qty times to space segments evenly
	for i := 0; i < qty; i++ {
		t := float64(i) / float64(qty-1)
		baseZ := pseudo.LerpF(startZ, endZ, t)
		// clamp baseZ to prevent overshoot
		if baseZ < startZ {
			baseZ = startZ
		}
		if baseZ > endZ {
			baseZ = endZ
		}

		// create a new segment
		segmentImg := ebiten.NewImage(1, 1)
		segmentImg.Fill(color.White)
		segment := entities.Entity{
			X:   0,              // center-line
			Y:   0,              // flat on the road
			Z:   float64(baseZ), // segment's depth in the world
			Img: segmentImg,     // 1x1 pixel, tinted later
		}
		// add the segment to the slice
		segments = append(segments, segment)
	}
	return segments
}
