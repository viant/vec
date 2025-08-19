package bitwise

import (
	"unsafe"

	"github.com/viant/vec/cpu"
)

//go:noescape
func _and_optimized(out, v1, v2 unsafe.Pointer, info uint64)

// And applies optimized AND: v1 & v2 -> o
func (o Uint64s) And(v1, v2 Uint64s) {
	_and_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _and_avx2(out, v1, v2 unsafe.Pointer, size int)

// AndAVX2 applies AVX2 vectorized AND: v1 & v2 -> o
func (o Uint64s) AndAVX2(v1, v2 Uint64s) {
	_and_avx2(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _and_avx512(out, v1, v2 unsafe.Pointer, size int)

// AndAVX512 applies AVX2 vectorized AND: v1 & v2 -> o
func (o Uint64s) AndAVX512(v1, v2 Uint64s) {
	_and_avx512(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _and_v3_avx2(out, v1, v2, v3 unsafe.Pointer, size int)

// AndV3AVX2 applies AVX2 vectorized and: v1 & v2 & v3 -> o
func (o Uint64s) AndV3AVX2(v1, v2, v3 Uint64s) {
	_and_v3_avx2(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), len(o))
}

//go:noescape
func _and_v3_avx512(out, v1, v2, v3 unsafe.Pointer, size int)

// AndV3AVX512 applies AVX2 vectorized and: v1 & v2 & v3 -> o
func (o Uint64s) AndV3AVX512(v1, v2, v3 Uint64s) {
	_and_v3_avx512(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), len(o))
}

//go:noescape
func _and_v3_optimized(out, v1, v2, v3 unsafe.Pointer, info uint64)

// AndV3 applies optimized AND: v1 & v2 & v3 -> o
func (o Uint64s) AndV3(v1, v2, v3 Uint64s) {
	_and_v3_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)|cpu.Info))
}

// AndV4 applies optimized AND: v1 & v2 & v3 & v4-> o
func (o Uint64s) AndV4(v1, v2, v3, v4 Uint64s) {
	o.AndV3(v1, v2, v3)
	o.And(o, v4)
}

// AndV5 applies optimized AND: v1 & v2 & v3 & v4 & v5-> o
func (o Uint64s) AndV5(v1, v2, v3, v4, v5 Uint64s) {
	o.AndV3(v1, v2, v3)
	o.AndV3(o, v4, v5)
}

// AndV6 applies optimized AND: v1 & v2 & v3 & v4 & v5 & v6 -> o
func (o Uint64s) AndV6(v1, v2, v3, v4, v5, v6 Uint64s) {
	o.AndV3(v1, v2, v3)
	o.AndV3(o, v4, v5)
	o.And(o, v6)
}

//go:noescape
func _or_optimized(out, v1, v2 unsafe.Pointer, info uint64)

// Or applies optimized OR: v1 | v2 -> o
func (o Uint64s) Or(v1, v2 Uint64s) {
	_or_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _or_avx2(out, v1, v2 unsafe.Pointer, size int)

// OrAVX2 applies AVX2 vectorized OR: v1 | v2 -> o
func (o Uint64s) OrAVX2(v1, v2 Uint64s) {
	_or_avx2(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _or_v3_avx2(out, v1, v2, v3 unsafe.Pointer, size int)

// OrV3AVX2 applies AVX2 vectorized or: v1 | v2 | v3 -> o
func (o Uint64s) OrV3AVX2(v1, v2, v3 Uint64s) {
	_or_v3_avx2(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), len(o))
}

//go:noescape
func _or_v3_avx512(out, v1, v2, v3 unsafe.Pointer, size int)

// OrV3AVX512 applies AVX2 vectorized or: v1 | v2 | v3 -> o
func (o Uint64s) OrV3AVX512(v1, v2, v3 Uint64s) {
	_or_v3_avx512(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), len(o))
}

//go:noescape
func _or_v3_optimized(out, v1, v2, v3 unsafe.Pointer, info uint64)

// OrV3 applies optimized OR: v1 | v2 | v3 -> o
func (o Uint64s) OrV3(v1, v2, v3 Uint64s) {
	_or_v3_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)|cpu.Info))
}

// OrV4 applies optimized OR: v1 | v2 | v3 | v4-> o
func (o Uint64s) OrV4(v1, v2, v3, v4 Uint64s) {
	o.OrV3(v1, v2, v3)
	o.Or(o, v4)
}

// OrV5 applies optimized OR: v1 | v2 | v3 | v4 | v5-> o
func (o Uint64s) OrV5(v1, v2, v3, v4, v5 Uint64s) {
	o.OrV3(v1, v2, v3)
	o.OrV3(o, v4, v5)
}

// OrV6 applies optimized OR: v1 | v2 | v3 | v4 | v5 | v6 -> o
func (o Uint64s) OrV6(v1, v2, v3, v4, v5, v6 Uint64s) {
	o.OrV3(v1, v2, v3)
	o.OrV3(o, v4, v5)
	o.Or(o, v6)
}

//go:noescape
func _xor_optimized(out, v1, v2 unsafe.Pointer, info uint64)

// Xor applies optimized XOR: v1 ^ v2 -> o
func (o Uint64s) Xor(v1, v2 Uint64s) {
	_xor_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _xor_avx2(out, v1, v2 unsafe.Pointer, size int)

// XorAVX2 applies AVX2 vectorized XOR: v1 ^ v2 -> o
func (o Uint64s) XorAVX2(v1, v2 Uint64s) {
	_xor_avx2(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _xor_v3_avx2(out, v1, v2, v3 unsafe.Pointer, size int)

// XorV3AVX2 applies AVX2 vectorized xor: v1 ^ v2 ^ v3 -> o
func (o Uint64s) XorV3AVX2(v1, v2, v3 Uint64s) {
	_xor_v3_avx2(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), len(o))
}

//go:noescape
func _xor_v3_avx512(out, v1, v2, v3 unsafe.Pointer, size int)

// XorV3AVX512 applies AVX2 vectorized xor: v1 ^ v2 ^ v3 -> o
func (o Uint64s) XorV3AVX512(v1, v2, v3 Uint64s) {
	_xor_v3_avx512(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), len(o))
}

//go:noescape
func _xor_v3_optimized(out, v1, v2, v3 unsafe.Pointer, info uint64)

// XorV3 applies optimized XOR: v1 ^ v2 ^ v3 -> o
func (o Uint64s) XorV3(v1, v2, v3 Uint64s) {
	_xor_v3_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)|cpu.Info))
}
