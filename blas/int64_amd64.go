package blas

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape

func _add_int64(input1, input2, output unsafe.Pointer, info uint64)
func (o Int64s) Add(input1, input2 []int64) {
	_add_int64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _sub_int64(input1, input2, output unsafe.Pointer, info uint64)
func (o Int64s) Sub(input1, input2 []int64) {
	_sub_int64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _mul_int64(input1, input2, output unsafe.Pointer, info uint64)
func (o Int64s) Mul(input1, input2 []int64) {
	_mul_int64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}

func _div_int64(input1, input2, output unsafe.Pointer, info uint64)
func (o Int64s) Div(input1, input2 []int64) {
	_div_int64(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)|cpu.Info))
}
