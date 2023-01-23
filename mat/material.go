package mat

import (
	"github.com/leetrout/raytracing/ray"
	"github.com/leetrout/raytracing/vec3"
)

type Material interface {
	Scatter(
		r *ray.Ray,
		h *Hit,
	) (
		doesScatter bool,
		scattered *ray.Ray,
		attenuation *vec3.Color,
	)
}

var _ Material = &Lambert{}
var _ Material = &Metal{}

type Lambert struct {
	Albedo *vec3.Color
}

func (m *Lambert) Scatter(r *ray.Ray, h *Hit) (doesScatter bool, scattered *ray.Ray, attenuation *vec3.Color) {
	scatterDirection := vec3.Add(h.N, vec3.RandomUnitVector())
	if scatterDirection.NearZero() {
		scatterDirection = h.N
	}
	scattered = &ray.Ray{h.P, scatterDirection}
	attenuation = m.Albedo
	return true, scattered, attenuation
}

type Metal struct {
	Albedo *vec3.Color
	Fuzz   float64
}

func (m *Metal) Scatter(r *ray.Ray, h *Hit) (doesScatter bool, scattered *ray.Ray, attenuation *vec3.Color) {
	reflected := vec3.Reflect(vec3.UnitVector(r.Direction), h.N)
	fuzzy := vec3.MultiplyFloat64(m.Fuzz, vec3.RandomInUnitSphere())
	scattered = &ray.Ray{h.P, vec3.Add(reflected, fuzzy)}
	attenuation = m.Albedo
	doesScatter = vec3.Dot(scattered.Direction, h.N) > 0
	return doesScatter, scattered, attenuation
}
