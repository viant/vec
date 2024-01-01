package bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const lLsb = 64

var dLsb []uint64

func init() {
	dLsb = make([]uint64, lLsb)
	for i := 0; i < lLsb-1; i++ {
		dLsb[i] = 0x000000000000000
	}
	dLsb[lLsb-1] = 0x8000000000000000
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
		actual := Lsb(dLsb)
		fmt.Println(actual)
		assert.EqualValues(t, testCase.pos, actual, testCase.description)
	}
}

func BenchmarkLsbNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lsb(dLsb)
	}
}

func BenchmarkLsb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lsb(dLsb)
	}
}
