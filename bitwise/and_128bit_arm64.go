//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import "unsafe"

//go:noescape
func _find_nonzero_128bit_words(v, mask unsafe.Pointer)

//go:noescape
func _and_masked_128bit_neon(v1, v2, out, mask unsafe.Pointer)

func FindNonZero128bitWords(v Uint64s) []uint32 {
	n := len(v)
	mask := make([]uint32, n/2+3)
	mask[0] = uint32(n / 2)
	mask[1] = 0
	_find_nonzero_128bit_words(
		unsafe.Pointer(&v[0]),
		unsafe.Pointer(&mask[0]),
	)
	return mask
}

func (o Uint64s) AndMasked128bitNeon(v1, v2 Uint64s, mask []uint32) {
	_and_masked_128bit_neon(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}

func _and_masked_128bit_unrolled_neon(v1, v2, out, mask unsafe.Pointer)

func (o Uint64s) AndMasked128bitUnrolledNeon(v1, v2 Uint64s, mask []uint32) {
	_and_masked_128bit_unrolled_neon(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}

//go:noescape
func _find_nonzero_spans(v, mask unsafe.Pointer)

//go:noescape
func _and_span_neon(v1, v2, out, mask unsafe.Pointer)

func FindNonZeroSpans(v Uint64s) []uint32 {
	n := len(v)
	mask := make([]uint32, n+3)
	mask[0] = uint32(n / 2)
	mask[1] = 0
	_find_nonzero_spans(
		unsafe.Pointer(&v[0]),
		unsafe.Pointer(&mask[0]),
	)
	return mask
}

func (o Uint64s) AndSpanNeon(v1, v2 Uint64s, mask []uint32) {
	_and_span_neon(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}

//go:noescape
func _find_nonzero_64spans(v, mask unsafe.Pointer)

//go:noescape
func _and_64span_neon(v1, v2, out, mask unsafe.Pointer)

func FindNonZero64Spans(v Uint64s) []uint32 {
	n := len(v)
	mask := make([]uint32, 2*n+3)
	mask[0] = uint32(n)
	mask[1] = 0
	_find_nonzero_64spans(
		unsafe.Pointer(&v[0]),
		unsafe.Pointer(&mask[0]),
	)
	return mask
}

func (o Uint64s) And64SpanNeon(v1, v2 Uint64s, mask []uint32) {
	_and_64span_neon(
		unsafe.Pointer(&v1[0]),
		unsafe.Pointer(&v2[0]),
		unsafe.Pointer(&o[0]),
		unsafe.Pointer(&mask[0]),
	)
}
