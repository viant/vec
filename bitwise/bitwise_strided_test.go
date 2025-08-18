//go:build arm64 && !noasm && !appengine
// +build arm64,!noasm,!appengine

package bitwise

import (
	"testing"
)

func mkStrides(spans ...[2]uint32) Strides {
	s := make(Strides, 3+2*len(spans))
	s[2] = uint32(len(spans))
	for i, sp := range spans {
		s[3+2*i+0] = sp[0]
		s[3+2*i+1] = sp[1]
	}
	return s
}

func genInputs(k, n int) []Uint64s {
	in := make([]Uint64s, k)
	for j := 0; j < k; j++ {
		v := make(Uint64s, n)
		// Deterministic, varied pattern per input j
		var a uint64 = 0x9e3779b97f4a7c15 ^ uint64(j)*0xbf58476d1ce4e5b9
		var b uint64 = 0x94d049bb133111eb ^ uint64(j*0x12345+7)
		for i := 0; i < n; i++ {
			x := uint64(i+1) * a
			y := (uint64(i) << (uint(j) % 7)) ^ b
			v[i] = x ^ y ^ uint64(j*0x5555555555555555)
		}
		in[j] = v
	}
	return in
}

func expectedAnd(outInit Uint64s, inputs []Uint64s, spans ...[2]uint32) Uint64s {
	out := make(Uint64s, len(outInit))
	copy(out, outInit)
	for _, sp := range spans {
		start, ln := int(sp[0]), int(sp[1])
		if ln == 0 {
			continue
		}
		for i := 0; i < ln; i++ {
			k := start + i
			acc := inputs[0][k]
			for j := 1; j < len(inputs); j++ {
				acc &= inputs[j][k]
			}
			out[k] = acc
		}
	}
	return out
}

func expectedOr(outInit Uint64s, inputs []Uint64s, spans ...[2]uint32) Uint64s {
	out := make(Uint64s, len(outInit))
	copy(out, outInit)
	for _, sp := range spans {
		start, ln := int(sp[0]), int(sp[1])
		if ln == 0 {
			continue
		}
		for i := 0; i < ln; i++ {
			k := start + i
			acc := inputs[0][k]
			for j := 1; j < len(inputs); j++ {
				acc |= inputs[j][k]
			}
			out[k] = acc
		}
	}
	return out
}

func checkEqual(t *testing.T, got, want Uint64s) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("length mismatch: got %d want %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("mismatch at i=%d: got=0x%x want=0x%x", i, got[i], want[i])
		}
	}
}

// Common test driver for AND/OR families.
func runAndCase(t *testing.T, k int, call func(out *Uint64s, ins []Uint64s, strides Strides)) {
	const N = 33 // odd on purpose
	ins := genInputs(k, N)

	// Sentinel output to ensure untouched indices remain unchanged.
	outInit := make(Uint64s, N)
	for i := range outInit {
		outInit[i] = 0xDEADBEEFCAFEBABE ^ uint64(i)*0x0101010101010101
	}

	// Mixture of spans: odd/even starts and lengths; includes a zero-length span.
	spans := [][2]uint32{
		{1, 7},  // odd start, odd length
		{8, 0},  // zero-length (should be ignored)
		{12, 9}, // even start, odd length
		{23, 2}, // odd start, even length
	}
	strides := mkStrides(spans...)

	// Compute expected and run DUT
	want := expectedAnd(outInit, ins, spans...)
	out := make(Uint64s, N)
	copy(out, outInit)

	call(&out, ins, strides)
	checkEqual(t, out, want)

	// Also test empty spans: no changes expected.
	empty := mkStrides()
	out2 := make(Uint64s, N)
	copy(out2, outInit)
	call(&out2, ins, empty)
	checkEqual(t, out2, outInit)
}

