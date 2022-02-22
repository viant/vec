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

func _hmax_float32(input, result unsafe.Pointer, info uint64)
func HmaxFloat32(input []float32) float32 {
	var max float32
	_hmax_float32(unsafe.Pointer(&input[0]), unsafe.Pointer(&max), uint64(len(input)|cpu.Info))
	return max
}

func _hmin_float32(input, result unsafe.Pointer, info uint64)
func HminFloat32(input []float32) float32 {
	var min float32
	_hmin_float32(unsafe.Pointer(&input[0]), unsafe.Pointer(&min), uint64(len(input)|cpu.Info))
	return min
}
