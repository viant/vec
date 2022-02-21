package bitwise

import (
	"github.com/viant/vec/cpu"
	"unsafe"
)

//go:noescape
func _and_neon(out, v1, v2 unsafe.Pointer, size uint64)

//AndNeon applies vectorized and: v1 & v2 -> o
func (o Uint64s) AndNeon(v1, v2 Uint64s) {
	_and_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)))
}

//go:noescape
func _and_sve(acc, v1, v2 unsafe.Pointer, size int)

//AndSVE applies vectorized and: v1 & v2 -> o
func (o Uint64s) AndSVE(v1, v2 Uint64s) {
	_and_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _and_optimized(acc, v1, v2 unsafe.Pointer, info uint64)

//And applies vectorized and: v1 & v2 -> o
func (o Uint64s) And(v1, v2 Uint64s) {
	_and_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _and_v3_neon(out, v1, v2, v3 unsafe.Pointer, size uint64)

//AndV3Neon applies vectorized and: v1 & v2  & v3-> o
func (o Uint64s) AndV3Neon(v1, v2, v3 Uint64s) {
	_and_v3_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)))
}

//go:noescape
func _and_v4_neon(out, v1, v2, v3, v4 unsafe.Pointer, size uint64)

//AndV4Neon applies vectorized and: v1 & v2  & v3 & v4 -> o
func (o Uint64s) AndV4Neon(v1, v2, v3, v4 Uint64s) {
	_and_v4_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), uint64(len(o)))
}

//go:noescape
func _and_v5_neon(out, v1, v2, v3, v4, v5 unsafe.Pointer, size uint64)

//AndV5Neon applies vectorized and: v1 & v2  & v3 & v4 & v5-> o
func (o Uint64s) AndV5Neon(v1, v2, v3, v4, v5 Uint64s) {
	_and_v5_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), uint64(len(o)))
}

//go:noescape
func _and_v6_neon(out, v1, v2, v3, v4, v5, v6 unsafe.Pointer, size uint64)

//AndV6Neon applies vectorized and: v1 & v2  & v3 & v4 & v5-> o
func (o Uint64s) AndV6Neon(v1, v2, v3, v4, v5, v6 Uint64s) {
	_and_v6_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), unsafe.Pointer(&v6[0]), uint64(len(o)))
}

//go:noescape
func _and_v3_sve(out, v1, v2, v3 unsafe.Pointer, size uint64)

//AndV3SVE applies vectorized and: v1 & v2  & v3 -> o
func (o Uint64s) AndV3SVE(v1, v2, v3 Uint64s) {
	_and_v3_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)))
}

//go:noescape
func _and_v4_sve(out, v1, v2, v3, v4 unsafe.Pointer, size uint64)

//AndV4SVE applies vectorized and: v1 & v2  & v3 & v4 -> o
func (o Uint64s) AndV4SVE(v1, v2, v3, v4 Uint64s) {
	_and_v4_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), uint64(len(o)))
}

//go:noescape
func _and_v5_sve(out, v1, v2, v3, v4, v5 unsafe.Pointer, size uint64)

//AndV5SVE applies vectorized and: v1 & v2  & v3 & v4 & v5 -> o
func (o Uint64s) AndV5SVE(v1, v2, v3, v4, v5 Uint64s) {
	_and_v5_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), uint64(len(o)))
}

//go:noescape
func _and_v6_sve(out, v1, v2, v3, v4, v5, v6 unsafe.Pointer, size uint64)

//AndV6SVE applies vectorized and: v1 & v2  & v3 & v4 & v5 & v6 -> o
func (o Uint64s) AndV6SVE(v1, v2, v3, v4, v5, v6 Uint64s) {
	_and_v6_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), unsafe.Pointer(&v6[0]), uint64(len(o)))
}

//go:noescape
func _and_v3_optimized(acc, v1, v2, v3 unsafe.Pointer, info uint64)

