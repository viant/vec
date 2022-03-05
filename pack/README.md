## Bit pack/unpack operations

### Introduction

#### The following operations are implemented:

#### 'Deltas':  delta encode an int32 vector

#### 'RecoverDeltas': decode an int32 delta encoded vector

#### 'PackedUint32Count': Helper function to determine a packed vector size

#### 'MaxBits': Helper function to determine maximum number of bits  to encode a number for an int32 vector

#### 'PackBits': pack an int32 into a smaller vector in which in32 are represented with fewer bits (determined by MaxBits over a given input vector)

#### 'UnpackBits': reverse the result of the PackBits operation

#### 'Shrink': condense a byte mask stored as a uint64 vector of not nore than 8 elemenys long (64 bytes) to a single uint64 bit pattern.

#### 'ShrinkValue': condense a byte mask stored as a uint64 vector of not nore than 8 elemenys long (64 bytes) to a single uint64 bit pattern. Here we check for a specific value to map to '1'.

#### 'Expand': reverse the result of the ShrinkValue operation

####    

### Usage:

#### Deltas/RecoverDeltas

```go
package mypkg;

import (
	"github.com/viant/vec/pack"
)

func Example_Deltas() {
	data := make([]int32, 1000)
	//...
	deltas := pack.Int32s(make([]int32, len(data)))
	deltas.Deltas(data)
	//...
	recoveredData := pack.Int32s(make([]int32, len(deltas)))
	recoveredData.RecoverDeltas(deltas)

}
```

#### MaxBits

```go
package mypkg;

import (
	"fmt"
	"github.com/viant/vec/pack"
)

func Example_MaxBits() {
	var data = make([]uint32, 1000)
	//....
	bitwidth := pack.MaxBits(data)
	fmt.Println(bitwidth)
}
```

#### PackBits/UnpackBits

```go
package mypkg;

import (
	"github.com/viant/vec/pack"
)

func Example_PackBits() {
	data := []uint32{20124, 8831, 2575, 1977, 15, 42441, 302690, 871222, 323452, 424532, 29434, 939141, 4244, 324314, 13, 1, 874255, 20124, 8831, 2575, 1977, 15, 42441, 302690, 871222, 323452, 424532, 29434, 939141, 4244, 324314, 13, 1, 874255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 42441, 302690, 871222, 323452, 424532, 29434, 939141, 4244, 324314, 13, 1, 874255, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	bits := pack.MaxBits(data)
	packedData := pack.Uint32s(make([]uint32, pack.PackedUint32Count(uint32(len(data)), bits)))
	packedData.PackBits(data, int(bits))

	unpackedData := make([]uint32, len(data))
	pack.Uint32s(unpackedData).UnpackBits(packedData, int(bits))
}
```

#### Shrink

```go
package mypkg;

import (
	"github.com/viant/vec/pack"
)

func Example_Shrink() {
	var ints = []uint64{
		0x8000800080008000, 0x8000800080008000, 0x8000800080008000, 0x8000800080008000,
		0x0080008000800080, 0x0080008000800080, 0x0080008000800080, 0x0080008000800080,
	}
	var shrunkValue = pack.Uint64(0)
	shrunkValue.Shrink(ints)

}
```

#### ShrinkValue

```go
package mypkg;

import (
	"github.com/viant/vec/pack"
)

func Example_ShrinkValue() {
	var ints = []uint64{
		0x8000800080008000, 0x8000800080008000, 0x8000800080008000, 0x8000800080008000,
		0x0080008000800080, 0x0080008000800080, 0x0080008000800080, 0x0080008000800080,
	}
	var shrunkValue = pack.Uint64(0)
	shrunkValue.ShrinkValue(0x80, ints)

}
```

#### Expand

```go
package mypkg;

import (
	"github.com/viant/vec/pack"
)

func Example_Expand() {
	var shrunkValue = uint64(0x55555555aaaaaaaa)
	out := pack.Uint64s(make([]uint64, 8))
	out.Expand(shrunkValue)
}
```


### Benchmarks

#### ARM64 (Neon)

