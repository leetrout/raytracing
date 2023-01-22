package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"

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

func render(fh io.Writer) {
	// Render
	samplesPerPixel := 100

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

	cam := scene.NewCamera()

	pixels := make([][3]int, imageHeight*imageWidth)

	for j := imageHeight - 1; j >= 0; j-- {
		fmt.Printf("\r %s Scanlines remaining: %d     ", time.Now().Format("02 Jan 06 15:04:05.9"), j)
		for i := 0; i < imageWidth; i++ {
			color := &vec3.Color{}

			for p := 0; p < samplesPerPixel; p++ {
				u := (float64(i) + rand.Float64()) / float64(imageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(imageHeight-1)

				r := cam.GetRay(u, v)
				color = vec3.Add(color, RayColor(r, s))
			}

			// We're walking up the image from bottom to top but we need to
			// write the pixels top to bottom so the current pixel is located
			// at the image height (e.g. 200) minus 1 to zero index (199)
			// and finally minus the current "row" (j) which starts at 199
			// assuming a 200px image
			pixelIdx := (imageWidth * (imageHeight - 1 - j)) + i
			pixels[pixelIdx] = img.Vec3AsRGB(color, samplesPerPixel)
		}
	}
	img.WritePPM(fh, imageWidth, imageHeight, pixels)
	fmt.Printf("\r\n")
}

func main() {
	name := "output.ppm"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fh, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	fmt.Printf("Starting render (%s)\n", name)
	render(fh)
	fmt.Println("Complete")
}
