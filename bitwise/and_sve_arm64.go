//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import "unsafe"

//go:noescape
func _find_nonzero_64bit_words_sve(v, mask unsafe.Pointer)

//go:noescape
func _and_masked_64bit_sve(v1, v2, out, mask unsafe.Pointer)

func FindNonZero64bitWordsSve(v Uint64s) Uint64s {
	n := len(v)
	mask := make(Uint64s, n+3)
	mask[0] = uint64(n)
	mask[1] = 0
	_find_nonzero_64bit_words_sve(
		unsafe.Pointer(&v[0]),
		unsafe.Pointer(&mask[0]),
	)
	return mask
}

func (o Uint64s) AndMasked64bitSve(v1, v2, mask Uint64s) {
	_and_masked_64bit_sve(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}
