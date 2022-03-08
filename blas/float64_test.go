package blas

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const nFloats64 = 255

var float64Data1 [nFloats64]float64
var expectedAddFloat64s Float64s
var expectedSubFloat64s Float64s
var expectedMulFloat64s Float64s
var expectedDivFloat64s Float64s
var float64Data2 [nFloats64]float64

func init() {
	for i := 0; i < len(float64Data1); i++ {
		float64Data1[i] = float64(i)
		float64Data2[i] = float64(2 * (i + 1))
	}
	expectedAddFloat64s = make([]float64, len(float64Data1))
	expectedSubFloat64s = make([]float64, len(float64Data1))
	expectedMulFloat64s = make([]float64, len(float64Data1))
	expectedDivFloat64s = make([]float64, len(float64Data1))
	expectedAddFloat64s.addFloat64(float64Data1[:], float64Data2[:])
	expectedSubFloat64s.subFloat64(float64Data1[:], float64Data2[:])
	expectedMulFloat64s.mulFloat64(float64Data1[:], float64Data2[:])
	expectedDivFloat64s.divFloat64(float64Data1[:], float64Data2[:])
	for i, v := range expectedDivFloat64s {
		expectedDivFloat64s[i] = float64(math.Round(float64(v+0.01) * 10))
	}

}

func TestAddlFloat64(t *testing.T) {
	out := Float64s(make([]float64, len(float64Data1)))
	out.AddFloat64(float64Data1[:], float64Data2[:])
	assert.EqualValues(t, expectedAddFloat64s, out, "TestAddFloat64")
}
func TestSubFloat64(t *testing.T) {
	out := Float64s(make([]float64, len(float64Data1)))
	out.SubFloat64(float64Data1[:], float64Data2[:])
	assert.EqualValues(t, expectedSubFloat64s, out, "TestSubFloat64")
}
func TestMulFloat64(t *testing.T) {
	out := Float64s(make([]float64, len(float64Data1)))
	out.MulFloat64(float64Data1[:], float64Data2[:])
	assert.EqualValues(t, expectedMulFloat64s, out, "TestMulFloat64")
}
func TestDivFloat64(t *testing.T) {
	out := Float64s(make([]float64, len(float64Data1)))
	out.DivFloat64(float64Data1[:], float64Data2[:])
	for i, v := range out {
		out[i] = float64(math.Round(float64(v+0.01) * 10))
	}
	assert.EqualValues(t, expectedDivFloat64s, out, "TestDivFloat64")
}

func BenchmarkAddFloat64Naive(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.addFloat64(float64Data1[:], float64Data2[:])
	}
}
func BenchmarkAddFloat64(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.AddFloat64(float64Data1[:], float64Data2[:])
	}
}

func BenchmarkSubFloat64Naive(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.subFloat64(float64Data1[:], float64Data2[:])
	}
}
func BenchmarkSubFloat64(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.SubFloat64(float64Data1[:], float64Data2[:])
	}
}

func BenchmarkMulFloat64Naive(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.mulFloat64(float64Data1[:], float64Data2[:])
	}
}
func BenchmarkMulFloat64(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.MulFloat64(float64Data1[:], float64Data2[:])
	}
}

func BenchmarkDivFloat64Naive(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.divFloat64(float64Data1[:], float64Data2[:])
	}
}
func BenchmarkDivFloat64(b *testing.B) {
	out := Float64s(make([]float64, len(float64Data1)))
	for i := 0; i < b.N; i++ {
		out.DivFloat64(float64Data1[:], float64Data2[:])
	}
}
