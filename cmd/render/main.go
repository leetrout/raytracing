package main

import (
	"math"
	"os"

	"github.com/leetrout/raytracing/geo"
	"github.com/leetrout/raytracing/img"
	"github.com/leetrout/raytracing/ray"
	"github.com/leetrout/raytracing/scene"
	"github.com/leetrout/raytracing/vec3"
)

func RayColor(r *ray.Ray, s *scene.Scene) *vec3.Vec3 {
	h := ray.NewHit()
	if s.Hit(r, 0, math.MaxFloat64, h) {
		N := vec3.Add(h.N, &vec3.Color{1, 1, 1})
		return vec3.MultiplyFloat64(0.5, N)
	}
	ud := vec3.UnitVector(r.Direction)
	t := 0.5 * (ud.Y + 1.0)

	// lerp white to black
	white := &vec3.Color{1.0, 1.0, 1.0}
	white = vec3.MultiplyFloat64(1.0-t, white)

	// lerp black to blue
	blue := &vec3.Color{0.5, 0.7, 1.0}
	blue = vec3.MultiplyFloat64(t, blue)

	// Add them to make a gradient
	return vec3.Add(white, blue)
}

func main() {
	// Render

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Scene
	s := &scene.Scene{}
	// Sphere
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{0, 0, -1}, 0.5})
	// Floor
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{0, -100.5, -1}, 100})

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

	pixels := make([][3]int, imageHeight*imageWidth)

	for j := imageHeight - 1; j >= 0; j-- {
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

			// We're walking up the image from bottom to top but we need to
			// write the pixels top to bottom so the current pixel is located
			// at the image height (e.g. 200) minus 1 to zero index (199)
			// and finally minus the current "row" (j) which starts at 199
			// assuming a 200px image
			pixelIdx := (imageWidth * (imageHeight - 1 - j)) + i
			pixels[pixelIdx] = img.Vec3AsRGB(RayColor(r, s))
		}
	}

	img.WritePPM(os.Stdout, imageWidth, imageHeight, pixels)
}
