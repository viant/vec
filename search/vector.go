package search

import "math"

// Float32s is a type alias for a slice of float32
type Float32s []float32

func (v Float32s) magnitude() float32 {
	var magnitudeSq float32
	for _, val := range v {
		magnitudeSq += val * val
	}
	return float32(math.Sqrt(float64(magnitudeSq)))
}

func (v Float32s) cosineDistance(vec2 []float32) float32 {
	vec1 := v
	if len(vec1) != len(vec2) {
		return 0.0
	}
	var dotProduct, magnitude1Sq, magnitude2Sq float32
	for i := 0; i < len(vec1); i++ {
		dotProduct += vec1[i] * vec2[i]
		magnitude1Sq += vec1[i] * vec1[i]
		magnitude2Sq += vec2[i] * vec2[i]
	}
	magnitude1 := math.Sqrt(float64(magnitude1Sq))
	magnitude2 := math.Sqrt(float64(magnitude2Sq))
	if magnitude1 == 0 || magnitude2 == 0 {
		return 1.0
	}
	return 1 - float32(float64(dotProduct)/(float64(magnitude1)*float64(magnitude2)))
}

func (v Float32s) cosineDistanceWithMagnitude(vec2 []float32, magnitude1, magnitude2 float32) float32 {
	vec1 := v
	if magnitude1 == 0 || magnitude2 == 0 {
		return 1.0
	}
	if len(vec1) != len(vec2) {
		return 0.0
	}
	var dotProduct float32
	for i := 0; i < len(vec1); i++ {
		dotProduct += vec1[i] * vec2[i]
	}
	return 1 - float32(float64(dotProduct)/(float64(magnitude1)*float64(magnitude2)))
}
