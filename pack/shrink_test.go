package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var ints = []uint64{
	0x8000800080008000, 0x8000800080008000, 0x8000800080008000, 0x8000800080008000,
	0x0080008000800080, 0x0080008000800080, 0x0080008000800080, 0x0080008000800080,
}
var expectedShrink = 0x55555555aaaaaaaa

func TestShrink(t *testing.T) {
	var shrunkValue = Uint64(0)
	shrunkValue.Shrink(ints)
	assert.EqualValues(t, expectedShrink, shrunkValue, "test")
}

func TestShrinkValue(t *testing.T) {
	var shrunkValue = Uint64(0)
	shrunkValue.ShrinkValue(0x80, ints)
	assert.EqualValues(t, expectedShrink, shrunkValue, "test")
}

func BenchmarkNaiveShrink(b *testing.B) {
	var shrunkValue Uint64 = 0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		shrunkValue.shrink(ints)
	}
}

func BenchmarkShrink(b *testing.B) {
	b.ReportAllocs()
	var shrunkValue Uint64 = 0
	for i := 0; i < b.N; i++ {
		shrunkValue.Shrink(ints)
	}
}

func BenchmarkNaiveShrinkValue(b *testing.B) {
	var shrunkValue Uint64 = 0
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		shrunkValue.shrinkValue(0, ints)
	}
}

func BenchmarkShrinkValue(b *testing.B) {
	b.ReportAllocs()
	var shrunkValue Uint64 = 0
	for i := 0; i < b.N; i++ {
		shrunkValue.ShrinkValue(0x80, ints)
	}
}
