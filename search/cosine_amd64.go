package search

import (
	"unsafe"
)

//go:noescape
func _cosine_distance_f32_avx2(input1, input2 unsafe.Pointer, size uint64, output unsafe.Pointer)

//go:noescape
func _cosine_distance_with_magnitudes_f32_avx2(input1, input2, magnitudes unsafe.Pointer, size uint64, output unsafe.Pointer)
func (v Float32s) cosineDistanceAvx2(vec []float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_cosine_distance_f32_avx2(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// cosineDistanceWithMagnitudesNeon calculates the cosine distance between two vectors with magnitudes, magnitudes are precalculated and passed as arguments
func (v Float32s) cosineDistanceWithMagnitudesAvx2(vec []float32, magnitude1, magnitude2 float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var magnitudes = [2]float32{magnitude1, magnitude2}
	var output float32
	_cosine_distance_with_magnitudes_f32_avx2(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), unsafe.Pointer(&magnitudes[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// CosineDistance calculates the cosine distance between two vectors
func (v Float32s) CosineDistance(vec []float32) float32 {
	return v.cosineDistanceAvx2(vec)
}

// CosineDistanceWithMagnitude calculates the cosine distance between two vectors, magnitudes are precalculated and passed as arguments
func (v Float32s) CosineDistanceWithMagnitude(vec []float32, magnitude1, magnitude2 float32) float32 {
	return v.cosineDistanceWithMagnitudesAvx2(vec, magnitude1, magnitude2)
}
