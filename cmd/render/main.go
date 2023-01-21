package main

import (
	"os"

	"github.com/leetrout/raytracing/img"
)

func main() {
	// Render
	img.WritePPM(os.Stdout, [][][3]int{{{255, 0, 0}, {0, 255, 0}, {0, 0, 255}}})
}
