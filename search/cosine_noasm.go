//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package search

// CosineDistance calculates the cosine distance between two vectors
func (v Float32s) CosineDistance(vec2 []float32) float32 {
	return v.cosineDistance(vec2)
}

// CosineDistanceWithMagnitude calculates the cosine distance between two vectors, magnitudes are precalculated and passed as arguments
func (v Float32s) CosineDistanceWithMagnitude(vec []float32, magnitude1, magnitude2 float32) float32 {
	return v.cosineDistanceWithMagnitude(vec, magnitude1, magnitude2)
}
