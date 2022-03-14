## Basic Linear Algebra Operations

### Introduction

#### The following operations are implemented using vectors as operands:

#### float32:
add, sub, mul, div

###
#### int32
add, sub, mul,

###
#### hfloat32/hint32
The "horizontal"(across a single vector) operations below are available for both int32 and float32:

hsum, hmin, hmax

###
#### float64
add, sub, mul, div

###
####int64
amd64: add, sub, mul
####
arm64: add, sub only. mul falls back to scalar

###
####hfloat64
hsum, hmin, hmax

###
####hint64
amd64: hsum requires AVX. hmin and hmax require AVX512, otherwise fall back to scalar ops.
####
arm64: hsum.  hmin and hmax fall back to scala ops due to unavailability of vector ops on neon.
##
### Usage:

#### Vectorized Add

```go
package mypkg;

import "github.com/viant/vec/blas"

func ExampleInt32_Add() {
	v1 := []int32{1, 2, 3, 4, 5, 6, 7, 8}
	v2 := []int32{1, 7, 3, 4, 3, 6, 7, 2}
	out := blas.Int32s(make([]int32, 8))
	out.AddInt32(v1, v2)
}


```

#### Vectorized Horizontal Sum

```go
package mypkg;

import (
	"fmt"
	"github.com/viant/vec/blas"
)

func ExampleInt32_HSum() {
	var data = make([]float32, 1000)
	//...
	sum := blas.HsumFloat32(data)
	fmt.Println(sum)
}

```

### Benchmarks

#### ARM64 (Neon)

```text
goarch: arm64
pkg: github.com/viant/vec/blas
BenchmarkAddFloat32Naive-16     	 8077633	       148.5 ns/op
BenchmarkAddFloat32-16          	40494718	        29.66 ns/op
BenchmarkSubFloat32Naive-16     	 8079802	       148.5 ns/op
BenchmarkSubFloat32-16          	40544472	        29.56 ns/op
BenchmarkMulFloat32Naive-16     	 8080488	       148.5 ns/op
BenchmarkMulFloat32-16          	37863837	        31.64 ns/op
BenchmarkDivFloat32Naive-16     	 7515062	       159.8 ns/op
BenchmarkDivFloat32-16          	15938955	        75.28 ns/op
BenchmarkHsumFloat32Naive-16    	 7127556	       168.4 ns/op
BenchmarkHsumFloat32-16         	24377919	        49.83 ns/op
BenchmarkHmaxFloat32Naive-16    	11152754	       107.6 ns/op
BenchmarkHmaxFloat32-16         	18596677	        63.38 ns/op
BenchmarkHminFloat32Naive-16    	 3211771	       373.6 ns/op
BenchmarkHminFloat32-16         	17684204	        66.94 ns/op
BenchmarkHsumInt32Naive-16      	10272612	       116.8 ns/op
BenchmarkHsumInt32-16           	24530227	        49.38 ns/op
BenchmarkHmaxInt32Naive-16      	 5949591	       201.7 ns/op
BenchmarkHmaxInt32-16           	18702337	        63.25 ns/op
BenchmarkHminInt32Naive-16      	 5949391	       201.7 ns/op
BenchmarkHminInt32-16           	15323509	        77.99 ns/op
BenchmarkAddInt32Naive-16       	 6029112	       199.1 ns/op
BenchmarkAddInt32-16            	39999256	        30.03 ns/op
BenchmarkSubInt32Naive-16       	 6029341	       199.1 ns/op
BenchmarkSubInt32-16            	39963598	        29.96 ns/op
BenchmarkMulInt32Naive-16       	 6028239	       199.1 ns/op
BenchmarkMulInt32-16            	38536497	        31.13 ns/op
```

### AMD64 (AVX2)

** MacBook2.4 GHz 8-Core Intel Core i9**

```text
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkAddFloat32Naive-16              9386047               125.3 ns/op
BenchmarkAddFloat32-16                  49576354                23.45 ns/op
BenchmarkSubFloat32Naive-16              9222256               126.7 ns/op
BenchmarkSubFloat32-16                  76574758                15.81 ns/op
BenchmarkMulFloat32Naive-16              9295201               130.4 ns/op
BenchmarkMulFloat32-16                  54382240                18.65 ns/op
BenchmarkDivFloat32Naive-16              6605150               172.0 ns/op
BenchmarkDivFloat32-16                  37387005                30.13 ns/op
BenchmarkHsumFloat32Naive-16             6167763               196.0 ns/op
BenchmarkHsumFloat32-16                 31841235                36.94 ns/op
BenchmarkHmaxFloat32Naive-16             8876424               127.9 ns/op
BenchmarkHmaxFloat32-16                 27600762                42.20 ns/op
BenchmarkHminFloat32Naive-16             8366142               136.5 ns/op
BenchmarkHminFloat32-16                 35269717                36.81 ns/op
BenchmarkHsumInt32Naive-16              13112025                94.23 ns/op
BenchmarkHsumInt32-16                   41688908                28.40 ns/op
BenchmarkHmaxInt32Naive-16               9809748               117.6 ns/op
BenchmarkHmaxInt32-16                   20696240                55.76 ns/op
BenchmarkHminInt32Naive-16               9347374               134.8 ns/op
BenchmarkHminInt32-16                   24188605                50.46 ns/op
BenchmarkAddInt32Naive-16                8154877               132.2 ns/op
BenchmarkAddInt32-16                    44763596                25.55 ns/op
BenchmarkSubInt32Naive-16                9035091               131.1 ns/op
BenchmarkSubInt32-16                    71036890                16.09 ns/op
BenchmarkMulInt32Naive-16                8255730               138.3 ns/op
BenchmarkMulInt32-16                    66136818                18.92 ns/op
```
