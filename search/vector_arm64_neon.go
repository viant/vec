//go:build arm64

package search

import "unsafe"

//go:noescape
func _magnitude_f32_neon(input unsafe.Pointer, size uint64, output unsafe.Pointer)

// MagnitudeNeon calculates the magnitude of a vector
func (v Float32s) MagnitudeNeon() float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_magnitude_f32_neon(unsafe.Pointer(&v[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}
