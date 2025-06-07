package pseudo

import (
	"image/color"
)

// lerp performs linear interpolation between two values
// a. starting value
// b. ending value
// t. interpolation factor (0.0 to 1.0)
func Lerp(a, b uint8, t float64) uint8 {
	return uint8(float64(a) + (float64(b)-float64(a))*t)
}

func LerpF(a, b, t float64) float64 {
	return a + (b-a)*t
}

// lerpColor performs linear interpolation between two colors
// c1. starting color
// c2. ending color
// t. interpolation factor (0.0 to 1.0)
func LerpColor(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		R: Lerp(c1.R, c2.R, t),
		G: Lerp(c1.G, c2.G, t),
		B: Lerp(c1.B, c2.B, t),
		A: Lerp(c1.A, c2.A, t),
	}
}

// ProjectPoint projects a 3D in-world point (X,Y,Z) into 2D screen space
// farthestZ is used to compute a minimal scale (such as pinning the farthest point at the horizon)
func ProjectPoint(
	worldX, worldY, worldZ float64,
	cameraX, cameraY, cameraZ float64,
	cameraDepth, cameraFocal float64,
	horizonY, farthestZ float64,
	screenWidth int,
) (screenX, screenY, scale float64) {
	// 1. compute the effective distance from the camera along Z
	// assume the player is at s.Player.Camera.Z along the track
	effectiveZ := EffectiveZ(worldZ, cameraZ)

	// 2. scale factor - nearer objects (small effectiveZ) get a larger scale
	scale = Scale(effectiveZ, cameraDepth)

	// 3. compute the minimum scale for the farthest visible Z
	minScale := MinScale(farthestZ, cameraDepth)

	// 4. compute baseY: the screen Y coordinate for a worldY == 0 point
	// if the camera height changes, everything shifts up/down proportionally
	baseY := BaseY(horizonY, cameraFocal, scale, minScale)

	// 5. adjust for the object's actual worldY (vertical offset above the ground)
	// if worldY == 0, the object is on the road, then it is exactly at baseY
	// if an object was above the road, worldY*scale will push it upward on the screen
	screenY = baseY - (worldY-cameraY)*scale

	// 6. compute screenX: center of the screen + (lateral worldX offset)*scale
	// cameraX is how far we panned left or right
	// worldX of 0 is the center of road, or screenWidth /2
	screenX = float64(screenWidth)/2 + (worldX-cameraX)*scale

	return screenX, screenY, scale
}

// EffectiveZ returns the distance along Z from camera to world point
func EffectiveZ(worldZ, camZ float64) float64 {
	ez := worldZ - camZ
	if ez < 0.1 {
		ez = 0.1
	}
	return ez
}

// Scale returns the scale factor given effectiveZ and camera depth
func Scale(effectiveZ, cameraDepth float64) float64 {
	return cameraDepth / effectiveZ
}

// MinScale returns the minimum scale for the farthest visible Z
func MinScale(farthestZ, cameraDepth float64) float64 {
	return cameraDepth / farthestZ
}

// BaseY computes the base screen Y for a ground-level point
func BaseY(horizonY, cameraHeight, scale, minScale float64) float64 {
	return horizonY + cameraHeight*(scale-minScale)
}
