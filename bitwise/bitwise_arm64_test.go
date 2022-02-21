package bitwise

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/sys/cpu"
	"math/rand"
	"testing"
)

func TestOrV3Neon(t *testing.T) {
	testBitwiseV3(t, testOrV3Neon)
}
func TestOrV4Neon(t *testing.T) {
	testBitwiseV4(t, testOrV4Neon)
}

func TestOrV5Neon(t *testing.T) {
	testBitwiseV5(t, testOrV5Neon)
}

func TestOrV6Neon(t *testing.T) {
	testBitwiseV6(t, testOrV6Neon)
}

func TestOrV3SVE(t *testing.T) {
	testBitwiseV3(t, testOrV3SVE)
}
func TestOrV4SVE(t *testing.T) {
	testBitwiseV4(t, testOrV4SVE)
}

func TestOrV5SVE(t *testing.T) {
	testBitwiseV5(t, testOrV5SVE)
}

func TestOrV6SVE(t *testing.T) {
	testBitwiseV6(t, testOrV6SVE)
}

func TestAndV3Neon(t *testing.T) {
	testBitwiseV3(t, testAndV3Neon)
}
func TestAndV4Neon(t *testing.T) {
	testBitwiseV4(t, testAndV4Neon)
}

func TestAndV5Neon(t *testing.T) {
	testBitwiseV5(t, testAndV5Neon)
}

func TestAndV6Neon(t *testing.T) {
	testBitwiseV6(t, testAndV6Neon)
}

func TestAndV3SVE(t *testing.T) {
	testBitwiseV3(t, testAndV3SVE)
}
func TestAndV4SVE(t *testing.T) {
	testBitwiseV4(t, testAndV4SVE)
}

func TestAndV5SVE(t *testing.T) {
	testBitwiseV5(t, testAndV5SVE)
}

func TestAndV6SVE(t *testing.T) {
	testBitwiseV6(t, testAndV6SVE)
}

func testOrV3Neon(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	{
		initOp(op1, op2, rnd)

		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i]
		}
		Uint64s(actual).OrV3Neon(op1, op2, op3)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV4Neon(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i]
		}
		Uint64s(actual).OrV4Neon(op1, op2, op3, op4)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV5Neon(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op1, op5, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i] | op5[i]
		}
		Uint64s(actual).OrV5Neon(op1, op2, op3, op4, op5)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV6Neon(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op5, op6, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i] | op5[i] | op6[i]
		}
		Uint64s(actual).OrV6Neon(op1, op2, op3, op4, op5, op6)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV3SVE(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)

		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i]
		}
		Uint64s(actual).OrV3SVE(op1, op2, op3)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV4SVE(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i]
		}
		Uint64s(actual).OrV4SVE(op1, op2, op3, op4)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV5SVE(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op1, op5, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i] | op5[i]
		}
		Uint64s(actual).OrV5SVE(op1, op2, op3, op4, op5)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV6SVE(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op5, op6, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i] | op5[i] | op6[i]
		}
		Uint64s(actual).OrV6SVE(op1, op2, op3, op4, op5, op6)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV3Neon(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	{
		initOp(op1, op2, rnd)

		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i]
		}
		Uint64s(actual).AndV3Neon(op1, op2, op3)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV4Neon(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i] & op4[i]
		}
		Uint64s(actual).AndV4Neon(op1, op2, op3, op4)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV5Neon(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op1, op5, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i] & op4[i] & op5[i]
		}
		Uint64s(actual).AndV5Neon(op1, op2, op3, op4, op5)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV6Neon(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op5, op6, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i] & op4[i] & op5[i] & op6[i]
		}
		Uint64s(actual).AndV6Neon(op1, op2, op3, op4, op5, op6)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV3SVE(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)

		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i]
		}
		Uint64s(actual).AndV3SVE(op1, op2, op3)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV4SVE(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i] & op4[i]
		}
		Uint64s(actual).AndV4SVE(op1, op2, op3, op4)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV5SVE(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op1, op5, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i] & op4[i] & op5[i]
		}
		Uint64s(actual).AndV5SVE(op1, op2, op3, op4, op5)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testAndV6SVE(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	if !cpu.ARM64.HasSVE {
		t.Skip("ARM64.SVE not supported")
	}
	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op5, op6, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] & op2[i] & op3[i] & op4[i] & op5[i] & op6[i]
		}
		Uint64s(actual).AndV6SVE(op1, op2, op3, op4, op5, op6)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func BenchmarkAnd_S_Arm64SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1S))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndSVE(op1S, op2S)
	}
}

func BenchmarkAnd_M_Arm64SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndSVE(op1M, op2M)
	}
}

func BenchmarkAnd_XL_Arm64SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op2XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndSVE(op1XL, op2XL)
	}
}

func BenchmarkAnd_S_Arm64Neon(b *testing.B) {
	if !cpu.ARM64.HasASIMD {
		b.Skip("ARM64.SIMD not supported")
	}
	out := make([]uint64, len(op1S))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndNeon(op1S, op2S)
	}
}

func BenchmarkAnd_M_Arm64Neon(b *testing.B) {
	if !cpu.ARM64.HasASIMD {
		b.Skip("ARM64.SIMD not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndNeon(op1M, op2M)
	}
}

func BenchmarkAnd_XL_Arm64Neon(b *testing.B) {
	if !cpu.ARM64.HasASIMD {
		b.Skip("ARM64.SIMD not supported")
	}
	out := make([]uint64, len(op2XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).AndNeon(op1XL, op2XL)
	}
}

func BenchmarkOr_M_V3_Neon(b *testing.B) {
	if !cpu.ARM64.HasASIMD {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV3Neon(op1M, op2M, op3M)
	}
}

func BenchmarkOr_M_V4_Neon(b *testing.B) {
	if !cpu.ARM64.HasASIMD {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV4Neon(op1M, op2M, op3M, op4M)
	}
}

func BenchmarkOr_M_V5_Neon(b *testing.B) {
	if !cpu.ARM64.HasASIMD {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV5Neon(op1M, op2M, op3M, op4M, op5M)
	}
}

func BenchmarkOr_M_V3_SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV3SVE(op1M, op2M, op3M)
	}
}

func BenchmarkOr_M_V4_SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV4SVE(op1M, op2M, op3M, op4M)
	}
}

func BenchmarkOr_M_V5_SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV5SVE(op1M, op2M, op3M, op4M, op5M)
	}
}

func BenchmarkOr_M_V6_SVE(b *testing.B) {
	if !cpu.ARM64.HasSVE {
		b.Skip("ARM64.SVE not supported")
	}
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).OrV6SVE(op1M, op2M, op3M, op4M, op5M, op6M)
	}
}
