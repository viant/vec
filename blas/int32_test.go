package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const nInts = 255

var int32Data1 = make([]int32, nInts)
var int32Data2 = make([]int32, nInts)
var expectedIncInt32s Int32s
var expectedAddInt32s Int32s
var expectedSubInt32s Int32s
var expectedMulInt32s Int32s

func init() {
	for i := 0; i < len(int32Data1); i++ {
		int32Data1[i] = int32(i)
		int32Data2[i] = int32(2 * (i + 1))
	}
	expectedIncInt32s = make([]int32, len(int32Data1))
	expectedAddInt32s = make([]int32, len(int32Data1))
	expectedSubInt32s = make([]int32, len(int32Data1))
	expectedMulInt32s = make([]int32, len(int32Data1))
	copy(expectedIncInt32s, int32Data1)
	expectedIncInt32s.inc(42)
	expectedAddInt32s.add(int32Data1[:], int32Data2[:])
	expectedSubInt32s.sub(int32Data1[:], int32Data2[:])
	expectedMulInt32s.mul(int32Data1[:], int32Data2[:])
}

func TestIncInt32(t *testing.T) {
	out := Int32s(make([]int32, len(int32Data1)))
	copy(out, int32Data1)
	out.Inc(42)
	assert.EqualValues(t, expectedIncInt32s, out, "TestIncInt32")
}
func TestAddInt32(t *testing.T) {
	out := Int32s(make([]int32, len(int32Data1)))
	out.Add(int32Data1[:], int32Data2[:])
	assert.EqualValues(t, expectedAddInt32s, out, "TestAddInt32")
}
func TestSubInt32(t *testing.T) {
	out := Int32s(make([]int32, len(int32Data1)))
	out.Sub(int32Data1[:], int32Data2[:])
	assert.EqualValues(t, expectedSubInt32s, out, "TestSubInt32")
}
func TestMulInt32(t *testing.T) {
	out := Int32s(make([]int32, len(int32Data1)))
	out.Mul(int32Data1[:], int32Data2[:])
	assert.EqualValues(t, expectedMulInt32s, out, "TestMulInt32")
}

func BenchmarkIncInt32Naive(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.inc(42)
	}
}
func BenchmarkIncInt32(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.Inc(42)
	}
}

func BenchmarkAddInt32Naive(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.add(int32Data1[:], int32Data2[:])
	}
}
func BenchmarkAddInt32(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.Add(int32Data1[:], int32Data2[:])
	}
}

func BenchmarkSubInt32Naive(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.sub(int32Data1[:], int32Data2[:])
	}
}
func BenchmarkSubInt32(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.Sub(int32Data1[:], int32Data2[:])
	}
}

func BenchmarkMulInt32Naive(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.mul(int32Data1[:], int32Data2[:])
	}
}
func BenchmarkMulInt32(b *testing.B) {
	out := Int32s(make([]int32, len(int32Data1)))
	for i := 0; i < b.N; i++ {
		out.Mul(int32Data1[:], int32Data2[:])
	}
}
