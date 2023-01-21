package main

import (
	"os"

	"github.com/leetrout/raytracing/img"
)

const RGBMax = 255.999

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
			row = append(row, [3]int{int(r * RGBMax), int(g * RGBMax), int(b * RGBMax)})
		}
		pixels = append(pixels, row)
	}

	img.WritePPM(os.Stdout, pixels)
}
