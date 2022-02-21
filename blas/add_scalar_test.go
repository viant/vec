package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data [255]int32
var expected Int32s

func init() {
	for i := 0; i < len(data); i++ {
		data[i] = int32(i)
	}
	expected = Int32s(make([]int32, len(data)))
	expected.addScalar(data[:], -4)
}

func TestAddScalar(t *testing.T) {
	out := Int32s(make([]int32, len(data)))
	out.AddScalar(data[:], -4)
	assert.EqualValues(t, expected, out, "TestAddScalar")
}
func TestAddScalarInPlace(t *testing.T) {
	out := Int32s(make([]int32, len(data)))
	copy(out, data[:])
	out.AddScalar(out, -4)
	assert.EqualValues(t, expected, out, "TestAddScalar in-place")
}

func BenchmarkAddScalarNaive(b *testing.B) {
	out := Int32s(make([]int32, len(data)))
	for i := 0; i < b.N; i++ {
		out.addScalar(data[:], -4)
	}
}
func BenchmarkAddScalar(b *testing.B) {
	out := Int32s(make([]int32, len(data)))
	for i := 0; i < b.N; i++ {
		out.AddScalar(data[:], -4)
	}
}
