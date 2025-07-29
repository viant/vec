//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

// Low-level assembly hooks generated from cpp/bitwise_strided_arm64.cpp

//go:noescape
func _and_strided_optimized(out, v1, v2, control unsafe.Pointer)

//go:noescape
func _or_strided_optimized(out, v1, v2, control unsafe.Pointer)

// AndStridedOptimized performs strided AND with dynamic stride adjustment.
//
// The behaviour matches andStrided (pure Go) but leverages NEON or SVE SIMD
// according to cpu.Info encoded in strides[1].
func (o Uint64s) AndStridedOptimized(v1, v2 Uint64s, strides Strides) {
	if len(o) == 0 {
		return
	}
	// Prepare control block expected by C++ (size, feature, start, inc1, inc2 ...)
	tmp := make([]uint32, len(strides)+2)
	tmp[0] = uint32(len(o))
	tmp[1] = uint32(cpu.Info >> 32) // low-level discriminator matches consts
	copy(tmp[2:], strides)

	_and_strided_optimized(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&tmp[0]),
	)

	// Copy back updated increments (skip header size+feature)
	copy(strides, tmp[2:])
}

// OrStridedOptimized performs strided OR with SIMD, widening stride on saturation.
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