//AndV3 applies vectorized and: v1 & v2 & v3 -> o
func (o Uint64s) AndV3(v1, v2, v3 Uint64s) {
	_and_v3_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _and_v4_optimized(acc, v1, v2, v3, v4 unsafe.Pointer, info uint64)

//AndV4 applies vectorized and: v1 & v2 &v3 & v4 -> o
func (o Uint64s) AndV4(v1, v2, v3, v4 Uint64s) {
	_and_v4_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _and_v5_optimized(acc, v1, v2, v3, v4, v5 unsafe.Pointer, info uint64)

//AndV5 applies vectorized and: v1 & v2 &v3 & v4 -> o
func (o Uint64s) AndV5(v1, v2, v3, v4, v5 Uint64s) {
	_and_v5_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _and_v6_optimized(acc, v1, v2, v3, v4, v5, v6 unsafe.Pointer, info uint64)

//AndV6 applies vectorized and: v1 & v2 &v3 & v4 & v6 -> o
func (o Uint64s) AndV6(v1, v2, v3, v4, v5, v6 Uint64s) {
	_and_v6_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), unsafe.Pointer(&v6[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _or_neon(out, v1, v2 unsafe.Pointer, size uint64)

//OrNeon applies vectorized or: v1 | v2 -> o
func (o Uint64s) OrNeon(v1, v2 Uint64s) {
	_or_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)))
}

//go:noescape
func _or_sve(acc, v1, v2 unsafe.Pointer, size int)

//OrSVE applies vectorized or: v1 | v2 -> o
func (o Uint64s) OrSVE(v1, v2 Uint64s) {
	_or_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _or_optimized(acc, v1, v2 unsafe.Pointer, info uint64)

//Or applies vectorized or: v1 | v2 -> o
func (o Uint64s) Or(v1, v2 Uint64s) {
	_or_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _or_v3_neon(out, v1, v2, v3 unsafe.Pointer, size uint64)

//OrV3Neon applies vectorized or: v1 | v2 | v3-> o
func (o Uint64s) OrV3Neon(v1, v2, v3 Uint64s) {
	_or_v3_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)))
}

//go:noescape
func _or_v4_neon(out, v1, v2, v3, v4 unsafe.Pointer, size uint64)

//OrV4Neon applies vectorized or: v1 | v2 | v3 | v4 -> o
func (o Uint64s) OrV4Neon(v1, v2, v3, v4 Uint64s) {
	_or_v4_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), uint64(len(o)))
}

//go:noescape
func _or_v5_neon(out, v1, v2, v3, v4, v5 unsafe.Pointer, size uint64)

//OrV5Neon applies vectorized or: v1 | v2 | v3-> o
func (o Uint64s) OrV5Neon(v1, v2, v3, v4, v5 Uint64s) {
	_or_v5_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), uint64(len(o)))
}

//go:noescape
func _or_v6_neon(out, v1, v2, v3, v4, v5, v6 unsafe.Pointer, size uint64)

//OrV6Neon applies vectorized or: v1 | v2 | v3-> o
func (o Uint64s) OrV6Neon(v1, v2, v3, v4, v5, v6 Uint64s) {
	_or_v6_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), unsafe.Pointer(&v6[0]), uint64(len(o)))
}

//go:noescape
func _or_v3_sve(out, v1, v2, v3 unsafe.Pointer, size uint64)

//OrV3SVE applies vectorized or: v1 | v2  | v3-> o
func (o Uint64s) OrV3SVE(v1, v2, v3 Uint64s) {
	_or_v3_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)))
}

//go:noescape
func _or_v4_sve(out, v1, v2, v3, v4 unsafe.Pointer, size uint64)

//OrV4SVE applies vectorized or: v1 | v2  | v3 | v4 -> o
func (o Uint64s) OrV4SVE(v1, v2, v3, v4 Uint64s) {
	_or_v4_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), uint64(len(o)))
}

//go:noescape
func _or_v5_sve(out, v1, v2, v3, v4, v5 unsafe.Pointer, size uint64)

