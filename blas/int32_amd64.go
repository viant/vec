package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape

func _add_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) AddInt32(input1, input2 []int32) {
	_add_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _sub_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) SubInt32(input1, input2 []int32) {
	_sub_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _mul_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) MulInt32(input1, input2 []int32) {
	_mul_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _div_int32(input1, input2, output unsafe.Pointer, info uint64)
func (o Int32s) DivInt32(input1, input2 []int32) {
	_div_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}
