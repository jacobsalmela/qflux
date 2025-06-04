package gameplay

import (
	"image/color"

	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
)

func (s *GameScene) Draw(screen *ebiten.Image) {
	s.drawSky(screen)
	s.drawGround(screen)
	s.drawRoad(screen)
	s.drawPlayer(screen)
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
