package main

import (
	"math"
)

func countTriples(n int) int {
	triplesCount := 0

	for c := 5; c <= n; c++ {
		cSquared := c * c

		for a := 1; a < c; a++ {
			aSquared := a * a

			// se a + b = c
			// então b = c - a
			bSquared := cSquared - aSquared
			b := int(math.Sqrt(float64(bSquared)))

			if b*b == bSquared && b >= a && b < c {
				triplesCount++
			}
		}
	}
	return triplesCount * 2
}
