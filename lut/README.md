## Lookup Tables

### Introduction

#### The following operations are implemented:

#### 'Lookup' maps an input byte vector to an output vector using a lookup table

#### 'LookupSubrange'

Maps an input byte vector to a collection of indexes using a subrange vector containing sorted uper limits defining the
partitioning. The subrange vector should be such that it would produce indexes ranging from 0 to 15. Then, an output
vector is created by mapping indexes using a 16 element lookup table. The 16 subrange limit is imposed due to
performance requirements related to vectorizing hardware architecture.

### Usage:

#### Lookup

```go
package mypkg;

import (
	"github.com/viant/vec/lut"
)

func Example_Lookup() {
	var table = make([]byte, 256)
	// .. populate the lookup table
	var input = make([]byte, 12345)
	// .. populate the input byte vector
	var output = make([]byte, len(input))
	lut.Lookup(input, table, &output)

}
```

#### LookupSubrange

```go
package mypkg;

import (
	"github.com/viant/vec/lut"
)

func Example_LookupSubrange() {
	var rangeUpper = []byte{10, 20, 100, 250}
	var rangeTable = []byte{0, 42, 43, 240, 211, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var input = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 255}
	var output = make([]byte, len(input))

	lut.LookupSubrange(input, rangeUpper, rangeTable, &output)
}
```

### Benchmarks

#### ARM64 (Neon)

```text
goarch: arm64
pkg: github.com/viant/vec/lut
Benchmark_LookupSubrangeNaive-16    	11207770	       112.7 ns/op
Benchmark_LookupSubrange-16         	100000000	        11.79 ns/op
Benchmark_Lookup_Naive-16           	 6004897	       199.8 ns/op
Benchmark_Lookup-16                 	21183451	        56.71 ns/op
```

### AMD64 (AVX2)

** MacBook2.4 GHz 8-Core Intel Core i9**

```text
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
Benchmark_LookupSubrangeNaive-16         7993898               145.4 ns/op
Benchmark_LookupSubrange-16             100000000               10.43 ns/op
Benchmark_Lookup_Naive-16                6125260               199.5 ns/op
Benchmark_Lookup-16                     11760783               103.1 ns/op
```
