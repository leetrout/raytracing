package generated

import (
	"math/rand"

	"github.com/leetrout/raytracing/geo"
	"github.com/leetrout/raytracing/mat"
	"github.com/leetrout/raytracing/scene"
	"github.com/leetrout/raytracing/vec3"
)

func GenerateScene() *scene.Scene {
	s := &scene.Scene{
		Camera: scene.NewCamera(
			&vec3.Vec3{13, 2, 3},
			&vec3.Vec3{0, 0, 0},
			&vec3.Vec3{0, 1, 0},
			20,
			3.0/2.0,
			0.1,
			10.0,
		),
	}

	// Ground
	s.Objects = append(s.Objects, &geo.Sphere{
		&vec3.Pt3{0, -1000, 0},
		1000,
		&mat.Lambert{&vec3.Color{0.5, 0.5, 0.5}},
	})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := &vec3.Pt3{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if vec3.Sub(center, &vec3.Pt3{4, 0.2, 0.0}).Length() > 0.9 {
				var sphereMat mat.Material
				albedo := vec3.MultiplyVec3(vec3.Random(0, 1), vec3.Random(0, 1))
				switch {
				case chooseMat < 0.8:
					// Lambert diffuse
					sphereMat = &mat.Lambert{albedo}
				case chooseMat < 0.95:
					// Metal
					sphereMat = &mat.Metal{albedo, vec3.GetRandomf(0, 0.5)}
				default:
					// Glass
					sphereMat = &mat.Dielectric{1.5}
				}
				s.Objects = append(s.Objects, &geo.Sphere{
					center,
					0.2,
					sphereMat,
				})
			}
		}
	}

	// 3 spheres in the middle
	s.Objects = append(s.Objects, &geo.Sphere{
		&vec3.Pt3{0, 1, 0},
		1.0,
		&mat.Dielectric{1.5},
	})

	s.Objects = append(s.Objects, &geo.Sphere{
		&vec3.Pt3{-4, 1, 0},
		1.0,
		&mat.Lambert{&vec3.Color{0.4, 0.2, 0.1}},
	})

	s.Objects = append(s.Objects, &geo.Sphere{
		&vec3.Pt3{4, 1, 0},
		1.0,
		&mat.Metal{&vec3.Color{0.7, 0.6, 0.5}, 0},
	})

	return s
}
