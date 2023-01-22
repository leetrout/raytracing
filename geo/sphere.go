package geo

import (
	"math"

	"github.com/leetrout/raytracing/ray"
	"github.com/leetrout/raytracing/vec3"
)

type Sphere struct {
	Center *vec3.Pt3
	Radius float64
}

var _ ray.Hittable = &Sphere{}

func (s *Sphere) Hit(r *ray.Ray, tMin float64, tMax float64, h *ray.Hit) bool {
	originToCenter := vec3.Sub(r.Origin, s.Center)
	a := r.Direction.LengthSquared()
	half_b := vec3.Dot(originToCenter, r.Direction)
	c := originToCenter.LengthSquared() - s.Radius*s.Radius

	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return false
	}

	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range.
	root := (-half_b - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-half_b + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}

	// TODO - I don't like mutating this here... author calls this out
	// as a design decision
	h.T = root
	h.P = r.At(h.T)
	outwardNormal := vec3.DivideFloat(vec3.Sub(h.P, s.Center), s.Radius)
	h.SetFaceNormal(r, outwardNormal)

	return true
}
