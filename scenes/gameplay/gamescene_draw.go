package gameplay

import (
	"image/color"
	"sort"

	"qflux/entities"
	"qflux/pkg/config"
	"qflux/pkg/pseudo"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.drawSky(screen)
	s.drawGround(screen)
	s.drawRoad(screen)
	s.drawObjects(screen)
	s.drawPlayer(screen)
}
func (s *GameScene) drawSky(screen *ebiten.Image) {
	skyImg := ebiten.NewImage(1, 1)
	skyImg.Fill(color.White)
	// top color, dark purple 80,0,80
	topR, topG, topB := float32(80)/255, float32(0)/255, float32(80)/255
	// bottom-left color, warm orange, 255,100,0
	bottomLeftR, bottomLeftG, bottomLeftB := float32(255)/255, float32(100)/255, float32(0)/255
	// bottom-right color, brighter orange, 255,200,0
	bottomRightR, bottomRightG, bottomRightB := float32(255)/255, float32(200)/255, float32(0)/255

	verticies := []ebiten.Vertex{
		{
			DstX:   0,
			DstY:   0,
			ColorR: topR,
			ColorG: topG,
			ColorB: topB,
			ColorA: 1,
		},
		{
			DstX:   float32(config.ScreenWidth),
			DstY:   0,
			ColorR: topR,
			ColorG: topG,
			ColorB: topB,
			ColorA: 1,
		},
		{
			DstX:   float32(config.ScreenWidth),
			DstY:   float32(config.ScreenHeight),
			ColorR: bottomRightR,
			ColorG: bottomRightG,
			ColorB: bottomRightB,
			ColorA: 1,
		},
		{
			DstX:   0,
			DstY:   float32(config.ScreenWidth),
			ColorR: bottomLeftR,
			ColorG: bottomLeftG,
			ColorB: bottomLeftB,
			ColorA: 1,
		},
	}

	indicies := []uint16{0, 1, 2, 0, 2, 3}

	opts := &ebiten.DrawTrianglesOptions{}
	screen.DrawTriangles(verticies, indicies, skyImg, opts)
}

func (s *GameScene) drawGround(screen *ebiten.Image) {
	horizonY := config.ScreenHeight / 2
	groundImg := ebiten.NewImage(1, 1)
	groundImg.Fill(color.White)

	// each vertex is a single corner in the mesh
	verticies := []ebiten.Vertex{
		{DstX: 0, DstY: float32(horizonY), ColorR: 0, ColorG: 0.2, ColorB: 0, ColorA: 1},
		{DstX: float32(config.ScreenWidth), DstY: float32(horizonY), ColorR: 0, ColorG: 0.2, ColorB: 0, ColorA: 1},
		{DstX: float32(config.ScreenWidth), DstY: float32(config.ScreenHeight), ColorR: 0, ColorG: 0.6, ColorB: 0, ColorA: 1},
		{DstX: 0, DstY: float32(config.ScreenHeight), ColorR: 0, ColorG: 0.6, ColorB: 0, ColorA: 1},
	}

	// the indicies tell ebitengine how to group them into triangles
	indicies := []uint16{0, 1, 2, 0, 2, 3}
	opts := &ebiten.DrawTrianglesOptions{}
	screen.DrawTriangles(verticies, indicies, groundImg, opts)

	// // visualize how triangles are drawn
	// // first triangle - blue
	// for _, i := range []uint16{0, 1, 3} {
	// 	verticies[i].ColorR, verticies[i].ColorG, verticies[i].ColorB, verticies[i].ColorA = 0.0, 0.447, 0.698, 1.0
	// }
	// screen.DrawTriangles(verticies, []uint16{0, 1, 3}, groundImg, opts)
	// // second triangle - orange
	// for _, i := range []uint16{1, 2, 3} {
	// 	verticies[i].ColorR, verticies[i].ColorG, verticies[i].ColorB, verticies[i].ColorA = 0.835, 0.369, 0.0, 1.0
	// }
	// screen.DrawTriangles(verticies, []uint16{1, 2, 3}, groundImg, opts)
}

func (s *GameScene) drawRoad(screen *ebiten.Image) {
	const trackLength = 1000.0
	const segmentLength = 1.0
	const roadWidth = 200.0

	segments := []entities.Entity{}
	for _, segment := range s.Road {
		if segment.Z > s.Player.Camera.Z && segment.Z <= trackLength {
			segments = append(segments, segment)
		}
	}

	sort.Slice(segments, func(i, j int) bool {
		return segments[i].Z > segments[j].Z
	})

	for i, segment := range segments {
		if segment.Z <= s.Player.Camera.Z {
			continue
		}
		if segment.Z > trackLength {
			break
		}

		sxTop, syTop, scaleTop := pseudo.ProjectPoint(
			0, 0, segment.Z,
			s.Player.Camera.X, s.Player.Camera.Y, s.Player.Camera.Z,
			s.Player.Camera.Depth, s.Player.Camera.Focal,
			horizonY, trackLength,
			config.ScreenWidth,
		)
		sxBot, syBot, scaleBot := pseudo.ProjectPoint(
			0, 0, segment.Z+segmentLength,
			s.Player.Camera.X, s.Player.Camera.Y, s.Player.Camera.Z,
			s.Player.Camera.Depth, s.Player.Camera.Focal,
			horizonY, trackLength,
			config.ScreenWidth,
		)

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

func (s *GameScene) drawObjects(screen *ebiten.Image) {
	const minZ = 0.1     // anything closer than this do not draw
	const maxScale = 5.0 // even at minZ, scale will not exceed this

	for _, billboard := range s.Billboards {
		// compute Z distance from camera
		dz := billboard.Z - s.Player.Camera.Z
		if dz < minZ {
			continue
		}
		opts := &ebiten.DrawImageOptions{}
		w := billboard.Img.Bounds().Dx()
		h := billboard.Img.Bounds().Dy()

		sx, sy, scale := pseudo.ProjectPoint(
			billboard.X, billboard.Y, billboard.Z,
			s.Player.Camera.X, s.Player.Camera.Y, s.Player.Camera.Z,
			s.Player.Camera.Depth, s.Player.Camera.Focal,
			horizonY, 1000.0, // FIXME
			config.ScreenWidth,
		)

		// clamp scale
		if scale > maxScale {
			scale = maxScale
		}
		opts.GeoM.Scale(scale, scale)
		opts.GeoM.Translate(
			sx-(float64(w)*scale)/2,
			sy-(float64(h)*scale),
		)
		screen.DrawImage(billboard.Img, opts)
	}
}

func (s *GameScene) drawPlayer(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	// get sprite's dimensions
	w := s.Player.Img.Bounds().Dx()
	h := s.Player.Img.Bounds().Dy()

	scale := 0.8

	// compute final w/h
	sw := float64(w) * scale
	sh := float64(h) * scale

	offsetX := float64(config.ScreenWidth)/2 - float64(sw)/2
	offsetY := float64(config.ScreenHeight) - float64(sh)

	opts.GeoM.Scale(scale, scale)
	opts.GeoM.Translate(offsetX, offsetY)
	screen.DrawImage(s.Player.Img, opts)
}
