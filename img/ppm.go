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

func WritePPM(w io.Writer, pixels [][][3]int) {
	height := len(pixels)
	width := len(pixels[0])
	fmt.Fprintln(w, "P3")
	fmt.Fprintf(w, "%d %d\n", width, height)
	fmt.Fprintln(w, "255")
	for i, row := range pixels {
		fmt.Fprintf(os.Stderr, "\rLines remaining: %d", height-1-i)
		for _, col := range row {
			fmt.Fprintf(w, "%d %d %d\n", col[0], col[1], col[2])
		}
	}
	fmt.Fprintf(os.Stderr, "\n Done")
}
