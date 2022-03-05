package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var shrunkValue = uint64(0x55555555aaaaaaaa)
var expectedExpansion = []uint64{
	0xff00ff00ff00ff00, 0xff00ff00ff00ff00, 0xff00ff00ff00ff00, 0xff00ff00ff00ff00,
	0x00ff00ff00ff00ff, 0x00ff00ff00ff00ff, 0x00ff00ff00ff00ff, 0x00ff00ff00ff00ff,
}

func TestExpand(t *testing.T) {
	out := Uint64s(make([]uint64, 8))
	out.Expand(shrunkValue)
	assert.EqualValues(t, Uint64s(expectedExpansion), out, "test")
}

func BenchmarkNaiveExpand(b *testing.B) {
	b.ReportAllocs()
	out := Uint64s(make([]uint64, 8))
	for i := 0; i < b.N; i++ {
		out.expand(shrunkValue)
	}
}

func BenchmarkExpand(b *testing.B) {
	b.ReportAllocs()
	out := Uint64s(make([]uint64, 8))
	for i := 0; i < b.N; i++ {
		out.Expand(shrunkValue)
	}
}
