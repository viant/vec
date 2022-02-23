package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape

func _hsum_int32(input, result unsafe.Pointer, info uint64)
func HsumInt32(input []int32) int32 {
	var sum int32
	_hsum_int32(unsafe.Pointer(&input[0]), unsafe.Pointer(&sum), uint64(len(input)|cpu.Info))
	return sum
}

func _hmax_int32(input, result unsafe.Pointer, info uint64)
func HmaxInt32(input []int32) int32 {
	var max int32
	_hmax_int32(unsafe.Pointer(&input[0]), unsafe.Pointer(&max), uint64(len(input)|cpu.Info))
	return max
}

func _hmin_int32(input, result unsafe.Pointer, info uint64)
func HminInt32(input []int32) int32 {
	var min int32
	_hmin_int32(unsafe.Pointer(&input[0]), unsafe.Pointer(&min), uint64(len(input)|cpu.Info))
	return min
}
