package pack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var deltaData [128]int32
var expectedDeltas Int32s

func init() {
	for i := 0; i < len(deltaData); i++ {
		deltaData[i] = int32(len(deltaData) - i)
	}
	expectedDeltas = make([]int32, len(deltaData))
	expectedDeltas.deltas(deltaData[:])
}

func TestDeltas(t *testing.T) {
	deltas := Int32s(make([]int32, len(deltaData)))
	deltas.Deltas(deltaData[:])
	assert.EqualValues(t, expectedDeltas, deltas, "Deltas test")
}
func TestRecoverDeltas(t *testing.T) {
	deltas := Int32s(make([]int32, len(deltaData)))
	deltas.Deltas(deltaData[:])
	recoveredData := Int32s(make([]int32, len(deltaData)))
	recoveredData.RecoverDeltas(deltas)
	assert.EqualValues(t, Int32s(deltaData[:]), recoveredData, "Recover from deltas test")
}
func TestDeltasInPlace(t *testing.T) {
	data := Int32s(make([]int32, len(deltaData)))
	copy(data, deltaData[:])
	data.Deltas(data)
	data.RecoverDeltas(data)
	assert.EqualValues(t, deltaData[:], data, "Deltas in-place test")
}

func BenchmarkDeltasNaive(b *testing.B) {
	deltas := Int32s(make([]int32, len(deltaData)))
	for i := 0; i < b.N; i++ {
		deltas.deltas(deltaData[:])
	}
}
func BenchmarkRecoverDeltasNaive(b *testing.B) {
	deltas := Int32s(make([]int32, len(deltaData)))
	deltas.Deltas(deltaData[:])
	recoveredData := Int32s(make([]int32, len(deltaData)))
	for i := 0; i < b.N; i++ {
		recoveredData.recoverDeltas(deltas)
	}
}

func BenchmarkDeltas(b *testing.B) {
	deltas := Int32s(make([]int32, len(deltaData)))
	for i := 0; i < b.N; i++ {
		deltas.Deltas(deltaData[:])
	}
}
func BenchmarkRecoverDeltas(b *testing.B) {
	deltas := Int32s(make([]int32, len(deltaData)))
	deltas.Deltas(deltaData[:])
	recoveredData := Int32s(make([]int32, len(deltaData)))
	for i := 0; i < b.N; i++ {
		recoveredData.RecoverDeltas(deltas)
	}
}
