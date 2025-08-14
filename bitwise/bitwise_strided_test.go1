package bitwise

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAndStrided(t *testing.T) {
	tests := []struct {
		name     string
		a, b     Uint64s
		expected Uint64s
	}{
		{
			name:     "only word 3 survives",
			a:        Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0},
			b:        Uint64s{0xffff_ffff_ffff_ffff, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0},
			expected: Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0},
		},
		{
			name:     "all zeros",
			a:        Uint64s{0, 0, 0, 0},
			b:        Uint64s{0, 0, 0, 0},
			expected: Uint64s{0, 0, 0, 0},
		},
		{
			name:     "all ones",
			a:        Uint64s{^uint64(0), ^uint64(0)},
			b:        Uint64s{^uint64(0), ^uint64(0)},
			expected: Uint64s{^uint64(0), ^uint64(0)},
		},
	}

	fmt.Println("andStrided")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(Uint64s, len(tt.a))
			strides := Strides{}
			strides.ensureStrides(tt.a)

			count := 0
			fmt.Println(strides)
			out.andStrided(tt.a, tt.b, strides, &count)
			fmt.Println(strides)

			for i := range tt.expected {
				if out[i] != tt.expected[i] {
					t.Fatalf("word %d: got %016x, want %016x", i, out[i], tt.expected[i])
				}
			}
		})
	}

	fmt.Println("AndStridedOptimized")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(Uint64s, len(tt.a))
			strides := Strides{}
			strides.ensureStrides(tt.a)

			out.AndStridedOptimized(tt.a, tt.b, strides)
			fmt.Println(strides)

			for i := range tt.expected {
				if out[i] != tt.expected[i] {
					t.Fatalf("word %d: got %016x, want %016x", i, out[i], tt.expected[i])
				}
			}
		})
	}
}

func TestOrStrided(t *testing.T) {
	tests := []struct {
		name     string
		a, b     Uint64s
		expected Uint64s
	}{
		{
			name:     "simple OR with saturation at word 0 and 3",
			a:        Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0},
			b:        Uint64s{0xffff_ffff_ffff_ffff, 0, 0, 0xffff_ffff_ffff_ffff, 0},
			expected: Uint64s{^uint64(0), 0, 0, ^uint64(0), 0},
		},
		{
			name:     "all zeros",
			a:        Uint64s{0, 0, 0},
			b:        Uint64s{0, 0, 0},
			expected: Uint64s{0, 0, 0},
		},
		{
			name:     "all ones + zeros",
			a:        Uint64s{^uint64(0), 0, ^uint64(0)},
			b:        Uint64s{0, ^uint64(0), 0},
			expected: Uint64s{^uint64(0), ^uint64(0), ^uint64(0)},
		},
	}

	fmt.Println("orStrided")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(Uint64s, len(tt.a))
			strides := Strides{}
			strides.ensureStrides(tt.a)

			out.orStrided(tt.a, tt.b, strides)
			fmt.Println(strides)

			for i := range tt.expected {
				if out[i] != tt.expected[i] {
					t.Fatalf("word %d: got %016x, want %016x", i, out[i], tt.expected[i])
				}
			}
		})
	}

	fmt.Println("OrStridedOptimized")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(Uint64s, len(tt.a))
			strides := Strides{}
			strides.ensureStrides(tt.a)

			out.OrStridedOptimized(tt.a, tt.b, strides)
			fmt.Println(strides)

			for i := range tt.expected {
				if out[i] != tt.expected[i] {
					t.Fatalf("word %d: got %016x, want %016x", i, out[i], tt.expected[i])
				}
			}
		})
	}
}

func TestAndV6StridedOptimized(t *testing.T) {
	// Setup test input length and stride (stride = 1 for simplicity)
	length := 8

	// Allocate input vectors
	v1 := Uint64s{0xFF, 0x0F, 0xF0, 0xAA, 0x55, 0xFF, 0xF0, 0x0F}
	v2 := Uint64s{0xF0, 0xFF, 0x0F, 0xAA, 0x55, 0x00, 0xF0, 0x0F}
	v3 := Uint64s{0x0F, 0x0F, 0x0F, 0xFF, 0xFF, 0x00, 0xF0, 0xF0}
	v4 := Uint64s{0xFF, 0xFF, 0xF0, 0x00, 0xFF, 0xFF, 0x00, 0xF0}
	v5 := Uint64s{0xFF, 0x00, 0xF0, 0xAA, 0x55, 0x00, 0xFF, 0x0F}
	v6 := Uint64s{0x0F, 0xF0, 0x0F, 0xAA, 0x55, 0xFF, 0xFF, 0x0F}

	strides := Strides{}
	strides.ensureStrides(v1)
	// Output slice
	out := make(Uint64s, length)

	// Call the function under test
	out.AndV6StridedOptimized(v1, v2, v3, v4, v5, v6, strides)

	// Compute expected result using pure Go for verification
	expected := make(Uint64s, length)
	for i := 0; i < length; i++ {
		expected[i] = v1[i] & v2[i] & v3[i] & v4[i] & v5[i] & v6[i]
	}

	// Validate result
	assert.Equal(t, expected, out, "AndV6StridedOptimized output mismatch")
}

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
