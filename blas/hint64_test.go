package blas

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//const nHsum = 255

var int64HsumData [nHsum]int64

func init() {
	for i := 0; i < len(int64HsumData); i++ {
		int64HsumData[i] = int64(i)
	}
}

func TestHsumInt64(t *testing.T) {
	sum := HsumInt64(int64HsumData[:])
	expSum := hsumInt64(int64HsumData[:])
	assert.EqualValues(t, expSum, sum, "TestHsumInt64")
}

func TestHmaxInt64(t *testing.T) {
	max := HmaxInt64(int64HsumData[:])
	expMax := hmaxInt64(int64HsumData[:])
	assert.EqualValues(t, expMax, max, "TestHmaxInt64")
}

func TestHminInt64(t *testing.T) {
	min := HminInt64(int64HsumData[:])
	expMin := hminInt64(int64HsumData[:])
	assert.EqualValues(t, expMin, min, "TestHminInt64")
}

func BenchmarkHsumInt64Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hsumInt64(int64HsumData[:])
	}
}
func BenchmarkHsumInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HsumInt64(int64HsumData[:])
	}
}

func BenchmarkHmaxInt64Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hmaxInt64(int64HsumData[:])
	}
}
func BenchmarkHmaxInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HmaxInt64(int64HsumData[:])
	}
}

func BenchmarkHminInt64Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = hminInt64(int64HsumData[:])
	}
}
func BenchmarkHminInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = HminInt64(int64HsumData[:])
	}
}
