package main

import (
	"math"
	"os"

	"github.com/leetrout/raytracing/img"
	"github.com/leetrout/raytracing/ray"
	"github.com/leetrout/raytracing/vec3"
)

func HitSphere(center *vec3.Vec3, radius float64, r *ray.Ray) float64 {
	originToCenter := vec3.Sub(r.Origin, center)
	a := vec3.Dot(r.Direction, r.Direction)
	b := vec3.Dot(originToCenter, r.Direction) * 2.0
	c := vec3.Dot(originToCenter, originToCenter) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - math.Sqrt(discriminant)) / (2.0 * a)
	}
}

func RayColor(r *ray.Ray) *vec3.Vec3 {
	sphereOrigin := &vec3.Vec3{0, 0, -1}
	t := HitSphere(sphereOrigin, 0.5, r)
	if t > 0.0 {
		N := vec3.UnitVector(vec3.Sub(r.At(t), sphereOrigin))
		return vec3.MultiplyFloat64(0.5, &vec3.Vec3{N.X + 1, N.Y + 1, N.Z + 1})
	}
	ud := vec3.UnitVector(r.Direction)
	t = 0.5 * (ud.Y + 1.0)
	white := &vec3.Color{1.0, 1.0, 1.0}
	white = vec3.MultiplyFloat64(1.0-t, white)
	blue := &vec3.Color{0.5, 0.7, 1.0}
	blue = vec3.MultiplyFloat64(t, blue)
	return vec3.Add(white, blue)
}

func main() {
	// Render

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Camera
	viewportHeight := 2.0
	viewportWidth := aspectRatio * viewportHeight
	focalLength := 1.0

	origin := &vec3.Pt3{}
	horizontal := &vec3.Vec3{viewportWidth, 0, 0}
	vertical := &vec3.Vec3{0, viewportHeight, 0}

	// Lower left corner is translated from the origin - start at 0,0,0
	lowerLeftCorner := &vec3.Vec3{}
	// Move to the left by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(horizontal, 2))
	// Move down by half the width
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, vec3.DivideFloat(vertical, 2))
	// Move back by the focal length
	lowerLeftCorner = vec3.Sub(lowerLeftCorner, &vec3.Vec3{0, 0, focalLength})

	pixels := [][][3]int{}

	for j := imageHeight - 1; j >= 0; j-- {
		row := [][3]int{}
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth-1)
			v := float64(j) / float64(imageHeight-1)

			// Initialize ray from origin to corner
			r := &ray.Ray{origin.Copy(), lowerLeftCorner.Copy()}
			// Move ray based on current pixel
			// First horizontally
			r.Direction = vec3.Add(r.Direction, vec3.MultiplyFloat64(u, horizontal))
			// Second vertically
			r.Direction = vec3.Add(r.Direction, vec3.MultiplyFloat64(v, vertical))
			// Subtract origin (remember negative Z is "forward")
			r.Direction = vec3.Sub(r.Direction, origin)

			row = append(row, img.Vec3AsRGB(RayColor(r)))
		}
		pixels = append(pixels, row)
	}

	img.WritePPM(os.Stdout, pixels)
}
