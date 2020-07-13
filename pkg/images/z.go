package images

import (
	"image"
	"log"
)

const (
	Square = 1
)

func MakeDistanceMap(bounds image.Rectangle, option int) [][]float32 {
	switch option {
	case Square:
		return makeSquare(bounds)
	}
	log.Fatalf("Got unknown option %d", option)
	return nil
}

// Returns a square in the inner third of the image
func makeSquare(bounds image.Rectangle) [][]float32 {
	var out [][]float32

	w := bounds.Dx()
	h := bounds.Dy()

	for x := 0; x < w; x++ {
		col := make([]float32, w)
		for y := 0; y < h; y++ {
			if inbounds(x, w) && inbounds(y, h) {
				col[y] = 0.9
				log.Println("INBOUND")
			} else {
				col[y] = 0
			}
		}
		out = append(out, col)
	}

	var sum float32
	sum = 0.0
	for _, row := range out {
		for _, val := range row {
			sum += val
		}
	}
	log.Printf("SUM: %f", sum)

	return out
}

func inbounds(index, bound int) bool {
	indexF := float64(index)
	boundF := float64(bound)
	return indexF > boundF/3.0 && indexF < 2*boundF/3.0
}
