package scene

import (
	"github.com/leetrout/raytracing/ray"
)

type Scene struct {
	Objects []ray.Hittable
}

var _ ray.Hittable = &Scene{}

func (s *Scene) Hit(r *ray.Ray, tMin float64, tMax float64) *ray.Hit {
	var anyHit *ray.Hit
	closestSoFar := tMax

	for _, o := range s.Objects {
		if h := o.Hit(r, tMin, closestSoFar); h != nil {
			closestSoFar = h.T
			anyHit = ray.NewHit()
			anyHit.CopyFrom(h)
		}
	}

	return anyHit
}
