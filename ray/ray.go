package ray

import "github.com/leetrout/raytracing/vec3"

type Ray struct {
	Origin    *vec3.Vec3
	Direction *vec3.Vec3
}

func (r *Ray) At(t float64) *vec3.Vec3 {
	return vec3.Add(r.Origin, vec3.MultiplyFloat64(t, r.Direction))
}

// class ray {
//     public:
//         ray() {}
//         ray(const point3& origin, const vec3& direction)
//             : orig(origin), dir(direction)
//         {}

//         point3 origin() const  { return orig; }
//         vec3 direction() const { return dir; }

//         point3 at(double t) const {
//             return orig + t*dir;
//         }

//     public:
//         point3 orig;
//         vec3 dir;
// };
