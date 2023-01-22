package img

import (
	"fmt"
	"io"

	"github.com/leetrout/raytracing/vec3"
)

const RGBMax = 256

func ClampRGB(f float64, samples int) int {
	scale := 1.0 / float64(samples)
	return int(RGBMax * (Clamp(f*scale, 0, 0.999)))
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func Vec3AsRGB(v *vec3.Vec3, samples int) [3]int {
	return [3]int{
		ClampRGB(v.X, samples),
		ClampRGB(v.Y, samples),
		ClampRGB(v.Z, samples),
	}
}

func WritePPM(w io.Writer, width, height int, pixels [][3]int) {
	fmt.Fprintln(w, "P3")
	fmt.Fprintf(w, "%d %d\n", width, height)
	fmt.Fprintln(w, "255")
	for _, px := range pixels {
		fmt.Fprintf(w, "%d %d %d\n", px[0], px[1], px[2])
	}
}
