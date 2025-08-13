package bitwise

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/viant/vec/cpu"
)

func TestAnd64Sve(t *testing.T) {
	if !cpu.CanUseSVE() {
		fmt.Println("no sve support")
		t.Skip()
	}
	length := 8

	// Allocate input vectors
	v1 := Uint64s{0xFF, 0x0F, 0xF0, 0x00, 0x55, 0xFF, 0x0F, 0x0F}
	v2 := Uint64s{0x00, 0x00, 0x0F, 0xAA, 0xFF, 0x00, 0x0F, 0x00}

	mask := FindNonZero64bitWordsSve(v2)
	out := make(Uint64s, length)
	out.AndMasked64bitSve(v1, v2, mask)
	fmt.Println(out)

	expected := make(Uint64s, length)
	for i := 0; i < length; i++ {
		expected[i] = v1[i] & v2[i]
	}

	// Validate result
	assert.Equal(t, expected, out, "AndMasked output mismatch")

}

func TestAndSpanned(t *testing.T) {

	length := 8

	// Allocate input vectors
	v1 := Uint64s{0xFF, 0x0F, 0xF0, 0x00, 0x55, 0xFF, 0x0F, 0x0F}
	v2 := Uint64s{0x00, 0x00, 0x0F, 0xAA, 0xFF, 0x00, 0x0F, 0x00}

	mask := FindNonZeroSpans(v2)
	out := make(Uint64s, length)
	out.AndSpanNeon(v1, v2, mask)

	expected := make(Uint64s, length)
	for i := 0; i < length; i++ {
		expected[i] = v1[i] & v2[i]
	}

	// Validate result
	assert.Equal(t, expected, out, "AndSpanned output mismatch")

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

func makeZeroPrefixedVector(n int, nonZeroFraction float64) Uint64s {
	vec := make(Uint64s, n)
	for i := range vec {
		if float64(i)/float64(len(vec)) < nonZeroFraction {
			vec[i] = rand.Uint64()
		}
	}
	return vec
}

// ///////
const N = 1 << 10
const sparsity = 0.5 // % of elements are non-zero
func BenchmarkAndProd(b *testing.B) {
	a := makeSparseVector(N, sparsity)
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
	a := makeSparseVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndSVE(a, bv)
	}
}

//func BenchmarkAndMaskedSparseNaive(b *testing.B) {
//	//a := makeSparseVector(N, sparsity)
//	a := makeZeroPrefixedVector(N, sparsity)
//	bv := makeSparseVector(N, 1.0) // fully populated
//	c := make(Uint64s, N)
//
//	mask := FindNonZeroWords(a)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		c.AndMaskedNaive(a, bv, mask)
//	}
//}

//func BenchmarkAndMaskedSparseNeon(b *testing.B) {
//	a := makeSparseVector(N, sparsity)
//	bv := makeSparseVector(N, 1.0) // fully populated
//	c := make(Uint64s, N)
//
//	mask := FindNonZeroWords(a)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		c.AndMaskedNeon(a, bv, mask)
//		//c.And(a, bv)
//	}
//}

func BenchmarkAndMaskedSparse128bitNeon(b *testing.B) {
	a := makeSparseVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	mask := FindNonZero128bitWords(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndMasked128bitNeon(a, bv, mask)
	}
}

func BenchmarkAndMaskedSparse128bitUnrolledNeon(b *testing.B) {
	a := makeSparseVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	mask := FindNonZero128bitWords(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndMasked128bitUnrolledNeon(a, bv, mask)
	}
}

func BenchmarkAndSpannedNeon(b *testing.B) {
	a := makeSparseVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	mask := FindNonZeroSpans(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndSpanNeon(a, bv, mask)
	}
}

func BenchmarkAndMaskedSparse64bitSve(b *testing.B) {
	if !cpu.CanUseSVE() {
		fmt.Println("no sve support")
		b.Skip()
	}

	a := makeSparseVector(N, sparsity)
	bv := makeSparseVector(N, 1.0) // fully populated
	c := make(Uint64s, N)

	mask := FindNonZero64bitWordsSve(a)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.AndMasked64bitSve(a, bv, mask)
	}
}

//func BenchmarkAndStridedNeon(b *testing.B) {
//	v1 := makeZeroPrefixedVector(N, sparsity)
//	v2 := makeSparseVector(N, 1.0) // fully populated
//	out := make(Uint64s, N)
//
//	mask := FindNonZeroWords(v1)
//	strides := Strides{}
//	strides.ensureStrides(v1)
//	strides.fromMask(mask)
//
//	b.ResetTimer()
//	for i := 0; i < b.N; i++ {
//		out.AndStridedNeon(v1, v2, strides)
//	}
//}
