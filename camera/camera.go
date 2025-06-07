package camera

import "math"

type Camera struct {
	X, Y, Z float64
	Focal   float64
	Depth   float64
}

func NewCamera(x, y, z, focal, depth float64) *Camera {
	return &Camera{
		X:     x,
		Y:     y,
		Z:     z,
		Focal: focal,
		Depth: depth,
	}
}

// sets the position of the camera based on the position of the target and the size of the screen
func (c *Camera) FollowTarget(targetX, targetY, screenWidth, screenHeight float64) {
	c.X = -targetX + screenWidth/2.0
	c.Y = -targetY + screenHeight/2.0
}

// stops the camera from showing past the boundaries of the tilemap
func (c *Camera) Constrain(
	tilemapWidthPixels, tilemapHeightPixels, screenWidth, screenHeight float64,
) {
	c.X = math.Min(c.X, 0.0)
	c.Y = math.Min(c.Y, 0.0)

	c.X = math.Max(c.X, screenWidth-tilemapWidthPixels)
	c.Y = math.Max(c.Y, screenHeight-tilemapHeightPixels)
}
