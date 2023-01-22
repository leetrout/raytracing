package scene

import (
	"github.com/leetrout/raytracing/ray"
)

type Scene struct {
	Objects []ray.Hittable
}

var _ ray.Hittable = &Scene{}

func (s *Scene) Hit(r *ray.Ray, tMin float64, tMax float64, h *ray.Hit) bool {
	tempRec := ray.NewHit()
	hitAny := false
	closestSoFar := tMax

	for _, o := range s.Objects {
		if o.Hit(r, tMin, closestSoFar, tempRec) {
			hitAny = true
			closestSoFar = tempRec.T
			// TODO I don't like this mutating
			h.CopyFrom(tempRec)
		}
	}

	return hitAny
}
