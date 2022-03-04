package pack

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var benchData1024 []uint32

func init() {
	rand.Seed(time.Now().Unix())
	tmp := rand.Perm(1024)
	benchData1024 = make([]uint32, 1024)
	for i := range tmp {
		benchData1024[i] = uint32(tmp[i])
	}
}

func TestTurboUnpackBits(t *testing.T) {
	data := []uint32{20124, 8831, 2575, 1977, 15, 42441, 302690, 871222, 323452, 424532, 29434, 939141, 4244, 324314, 13, 1, 874255, 20124, 8831, 2575, 1977, 15, 42441, 302690, 871222, 323452, 424532, 29434, 939141, 4244, 324314, 13, 1, 874255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42441, 302690, 871222, 323452, 424532, 29434, 939141, 4244, 324314, 13, 1, 874255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	bits := MaxBits(data)
	packedData := Uint32s(make([]uint32, PackedUint32Count(uint32(len(data)), bits)))
	packedData.PackBits(data, int(bits))

	unpackedData := make([]uint32, len(data))
	Uint32s(unpackedData).UnpackBits(packedData, int(bits))

	assert.EqualValues(t, data, unpackedData, "TestPackBits")
}

///
func BenchmarkPackBitsNaive128(b *testing.B) {
	packedData := Uint32s(make([]uint32, 128))
	for i := 0; i < b.N; i++ {
		packedData.packBits(benchData1024[:len(packedData)], 21)
	}
}
func BenchmarkPackBitsNaive256(b *testing.B) {
	packedData := Uint32s(make([]uint32, 256))
	for i := 0; i < b.N; i++ {
		packedData.packBits(benchData1024[:len(packedData)], 21)
	}
}
func BenchmarkPackBitsNaive1024(b *testing.B) {
	packedData := Uint32s(make([]uint32, 1024))
	for i := 0; i < b.N; i++ {
		packedData.packBits(benchData1024, 21)
	}
}

///
func BenchmarkPackBitsTurbo128(b *testing.B) {
	packedData := Uint32s(make([]uint32, 128))
	for i := 0; i < b.N; i++ {
		packedData.PackBits(benchData1024[:len(packedData)], 21)
	}
}
func BenchmarkPackBitsTurbo256(b *testing.B) {
	packedData := Uint32s(make([]uint32, 256))
	for i := 0; i < b.N; i++ {
		packedData.PackBits(benchData1024[:len(packedData)], 21)
	}
}
func BenchmarkPackBitsTurbo1024(b *testing.B) {
	packedData := Uint32s(make([]uint32, 1024))
	for i := 0; i < b.N; i++ {
		packedData.PackBits(benchData1024, 21)
	}
}

func BenchmarkUnpackBitsNaive128(b *testing.B) {
	packedData := Uint32s(make([]uint32, 128))
	packedData.packBits(benchData1024[:len(packedData)], 21)
	unpackedData := make([]uint32, len(packedData))
	for i := 0; i < b.N; i++ {
		Uint32s(unpackedData).unpackBits(packedData, 21)
	}
}
func BenchmarkUnpackBitsNaive256(b *testing.B) {
	packedData := Uint32s(make([]uint32, 256))
	packedData.packBits(benchData1024[:len(packedData)], 21)
	unpackedData := make([]uint32, len(packedData))
	for i := 0; i < b.N; i++ {
		Uint32s(unpackedData).unpackBits(packedData, 21)
	}
}
func BenchmarkUnpackBitsNaive1024(b *testing.B) {
	packedData := Uint32s(make([]uint32, 1024))
	packedData.packBits(benchData1024[:len(packedData)], 21)
	unpackedData := make([]uint32, len(packedData))
	for i := 0; i < b.N; i++ {
		Uint32s(unpackedData).unpackBits(packedData, 21)
	}
}

func BenchmarkUnpackBitsTurbo128(b *testing.B) {
	packedData := Uint32s(make([]uint32, 128))
	packedData.packBits(benchData1024[:len(packedData)], 21)
	unpackedData := make([]uint32, len(packedData))
	for i := 0; i < b.N; i++ {
		Uint32s(unpackedData).UnpackBits(packedData, 21)
	}
}
func BenchmarkUnpackBitsTurbo256(b *testing.B) {
	packedData := Uint32s(make([]uint32, 256))
	packedData.packBits(benchData1024[:len(packedData)], 21)
	unpackedData := make([]uint32, len(packedData))
	for i := 0; i < b.N; i++ {
		Uint32s(unpackedData).UnpackBits(packedData, 21)
	}
}
func BenchmarkUnpackBitsTurbo1024(b *testing.B) {
	packedData := Uint32s(make([]uint32, 1024))
	packedData.packBits(benchData1024[:len(packedData)], 21)
	unpackedData := make([]uint32, len(packedData))
	for i := 0; i < b.N; i++ {
		Uint32s(unpackedData).UnpackBits(packedData, 21)
	}
}
