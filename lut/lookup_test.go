package lut

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_Lookup(t *testing.T) {
	var testCases = []struct {
		description string
		size        int
		lookup      map[uint8]uint8
		expect      []uint8
	}{
		{
			description: "64 bits lookup",
			size:        64,
			lookup: map[uint8]uint8{
				12: 1,
				15: 2,
				6:  62,
			},
			expect: []byte{
				62, 1, 2,
			},
		},
		{
			description: "128 bits lookup",
			size:        128,
			lookup: map[uint8]uint8{
				12: 1,
				15: 124,
				6:  62,
			},
			expect: []byte{
				62, 1, 124,
			},
		},
	}

	for _, testCase := range testCases {
		var table = make([]byte, testCase.size)
		var input = make([]byte, testCase.size)
		var output = make([]byte, testCase.size)
		i := 0
		var keys = make([]int, 0)
		for k := range testCase.lookup {
			keys = append(keys, int(k))
		}
		sort.Ints(keys)
		for _, k := range keys {
			v := testCase.lookup[uint8(k)]
			table[k] = v
			input[i] = uint8(k)
			i++
		}
		Lookup(input, table, &output)
		assert.EqualValues(t, testCase.expect, output[:len(testCase.lookup)], testCase.description)
	}
}

var bmTable = make([]byte, 256)
var bmInput = make([]byte, 256)
var bmOutput = make([]byte, 256)
var bmExpect = make([]byte, 256)

func init() {
	bmTable[1] = 253
	bmTable[24] = 21
	bmTable[102] = 1

	bmInput[0] = 1
	bmInput[1] = 102
	bmInput[2] = 24
	bmInput[3] = 253

	bmExpect[0] = 253
	bmExpect[1] = 1
	bmExpect[2] = 21
	bmExpect[3] = 0

}

func Benchmark_Lookup(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lookup(bmInput, bmTable, &bmOutput)
	}
	assert.EqualValues(b, bmExpect, bmOutput)
}

func Benchmark_Lookup_Naive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lookup(bmInput, bmTable, &bmOutput)
	}
	assert.EqualValues(b, bmExpect, bmOutput)
}
