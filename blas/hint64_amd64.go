package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape

func _hsum_int64(input, result unsafe.Pointer, info uint64)
func HsumInt64(input []int64) int64 {
	var sum int64
	_hsum_int64(unsafe.Pointer(&input[0]), unsafe.Pointer(&sum), uint64(len(input)|cpu.Info))
	return sum
}

func _hmax_int64(input, result unsafe.Pointer, info uint64)
func HmaxInt64(input []int64) int64 {
	var max int64
	_hmax_int64(unsafe.Pointer(&input[0]), unsafe.Pointer(&max), uint64(len(input)|cpu.Info))
	return max
}

func _hmin_int64(input, result unsafe.Pointer, info uint64)
func HminInt64(input []int64) int64 {
	var min int64
	_hmin_int64(unsafe.Pointer(&input[0]), unsafe.Pointer(&min), uint64(len(input)|cpu.Info))
	return min
}
