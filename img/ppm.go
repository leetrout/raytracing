package img

import (
	"fmt"
	"io"
)

func WritePPM(w io.Writer, pixels [][][3]int) {
	height := len(pixels)
	width := len(pixels[0])
	fmt.Fprintln(w, "P3")
	fmt.Fprintf(w, "%d %d\n", width, height)
	fmt.Fprintln(w, "255")
	for _, row := range pixels {
		for _, col := range row {
			fmt.Fprintf(w, "%d %d %d\n", col[0], col[1], col[2])
		}
	}
}
