//go:build arm64 && !noasm

package bitwise

import (
	"math/rand"
	"testing"
)

const (
	dataSize = 1024 * 1024 // 1 million elements for realistic benchmark
)

// generateTestData initializes Uint64 slices with pseudo-random values.
func generateTestData(size int) (Uint64s, Uint64s, Uint64s) {
	v1 := make(Uint64s, size)
	v2 := make(Uint64s, size)
	out := make(Uint64s, size)
	for i := 0; i < size; i++ {
		v1[i] = rand.Uint64()
		v2[i] = rand.Uint64()
	}
	return out, v1, v2
}

func BenchmarkAnd(b *testing.B) {
	out, v1, v2 := generateTestData(dataSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out.And(v1, v2)
	}
}

func BenchmarkAndStridedOptimized(b *testing.B) {
	out, v1, v2 := generateTestData(dataSize)
	strides := Strides{}
	strides.ensureStrides(v1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//strideCopy := append(Strides(nil), strides...)
		out.AndStridedOptimized(v1, v2, strides)
	}
}

func BenchmarkOr(b *testing.B) {
	out, v1, v2 := generateTestData(dataSize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out.Or(v1, v2)
	}
}

func BenchmarkOrStridedOptimized(b *testing.B) {
	out, v1, v2 := generateTestData(dataSize)
	strides := Strides{}
	strides.ensureStrides(v1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//strideCopy := append(Strides(nil), strides...)
		out.OrStridedOptimized(v1, v2, strides)
	}
}
