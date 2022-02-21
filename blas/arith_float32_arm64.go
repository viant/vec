package blas

import "unsafe"

//go:noescape

func _add_float32(input1, input2, output unsafe.Pointer, size uint64)
func (o Float32s) AddFloat32(input1, input2 []float32) {
	_add_float32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}

func _sub_float32(input1, input2, output unsafe.Pointer, size uint64)
func (o Float32s) SubFloat32(input1, input2 []float32) {
	_sub_float32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}

func _mul_float32(input1, input2, output unsafe.Pointer, size uint64)
func (o Float32s) MulFloat32(input1, input2 []float32) {
	_mul_float32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}

func _div_float32(input1, input2, output unsafe.Pointer, size uint64)
func (o Float32s) DivFloat32(input1, input2 []float32) {
	_div_float32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}