```text
goarch: arm64
pkg: github.com/viant/vec/pack
BenchmarkBitwidthNaive-16          	20608507	        58.19 ns/op
BenchmarkBitwidth-16               	100000000	        10.87 ns/op
BenchmarkDeltasNaive-16            	16701920	        71.89 ns/op
BenchmarkRecoverDeltasNaive-16     	19388534	        61.83 ns/op
BenchmarkDeltas-16                 	49066635	        24.59 ns/op
BenchmarkRecoverDeltas-16          	32289895	        37.16 ns/op
BenchmarkExpand-16                 	310915053	         3.864 ns/op	       0 B/op	       0 allocs/op
BenchmarkNaiveExpand-16            	47216738	        25.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkPackBitsNaive128-16       	 5322465	       225.0 ns/op
BenchmarkPackBitsNaive256-16       	 2706843	       444.4 ns/op
BenchmarkPackBitsNaive1024-16      	  680054	      1763 ns/op
BenchmarkPackBitsTurbo128-16       	39110784	        30.76 ns/op
BenchmarkPackBitsTurbo256-16       	21159951	        56.62 ns/op
BenchmarkPackBitsTurbo1024-16      	 5481162	       219.2 ns/op
BenchmarkUnpackBitsNaive128-16     	 4581700	       263.9 ns/op
BenchmarkUnpackBitsNaive256-16     	 2313916	       521.2 ns/op
BenchmarkUnpackBitsNaive1024-16    	  583735	      2056 ns/op
BenchmarkUnpackBitsTurbo128-16     	41587406	        28.84 ns/op
BenchmarkUnpackBitsTurbo256-16     	22318448	        53.76 ns/op
BenchmarkUnpackBitsTurbo1024-16    	 5826024	       206.0 ns/op
BenchmarkNaiveShrink-16            	19310388	        61.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkShrink-16                 	136003323	         8.823 ns/op	       0 B/op	       0 allocs/op
BenchmarkNaiveShrinkValue-16       	19474821	        61.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkShrinkValue-16            	136449284	         8.793 ns/op	       0 B/op	       0 allocs/op
```

### AMD64 (AVX2)

** MacBook2.4 GHz 8-Core Intel Core i9**

```text
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkBitwidthNaive-16               21929083                48.18 ns/op
BenchmarkBitwidth-16                    182549624                6.512 ns/op
BenchmarkDeltasNaive-16                 12633043                91.23 ns/op
BenchmarkDeltas-16                      75425048                15.13 ns/op
BenchmarkRecoverDeltasNaive-16          18658293                63.07 ns/op
BenchmarkRecoverDeltas-16               46790792                25.17 ns/op
BenchmarkNaiveExpand-16                 37795432                31.88 ns/op            0 B/op          0 allocs/op
BenchmarkExpand-16                      362312720                3.157 ns/op           0 B/op          0 allocs/op
BenchmarkPackBitsNaive128-16             4754596               236.6 ns/op
BenchmarkPackBitsNaive256-16             2571043               470.3 ns/op
BenchmarkPackBitsNaive1024-16             586951              1845 ns/op
BenchmarkPackBitsTurbo128-16            30742057                39.71 ns/op
BenchmarkPackBitsTurbo256-16            14866566                76.22 ns/op
BenchmarkPackBitsTurbo1024-16            3894434               304.9 ns/op
BenchmarkUnpackBitsNaive128-16           4204603               301.4 ns/op
BenchmarkUnpackBitsNaive256-16           2077891               574.8 ns/op
BenchmarkUnpackBitsNaive1024-16           514095              2273 ns/op
BenchmarkUnpackBitsTurbo128-16          28993444                36.30 ns/op
BenchmarkUnpackBitsTurbo256-16          15819532                70.37 ns/op
BenchmarkUnpackBitsTurbo1024-16          4241524               284.9 ns/op
BenchmarkNaiveShrink-16                 17989338                65.47 ns/op            0 B/op          0 allocs/op
BenchmarkShrink-16                      449994657                2.630 ns/op           0 B/op          0 allocs/op
BenchmarkNaiveShrinkValue-16            18004347                68.33 ns/op            0 B/op          0 allocs/op
BenchmarkShrinkValue-16                 409808895                2.927 ns/op           0 B/op          0 allocs/op
```
