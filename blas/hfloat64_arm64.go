package blas

import (
	"unsafe"
)

//go:noescape

func _hsum_float64(input, result unsafe.Pointer, size uint64)
func HsumFloat64(input []float64) float64 {
	var sum float64
	_hsum_float64(unsafe.Pointer(&input[0]), unsafe.Pointer(&sum), uint64(len(input)))
	return sum
}

func _hmax_float64(input, result unsafe.Pointer, size uint64)
func HmaxFloat64(input []float64) float64 {
	var max float64
	_hmax_float64(unsafe.Pointer(&input[0]), unsafe.Pointer(&max), uint64(len(input)))
	return max
}

func _hmin_float64(input, result unsafe.Pointer, size uint64)
func HminFloat64(input []float64) float64 {
	var min float64
	_hmin_float64(unsafe.Pointer(&input[0]), unsafe.Pointer(&min), uint64(len(input)))
	return min
}
