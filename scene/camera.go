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
	u               *vec3.Vec3
	v               *vec3.Vec3
	lensRadius      float64
}

func DegToRad(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func NewCamera(lookFrom, lookAt *vec3.Pt3, up *vec3.Vec3, vfovDegrees, aspectRatio, aperature, focusDist float64) *Camera {
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
	h := vec3.MultiplyFloat64(focusDist*vpw, hor)
	v := vec3.MultiplyFloat64(focusDist*vph, vert)

	// Lower left corner is translated from the origin - start at 0,0,0
	lowerLeftCorner := origin.Copy()
	// Move to the left by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(h, 2))
	// Move down by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(v, 2))
	// Move back by the focal length
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.MultiplyFloat64(focusDist, w))

	lensRadius := aperature / 2

	return &Camera{
		AspectRatio:     aspectRatio,
		ViewportHeight:  vph,
		ViewportWidth:   vpw,
		FocalLength:     f,
		Origin:          origin,
		horizontal:      h,
		vertical:        v,
		lowerLeftCorner: lowerLeftCorner,
		u:               hor,
		v:               vert,
		lensRadius:      lensRadius,
	}
}

func (c *Camera) GetRay(u, v float64) *ray.Ray {
	rd := vec3.MultiplyFloat64(c.lensRadius, vec3.RandomInUnitDisk())
	offset := vec3.Add(vec3.MultiplyFloat64(rd.X, c.u), vec3.MultiplyFloat64(rd.Y, c.v))

	// Initialize ray from origin to corner
	r := &ray.Ray{vec3.Add(c.Origin, offset), c.lowerLeftCorner.Copy()}
	// Move ray based on current pixel
	// First horizontally
	r.Direction = vec3.Add(r.Direction, vec3.MultiplyFloat64(u, c.horizontal))
	// Second vertically
	r.Direction = vec3.Add(r.Direction, vec3.MultiplyFloat64(v, c.vertical))
	// Subtract origin (remember negative Z is "forward")
	r.Direction = vec3.Sub(r.Direction, c.Origin)
	r.Direction = vec3.Sub(r.Direction, offset)
	return r
}
