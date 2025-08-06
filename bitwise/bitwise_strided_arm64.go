//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import (
	"unsafe"

	"github.com/viant/vec/cpu"
)

// Low-level assembly hooks generated from cpp/bitwise_strided_arm64.cpp

// AND
//
//go:noescape
func _and_strided_optimized(out, v1, v2, control unsafe.Pointer)

//go:noescape
func _and_v3_strided_optimized(out, v1, v2, v3, control unsafe.Pointer)

//go:noescape
func _and_v4_strided_optimized(out, v1, v2, v3, v4, control unsafe.Pointer)

//go:noescape
func _and_v5_strided_optimized(out, v1, v2, v3, v4, v5, control unsafe.Pointer)

//go:noescape
func _and_v6_strided_optimized(out, v1, v2, v3, v4, v5, v6, control unsafe.Pointer)

// OR
//
//go:noescape
func _or_strided_optimized(out, v1, v2, control unsafe.Pointer)

//go:noescape
func _or_v3_strided_optimized(out, v1, v2, v3, control unsafe.Pointer)

//go:noescape
func _or_v4_strided_optimized(out, v1, v2, v3, v4, control unsafe.Pointer)

//go:noescape
func _or_v5_strided_optimized(out, v1, v2, v3, v4, v5, control unsafe.Pointer)

//go:noescape
func _or_v6_strided_optimized(out, v1, v2, v3, v4, v5, v6, control unsafe.Pointer)

// AND implementations

func (o Uint64s) AndStridedOptimized(v1, v2 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_and_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) AndV3StridedOptimized(v1, v2, v3 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_and_v3_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) AndV4StridedOptimized(v1, v2, v3, v4 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_and_v4_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) AndV5StridedOptimized(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_and_v5_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) AndV6StridedOptimized(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_and_v6_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&v6[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

// OR implementations

func (o Uint64s) OrStridedOptimized(v1, v2 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_or_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) OrV3StridedOptimized(v1, v2, v3 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_or_v3_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) OrV4StridedOptimized(v1, v2, v3, v4 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_or_v4_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) OrV5StridedOptimized(v1, v2, v3, v4, v5 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_or_v5_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}

func (o Uint64s) OrV6StridedOptimized(v1, v2, v3, v4, v5, v6 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32)
	copy(tmp[2:], strides)

	_or_v6_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&v3[0]),
		unsafe.Pointer(&v4[0]),
		unsafe.Pointer(&v5[0]),
		unsafe.Pointer(&v6[0]),
		unsafe.Pointer(&tmp[0]),
	)

	copy(strides, tmp[2:])
}
