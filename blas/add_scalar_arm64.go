package blas

import (
	"unsafe"
)

//go:noescape

func _add_scalar(input, output unsafe.Pointer, value int32, info uint64)

func (o Int32s) AddScalar(input []int32, value int32) {
	_add_scalar(unsafe.Pointer(&input[0]), unsafe.Pointer(&(o)[0]), value, uint64(len(input)))
}
