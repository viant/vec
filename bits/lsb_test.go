package bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const L1 = 64

func init() {
	d = make([]uint64, L1)
	for i := 0; i < L1-1; i++ {
		d[i] = 0x000000000000000
	}
	d[L1-1] = 0x8000000000000000
}

func TestLsb(t *testing.T) {

	var testCases = []struct {
		description string
		pos         int
	}{
		{
			description: "first test",
			pos:         4095,
		},
	}

	for _, testCase := range testCases {
		actual := Lsb(d)
		fmt.Println(actual)
		assert.EqualValues(t, testCase.pos, actual, testCase.description)
		assert.EqualValues(t, testCase.pos, actual, testCase.description)
	}
}

func BenchmarkLsbNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lsb(d)
	}
}

func BenchmarkLsb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lsb(d)
	}
}
