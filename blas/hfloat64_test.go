package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//const nHsum = 255

var float64HsumData [nHsum]float64

func init() {
	for i := 0; i < len(float64HsumData); i++ {
		float64HsumData[i] = float64(i)
	}
}

func TestHsumFloat64(t *testing.T) {
	sum := HsumFloat64(float64HsumData[:])
	expSum := hsumFloat64(float64HsumData[:])
	assert.EqualValues(t, expSum, sum, "TestHsumFloat64")
}

func TestHmaxFloat64(t *testing.T) {
	max := HmaxFloat64(float64HsumData[:])
	expMax := hmaxFloat64(float64HsumData[:])
	assert.EqualValues(t, expMax, max, "TestHmaxFloat64")
}

func TestHminFloat64(t *testing.T) {
	min := HminFloat64(float64HsumData[:])
	expMin := hminFloat64(float64HsumData[:])
	assert.EqualValues(t, expMin, min, "TestHminFloat64")
}

func BenchmarkHsumFloat64Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hsumFloat64(float64HsumData[:])
	}
}
func BenchmarkHsumFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HsumFloat64(float64HsumData[:])
	}
}

func BenchmarkHmaxFloat64Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hmaxFloat64(float64HsumData[:])
	}
}
func BenchmarkHmaxFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HmaxFloat64(float64HsumData[:])
	}
}

func BenchmarkHminFloat64Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hminFloat64(float64HsumData[:])
	}
}
func BenchmarkHminFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HminFloat64(float64HsumData[:])
	}
}
