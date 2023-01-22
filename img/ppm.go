package img

import (
	"fmt"
	"io"
	"os"

	"github.com/leetrout/raytracing/vec3"
)

const RGBMax = 255.999

func ClampRGB(f float64) int {
	return int(f * RGBMax)
}

func Vec3AsRGB(v *vec3.Vec3) [3]int {
	return [3]int{
		ClampRGB(v.X),
		ClampRGB(v.Y),
		ClampRGB(v.Z),
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
