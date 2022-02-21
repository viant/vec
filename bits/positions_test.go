package bits

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/bits"
	"math/rand"
	"testing"
	"time"
)

var benchInput = uint64(0xf55555555555555f)
var out Positions
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func TestPositions(t *testing.T) {
	var expectedOut, actualOut Positions

	for i := 0; i < 100; i++ {
		input := rnd.Uint64()
		expectedCount := expectedOut.populate(input)
		actualCount := actualOut.Populate(input)
		expect := bits.OnesCount(uint(input))
		assert.EqualValues(t, expect, actualCount)
		for k := 0; k < int(actualCount); k++ {
			assert.True(t, input&(uint64(1)<<expectedOut[k]) != 0, fmt.Sprintf("should have position set at %v %b", expectedOut[k], input))
		}
		assert.EqualValues(t, expectedOut[0:expectedCount], actualOut[0:actualCount], "TestPositions")

	}
}

func BenchmarkPositionsNaive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		out.populate(benchInput)
	}
}

func BenchmarkPositions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		out.Populate(benchInput)
	}
}
