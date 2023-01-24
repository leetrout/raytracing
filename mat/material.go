package mat

import (
	"math"
	"math/rand"

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

type Dielectric struct {
	// Index of refraction
	IR float64
}

func (m *Dielectric) Reflectance(cos, refIdx float64) float64 {
	// Schlick's approximation for reflectance
	// https://en.wikipedia.org/wiki/Schlick%27s_approximation
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cos, 5)
}

func (m *Dielectric) Scatter(r *ray.Ray, h *Hit) (doesScatter bool, scattered *ray.Ray, attenuation *vec3.Color) {
	attenuation = &vec3.Color{1, 1, 1}
	refractionRatio := m.IR
	if h.FrontFace {
		refractionRatio = 1.0 / m.IR
	}

	ud := vec3.UnitVector(r.Direction)

	cosTheta := math.Min(vec3.Dot(vec3.Invert(ud), h.N), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0
	maybeReflect := m.Reflectance(cosTheta, refractionRatio) > rand.Float64()

	var direction *vec3.Vec3
	if cannotRefract || maybeReflect {
		direction = vec3.Reflect(ud, h.N)
	} else {
		direction = vec3.Refract(ud, h.N, refractionRatio)
	}

	scattered = &ray.Ray{h.P, direction}

	return true, scattered, attenuation
}
