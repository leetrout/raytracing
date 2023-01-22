package vec3

func Add(u, v *Vec3) *Vec3 {
	return &Vec3{
		u.X + v.X,
		u.Y + v.Y,
		u.Z + v.Z,
	}
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

func Dot(u, v Vec3) float64 {
	return u.X*v.X + u.Y*v.Y + u.Z*v.Z
}

func Cross(u, v Vec3) *Vec3 {
	return &Vec3{
		u.Y*v.Z - u.Z*v.Y,
		u.Z*v.X - u.X*v.Z,
		u.X*v.Y - u.Y*v.X,
	}
}

func DivideVec3(u, v Vec3) *Vec3 {
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
