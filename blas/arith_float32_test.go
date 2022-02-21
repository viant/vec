package blas

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const nFloats = 255

var float32Data1 [nFloats]float32
var expectedAddFloat32s Float32s
var expectedSubFloat32s Float32s
var expectedMulFloat32s Float32s
var expectedDivFloat32s Float32s
var float32Data2 [nFloats]float32

func init() {
	for i := 0; i < len(float32Data1); i++ {
		float32Data1[i] = float32(i)
		float32Data2[i] = float32(2 * (i + 1))
	}
	expectedAddFloat32s = make([]float32, len(float32Data1))
	expectedSubFloat32s = make([]float32, len(float32Data1))
	expectedMulFloat32s = make([]float32, len(float32Data1))
	expectedDivFloat32s = make([]float32, len(float32Data1))
	expectedAddFloat32s.addFloat32(float32Data1[:], float32Data2[:])
	expectedSubFloat32s.subFloat32(float32Data1[:], float32Data2[:])
	expectedMulFloat32s.mulFloat32(float32Data1[:], float32Data2[:])
	expectedDivFloat32s.divFloat32(float32Data1[:], float32Data2[:])
	for i, v := range expectedDivFloat32s {
		expectedDivFloat32s[i] = float32(math.Round(float64(v+0.01) * 10))
	}

}

func TestAddlFloat32(t *testing.T) {
	out := Float32s(make([]float32, len(float32Data1)))
	out.AddFloat32(float32Data1[:], float32Data2[:])
	assert.EqualValues(t, expectedAddFloat32s, out, "TestAddFloat32")
}
func TestSubFloat32(t *testing.T) {
	out := Float32s(make([]float32, len(float32Data1)))
	out.SubFloat32(float32Data1[:], float32Data2[:])
	assert.EqualValues(t, expectedSubFloat32s, out, "TestSubFloat32")
}
func TestMulFloat32(t *testing.T) {
	out := Float32s(make([]float32, len(float32Data1)))
	out.MulFloat32(float32Data1[:], float32Data2[:])
	assert.EqualValues(t, expectedMulFloat32s, out, "TestMulFloat32")
}
func TestDivFloat32(t *testing.T) {
	out := Float32s(make([]float32, len(float32Data1)))
	out.DivFloat32(float32Data1[:], float32Data2[:])
	for i, v := range out {
		out[i] = float32(math.Round(float64(v+0.01) * 10))
	}
	assert.EqualValues(t, expectedDivFloat32s, out, "TestDivFloat32")
}

func BenchmarkAddFloat32Naive(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.addFloat32(float32Data1[:], float32Data2[:])
	}
}
func BenchmarkAddFloat32(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.AddFloat32(float32Data1[:], float32Data2[:])
	}
}

func BenchmarkSubFloat32Naive(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.subFloat32(float32Data1[:], float32Data2[:])
	}
}
func BenchmarkSubFloat32(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.SubFloat32(float32Data1[:], float32Data2[:])
	}
}

func BenchmarkMulFloat32Naive(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.mulFloat32(float32Data1[:], float32Data2[:])
	}
}
func BenchmarkMulFloat32(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.MulFloat32(float32Data1[:], float32Data2[:])
	}
}

func BenchmarkDivFloat32Naive(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.divFloat32(float32Data1[:], float32Data2[:])
	}
}
func BenchmarkDivFloat32(b *testing.B) {
	out := Float32s(make([]float32, len(float32Data1)))
	for i := 0; i < b.N; i++ {
		out.DivFloat32(float32Data1[:], float32Data2[:])
	}
}
