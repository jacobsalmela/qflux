package scenes

import (
	"image/color"
	"math"
	"qflux/pkg/config"
)

func (s *TitleScene) Update() error {
	s.next = TitleSceneId // default to this scene
	for _, star := range s.stars {
		if err := star.Update(config.ScreenWidth, config.ScreenHeight); err != nil {
			return err
		}
	}
	if err := s.menu.Update(); err != nil {
		return err
	}

	s.elapsed += 0.01
	return nil
}

func (s *TitleScene) calculateBgColor() color.RGBA {
	colors := []color.RGBA{
		{R: 0, G: 0, B: 255, A: 255},     //blue
		{R: 128, G: 0, B: 128, A: 255},   //purple
		{R: 255, G: 105, B: 180, A: 255}, //pink
	}

	segmentDuration := 2.0 // duration for each color transition
	totalDuration := segmentDuration * float64(len(colors))

	// determine current position in the color cycle based on elapsed time
	tCycle := math.Mod(s.elapsed, totalDuration)  // time within current cycle
	segmentIndex := int(tCycle / segmentDuration) // current index
	nextIndex := (segmentIndex + 1) % len(colors) // index of next color

	// calculate interpolation factor
	t := (tCycle - float64(segmentIndex)*segmentDuration) / segmentDuration

	// interpolate between the current and next color
	return lerpColor(colors[segmentIndex], colors[nextIndex], t)
}
