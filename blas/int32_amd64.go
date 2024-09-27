package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func (o Int32s) Inc(constant int32) {
	o.inc(constant) // This is a fallback implementation, not a real one
}

func _add_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) Add(input1, input2 []int32) {
	_add_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _sub_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) Sub(input1, input2 []int32) {
	_sub_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func (o Int32s) ScalarMul(constant int32) {
	o.scalarMul(constant) // This is a fallback implementation, not a real one
}

func _mul_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) Mul(input1, input2 []int32) {
	_mul_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _div_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) Div(input1, input2 []int32) {
	_div_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}
