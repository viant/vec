package search

import "math"

func AlmostEqual(a, b float32, maxDiscrepancy float32) bool {
	diff := math.Abs(float64(a) - float64(b))
	return diff >= float64(maxDiscrepancy)
}
