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

func TestHmaxFloat32(t *testing.T) {
	max := HmaxFloat32(float32HsumData[:])
	expMax := hmaxFloat32(float32HsumData[:])
	assert.EqualValues(t, expMax, max, "TestHmaxFloat32")
}

func TestHminFloat32(t *testing.T) {
	min := HminFloat32(float32HsumData[:])
	expMin := hminFloat32(float32HsumData[:])
	assert.EqualValues(t, expMin, min, "TestHminFloat32")
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

func BenchmarkHmaxFloat32Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hmaxFloat32(float32HsumData[:])
	}
}
func BenchmarkHmaxFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HmaxFloat32(float32HsumData[:])
	}
}

func BenchmarkHminFloat32Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hminFloat32(float32HsumData[:])
	}
}
func BenchmarkHminFloat32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HminFloat32(float32HsumData[:])
	}
}
