package bitwise

import (
	"unsafe"

	"github.com/viant/vec/cpu"
)

/////////////////////////////
// AND implementations
/////////////////////////////

// Hot blocks
//
//go:noescape
func _set_and_hot_blocks(v, mask unsafe.Pointer)

func (s Strides) SetActiveStrides(set Uint64s) {
	s[0] = uint32(len(set))
	s[1] = uint32(cpu.Info >> 32)

	_set_and_hot_blocks(
		unsafe.Pointer(&set[0]),
		unsafe.Pointer(&s[0]),
	)
}

// 2 args
// Decls match the TEXT headers (order: v..., out, control).
//
//go:noescape
func _and_strided(v1, v2, out, control unsafe.Pointer)

func (o Uint64s) AndStrided(v1, v2 Uint64s, strides Strides) {
	_and_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 3 args
//
//go:noescape
func _and3_strided(v1, v2, v3, out, control unsafe.Pointer)

func (o Uint64s) And3Strided(v1, v2, v3 Uint64s, strides Strides) {
	_and3_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 4 args
//
//go:noescape
func _and4_strided(v1, v2, v3, v4, out, control unsafe.Pointer)

func (o Uint64s) And4Strided(v1, v2, v3, v4 Uint64s, strides Strides) {
	_and4_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 5 args
//
//go:noescape
func _and5_strided(v1, v2, v3, v4, v5, out, control unsafe.Pointer)

func (o Uint64s) And5Strided(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	_and5_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 6 args
//
//go:noescape
func _and6_strided(v1, v2, v3, v4, v5, v6, out, control unsafe.Pointer)

func (o Uint64s) And6Strided(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	_and6_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&v6[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

/////////////////////////////
// OR implementations
/////////////////////////////

// Hot blocks
//

// 2 args
// Decls match the TEXT headers (order: v..., out, control).
//
//go:noescape
func _or_strided(v1, v2, out, control unsafe.Pointer)

func (o Uint64s) OrStrided(v1, v2 Uint64s, strides Strides) {
	_or_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 3 args
//
//go:noescape
func _or3_strided(v1, v2, v3, out, control unsafe.Pointer)

func (o Uint64s) Or3Strided(v1, v2, v3 Uint64s, strides Strides) {
	_or3_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 4 args
//
//go:noescape
func _or4_strided(v1, v2, v3, v4, out, control unsafe.Pointer)

func (o Uint64s) Or4Strided(v1, v2, v3, v4 Uint64s, strides Strides) {
	_or4_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 5 args
//
//go:noescape
func _or5_strided(v1, v2, v3, v4, v5, out, control unsafe.Pointer)

func (o Uint64s) Or5Strided(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	_or5_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

// 6 args
//
//go:noescape
func _or6_strided(v1, v2, v3, v4, v5, v6, out, control unsafe.Pointer)

func (o Uint64s) Or6Strided(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	_or6_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&v6[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}
