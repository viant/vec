package search

import "math"

// EuclideanDistance calculates the Euclidean distance between two vectors
func (v Float32s) EuclideanDistance(vec Float32s) float32 {
	sum := float32(0.0)
	for i := 0; i < len(vec); i++ {
		diff := v[i] - vec[i]
		sum += diff * diff
	}
	return float32(math.Sqrt(float64(sum)))
}
