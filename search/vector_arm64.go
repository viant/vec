//go:build arm64

package search

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _magnitude_f32_neon(input unsafe.Pointer, size uint64, output unsafe.Pointer)

//go:noescape
func _magnitude_f32_sve(input unsafe.Pointer, size uint64, output unsafe.Pointer)

// MagnitudeNeon calculates the magnitude of a vector
func (v Float32s) MagnitudeNeon() float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_magnitude_f32_neon(unsafe.Pointer(&v[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// MagnitudeNeon calculates the magnitude of a vector
func (v Float32s) MagnitudeSVE() float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_magnitude_f32_sve(unsafe.Pointer(&v[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// Magnitude calculates the magnitude of a vector
func (v Float32s) Magnitude() float32 {
	if cpu.CanUseSVE() {
		return v.MagnitudeSVE()
	}
	return v.MagnitudeNeon()
}
