package vec3

import (
	"math"
	"math/rand"
)

func Add(u, v *Vec3) *Vec3 {
	return &Vec3{
		u.X + v.X,
		u.Y + v.Y,
		u.Z + v.Z,
	}
}

func Sum(vs ...*Vec3) *Vec3 {
	v := &Vec3{}
	for _, u := range vs {
		v = Add(v, u)
	}
	return v
}

func Sub(u, v *Vec3) *Vec3 {
	return &Vec3{
		u.X - v.X,
		u.Y - v.Y,
		u.Z - v.Z,
	}
}

func MultiplyVec3(u, v *Vec3) *Vec3 {
	return &Vec3{
		u.X * v.X,
		u.Y * v.Y,
		u.Z * v.Z,
	}
}

func MultiplyFloat64(t float64, v *Vec3) *Vec3 {
	return &Vec3{
		v.X * t,
		v.Y * t,
		v.Z * t,
	}
}

func Dot(u, v *Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func Cross(u, v *Vec3) *Vec3 {
	return &Vec3{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}
}

func DivideVec3(u, v *Vec3) *Vec3 {
	return &Vec3{
		u.X / v.X,
		u.Y / v.Y,
		u.Z / v.Z,
	}
}

func DivideFloat(v *Vec3, f float64) *Vec3 {
	return MultiplyFloat64(1/f, v)
}

func UnitVector(v *Vec3) *Vec3 {
	return DivideFloat(v, v.Length())
}

func Invert(v *Vec3) *Vec3 {
	return &Vec3{
		-v.X,
		-v.Y,
		-v.Z,
	}
}

func GetRandom(min, max int) float64 {
	return rand.Float64()*float64(max-min) + float64(min)
}

func GetRandomf(min, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

func Random(min, max int) *Vec3 {
	return &Vec3{
		GetRandom(min, max),
		GetRandom(min, max),
		GetRandom(min, max),
	}
}

func RandomInUnitSphere() *Vec3 {
	for {
		// Random point in unit cube
		p := Random(-1, 1)
		if p.LengthSquared() >= 1 {
			// p is outside unit sphere
			continue
		}
		return p
	}
}

func RandomInUnitDisk() *Vec3 {
	for {
		p := &Vec3{
			GetRandom(-1, 1),
			GetRandom(-1, 1),
			0,
		}
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func RandomUnitVector() *Vec3 {
	return UnitVector(RandomInUnitSphere())
}

func Reflect(v, n *Vec3) *Vec3 {
	return Sub(v, MultiplyFloat64(2*Dot(v, n), n))
}

func Refract(uv *Vec3, n *Vec3, etaiOverEtat float64) *Vec3 {
	cosTheta := math.Min(Dot(Invert(uv), n), 1.0)
	rPerpendicular := MultiplyFloat64(etaiOverEtat, Add(uv, MultiplyFloat64(cosTheta, n)))
	rParallel := MultiplyFloat64(math.Sqrt(math.Abs(1.0-rPerpendicular.LengthSquared()))*-1, n)
	return Add(rPerpendicular, rParallel)
}
