package vec3

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3) Copy() *Vec3 {
	return &Vec3{v.X, v.Y, v.Z}
}

func (v *Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v *Vec3) NearZero() bool {
	s := 1e-8
	return math.Abs(v.X) < s &&
		math.Abs(v.Y) < s &&
		math.Abs(v.Z) < s

}

type Color = Vec3
type Pt3 = Vec3
