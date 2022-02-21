package bits

import (
	"github.com/stretchr/testify/assert"
	"math/bits"
	"math/rand"
	"testing"
	"time"
)

var d []uint64

const L = 64

func init() {
	d = make([]uint64, L)
	for i := 0; i < L; i++ {
		d[i] = 0xffffffffffffffff
	}
}

func TestCount(t *testing.T) {

	var testCases = []struct {
		description string
		size        int
	}{
		{
			description: "4096 bit",
			size:        64,
		},
		{
			description: "256 bit",
			size:        4,
		},
		{
			description: "1024 bit",
			size:        16,
		},
		{
			description: "320 bit",
			size:        5,
		},
		{
			description: "128 bit",
			size:        2,
		},
		{
			description: "75x64 bit",
			size:        75,
		},
		{
			description: "1232x64 bit",
			size:        1232,
		},
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, testCase := range testCases {
		var vec = make([]uint64, testCase.size)
		expect := 0
		for i := range vec {
			vec[i] = rnd.Uint64()
			expect += bits.OnesCount(uint(vec[i]))
		}
		actual := Count(vec)
		assert.EqualValues(t, expect, actual, testCase.description)
	}
}

func BenchmarkCountOptimized(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Count(d)
	}
}

func BenchmarkNaiveCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		count(d)
	}
}
