package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var data [128]uint32

func init() {
	for i := 0; i < len(data); i++ {
		data[i] = uint32(i)
	}
}

func TestBitwidth(t *testing.T) {
	actual := MaxBits(data[:])
	assert.EqualValues(t, 7, actual, "TestBitwidth")
}

func BenchmarkBitwidthNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = maxBits(data[:])
	}
}
func BenchmarkBitwidth(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MaxBits(data[:])
	}
}
