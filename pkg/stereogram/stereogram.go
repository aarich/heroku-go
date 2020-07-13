package stereogram

import (
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
)

const (
	DEPTH_OF_FIELD = 0.3 // fraction space between back pane and screen
	EYE_SEPARATION = 200 // pixels
)

func Generate(z [][]float32) *image.RGBA {
	height := len(z)
	width := len(z[0])

	out := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		constraintsFound := ""
		constraints := makeConstraints(y, z)

		for x := 0; x < width; x++ {
			var val color.Color
			if constraints[x] == x {
				// There is no constraint. Use a random number

				val = randomDot()
			} else {
				// We will have already written the previous value, so use that
				val = out.At(constraints[x], y)
				constraintsFound += "."
			}
			out.Set(x, y, val)
		}

		// log.Printf("constraints: %s", constraintsFound)
	}

	return out
}

func makeConstraints(y int, zImage [][]float32) []int {
	width := len(zImage[0])
	constraints := make([]int, width)

	// Initially, there are no constraints
	for x := 0; x < width; x++ {
		constraints[x] = x
	}

	for x := 0; x < width; x++ {
		z := zImage[y][x]
		sep := separation(z)

		// add constraints
		p2 := x + sep/2
		p1 := p2 - sep

		// If both points are in the image
		if p1 >= 0 && p2 < width {
			// Constraint of right is the left, so we can look back traversing left to right
			// we want to know about a constraint the second time we hit it
			constraints[p2] = p1
		}
	}

	return constraints
}

// Todo use float to potentially set neighboring pixels
func separation(z float32) int {
	sep := ((1.0 - DEPTH_OF_FIELD*z) / (2.0 - DEPTH_OF_FIELD*z)) * EYE_SEPARATION
	return int(math.Round(float64(sep)))
}

func randomDot() color.Color {
	// For now just return white/black
	value := rand.Intn(2)

	alpha := uint8(255)
	switch value {
	case 0:
		return color.RGBA{0, 0, 0, alpha}
	case 1:
		return color.RGBA{255, 255, 255, alpha}
	case 2:
		return color.RGBA{122, 122, 122, alpha}
	case 3:
		return color.RGBA{0, 255, 255, alpha}
	case 4:
		return color.RGBA{255, 255, 0, alpha}
	case 5:
		return color.RGBA{255, 0, 255, alpha}
	}
	log.Fatalf("Got unexpected random number %d", value)
	return nil
}
