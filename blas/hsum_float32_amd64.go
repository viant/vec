package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape

func _hsum_float32(input, result unsafe.Pointer, info uint64)
func HsumFloat32(input []float32) float32 {
	var sum float32
	_hsum_float32(unsafe.Pointer(&input[0]), unsafe.Pointer(&sum), uint64(len(input)|cpu.Info))
	return sum
}
