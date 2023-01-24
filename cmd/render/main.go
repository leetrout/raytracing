package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/leetrout/raytracing/geo"
	"github.com/leetrout/raytracing/img"
	"github.com/leetrout/raytracing/mat"
	"github.com/leetrout/raytracing/ray"
	"github.com/leetrout/raytracing/scene"
	"github.com/leetrout/raytracing/vec3"
)

func RayColor(r *ray.Ray, s *scene.Scene, depth int) *vec3.Vec3 {
	if depth <= 0 {
		// Reached the limit, return black
		return &vec3.Color{}
	}
	if h := s.Hit(r, 0.001, math.MaxFloat64); h != nil {
		// Render as normals
		// N := vec3.Add(h.N, &vec3.Color{1, 1, 1})
		// return vec3.MultiplyFloat64(0.5, N)

		// Hardcoded lambert
		// randomBounce := vec3.RandomUnitVector()
		// target := vec3.Sum(h.P, h.N, randomBounce)
		// newRay := &ray.Ray{
		// 	Origin:    h.P,
		// 	Direction: vec3.Sub(target, h.P),
		// }
		// return vec3.MultiplyFloat64(0.5, RayColor(newRay, s, depth-1))

		doesScatter, scattered, attenuation := h.Mat.Scatter(r, h)
		if doesScatter {
			return vec3.MultiplyVec3(attenuation, RayColor(scattered, s, depth-1))
		}

		return &vec3.Color{}
	}

	// No hit, draw the background as a sky
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

type Pixel struct {
	id int
	c  *vec3.Color
}

func RenderPixel(out chan *Pixel, s *scene.Scene, id, samplesPerPixel, maxDepth, width, height, row, col int) {
	color := &vec3.Color{}
	for p := 0; p < samplesPerPixel; p++ {
		u := (float64(row) + rand.Float64()) / float64(width-1)
		v := (float64(col) + rand.Float64()) / float64(height-1)

		r := s.Camera.GetRay(u, v)
		color = vec3.Add(color, RayColor(r, s, maxDepth))
	}
	out <- &Pixel{id, color}
}

func render(fh io.Writer) {
	// Render
	samplesPerPixel := 100
	maxDepth := 25

	// Image
	aspectRatio := 16.0 / 9.0
	imageWidth := 400
	imageHeight := int(float64(imageWidth) / aspectRatio)

	// Scene
	s := &scene.Scene{
		Camera: scene.NewCamera(),
	}

	// Materials
	mGround := &mat.Lambert{&vec3.Color{0.8, 0.8, 0.0}}
	mCenter := &mat.Lambert{&vec3.Color{0.1, 0.2, 0.5}}
	mLeft := &mat.Dielectric{1.5}
	mRight := &mat.Metal{&vec3.Color{0.8, 0.6, 0.2}, 0.0}

	// Objects
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{0, -100.5, -1}, 100, mGround})
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{0, 0, -1}, 0.5, mCenter})
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{-1, 0, -1}, 0.5, mLeft})
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{-1, 0, -1}, -0.4, mLeft})
	s.Objects = append(s.Objects, &geo.Sphere{&vec3.Pt3{1, 0, -1}, 0.5, mRight})

	pixels := make([][3]int, imageHeight*imageWidth)
	pxLock := &sync.Mutex{}
	pxOut := make(chan *Pixel)

	totalWork := imageHeight * imageWidth

	wg := &sync.WaitGroup{}
	wg.Add(totalWork)

	go (func(wg *sync.WaitGroup, out chan *Pixel, lock *sync.Mutex) {
		fmt.Println("Pixel worker available")
		for i := 0; i < totalWork; i++ {
			if i > 0 && i%1000 == 0 {
				progress := int(math.Round(float64(i) / float64(totalWork) * 100))
				fmt.Printf("\rProgress (%d%%)", progress)
			}
			px := <-out
			lock.Lock()
			pixels[px.id] = img.Vec3AsRGB(px.c, samplesPerPixel)
			wg.Done()
			lock.Unlock()
		}
		fmt.Printf("\r\n")
	})(wg, pxOut, pxLock)

	fmt.Println("Enqueueing render jobs")
	for j := imageHeight - 1; j >= 0; j-- {
		// fmt.Printf("\r %s Scanlines remaining: %d     ", time.Now().Format("02 Jan 06 15:04:05.9"), j)
		for i := 0; i < imageWidth; i++ {
			// We're walking up the image from bottom to top but we need to
			// write the pixels top to bottom so the current pixel is located
			// at the image height (e.g. 200) minus 1 to zero index (199)
			// and finally minus the current "row" (j) which starts at 199
			// assuming a 200px image
			pixelIdx := (imageWidth * (imageHeight - 1 - j)) + i

			go RenderPixel(pxOut, s, pixelIdx, samplesPerPixel, maxDepth, imageWidth, imageHeight, i, j)
		}
	}

	wg.Wait()
	fmt.Println("Writing output")
	img.WritePPM(fh, imageWidth, imageHeight, pixels)
}

func main() {
	rand.Seed(time.Now().UnixMicro())
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
