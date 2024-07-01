package search

import (
	"unsafe"
)

//go:noescape
func _magnitude_f32_avx2(input unsafe.Pointer, size uint64, output unsafe.Pointer)

func (v Float32s) MagnitudeAvx2() float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_magnitude_f32_avx2(unsafe.Pointer(&v[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// Magnitude calculates the magnitude of a vector
func (v Float32s) Magnitude() float32 {
	return v.MagnitudeAvx2()
}
