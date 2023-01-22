package scene

import (
	"github.com/leetrout/raytracing/ray"
	"github.com/leetrout/raytracing/vec3"
)

type Camera struct {
	AspectRatio    float64
	ViewportHeight float64
	ViewportWidth  float64
	FocalLength    float64

	Origin          *vec3.Pt3
	horizontal      *vec3.Vec3
	vertical        *vec3.Vec3
	lowerLeftCorner *vec3.Vec3
}

func NewCamera() *Camera {
	f := 1.0
	ar := 16.0 / 9.0
	vph := 2.0
	vpw := vph * ar

	h := &vec3.Vec3{vpw, 0, 0}
	v := &vec3.Vec3{0, vph, 0}

	// Lower left corner is translated from the origin - start at 0,0,0
	lowerLeftCorner := &vec3.Vec3{}
	// Move to the left by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(h, 2))
	// Move down by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(v, 2))
	// Move back by the focal length
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, &vec3.Vec3{0, 0, f})

	return &Camera{
		AspectRatio:     ar,
		ViewportHeight:  vph,
		ViewportWidth:   vpw,
		FocalLength:     f,
		Origin:          &vec3.Pt3{},
		horizontal:      h,
		vertical:        v,
		lowerLeftCorner: lowerLeftCorner,
	}
}

func (c *Camera) GetRay(u, v float64) *ray.Ray {
	// Initialize ray from origin to corner
	r := &ray.Ray{c.Origin.Copy(), c.lowerLeftCorner.Copy()}
	// Move ray based on current pixel
	// First horizontally
	r.Direction = vec3.Add(r.Direction, vec3.MultiplyFloat64(u, c.horizontal))
	// Second vertically
	r.Direction = vec3.Add(r.Direction, vec3.MultiplyFloat64(v, c.vertical))
	// Subtract origin (remember negative Z is "forward")
	r.Direction = vec3.Sub(r.Direction, c.Origin)
	return r
}
