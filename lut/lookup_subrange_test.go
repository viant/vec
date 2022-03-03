package lut

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LookupSubrange(t *testing.T) {
	var testCases = []struct {
		description    string
		in             []byte
		rangeUpper     []byte
		table          []byte
		lookupSubrange func(input, range_upper, table []byte, output *[]byte)
		expect         []byte
	}{
		{
			description:    "naive 16 byte lookup",
			rangeUpper:     []byte{5, 10, 15, 250},
			table:          []byte{0, 42, 43, 240, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			in:             []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255},
			lookupSubrange: lookupSubrange,
			expect:         []byte{0, 0, 0, 0, 0, 42, 42, 42, 42, 42, 43, 43, 43, 43, 43, 211},
		},
		{
			description:    "vector 16 byte lookup",
			rangeUpper:     []byte{5, 10, 15, 250},
			table:          []byte{0, 42, 43, 240, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			in:             []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255},
			lookupSubrange: LookupSubrange,
			expect:         []byte{0, 0, 0, 0, 0, 42, 42, 42, 42, 42, 43, 43, 43, 43, 43, 211},
		},
		{
			description:    "vector 32 byte lookup",
			rangeUpper:     []byte{10, 20, 100, 250},
			table:          []byte{0, 42, 43, 240, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			in:             []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255},
			lookupSubrange: LookupSubrange,
			expect:         []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 42, 42, 42, 42, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42, 42, 42, 42, 42, 211},
		},
	}

	for _, testCase := range testCases {
		fmt.Println(testCase.description)
		output := make([]byte, len(testCase.in))
		testCase.lookupSubrange(testCase.in, testCase.rangeUpper, testCase.table, &output)
		assert.EqualValues(t, testCase.expect, output, testCase.description)
	}
}

var benchRangeUpper = []byte{10, 20, 100, 250}
var benchTable = []byte{0, 42, 43, 240, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var benchIn = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255}

func Benchmark_LookupSubrangeNaive(b *testing.B) {
	output := make([]byte, len(benchIn))
	for i := 0; i < b.N; i++ {
		lookupSubrange(benchIn, benchRangeUpper, benchIn, &output)
	}
}

func Benchmark_LookupSubrange(b *testing.B) {
	output := make([]byte, len(benchIn))
	for i := 0; i < b.N; i++ {
		LookupSubrange(benchIn, benchRangeUpper, benchIn, &output)
	}
}
