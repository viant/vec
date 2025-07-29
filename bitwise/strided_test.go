package bitwise

import (
	"testing"
)

func TestAndStrided(t *testing.T) {
	tests := []struct {
		name     string
		a, b     Uint64s
		expected Uint64s
	}{
		{
			name:     "only word 3 survives",
			a:        Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0},
			b:        Uint64s{0xffff_ffff_ffff_ffff, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0},
			expected: Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0, 0, 0, 0},
		},
		{
			name:     "all zeros",
			a:        Uint64s{0, 0, 0, 0},
			b:        Uint64s{0, 0, 0, 0},
			expected: Uint64s{0, 0, 0, 0},
		},
		{
			name:     "all ones",
			a:        Uint64s{^uint64(0), ^uint64(0)},
			b:        Uint64s{^uint64(0), ^uint64(0)},
			expected: Uint64s{^uint64(0), ^uint64(0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(Uint64s, len(tt.a))
			strides := Strides{}
			strides.ensureStrides(tt.a)

			out.andStrided(tt.a, tt.b, strides)

			for i := range tt.expected {
				if out[i] != tt.expected[i] {
					t.Fatalf("word %d: got %016x, want %016x", i, out[i], tt.expected[i])
				}
			}
		})
	}
}

func TestOrStrided(t *testing.T) {
	tests := []struct {
		name     string
		a, b     Uint64s
		expected Uint64s
	}{
		{
			name:     "simple OR with saturation at word 0 and 3",
			a:        Uint64s{0, 0, 0, 0xffff_ffff_ffff_ffff, 0},
			b:        Uint64s{0xffff_ffff_ffff_ffff, 0, 0, 0xffff_ffff_ffff_ffff, 0},
			expected: Uint64s{^uint64(0), 0, 0, ^uint64(0), 0},
		},
		{
			name:     "all zeros",
			a:        Uint64s{0, 0, 0},
			b:        Uint64s{0, 0, 0},
			expected: Uint64s{0, 0, 0},
		},
		{
			name:     "all ones + zeros",
			a:        Uint64s{^uint64(0), 0, ^uint64(0)},
			b:        Uint64s{0, ^uint64(0), 0},
			expected: Uint64s{^uint64(0), ^uint64(0), ^uint64(0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := make(Uint64s, len(tt.a))
			strides := Strides{}
			strides.ensureStrides(tt.a)

			out.orStrided(tt.a, tt.b, strides)

			for i := range tt.expected {
				if out[i] != tt.expected[i] {
					t.Fatalf("word %d: got %016x, want %016x", i, out[i], tt.expected[i])
				}
			}
		})
	}
}
