# Vec - Vectorization toolkit for go

[![GoReportCard](https://goreportcard.com/badge/github.com/viant/vec)](https://goreportcard.com/report/github.com/viant/vec)
[![GoDoc](https://godoc.org/github.com/viant/vec?status.svg)](https://godoc.org/github.com/viant/vec)

This library is compatible with Go 1.17+

Please refer to [`CHANGELOG.md`](CHANGELOG.md) if you encounter breaking changes.

- [Motivation](#motivation)
- [Usage](#usage)
- [License](#license)

## Motivation

The goal of this project is to take advantages of the SMID (Single Instruction Multiple Data) CPU extension to perform
128/256/512 bits wide operation per CPU cycle.  
This library uses the following extension: AVX2, AVX512, Neon, SVE. The operation are grouped in the following packages:

- [bitewise](bitwise): bitwise operations
- [lut](lut): lookup operations
- [blas](blas): basic linear algebra operations
- [bits](pack): bitset/masking/packing/unpacking operations
- 

## Usage

#### Bitwise operation:

```go
    out := make([]uint64, 8)
    v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
    v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
    bitwise.Uint64s(out).And(v1, v2)
```




## Contributing to vec

Vec is an open source project and contributors are welcome!

See [TODO](TODO.md) list

## License

The source code is made available under the terms of the Apache License, Version 2, as stated in the file `LICENSE`.

Individual files may be made available under their own specific license, all compatible with Apache License, Version 2.
Please see individual files for details.

<a name="Credits-and-Acknowledgements"></a>

## Credits and Acknowledgements

**Library Authors:**

- Valery Carey
- Adrian Witas
 