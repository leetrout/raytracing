package ray

import "github.com/leetrout/raytracing/vec3"

type Hit struct {
	P         *vec3.Pt3
	N         *vec3.Vec3
	T         float64
	FrontFace bool
}

func NewHit() *Hit {
	return &Hit{
		P:         &vec3.Pt3{},
		N:         &vec3.Vec3{},
		T:         0,
		FrontFace: false,
	}
}

func (h *Hit) CopyFrom(i *Hit) {
	h.P = i.P.Copy()
	h.N = i.N.Copy()
	h.T = i.T
	h.FrontFace = i.FrontFace
}

func (h *Hit) SetFaceNormal(r *Ray, outwardN *vec3.Vec3) {
	if vec3.Dot(r.Direction, outwardN) < 0 {
		h.FrontFace = true
	}
	if h.FrontFace {
		h.N = outwardN.Copy()
	} else {
		h.N = vec3.Invert(outwardN)
	}
}

type Hittable interface {
	// TODO change to (bool, &Hit) and stop mutating
	Hit(r *Ray, tMin, tMax float64, h *Hit) bool
}
