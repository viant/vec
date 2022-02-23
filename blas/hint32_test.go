package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//const nHsum = 255

var int32HsumData [nHsum]int32

func init() {
	for i := 0; i < len(int32HsumData); i++ {
		int32HsumData[i] = int32(i)
	}
}

func TestHsumInt32(t *testing.T) {
	sum := HsumInt32(int32HsumData[:])
	expSum := hsumInt32(int32HsumData[:])
	assert.EqualValues(t, expSum, sum, "TestHsumInt32")
}

func TestHmaxInt32(t *testing.T) {
	max := HmaxInt32(int32HsumData[:])
	expMax := hmaxInt32(int32HsumData[:])
	assert.EqualValues(t, expMax, max, "TestHmaxInt32")
}

func TestHminInt32(t *testing.T) {
	min := HminInt32(int32HsumData[:])
	expMin := hminInt32(int32HsumData[:])
	assert.EqualValues(t, expMin, min, "TestHminInt32")
}

func BenchmarkHsumInt32Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hsumInt32(int32HsumData[:])
	}
}
func BenchmarkHsumInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HsumInt32(int32HsumData[:])
	}
}

func BenchmarkHmaxInt32Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hmaxInt32(int32HsumData[:])
	}
}
func BenchmarkHmaxInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HmaxInt32(int32HsumData[:])
	}
}

func BenchmarkHminInt32Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hminInt32(int32HsumData[:])
	}
}
func BenchmarkHminInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HminInt32(int32HsumData[:])
	}
}
