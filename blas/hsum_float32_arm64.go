package blas

import (
	"unsafe"
)

//go:noescape

func _hsum_float32(input, result unsafe.Pointer, size uint64)
func HsumFloat32(input []float32) float32 {
	var sum float32
	_hsum_float32(unsafe.Pointer(&input[0]), unsafe.Pointer(&sum), uint64(len(input)))
	return sum
}
