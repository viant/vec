## Vectorized Bitwise Operations

### Introduction

Vectorized bitwise operations use SIMD (Single Instruction Multiple Data) CPU extension to perform 128/256/512 bits wide
operation per CPU cycle. This library uses the following extension: AVX2, AVX512, Neon, SVE. All implemented cases
provide substantial performance benefit over go lang '&', '|',  '^' operators.

### Usage:

#### Vectorized AND

```go
package mypkg;

import "github.com/viant/vec/bitwise"

func ExampleUint64s_And() {
	{
		out := bitwise.Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		out.And(v1, v2)
	}
	{
		out := make([]uint64, 8)
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		v3 := []uint64{1, 1, 0, 4, 1, 6, 7, 2}
		bitwise.Uint64s(out).AndV3(v1, v2, v3)
	}
}

```

#### Vectorized OR

```go
package mypkg;

import "github.com/viant/vec/bitwise"

func ExampleUint64s_Or() {
	{
		out := make([]uint64, 8)
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		bitwise.Uint64s(out).Or(v1, v2)
	}
	{
		out := bitwise.Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		v3 := []uint64{1, 1, 0, 4, 1, 6, 7, 2}
		out.OrV3(v1, v2, v3)
	}
}

```

#### Vectorized XOR

```go
package mypkg;

import "github.com/viant/vec/bitwise"

func ExampleUint64s_XOr() {
	{
		out := bitwise.Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		out.Xor(v1, v2)
	}
	{
		out := bitwise.Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		v3 := []uint64{1, 1, 0, 4, 1, 6, 7, 2}
		out.XorV3(v1, v2, v3)
	}
}
```

### Benchmarks

The benchmark uses 3 sets of []uints
**S**mall -  [8]uint64
**M**edium - [32]uint64
**XL**arge - [128]uint64


#### ARM64 (Neon/SVE)

**Apple M1**

```text
goos: darwin
goarch: arm64
pkg: github.com/viant/vec/bitwise
BenchmarkAnd_S_Arm64Neon-8      362116297                3.131 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_M_Arm64Neon-8      100000000               11.54 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_XL_Arm64Neon-8     48047887                24.22 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_M_V3_Neon-8        100000000               11.53 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_S-8                348350084                3.449 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_M-8                150012367                7.994 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_XL-8               38915920                30.50 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_S_Naive-8          186783776                6.564 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_M_Naive-8          70419261                17.00 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_XL_Naive-8         17308191                68.91 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_M_V3-8             100000000               11.38 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_XL_V3-8            36936477                32.07 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_M_V3_Naive-8       52957920                22.63 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_XL_V3_Naive-8      13048790                89.38 ns/op            0 B/op          0 allocs/op
```

**Graviton 2**

```text
goos: linux
goarch: arm64
pkg: github.com/viant/vec/bitwise
BenchmarkAnd_S_Arm64Neon-16     	148614336	         8.077 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M_Arm64Neon-16     	68040201	        17.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_Arm64Neon-16    	21396996	        56.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_V3_Neon-16       	22370042	        53.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_S-16               	136180149	         8.812 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M-16               	65123619	        18.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL-16              	21087504	        56.88 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_S_Naive-16         	75354376	        15.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_Naive-16         	26723798	        44.91 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_Naive-16        	 7453200	       161.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_V3-16            	22334678	        53.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_V3-16           	 6257778	       192.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_V3_Naive-16      	20654671	        58.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_V3_Naive-16     	 5529211	       221.2 ns/op	       0 B/op	       0 allocs/op
```


### AMD64 (AVX2/AVX512)

**Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz**

```text
BenchmarkAnd_S_AVX2-16          296558868                4.015 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_M_AVX2-16          164493822                6.343 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_XL_AVX2-16         56744065                18.01 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_S-16               269001751                4.854 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_M-16               148436960                8.201 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_XL-16              59192054                21.08 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_S_Naive-16         146186683                7.764 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_M_Naive-16         61259191                20.51 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_XL_Naive-16        13925252                76.87 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_M_V3-16            141754650                8.224 ns/op           0 B/op          0 allocs/op
BenchmarkAnd_XL_V3-16           55797906                23.16 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_M_V3_Naive-16      48094837                25.76 ns/op            0 B/op          0 allocs/op
BenchmarkAnd_XL_V3_Naive-16     11524884                97.60 ns/op            0 B/op          0 allocs/op
```

**Intel(R) Xeon(R) Platinum 8124M CPU @ 3.00GHz**

```text
BenchmarkAnd_S_AVX2-16         	237390595	         5.024 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M_AVX2-16         	150596898	         7.963 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_XL_AVX2-16        	53840293	        22.12 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_S_AVX512-16       	141063717	         8.508 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M_AVX512-16       	51851728	        23.09 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_AVX512-16      	13561696	        88.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_V3_AVX2-16      	127001692	         9.449 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M_V3_AVX512-16    	49551648	        24.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_S-16              	204319275	         5.795 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M-16              	135662316	         8.845 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_XL-16             	52037328	        23.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_S_Naive-16        	155402485	         7.597 ns/op	   0 B/op	       0 allocs/op
BenchmarkAnd_M_Naive-16        	53635558	        22.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_Naive-16       	13346347	        90.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_V3-16           	100000000	        10.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_V3-16          	49549333	        24.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_M_V3_Naive-16     	40608735	        29.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnd_XL_V3_Naive-16    	10238077	       117.1 ns/op	       0 B/op	       0 allocs/op
```