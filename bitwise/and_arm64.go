//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import "unsafe"

//go:noescape
func _find_nonzero_words(v, mask unsafe.Pointer)

//go:noescape
func _and_masked_naive(v1, v2, out, mask unsafe.Pointer)

//go:noescape
func _and_masked_neon(v1, v2, out, mask unsafe.Pointer)

//go:noescape
func _and_strided_neon(out, v1, v2, mask unsafe.Pointer)

func FindNonZeroWords(v Uint64s) []uint32 {
	n := len(v)
	mask := make([]uint32, n+3)
	mask[0] = uint32(n)
	mask[1] = 0
	_find_nonzero_words(
		unsafe.Pointer(&v[0]),
		unsafe.Pointer(&mask[0]),
	)
	return mask
}

func (o Uint64s) AndMaskedNaive(v1, v2 Uint64s, mask []uint32) {
	_and_masked_naive(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}

func (o Uint64s) AndMaskedNeon(v1, v2 Uint64s, mask []uint32) {
	_and_masked_neon(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}

func (o Uint64s) AndStridedNeon(v1, v2 Uint64s, strides Strides) {
	_and_strided_neon(
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&strides[0]),
	)
}
