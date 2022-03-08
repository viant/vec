package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape

func _add_float64(input1, input2, output unsafe.Pointer, info uint64)
func (o Float64s) AddFloat64(input1, input2 []float64) {
	_add_float64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _sub_float64(input1, input2, output unsafe.Pointer, info uint64)
func (o Float64s) SubFloat64(input1, input2 []float64) {
	_sub_float64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _mul_float64(input1, input2, output unsafe.Pointer, info uint64)
func (o Float64s) MulFloat64(input1, input2 []float64) {
	_mul_float64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _div_float64(input1, input2, output unsafe.Pointer, info uint64)
func (o Float64s) DivFloat64(input1, input2 []float64) {
	_div_float64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}
