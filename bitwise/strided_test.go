package bitwise

import (
	"fmt"
	"testing"
)

func TestAddStrided(t *testing.T) {
	// Arrange: 8 words -> 512 bits. Make only word 3 survive.
	a := Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0}
	b := Uint64s{0xffff_ffff_ffff_ffff, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0}
	out := make(Uint64s, len(a))
	strides := Strides{}
	strides.ensureStrides(a)

	// Act
	out.addStrided(a, b, strides)

	// Assert result bits
	expected := Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0}
	for i := range expected {
		if out[i] != expected[i] {
			t.Fatalf("word %d: have %016x want %016x", i, out[i], expected[i])
		}
	}

	fmt.Printf("%v\n", strides)
}