func runOrCase(t *testing.T, k int, call func(out *Uint64s, ins []Uint64s, strides Strides)) {
	const N = 33
	ins := genInputs(k, N)

	outInit := make(Uint64s, N)
	for i := range outInit {
		outInit[i] = 0x0123456789ABCDEF ^ uint64(i)*0x1111111111111111
	}

	spans := [][2]uint32{
		{0, 1},  // first element
		{2, 6},  // even start, even length
		{13, 7}, // odd start, odd length
		{25, 4}, // tail region, even length
	}
	strides := mkStrides(spans...)

	want := expectedOr(outInit, ins, spans...)
	out := make(Uint64s, N)
	copy(out, outInit)

	call(&out, ins, strides)
	checkEqual(t, out, want)

	// Empty spans: unchanged
	empty := mkStrides()
	out2 := make(Uint64s, N)
	copy(out2, outInit)
	call(&out2, ins, empty)
	checkEqual(t, out2, outInit)
}

// callAnd dispatches to the proper AndXStrided based on k inputs.
func callAnd(t *testing.T, k int, out *Uint64s, ins []Uint64s, strides Strides) {
	t.Helper()
	switch k {
	case 2:
		(*out).AndStrided(ins[0], ins[1], strides)
	case 3:
		(*out).And3Strided(ins[0], ins[1], ins[2], strides)
	case 4:
		(*out).And4Strided(ins[0], ins[1], ins[2], ins[3], strides)
	case 5:
		(*out).And5Strided(ins[0], ins[1], ins[2], ins[3], ins[4], strides)
	case 6:
		(*out).And6Strided(ins[0], ins[1], ins[2], ins[3], ins[4], ins[5], strides)
	default:
		t.Fatalf("unsupported k=%d", k)
	}
}

// callOr dispatches to the proper OrXStrided based on k inputs.
func callOr(t *testing.T, k int, out *Uint64s, ins []Uint64s, strides Strides) {
	t.Helper()
	switch k {
	case 2:
		(*out).OrStrided(ins[0], ins[1], strides)
	case 3:
		(*out).Or3Strided(ins[0], ins[1], ins[2], strides)
	case 4:
		(*out).Or4Strided(ins[0], ins[1], ins[2], ins[3], strides)
	case 5:
		(*out).Or5Strided(ins[0], ins[1], ins[2], ins[3], ins[4], strides)
	case 6:
		(*out).Or6Strided(ins[0], ins[1], ins[2], ins[3], ins[4], ins[5], strides)
	default:
		t.Fatalf("unsupported k=%d", k)
	}
}

// expectedBySet computes expected output using the set mask (hot blocks) and op selector.
func expectedBySet(outInit Uint64s, ins []Uint64s, set Uint64s, useAnd bool) Uint64s {
	out := make(Uint64s, len(outInit))
	copy(out, outInit)
	for i := 0; i < len(out); i++ {
		if set[i] == 0 {
			continue
		}
		acc := ins[0][i]
		if useAnd {
			for j := 1; j < len(ins); j++ {
				acc &= ins[j][i]
			}
		} else {
			for j := 1; j < len(ins); j++ {
				acc |= ins[j][i]
			}
		}
		out[i] = acc
	}
	return out
}

// buildSet builds a set (hot blocks) vector of length n from provided spans.
func buildSet(n int, spans ...[2]uint32) Uint64s {
	set := make(Uint64s, n)
	for _, sp := range spans {
		start, ln := int(sp[0]), int(sp[1])
		for i := 0; i < ln && start+i < n; i++ {
			set[start+i] = 1
		}
	}
	return set
}

