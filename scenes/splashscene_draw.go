package scenes

import (
	"math"
	"qflux/pkg/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/colorm"
)

func (s *SplashScene) Draw(screen *ebiten.Image) {
	// always draw the ebitengine logo, drawn from vector and text packages
	// s.drawPoweredByEbitengine(screen)

	for _, splash := range s.images {
		// calculate the scaling factor since images may be different sizes
		imgWidth, imgHeight := splash.Bounds().Dx(), splash.Bounds().Dy()
		scaleX := float64(config.ScreenWidth) / float64(imgWidth)
		scaleY := float64(config.ScreenHeight) / float64(imgHeight)

		// use the smaller of the two to maintain aspect ratio
		scale := math.Min(scaleX, scaleY)

		// create draw options and scale the image with the above calculations
		op := &colorm.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)

		// center the image on the screen
		tx := (float64(config.ScreenWidth) - float64(imgWidth)*scale) / 2
		ty := (float64(config.ScreenHeight) - float64(imgHeight)*scale) / 2
		op.GeoM.Translate(tx, ty)

		// calculate a normalized alpha value based on how much time has passed relative to the total fade-in duration
		//s.elapsed represents the time that has passed since the splash started
		// s.fadeInTime.Seconds() converts the total fade-in duration into seconds
		// this results in a value from 0 to 1+
		alpha := s.elapsed / s.fadeInTime.Seconds()
		// clamp the ratio so the fade-in stops and the image is visible
		if alpha > 1 {
			alpha = 1
		}
		// scale RGB by 1 (leaving them unchanged)
		// scale alpha channel by the computed alpha
		// this draws the image with a gradually increasing opacity
		cm := colorm.ColorM{}
		cm.Scale(1, 1, 1, alpha)
		colorm.DrawImage(screen, splash, cm, op)
	}
}

// func (s *SplashScene) drawPoweredByEbitengine(screen *ebiten.Image) {
// 	// Decide how big to scale your 96×96 logo.
// 	scale := 1.0

// 	// 1) Render the logo to its own image.
// 	logoImg := s.drawEbitenLogoImage(false, false, scale)
// 	logoW, logoH := logoImg.Bounds().Dx(), logoImg.Bounds().Dy()

// 	// 2) Measure the text sizes (using your chosen font).
// 	textAbove := "Powered by"
// 	textBelow := "Ebitengine™"

// 	// Width/height in pixels for each text string
// 	aboveW, aboveH := text.Measure(textAbove, fonts.RobotoMediumFontFace, 24)
// 	belowW, belowH := text.Measure(textBelow, fonts.RobotoBoldFontFace, 36)

// 	// 3) Figure out the composite image size with extra padding
// 	topPadding := int(5 * scale)
// 	bottomPadding := int(5 * scale)
// 	spacing := int(10 * scale)
// 	margin := 20
// 	totalW := max(logoW, int(aboveW), int(belowW)) + margin
// 	totalH := topPadding + int(aboveH) + spacing + logoH + spacing + int(belowH) + bottomPadding

// 	composite := ebiten.NewImage(totalW, totalH)
// 	composite.Fill(color.Black) // or color.Transparent if you want no background

// 	// 4) Draw textAbove at the top
// 	{
// 		opts := &text.DrawOptions{}
// 		opts.PrimaryAlign = text.AlignCenter // center horizontally
// 		// We want the text's baseline near aboveH
// 		// (We add half of aboveH if you want to center it, or just aboveH if you want it flush)
// 		opts.GeoM.Translate(float64(totalW)/2, float64(topPadding))
// 		text.Draw(composite, textAbove, fonts.RobotoMediumFontFace, opts)
// 	}

// 	// Draw the logo in the middle
// 	{
// 		op := &ebiten.DrawImageOptions{}
// 		op.GeoM.Translate(float64(totalW-logoW)/2, float64(topPadding+int(aboveH)+spacing))
// 		composite.DrawImage(logoImg, op)
// 	}

// 	// Draw textBelow at the bottom
// 	{
// 		opts := &text.DrawOptions{}
// 		opts.PrimaryAlign = text.AlignCenter
// 		// place below the logo
// 		opts.GeoM.Translate(float64(totalW)/2, float64(topPadding+int(aboveH)+spacing+logoH+spacing))
// 		text.Draw(composite, textBelow, fonts.RobotoBoldFontFace, opts)
// 	}

// 	// 5) Finally, center the composite image on the screen
// 	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
// 	finalOp := &ebiten.DrawImageOptions{}
// 	finalOp.GeoM.Translate(float64(sw-totalW)/2, float64(sh-totalH)/2)
// 	screen.DrawImage(composite, finalOp)
// }

// func (s *SplashScene) drawEbitenLogoImage(aa bool, line bool, scale float64) *ebiten.Image {
// 	const unit = 16
// 	const originalSize = 96.0 // original bounding box: 96x96
// 	width := int(originalSize * scale)
// 	height := int(originalSize * scale)
// 	logoImg := ebiten.NewImage(width, height)

// 	// Build the vector path for the logo.
// 	var path vector.Path
// 	path.MoveTo(0, 4*unit)
// 	path.LineTo(0, 6*unit)
// 	path.LineTo(2*unit, 6*unit)
// 	path.LineTo(2*unit, 5*unit)
// 	path.LineTo(3*unit, 5*unit)
// 	path.LineTo(3*unit, 4*unit)
// 	path.LineTo(4*unit, 4*unit)
// 	path.LineTo(4*unit, 2*unit)
// 	path.LineTo(6*unit, 2*unit)
// 	path.LineTo(6*unit, 1*unit)
// 	path.LineTo(5*unit, 1*unit)
// 	path.LineTo(5*unit, 0)
// 	path.LineTo(4*unit, 0)
// 	path.LineTo(4*unit, 2*unit)
// 	path.LineTo(2*unit, 2*unit)
// 	path.LineTo(2*unit, 3*unit)
// 	path.LineTo(unit, 3*unit)
// 	path.LineTo(unit, 4*unit)
// 	path.Close()

// 	// Prepare transformation for scaling.
// 	var geoM ebiten.GeoM
// 	geoM.Scale(scale, scale)

// 	// Draw the logo into the image.
// 	if line {
// 		op := &vector.StrokeOptions{
// 			Width:    float32(5 * scale),
// 			LineJoin: vector.LineJoinRound,
// 		}
// 		vector.StrokePath(logoImg, path.ApplyGeoM(geoM), color.RGBA{0xdb, 0x56, 0x20, 0xff}, aa, op)
// 	} else {
// 		vector.DrawFilledPath(logoImg, path.ApplyGeoM(geoM), color.RGBA{0xdb, 0x56, 0x20, 0xff}, aa, vector.FillRuleNonZero)
// 	}
// 	return logoImg
// }
