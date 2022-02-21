package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const nHsum = 255

var float32HsumData [nHsum]float32

func init() {
	for i := 0; i < len(float32HsumData); i++ {
		float32HsumData[i] = float32(i)
	}
}

func TestHsumFloat32(t *testing.T) {
	sum := HsumFloat32(float32HsumData[:])
	expSum := hsumFloat32(float32HsumData[:])
	assert.EqualValues(t, expSum, sum, "TestHsumFloat32")
}

func BenchmarkHsumFloat32Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hsumFloat32(float32HsumData[:])
	}
}

func BenchmarkHsumFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HsumFloat32(float32HsumData[:])
	}
}
