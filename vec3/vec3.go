package vec3

import "math"

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func (v *Vec3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

type Color = Vec3
type Pt3 = Vec3
