## Miscellaneous Bit Operations

### Introduction

#### The following operations are implemented:

#### 'Count' counts number of '1' (set) bits ia a vector

#### 'Populate' populates a preallocated [64]int32 array with bit positions in a int64 input

#### 'Lsb/Msb' least significant/most significan bit position in a vector
### Usage:

#### Count

```go
package mypkg;

import (
	"fmt"
	"github.com/viant/vec/bits"
	"math/rand"
	"time"
)

func Example_Count() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var vec = make([]uint64, 267)
	for i := range vec {
		vec[i] = rnd.Uint64()
	}
	cnt := bits.Count(vec)
	fmt.Println(cnt)
}
```

#### Populate

```go
package mypkg;

import (
	"fmt"
	"github.com/viant/vec/bits"
	"math/rand"
	"time"
)

func Example_Count() {
	var out bits.Positions
	input := rand.New(rand.NewSource(time.Now().UnixNano())).Uint64()
	lastIndex := out.Populate(input)
	fmt.Println(out[0:lastIndex])
}
```


#### LSb/Msb

```go
package mypkg;

import (
	"fmt"
	"github.com/viant/vec/bits"
	"math/rand"
	"time"
)

func Example_Lsb() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var vec = make([]uint64, 267)
	for i := range vec {
		vec[i] = rnd.Uint64()
	}
	position := bits.Lsb(vec)
	fmt.Println(position)}
```

### Benchmarks
#### ARM64 (Neon)

```text
goarch: arm64
pkg: github.com/viant/vec/bits
BenchmarkCountNaive-16        	60427422	        19.49 ns/op
BenchmarkCount-16             	176838339	         6.719 ns/op
BenchmarkLsbNaive-16          	63967446	        19.11 ns/op
BenchmarkLsb-16               	160451912	         7.408 ns/op
BenchmarkMsbNaive-16          	62994744	        19.14 ns/op
BenchmarkMsb-16               	136166562	         8.788 ns/op
BenchmarkPositionsNaive-16    	50261076	        23.71 ns/op
BenchmarkPositions-16         	199557674	         5.987 ns/op
```

#### AMD64 (AVX2)

** MacBook2.4 GHz 8-Core Intel Core i9**

```text
goarch: amd64
pkg: github.com/viant/vec/bits
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkCountNaive-16          34619952                31.37 ns/op
BenchmarkCount-16               82600533                13.84 ns/op
BenchmarkPositionsNaive-16      43401525                26.80 ns/op
BenchmarkPositions-16           170968920                6.977 ns/op
```
