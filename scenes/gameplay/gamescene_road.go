package gameplay

import (
	"image/color"
	"qflux/entities"
	"qflux/pkg/pseudo"
	"sort"

	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cameraDepth  = 4.8 * 1.0 // looks good here, but leave adjustable
	cameraX      = 0
	cameraHeight = 68 // hardcoded starting value
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

// projectPoint projects a 3D in-world point (X,Y,Z) into 2D screen space
// farthestZ is used to compute a minimal scale (such as pinning the farthest point at the horizon)
func (s *GameScene) projectPoint(worldX, worldY, worldZ, farthestZ float64) (screenX, screenY, scale float64) {
	// 1. compute the effective distance from the camera along Z
	// assume the player is at s.Player.Z along the track
	effectiveZ := worldZ - s.Player.Z
	// if worldZ == s.Player.Z, then the point is "at the camera", so clamp it to 0.1 to avoid division by 0
	if effectiveZ < 0.1 {
		effectiveZ = 0.1
	}

	// 2. scale factor - nearer objects (small effectiveZ) get a larger scale
	scale = cameraDepth / effectiveZ

	// 3. compute the minimum scale for the farthest visible Z
	minScale := cameraDepth / farthestZ

	// 4. compute baseY: the screen Y coordinate for a worldY == 0 point
	// if the camera height changes, everything shifts up/down proportionally
	baseY := horizonY + cameraHeight*(scale-minScale)

	// 5. adjust for the object's actual worldY (vertical offset above the ground)
	// if worldY == 0, the object is on the road, then it is exactly at baseY
	// if an object was above the road, worldY*scale will push it upward on the screen
	screenY = baseY - worldY*scale

	// 6. compute screenX: center of the screen + (lateral worldX offset)*scale
	// cameraX is how far we panned left or right
	// worldX of 0 is the center of road, or screenWidth /2
	screenX = float64(config.ScreenWidth)/2 + (worldX-cameraX)*scale

	return screenX, screenY, scale
}

func (s *GameScene) drawRoad(screen *ebiten.Image) {
	const trackLength = 1000.0
	const segmentLength = 1.0
	const roadWidth = 200.0

	segments := []entities.Entity{}
	for _, segment := range s.Road {
		if segment.Z > s.Player.Z && segment.Z <= trackLength {
			segments = append(segments, segment)
		}
	}

	sort.Slice(segments, func(i, j int) bool {
		return segments[i].Z > segments[j].Z
	})

	for i, segment := range segments {
		if segment.Z <= s.Player.Z {
			continue
		}
		if segment.Z > trackLength {
			break
		}

		sxTop, syTop, scaleTop := s.projectPoint(0, 0, segment.Z, trackLength)
		sxBot, syBot, scaleBot := s.projectPoint(0, 0, segment.Z+segmentLength, trackLength)

		leftTop := sxTop - (roadWidth / 2 * scaleTop)
		rightTop := sxTop + (roadWidth / 2 * scaleTop)
		leftBot := sxBot - (roadWidth / 2 * scaleBot)
		rightBot := sxBot + (roadWidth / 2 * scaleBot)

		var segmentColor color.RGBA
		if i%2 == 0 {
			segmentColor = color.RGBA{192, 192, 192, 255}
		} else {
			segmentColor = color.RGBA{80, 80, 80, 255}
		}
		verticies := []ebiten.Vertex{
			{
				DstX:   float32(leftTop),
				DstY:   float32(syTop),
				ColorR: float32(segmentColor.R) / 255,
				ColorG: float32(segmentColor.G) / 255,
				ColorB: float32(segmentColor.B) / 255,
				ColorA: 1,
			},
			{
				DstX:   float32(rightTop),
				DstY:   float32(syTop),
				ColorR: float32(segmentColor.R) / 255,
				ColorG: float32(segmentColor.G) / 255,
				ColorB: float32(segmentColor.B) / 255,
				ColorA: 1,
			},
			{
				DstX:   float32(rightBot),
				DstY:   float32(syBot),
				ColorR: float32(segmentColor.R) / 255,
				ColorG: float32(segmentColor.G) / 255,
				ColorB: float32(segmentColor.B) / 255,
				ColorA: 1,
			},
			{
				DstX:   float32(leftBot),
				DstY:   float32(syBot),
				ColorR: float32(segmentColor.R) / 255,
				ColorG: float32(segmentColor.G) / 255,
				ColorB: float32(segmentColor.B) / 255,
				ColorA: 1,
			},
		}

		indicies := []uint16{0, 1, 2, 0, 2, 3}
		opts := &ebiten.DrawTrianglesOptions{}
		screen.DrawTriangles(verticies, indicies, segment.Img, opts)
	}
}
