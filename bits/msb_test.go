package bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var dMsb []uint64

const lMsb = 64

func init() {
	dMsb = make([]uint64, lMsb)
	for i := 0; i < lMsb-1; i++ {
		dMsb[i] = 0x000000000000000
	}
	dMsb[0] = 0x8000000000000000
}

func TestMsb(t *testing.T) {

	var testCases = []struct {
		description string
		pos         int
	}{
		{
			description: "first test",
			pos:         63,
		},
	}

	for _, testCase := range testCases {
		actual := Msb(dMsb)
		fmt.Println(actual)
		assert.EqualValues(t, testCase.pos, actual, testCase.description)
	}
}

func BenchmarkMsbNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		msb(dMsb)
	}
}

func BenchmarkMsb(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Msb(dMsb)
	}
}
