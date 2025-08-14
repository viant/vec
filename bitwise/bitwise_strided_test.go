package bitwise

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viant/vec/cpu"
)

//func TestAnd64Sve(t *testing.T) {
//	if !cpu.CanUseSVE() {
//		fmt.Println("no sve support")
//		t.Skip()
//	}
//	length := 8
//
//	// Allocate input vectors
//	v1 := Uint64s{0xFF, 0x0F, 0xF0, 0x00, 0x55, 0xFF, 0x0F, 0x0F}
//	v2 := Uint64s{0x00, 0x00, 0x0F, 0xAA, 0xFF, 0x00, 0x0F, 0x00}
//
//	mask := FindNonZero64bitWordsSve(v2)
//	out := make(Uint64s, length)
//	out.AndMasked64bitSve(v1, v2, mask)
//	fmt.Println(out)
//
//	expected := make(Uint64s, length)
//	for i := 0; i < length; i++ {
//		expected[i] = v1[i] & v2[i]
//	}
//
//	// Validate result
//	assert.Equal(t, expected, out, "AndMasked output mismatch")
//
//}

func TestAndStrided(t *testing.T) {

	length := 8

	// Allocate input vectors
	v1 := Uint64s{0xFF, 0x0F, 0xF0, 0x00, 0x55, 0xFF, 0x0F, 0x0F}
	v2 := Uint64s{0x00, 0x00, 0x0F, 0xAA, 0xFF, 0x00, 0x0F, 0x00}

	strides := make(Strides, 2*length+3)
	strides.SetAndHotBlocks(v2)

	out := append(Uint64s(nil), v2...)
	out.AndStrided(v1, v2, strides)

	expected := make(Uint64s, length)
	for i := 0; i < length; i++ {
		expected[i] = v1[i] & v2[i]
	}

	// Validate result
	assert.Equal(t, expected, out, "AndStrided output mismatch")

}

func TestOrStrided(t *testing.T) {

	length := 8

	// Allocate input vectors
	v1 := Uint64s{^uint64(0), 0x0F, 0xF0, 0x00, 0x55, ^uint64(0), 0x0F, 0x0F}
	v2 := Uint64s{0x00, 0x00, 0x0F, 0xAA, 0xFF, 0x00, 0x0F, 0x00}

	strides := make(Strides, 2*length+3)
	strides.SetOrHotBlocks(v1)

	out := append(Uint64s(nil), v1...)
	out.OrStrided(v1, v2, strides)

	expected := make(Uint64s, length)
	for i := 0; i < length; i++ {
		expected[i] = v1[i] | v2[i]
	}

	// Validate result
	assert.Equal(t, expected, out, "OrStrided output mismatch")

}

func TestOr3Strided(t *testing.T) {

	length := 8

	// Allocate input vectors
	v1 := Uint64s{^uint64(0), 0x01, 0x00, 0x00, 0x00, ^uint64(0), 0x00, 0x00}
	v2 := Uint64s{0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00}
	v3 := Uint64s{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03}

	strides := make(Strides, 2*length+3)
	strides.SetOrHotBlocks(v1)

	out := append(Uint64s(nil), v1...)
	out.Or3Strided(v1, v2, v3, strides)

	expected := make(Uint64s, length)
	for i := 0; i < length; i++ {
		expected[i] = v1[i] | v2[i] | v3[i]
	}

	// Validate result
	assert.Equal(t, expected, out, "Or3Strided output mismatch")

}

//////////

func makeSparseVector(n int, nonZeroFraction float64) Uint64s {
	vec := make(Uint64s, n)
	for i := range vec {
		if rand.Float64() < nonZeroFraction {
			vec[i] = rand.Uint64()
		}
	}
	return vec
}

func makeZeroPrefixedVector(n int, fraction float64) Uint64s {
	vec := make(Uint64s, n)
	for i := range vec {
		if float64(i)/float64(len(vec)) < fraction {
			vec[i] = 0
		} else {
			vec[i] = ^uint64(0)
		}
	}
	return vec
}

func makeOnePrefixedVector(n int, fraction float64) Uint64s {
	vec := make(Uint64s, n)
	for i := range vec {
		if float64(i)/float64(len(vec)) < fraction {
			vec[i] = ^uint64(0)

		} else {
			vec[i] = 0
		}
	}
	return vec
}

// ///////
const N = 1 << 10
const sparsity = 0.50 // % of elements are non-zero
func BenchmarkAndProd(b *testing.B) {
	a := makeZeroPrefixedVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.And(a, bv)
	}
}

func BenchmarkAndSveProd(b *testing.B) {
	if !cpu.CanUseSVE() {
		fmt.Println("no sve support")
		b.Skip()
	}
	a := makeZeroPrefixedVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndSVE(a, bv)
	}
}

func BenchmarkAndStrided(b *testing.B) {
	a := makeZeroPrefixedVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	strides := make(Strides, 2*len(a)+3)
	strides.SetAndHotBlocks(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndStrided(a, bv, strides)
	}
}

func BenchmarkOrStrided(b *testing.B) {
	a := makeOnePrefixedVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	strides := make(Strides, 2*len(a)+3)
	strides.SetAndHotBlocks(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.OrStrided(a, bv, strides)
	}
}

func BenchmarkOr3Strided(b *testing.B) {
	v1 := makeOnePrefixedVector(N, sparsity)
	v2 := makeSparseVector(N, 1.0) // fully populated
	v3 := makeSparseVector(N, 1.0) // fully populated
	out := make(Uint64s, N)

	strides := make(Strides, 2*len(v1)+3)
	strides.SetAndHotBlocks(v1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out.Or3Strided(v1, v2, v3, strides)
	}
}
