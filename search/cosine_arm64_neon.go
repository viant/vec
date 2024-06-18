//go:build arm64

package search

import (
	"unsafe"
)

//go:noescape
func _cosine_distance_f32_neon(input1, input2 unsafe.Pointer, size uint64, output unsafe.Pointer)

//go:noescape
func _cosine_distance_with_magnitude_f32_neon(input1, input2, magnitudes unsafe.Pointer, size uint64, output unsafe.Pointer)

// CosineDistanceNeon calculates the cosine distance between two vectors
func (v Float32s) CosineDistanceNeon(vec []float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_cosine_distance_f32_neon(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// CosineDistanceWithMagnitudesNeon calculates the cosine distance between two vectors with magnitudes, magnitudes are precalculated and passed as arguments
func (v Float32s) CosineDistanceWithMagnitudesNeon(vec []float32, magnitude1, magnitude2 float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var magnitudes = [2]float32{magnitude1, magnitude2}
	var output float32
	_cosine_distance_with_magnitude_f32_neon(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), unsafe.Pointer(&magnitudes[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// CosineDistance calculates the cosine distance between two vectors
func (v Float32s) CosineDistance(vec []float32) float32 {
	return v.CosineDistanceNeon(vec)
}

// CosineDistanceWithMagnitude calculates the cosine distance between two vectors, magnitudes are precalculated and passed as arguments
func (v Float32s) CosineDistanceWithMagnitude(vec []float32, magnitude1, magnitude2 float32) float32 {
	return v.CosineDistanceWithMagnitudesNeon(vec, magnitude1, magnitude2)
}
