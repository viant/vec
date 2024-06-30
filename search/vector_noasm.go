//go:build (!amd64 && !arm64) || noasm
// +build !amd64,!arm64 noasm

package search

// Magnitude calculates the magnitude of a vector
func (v Float32s) Magnitude() float32 {
	return v.magnitude()
}
