//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import (
	"unsafe"

	"github.com/viant/vec/cpu"
)

//// AND implementations

//go:noescape
func _set_and_hot_blocks(v, mask unsafe.Pointer)

func (s Strides) SetAndHotBlocks(set Uint64s) {
	s[0] = uint32(len(set))
	s[1] = uint32(cpu.Info >> 32)

	_set_and_hot_blocks(
		unsafe.Pointer(&set[0]),
		unsafe.Pointer(&s[0]),
	)
}

//go:noescape
func _and_strided(out, v1, v2, control unsafe.Pointer)

func (o Uint64s) AndStrided(v1, v2 Uint64s, strides Strides) {
	_and_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

//go:noescape
func _and3_strided(out, v1, v2, v3, control unsafe.Pointer)

func (o Uint64s) And3Strided(v1, v2, v3 Uint64s, strides Strides) {
	_and3_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

//// OR implementations

//go:noescape
func _set_or_hot_blocks(v, mask unsafe.Pointer)

func (s Strides) SetOrHotBlocks(set Uint64s) {
	s[0] = uint32(len(set))
	s[1] = uint32(cpu.Info >> 32)

	_set_or_hot_blocks(
		unsafe.Pointer(&set[0]),
		unsafe.Pointer(&s[0]),
	)
}

//go:noescape
func _or_strided(out, v1, v2, control unsafe.Pointer)

func (o Uint64s) OrStrided(v1, v2 Uint64s, strides Strides) {
	_or_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}

//go:noescape
func _or3_strided(out, v1, v2, v3, control unsafe.Pointer)

func (o Uint64s) Or3Strided(v1, v2, v3 Uint64s, strides Strides) {
	_or3_strided(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&strides[0]),
	)
}
