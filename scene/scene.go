package scene

import (
	"github.com/leetrout/raytracing/mat"
	"github.com/leetrout/raytracing/ray"
)

type Scene struct {
	Objects []mat.Hittable
}

var _ mat.Hittable = &Scene{}

func (s *Scene) Hit(r *ray.Ray, tMin float64, tMax float64) *mat.Hit {
	var anyHit *mat.Hit
	closestSoFar := tMax

	for _, o := range s.Objects {
		if h := o.Hit(r, tMin, closestSoFar); h != nil {
			closestSoFar = h.T
			anyHit = mat.NewHit()
			anyHit.CopyFrom(h)
		}
	}

	return anyHit
}
