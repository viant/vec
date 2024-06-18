//go:build !arm64

package search

// Magnitude calculates the magnitude of a vector
func (v Float32s) Magnitude() float32 {
	return v.magnitude()
}