// Data-driven tests that exercise various hot-block layouts: beginning, end, middle, and random places.
func TestStrided_DataDriven_ManualSpans(t *testing.T) {
	// Define explicit inputs for k=2..6; length n=8 for concise expectations.
	v1 := Uint64s{1, 2, 4, 8, 1, 2, 4, 8}
	v2 := Uint64s{3, 3, 3, 3, 12, 12, 12, 12}
	v3 := Uint64s{5, 5, 5, 5, 5, 5, 5, 5}
	v4 := Uint64s{14, 14, 14, 14, 14, 14, 14, 14}
	v5 := Uint64s{9, 9, 9, 9, 9, 9, 9, 9}
	v6 := Uint64s{6, 6, 6, 6, 6, 6, 6, 6}

	type opKind int
	const (
		opAND opKind = iota
		opOR
	)

	type manualCase struct {
		name     string
		op       opKind
		spans    [][2]uint32
		ins      []Uint64s
		outInit  Uint64s
		expected Uint64s
	}

	cases := []manualCase{
		// OR 2 inputs, begin span
		{
			name:     "OR/k2/begin",
			op:       opOR,
			spans:    [][2]uint32{{0, 4}},
			ins:      []Uint64s{v1, v2},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{3, 3, 7, 11, 204, 205, 206, 207},
		},
		// AND 2 inputs, multiple disjoint spans
		{
			name:     "AND/k2/multi-spans",
			op:       opAND,
			spans:    [][2]uint32{{0, 2}, {4, 2}, {7, 1}},
			ins:      []Uint64s{v1, v2},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{1, 2, 102, 103, 0, 0, 106, 8},
		},
		// AND 2 inputs, begin span
		{
			name:     "AND/k2/begin",
			op:       opAND,
			spans:    [][2]uint32{{0, 4}},
			ins:      []Uint64s{v1, v2},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{1, 2, 0, 0, 104, 105, 106, 107},
		},

		// AND 3 inputs, middle span
		{
			name:     "AND/k3/middle",
			op:       opAND,
			spans:    [][2]uint32{{2, 4}},
			ins:      []Uint64s{v1, v2, v3},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{100, 101, 0, 0, 0, 0, 106, 107},
		},
		// AND 3 inputs, multiple disjoint spans
		{
			name:     "AND/k3/multi-spans",
			op:       opAND,
			spans:    [][2]uint32{{1, 2}, {5, 2}},
			ins:      []Uint64s{v1, v2, v3},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{100, 0, 0, 103, 104, 0, 4, 107},
		},
		// AND 4 inputs, tail span
		{
			name:     "AND/k4/tail",
			op:       opAND,
			spans:    [][2]uint32{{5, 3}},
			ins:      []Uint64s{v1, v2, v3, v4},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{100, 101, 102, 103, 104, 0, 4, 0},
		},
		// AND 4 inputs, multiple disjoint spans
		{
			name:     "AND/k4/multi-spans",
			op:       opAND,
			spans:    [][2]uint32{{0, 1}, {2, 2}, {6, 1}},
			ins:      []Uint64s{v1, v2, v3, v4},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{0, 101, 0, 0, 104, 105, 4, 107},
		},
		// AND 5 inputs, large span
		{
			name:     "AND/k5/large",
			op:       opAND,
			spans:    [][2]uint32{{1, 6}},
			ins:      []Uint64s{v1, v2, v3, v4, v5},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{100, 0, 0, 0, 0, 0, 0, 107},
		},
		// OR 2 inputs, multiple disjoint spans
		{
			name:     "OR/k2/multi-spans",
			op:       opOR,
			spans:    [][2]uint32{{0, 2}, {4, 2}, {7, 1}},
			ins:      []Uint64s{v1, v2},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{3, 3, 202, 203, 13, 14, 206, 12},
		},
		// AND 6 inputs, full span
		{
			name:     "AND/k6/all",
			op:       opAND,
			spans:    [][2]uint32{{0, 8}},
			ins:      []Uint64s{v1, v2, v3, v4, v5, v6},
			outInit:  Uint64s{100, 101, 102, 103, 104, 105, 106, 107},
			expected: Uint64s{0, 0, 0, 0, 0, 0, 0, 0},
		},
		// OR 3 inputs, multiple disjoint spans
		{
			name:     "OR/k3/multi-spans",
			op:       opOR,
			spans:    [][2]uint32{{0, 3}, {6, 2}},
			ins:      []Uint64s{v1, v2, v3},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{7, 7, 7, 203, 204, 205, 13, 13},
		},

		// OR 4 inputs, multiple disjoint spans
		{
			name:     "OR/k4/multi-spans",
			op:       opOR,
			spans:    [][2]uint32{{1, 3}, {5, 3}},
			ins:      []Uint64s{v1, v2, v3, v4},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{200, 15, 15, 15, 204, 15, 15, 15},
		},
		// OR 3 inputs, middle span
		{
			name:     "OR/k3/middle",
			op:       opOR,
			spans:    [][2]uint32{{2, 4}},
			ins:      []Uint64s{v1, v2, v3},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{200, 201, 7, 15, 13, 15, 206, 207},
		},
		// OR 4 inputs, tail span
		{
			name:     "OR/k4/tail",
			op:       opOR,
			spans:    [][2]uint32{{5, 3}},
			ins:      []Uint64s{v1, v2, v3, v4},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{200, 201, 202, 203, 204, 15, 15, 15},
		},
		// OR 5 inputs, large span
		{
			name:     "OR/k5/large",
			op:       opOR,
			spans:    [][2]uint32{{1, 6}},
			ins:      []Uint64s{v1, v2, v3, v4, v5},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{200, 15, 15, 15, 15, 15, 15, 207},
		},
		// OR 6 inputs, full span
		{
			name:     "OR/k6/all",
			op:       opOR,
			spans:    [][2]uint32{{0, 8}},
			ins:      []Uint64s{v1, v2, v3, v4, v5, v6},
			outInit:  Uint64s{200, 201, 202, 203, 204, 205, 206, 207},
			expected: Uint64s{15, 15, 15, 15, 15, 15, 15, 15},
		},
	}

	callOp := func(t *testing.T, out *Uint64s, ins []Uint64s, op opKind, strides Strides) {
		t.Helper()
		k := len(ins)
		switch op {
		case opAND:
			callAnd(t, k, out, ins, strides)
		case opOR:
			callOr(t, k, out, ins, strides)
		default:
			t.Fatalf("unknown op")
		}
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			n := len(c.outInit)
			if n == 0 || len(c.expected) != n {
				t.Fatalf("invalid case %q lengths", c.name)
			}
			out := make(Uint64s, n)
			copy(out, c.outInit)
			strides := mkStrides(c.spans...)
			callOp(t, &out, c.ins, c.op, strides)
			checkEqual(t, out, c.expected)
		})
	}
}

