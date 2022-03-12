package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//const nInts = 255

var int64Data1 [nInts]int64
var expectedAddInt64s Int64s
var expectedSubInt64s Int64s
var expectedMulInt64s Int64s
var expectedDivInt64s Int64s
var int64Data2 [nInts]int64

func init() {
	for i := 0; i < len(int64Data1); i++ {
		int64Data1[i] = int64(i)
		int64Data2[i] = int64(2 * (i + 1))
	}
	expectedAddInt64s = make([]int64, len(int64Data1))
	expectedSubInt64s = make([]int64, len(int64Data1))
	expectedMulInt64s = make([]int64, len(int64Data1))
	expectedDivInt64s = make([]int64, len(int64Data1))
	expectedAddInt64s.addInt64(int64Data1[:], int64Data2[:])
	expectedSubInt64s.subInt64(int64Data1[:], int64Data2[:])
	expectedMulInt64s.mulInt64(int64Data1[:], int64Data2[:])
}

func TestAddlInt64(t *testing.T) {
	out := Int64s(make([]int64, len(int64Data1)))
	out.AddInt64(int64Data1[:], int64Data2[:])
	assert.EqualValues(t, expectedAddInt64s, out, "TestAddInt64")
}
func TestSubInt64(t *testing.T) {
	out := Int64s(make([]int64, len(int64Data1)))
	out.SubInt64(int64Data1[:], int64Data2[:])
	assert.EqualValues(t, expectedSubInt64s, out, "TestSubInt64")
}
func TestMulInt64(t *testing.T) {
	out := Int64s(make([]int64, len(int64Data1)))
	out.MulInt64(int64Data1[:], int64Data2[:])
	assert.EqualValues(t, expectedMulInt64s, out, "TestMulInt64")
}

func BenchmarkAddInt64Naive(b *testing.B) {
	out := Int64s(make([]int64, len(int64Data1)))
	for i := 0; i < b.N; i++ {
		out.addInt64(int64Data1[:], int64Data2[:])
	}
}
func BenchmarkAddInt64(b *testing.B) {
	out := Int64s(make([]int64, len(int64Data1)))
	for i := 0; i < b.N; i++ {
		out.AddInt64(int64Data1[:], int64Data2[:])
	}
}

func BenchmarkSubInt64Naive(b *testing.B) {
	out := Int64s(make([]int64, len(int64Data1)))
	for i := 0; i < b.N; i++ {
		out.subInt64(int64Data1[:], int64Data2[:])
	}
}
func BenchmarkSubInt64(b *testing.B) {
	out := Int64s(make([]int64, len(int64Data1)))
	for i := 0; i < b.N; i++ {
		out.SubInt64(int64Data1[:], int64Data2[:])
	}
}

func BenchmarkMulInt64Naive(b *testing.B) {
	out := Int64s(make([]int64, len(int64Data1)))
	for i := 0; i < b.N; i++ {
		out.mulInt64(int64Data1[:], int64Data2[:])
	}
}
func BenchmarkMulInt64(b *testing.B) {
	out := Int64s(make([]int64, len(int64Data1)))
	for i := 0; i < b.N; i++ {
		out.MulInt64(int64Data1[:], int64Data2[:])
	}
}
