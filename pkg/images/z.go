package images

import (
	"image"
	"log"
)

const (
	Square = "square"
	Steps  = "steps"
)

type distanceFunction func(x, y, w, h int) float32

func MakeDistanceMap(bounds image.Rectangle, option string) [][]float32 {
	var fn distanceFunction

	switch option {
	case Square:
		fn = squareFn
		break
	case Steps:
		fn = stepsFn
		break
	default:
		log.Fatalf("Got unknown option %s", option)
	}

	return makeInner(bounds, fn)
}

func makeInner(bounds image.Rectangle, d distanceFunction) [][]float32 {
	var out [][]float32

	w := bounds.Dx()
	h := bounds.Dy()

	for x := 0; x < w; x++ {
		col := make([]float32, w)
		for y := 0; y < h; y++ {
			col[y] = d(x, y, w, h)
		}
		out = append(out, col)
	}
	return out
}

func stepsFn(x, y, w, h int) float32 {
	if !inbounds(y, h) {
		// Keep the top and bottom third as background
		return 0
	}

	return float32(x*5/w) / 5
}

func squareFn(x, y, w, h int) float32 {
	if inbounds(x, w) && inbounds(y, h) {
		return 0.9
	} else {
		return 0
	}
}

func inbounds(index, bound int) bool {
	indexF := float64(index)
	boundF := float64(bound)
	return indexF > boundF/3.0 && indexF < 2*boundF/3.0
}
