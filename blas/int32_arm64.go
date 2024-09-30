package blas

import "unsafe"

//go:noescape

func _inc_int32(vec unsafe.Pointer, constant int32, size uint64)
func (o Int32s) Inc(constant int32) {
	_inc_int32(unsafe.Pointer(&o[0]), constant, uint64(len(o)))
}

func _add_int32(input1, input2, output unsafe.Pointer, size uint64)
func (o Int32s) Add(input1, input2 []int32) {
	_add_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}

func _sub_int32(input1, input2, output unsafe.Pointer, size uint64)
func (o Int32s) Sub(input1, input2 []int32) {
	_sub_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}

func _scalar_mul_int32(vec unsafe.Pointer, constant int32, size uint64)
func (o Int32s) ScalarMul(constant int32) {
	_scalar_mul_int32(unsafe.Pointer(&o[0]), constant, uint64(len(o)))
}

func _mul_int32(input1, input2, output unsafe.Pointer, size uint64)
func (o Int32s) Mul(input1, input2 []int32) {
	_mul_int32(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(o)[0]), uint64(len(o)))
}

func _inc_int32_sve(vec unsafe.Pointer, constant int32, size uint64)
func (o Int32s) IncSVE(constant int32) {
	_inc_int32_sve(unsafe.Pointer(&o[0]), constant, uint64(len(o)))
}

func _scalar_mul_int32_sve(vec unsafe.Pointer, constant int32, size uint64)
func (o Int32s) ScalarMulSVE(constant int32) {
	_scalar_mul_int32_sve(unsafe.Pointer(&o[0]), constant, uint64(len(o)))
}