//OrV5SVE applies vectorized or: v1 | v2  | v3-> o
func (o Uint64s) OrV5SVE(v1, v2, v3, v4, v5 Uint64s) {
	_or_v5_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), uint64(len(o)))
}

//go:noescape
func _or_v6_sve(out, v1, v2, v3, v4, v5, v6 unsafe.Pointer, size uint64)

//OrV6SVE applies vectorized or: v1 | v2  | v3-> o
func (o Uint64s) OrV6SVE(v1, v2, v3, v4, v5, v6 Uint64s) {
	_or_v6_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), unsafe.Pointer(&v6[0]), uint64(len(o)))
}

//go:noescape
func _or_v3_optimized(acc, v1, v2, v3 unsafe.Pointer, info uint64)

//OrV3 applies vectorized or: v1 | v2 | v3 -> o
func (o Uint64s) OrV3(v1, v2, v3 Uint64s) {
	_or_v3_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _or_v4_optimized(acc, v1, v2, v3, v4 unsafe.Pointer, info uint64)

//OrV4 applies vectorized or: v1 | v2 | v3 | v4 -> o
func (o Uint64s) OrV4(v1, v2, v3, v4 Uint64s) {
	_or_v4_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _or_v5_optimized(acc, v1, v2, v3, v4, v5 unsafe.Pointer, info uint64)

//OrV5 applies vectorized or: v1 | v2 | v3 | v4 | v5 -> o
func (o Uint64s) OrV5(v1, v2, v3, v4, v5 Uint64s) {
	_or_v5_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _or_v6_optimized(acc, v1, v2, v3, v4, v5, v6 unsafe.Pointer, info uint64)

//OrV6 applies vectorized or: v1 | v2 | v3 | v4 | v5 | v6 -> o
func (o Uint64s) OrV6(v1, v2, v3, v4, v5, v6 Uint64s) {
	_or_v6_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), unsafe.Pointer(&v4[0]), unsafe.Pointer(&v5[0]), unsafe.Pointer(&v6[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _xor_neon(out, v1, v2 unsafe.Pointer, size uint64)

//XorNeon applies vectorized xor: v1 ^ v2 -> o
func (o Uint64s) XorNeon(v1, v2 Uint64s) {
	_xor_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)))
}

//go:noescape
func _xor_sve(acc, v1, v2 unsafe.Pointer, size int)

//Xor applies vectorized XOR: v1 ^ v2 -> o
func (o Uint64s) XorSVE(v1, v2 Uint64s) {
	_xor_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), len(o))
}

//go:noescape
func _xor_optimized(acc, v1, v2 unsafe.Pointer, info uint64)

//Xor applies vectorized XOR: v1 ^ v2 -> o
func (o Uint64s) Xor(v1, v2 Uint64s) {
	_xor_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), uint64(len(o)|cpu.Info))
}

//go:noescape
func _xor_v3_neon(out, v1, v2, v3 unsafe.Pointer, size uint64)

//Xor applies vectorized XOR: v1 ^ v2  ^ v3-> o
func (o Uint64s) XorV3Neon(v1, v2, v3 Uint64s) {
	_xor_v3_neon(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)))
}

//go:noescape
func _xor_v3_sve(out, v1, v2, v3 unsafe.Pointer, size uint64)

//Xor applies vectorized XOR: v1 ^ v2  ^ v3-> o
func (o Uint64s) XorV3SVE(v1, v2, v3 Uint64s) {
	_xor_v3_sve(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)))
}

//go:noescape
func _xor_v3_optimized(acc, v1, v2, v3 unsafe.Pointer, info uint64)

//Xor applies vectorized XOR: v1 ^ v2 ^ v3 -> o
func (o Uint64s) XorV3(v1, v2, v3 Uint64s) {
	_xor_v3_optimized(unsafe.Pointer(&o[0]), unsafe.Pointer(&v1[0]), unsafe.Pointer(&v2[0]), unsafe.Pointer(&v3[0]), uint64(len(o)|cpu.Info))
}
