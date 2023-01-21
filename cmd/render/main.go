package main

import (
	"os"

	"github.com/leetrout/raytracing/img"
)

func RayColor() {
	// TODO https://raytracing.github.io/books/RayTracingInOneWeekend.html#rays,asimplecamera,andbackground/sendingraysintothescene
	// color ray_color(const ray& r) {
	// 	vec3 unit_direction = unit_vector(r.direction());
	// 	auto t = 0.5*(unit_direction.y() + 1.0);
	// 	return (1.0-t)*color(1.0, 1.0, 1.0) + t*color(0.5, 0.7, 1.0);
	// }
}

func main() {
	// Render
	imageWidth := 256
	imageHeight := 256

	pixels := [][][3]int{}

	for j := imageHeight - 1; j >= 0; j-- {
		row := [][3]int{}
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth-1)
			g := float64(j) / float64(imageHeight-1)
			b := 0.25
			row = append(row, [3]int{img.ClampRGB(r), img.ClampRGB(g), img.ClampRGB(b)})
		}
		pixels = append(pixels, row)
	}

	img.WritePPM(os.Stdout, pixels)
}
