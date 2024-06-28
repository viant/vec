//go:build arm64

package search

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _cosine_distance_f32_neon(input1, input2 unsafe.Pointer, size uint64, output unsafe.Pointer)

//go:noescape
func _cosine_distance_with_magnitude_f32_neon(input1, input2, magnitudes unsafe.Pointer, size uint64, output unsafe.Pointer)

//go:noescape
func _cosine_distance_f32_sve(input1, input2 unsafe.Pointer, size uint64, output unsafe.Pointer)

//go:noescape
func _cosine_distance_with_magnitude_f32_sve(input1, input2, magnitudes unsafe.Pointer, size uint64, output unsafe.Pointer)

// cosineDistanceNeon calculates the cosine distance between two vectors
func (v Float32s) cosineDistanceNeon(vec []float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_cosine_distance_f32_neon(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// cosineDistanceWithMagnitudesNeon calculates the cosine distance between two vectors with magnitudes, magnitudes are precalculated and passed as arguments
func (v Float32s) cosineDistanceWithMagnitudesNeon(vec []float32, magnitude1, magnitude2 float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var magnitudes = [2]float32{magnitude1, magnitude2}
	var output float32
	_cosine_distance_with_magnitude_f32_neon(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), unsafe.Pointer(&magnitudes[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// cosineDistanceSVE calculates the cosine distance between two vectors
func (v Float32s) cosineDistanceSVE(vec []float32) float32 {
	if len(v) != len(v) {
		return 0.0
	}
	var output float32
	_cosine_distance_f32_sve(unsafe.Pointer(&v[0]), unsafe.Pointer(&vec[0]), uint64(len(v)), unsafe.Pointer(&output))
	return output
}

// cosineDistanceWithMagnitudesSVE calculates the cosine distance between two vectors with magnitudes, magnitudes are precalculated and passed as arguments
func (v Float32s) cosineDistanceWithMagnitudesSVE(vec []float32, magnitude1, magnitude2 float32) float32 {
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
	if cpu.CanUseSVE() {
		return v.cosineDistanceSVE(vec)
	}
	return v.cosineDistanceNeon(vec)
}

// CosineDistanceWithMagnitude calculates the cosine distance between two vectors, magnitudes are precalculated and passed as arguments
func (v Float32s) CosineDistanceWithMagnitude(vec []float32, magnitude1, magnitude2 float32) float32 {
	if cpu.CanUseSVE() {
		return v.cosineDistanceWithMagnitudesSVE(vec, magnitude1, magnitude2)
	}
	return v.cosineDistanceWithMagnitudesNeon(vec, magnitude1, magnitude2)
}
