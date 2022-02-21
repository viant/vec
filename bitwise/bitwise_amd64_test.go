package bitwise

import (
	"golang.org/x/sys/cpu"
	"testing"
)

func BenchmarkAnd_S_AVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.Skip("AVX2 not supported")
	}
	out := make([]uint64, len(op1S))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndAVX2(op1S, op2S)
	}
}

func BenchmarkAnd_M_AVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.Skip("AVX2 not present")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndAVX2(op1M, op2M)
	}
}

func BenchmarkAnd_XL_AVX2(b *testing.B) {
	if !cpu.X86.HasAVX2 {
		b.Skip("AVX2 not present")
	}
	out := make([]uint64, len(op2XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndAVX2(op1XL, op2XL)
	}
}

func BenchmarkAnd_S_AVX512(b *testing.B) {
	if !cpu.X86.HasAVX512 {
		b.Skip("AVX2 not supported")
	}
	out := make([]uint64, len(op1S))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndAVX512(op1S, op2S)
	}
}

func BenchmarkAnd_M_AVX512(b *testing.B) {
	if !cpu.X86.HasAVX512 {
		b.Skip("AVX2 not present")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndAVX512(op1M, op2M)
	}
}

func BenchmarkAnd_XL_AVX512(b *testing.B) {
	if !cpu.X86.HasAVX512 {
		b.Skip("AVX512 not present")
	}
	out := make([]uint64, len(op2XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndAVX512(op1XL, op2XL)
	}
}

func BenchmarkAnd_M_V3_AVX2(b *testing.B) {
	if !cpu.X86.HasAVX512 {
		b.Skip("AVX512 not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndV3AVX2(op1M, op2M, op3M)
	}
}

func BenchmarkAnd_M_V3_AVX512(b *testing.B) {
	if !cpu.X86.HasAVX512 {
		b.Skip("AVX512 not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndV3AVX512(op1M, op2M, op3M)
	}
}
