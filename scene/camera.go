package scene

import (
	"math"

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

func DegToRad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func NewCamera(lookFrom, lookAt *vec3.Pt3, up *vec3.Vec3, vfovDegrees, aspectRatio float64) *Camera {
	theta := DegToRad(vfovDegrees)
	height := math.Tan(theta / 2)
	f := 1.0
	vph := 2.0 * height
	vpw := vph * aspectRatio

	// Position camera
	// TODO is w is focal length?
	// Book says 'complete orthonormal basis (u,v,w)'
	w := vec3.UnitVector(vec3.Sub(lookFrom, lookAt))
	hor := vec3.UnitVector(vec3.Cross(up, w))
	vert := vec3.Cross(w, hor)

	origin := lookFrom.Copy()
	h := vec3.MultiplyFloat64(vpw, hor)
	v := vec3.MultiplyFloat64(vph, vert)

	// Lower left corner is translated from the origin - start at 0,0,0
	lowerLeftCorner := origin.Copy()
	// Move to the left by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(h, 2))
	// Move down by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(v, 2))
	// Move back by the focal length
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, w)

	return &Camera{
		AspectRatio:     aspectRatio,
		ViewportHeight:  vph,
		ViewportWidth:   vpw,
		FocalLength:     f,
		Origin:          origin,
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
