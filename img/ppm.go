package img

import (
	"fmt"
	"io"
	"os"

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
	for i, px := range pixels {
		// The number of lines is the height of the image but we
		// have to divide by the width to determine where we are in
		// the array
		if i%width == 0 {
			fmt.Fprintf(os.Stderr, "\rLines remaining: %d", height-1-i/width)
		}
		fmt.Fprintf(w, "%d %d %d\n", px[0], px[1], px[2])
	}
	fmt.Fprint(os.Stderr, "\n Done")
}
