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