//
// AND tests
//

func TestAnd2Strided(t *testing.T) {
	runAndCase(t, 2, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).AndStrided(ins[0], ins[1], strides)
	})
}

func TestAnd3Strided(t *testing.T) {
	runAndCase(t, 3, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).And3Strided(ins[0], ins[1], ins[2], strides)
	})
}

func TestAnd4Strided(t *testing.T) {
	runAndCase(t, 4, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).And4Strided(ins[0], ins[1], ins[2], ins[3], strides)
	})
}

func TestAnd5Strided(t *testing.T) {
	runAndCase(t, 5, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).And5Strided(ins[0], ins[1], ins[2], ins[3], ins[4], strides)
	})
}

func TestAnd6Strided(t *testing.T) {
	runAndCase(t, 6, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).And6Strided(ins[0], ins[1], ins[2], ins[3], ins[4], ins[5], strides)
	})
}

//
// OR tests
//

func TestOr2Strided(t *testing.T) {
	runOrCase(t, 2, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).OrStrided(ins[0], ins[1], strides)
	})
}

func TestOr3Strided(t *testing.T) {
	runOrCase(t, 3, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).Or3Strided(ins[0], ins[1], ins[2], strides)
	})
}

func TestOr4Strided(t *testing.T) {
	runOrCase(t, 4, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).Or4Strided(ins[0], ins[1], ins[2], ins[3], strides)
	})
}

func TestOr5Strided(t *testing.T) {
	runOrCase(t, 5, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).Or5Strided(ins[0], ins[1], ins[2], ins[3], ins[4], strides)
	})
}

func TestOr6Strided(t *testing.T) {
	runOrCase(t, 6, func(out *Uint64s, ins []Uint64s, strides Strides) {
		(*out).Or6Strided(ins[0], ins[1], ins[2], ins[3], ins[4], ins[5], strides)
	})
}
