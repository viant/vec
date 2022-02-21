package bitwise

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

type testCase struct {
	description string
	size        int
}

func bitwiseTestCases() []testCase {
	var testCases = []testCase{
		{
			description: "4096 bit",
			size:        64,
		},
		{
			description: "256 bit",
			size:        4,
		},
		{
			description: "1024 bit",
			size:        16,
		},
		{
			description: "320 bit",
			size:        5,
		},
		{
			description: "128 bit",
			size:        2,
		},
		{
			description: "128 bit",
			size:        1023,
		},
	}
	return testCases
}

func testAnd(t *testing.T, op1 []uint64, op2 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] & op2[i]
	}
	Uint64s(actual).And(op1, op2)
	assert.Equal(t, expect, actual, testCase.description)
}

func testAndV3(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	initOp(op1, op3, rnd)

	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] & op2[i] & op3[i]
	}
	Uint64s(actual).AndV3(op1, op2, op3)
	assert.Equal(t, expect, actual, testCase.description)
}

func testAndV4(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	initOp(op3, op4, rnd)
	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] & op2[i] & op3[i] & op4[i]
	}
	Uint64s(actual).AndV4(op1, op2, op3, op4)
	assert.Equal(t, expect, actual, testCase.description)
}

func testAndV5(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	initOp(op3, op4, rnd)
	initOp(op5, op4, rnd)

	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] & op2[i] & op3[i] & op4[i] & op5[i]
	}
	Uint64s(actual).AndV5(op1, op2, op3, op4, op5)
	assert.Equal(t, expect, actual, testCase.description)
}

func testAndV6(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	initOp(op3, op4, rnd)
	initOp(op5, op6, rnd)

	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] & op2[i] & op3[i] & op4[i] & op5[i] & op6[i]
	}
	Uint64s(actual).AndV6(op1, op2, op3, op4, op5, op6)
	assert.Equal(t, expect, actual, testCase.description)
}

func TestAnd(t *testing.T) {
	testBitwise(t, testAnd)
}

func TestAndV3(t *testing.T) {
	testBitwiseV3(t, testAndV3)
}

func TestAndV4(t *testing.T) {
	testBitwiseV4(t, testAndV4)
}

func TestAndV5(t *testing.T) {
	testBitwiseV5(t, testAndV5)
}

func TestAndV6(t *testing.T) {
	testBitwiseV6(t, testAndV6)
}

func TestOr(t *testing.T) {
	testBitwise(t, testOr)
}

func TestOrV3(t *testing.T) {
	testBitwiseV3(t, testOrV3)
}

func TestOrV4(t *testing.T) {
	testBitwiseV4(t, testOrV4)
}

func TestOrV5(t *testing.T) {
	testBitwiseV5(t, testOrV5)
}

func TestOrV6(t *testing.T) {
	testBitwiseV6(t, testOrV6)
}

func TestXor(t *testing.T) {
	testBitwise(t, testXOr)
}

func TestXorV3(t *testing.T) {
	testBitwiseV3(t, testXOrV3)
}

func testBitwise(t *testing.T, tester func(t *testing.T, op1 []uint64, op2 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64)) {
	testCases := bitwiseTestCases()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, testCase := range testCases {
		var op1 = make([]uint64, testCase.size)
		var op2 = make([]uint64, testCase.size)
		var expect = make([]uint64, testCase.size)
		var actual = make([]uint64, testCase.size)
		tester(t, op1, op2, rnd, testCase, expect, actual)
	}
}

func testBitwiseV3(t *testing.T, tester func(t *testing.T, op1 []uint64, op2 []uint64, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64)) {
	testCases := bitwiseTestCases()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, testCase := range testCases {
		var op1 = make([]uint64, testCase.size)
		var op2 = make([]uint64, testCase.size)
		var op3 = make([]uint64, testCase.size)
		var expect = make([]uint64, testCase.size)
		var actual = make([]uint64, testCase.size)
		tester(t, op1, op2, op3, rnd, testCase, expect, actual)
	}
}

func testBitwiseV4(t *testing.T, tester func(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64)) {
	testCases := bitwiseTestCases()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, testCase := range testCases {
		var op1 = make([]uint64, testCase.size)
		var op2 = make([]uint64, testCase.size)
		var op3 = make([]uint64, testCase.size)
		var op4 = make([]uint64, testCase.size)
		var expect = make([]uint64, testCase.size)
		var actual = make([]uint64, testCase.size)
		tester(t, op1, op2, op3, op4, rnd, testCase, expect, actual)
	}
}

func testBitwiseV6(t *testing.T, tester func(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64)) {
	testCases := bitwiseTestCases()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, testCase := range testCases {
		var op1 = make([]uint64, testCase.size)
		var op2 = make([]uint64, testCase.size)
		var op3 = make([]uint64, testCase.size)
		var op4 = make([]uint64, testCase.size)
		var op5 = make([]uint64, testCase.size)
		var op6 = make([]uint64, testCase.size)
		var expect = make([]uint64, testCase.size)
		var actual = make([]uint64, testCase.size)
		tester(t, op1, op2, op3, op4, op5, op6, rnd, testCase, expect, actual)
	}
}

func testBitwiseV5(t *testing.T, tester func(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64)) {
	testCases := bitwiseTestCases()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, testCase := range testCases {
		var op1 = make([]uint64, testCase.size)
		var op2 = make([]uint64, testCase.size)
		var op3 = make([]uint64, testCase.size)
		var op4 = make([]uint64, testCase.size)
		var op5 = make([]uint64, testCase.size)
		var expect = make([]uint64, testCase.size)
		var actual = make([]uint64, testCase.size)
		tester(t, op1, op2, op3, op4, op5, rnd, testCase, expect, actual)
	}
}

func testOr(t *testing.T, op1 []uint64, op2 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] | op2[i]
	}
	Uint64s(actual).Or(op1, op2)
	assert.Equal(t, expect, actual, testCase.description)
}

func testOrV3(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)

		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i]
		}
		Uint64s(actual).OrV3(op1, op2, op3)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV4(t *testing.T, op1, op2, op3, op4 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i]
		}
		Uint64s(actual).OrV4(op1, op2, op3, op4)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV5(t *testing.T, op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op1, op5, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i] | op5[i]
		}
		Uint64s(actual).OrV5(op1, op2, op3, op4, op5)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testOrV6(t *testing.T, op1, op2, op3, op4, op5, op6 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {

	{
		initOp(op1, op2, rnd)
		initOp(op3, op4, rnd)
		initOp(op5, op6, rnd)
		for i := 0; i < testCase.size; i++ {
			expect[i] = op1[i] | op2[i] | op3[i] | op4[i] | op5[i] | op6[i]
		}
		Uint64s(actual).OrV6(op1, op2, op3, op4, op5, op6)
		assert.Equal(t, expect, actual, testCase.description)
	}
}

func testXOr(t *testing.T, op1 []uint64, op2 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] ^ op2[i]
	}
	Uint64s(actual).Xor(op1, op2)
	assert.Equal(t, expect, actual, testCase.description)
}

func testXOrV3(t *testing.T, op1, op2, op3 []uint64, rnd *rand.Rand, testCase testCase, expect []uint64, actual []uint64) {
	initOp(op1, op2, rnd)
	for i := 0; i < testCase.size; i++ {
		expect[i] = op1[i] ^ op2[i] ^ op3[i]
	}
	Uint64s(actual).XorV3(op1, op2, op3)
	assert.Equal(t, expect, actual, testCase.description)
}

func initOp(op1, op2 []uint64, rnd *rand.Rand) {
	for i := range op1 {
		op1[i] = rnd.Uint64()
		op2[i] = rnd.Uint64()
	}
}

func initOpV3(op1, op2, op3 []uint64, rnd *rand.Rand) {
	for i := range op1 {
		op1[i] = rnd.Uint64()
		op2[i] = rnd.Uint64()
		op3[i] = rnd.Uint64()
	}
}

func initOpV4(op1, op2, op3, op4 []uint64, rnd *rand.Rand) {
	for i := range op1 {
		op1[i] = rnd.Uint64()
		op2[i] = rnd.Uint64()
		op3[i] = rnd.Uint64()
		op4[i] = rnd.Uint64()

	}
}

func initOpV5(op1, op2, op3, op4, op5 []uint64, rnd *rand.Rand) {
	for i := range op1 {
		op1[i] = rnd.Uint64()
		op2[i] = rnd.Uint64()
		op3[i] = rnd.Uint64()
		op4[i] = rnd.Uint64()
		op5[i] = rnd.Uint64()
	}
}

var op1S, op2S, op3S, op4S, op5S, op6S []uint64
var op1M, op2M, op3M, op4M, op5M, op6M, op7M, op8M, op9M []uint64
var op1XL, op2XL, op3XL, op4XL, op5XL, op6XL []uint64

func init() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	op1S = make([]uint64, 8)
	op2S = make([]uint64, 8)
	op3S = make([]uint64, 8)
	op4S = make([]uint64, 8)
	op5S = make([]uint64, 8)
	op6S = make([]uint64, 8)

	initOpV3(op1S, op2S, op3S, rnd)
	initOpV3(op4S, op5S, op6S, rnd)

	op1M = make([]uint64, 32)
	op2M = make([]uint64, 32)
	op3M = make([]uint64, 32)
	op4M = make([]uint64, 32)
	op5M = make([]uint64, 32)
	op6M = make([]uint64, 32)
	op7M = make([]uint64, 32)
	op8M = make([]uint64, 32)
	op9M = make([]uint64, 32)

	initOpV3(op1M, op2M, op3M, rnd)
	initOpV3(op4M, op5M, op6M, rnd)
	initOpV3(op7M, op8M, op9M, rnd)

	op1XL = make([]uint64, 128)
	op2XL = make([]uint64, 128)
	op3XL = make([]uint64, 128)
	op4XL = make([]uint64, 128)
	op5XL = make([]uint64, 128)
	op6XL = make([]uint64, 128)

	initOpV3(op1XL, op2XL, op3XL, rnd)
	initOpV3(op4XL, op5XL, op6XL, rnd)

}

func BenchmarkAnd_S(b *testing.B) {
	out := make(Uint64s, len(op1S))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out.And(op1S, op2S)
	}
}

func BenchmarkAnd_M(b *testing.B) {
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).And(op1M, op2M)
	}
}

func BenchmarkAnd_XL(b *testing.B) {
	out := make([]uint64, len(op2XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).And(op1XL, op2XL)
	}
}

func BenchmarkAnd_S_Naive(b *testing.B) {
	out := make(Uint64s, len(op1S))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out.and(op1S, op2S)
	}
}

func BenchmarkAnd_M_Naive(b *testing.B) {
	out := make([]uint64, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).and(op1M, op2M)
	}
}

func BenchmarkAnd_XL_Naive(b *testing.B) {
	out := make([]uint64, len(op2XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Uint64s(out).and(op1XL, op2XL)
	}
}

func BenchmarkAnd_M_V3(b *testing.B) {
	out := make(Uint64s, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out.AndV3(op1M, op2M, op3M)
	}
}

func BenchmarkAnd_XL_V3(b *testing.B) {
	out := make(Uint64s, len(op1XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out.AndV3(op1XL, op2XL, op3XL)
	}
}

func BenchmarkAnd_M_V3_Naive(b *testing.B) {
	out := make(Uint64s, len(op1M))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out.andV3(op1M, op2M, op3M)
	}
}

func BenchmarkAnd_XL_V3_Naive(b *testing.B) {
	out := make(Uint64s, len(op1XL))
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		out.andV3(op1XL, op2XL, op3XL)
	}
}

func ExampleUint64s_And() {
	{
		out := Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		out.And(v1, v2)
	}
	{
		out := Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		v3 := []uint64{1, 1, 0, 4, 1, 6, 7, 2}
		out.AndV3(v1, v2, v3)
	}
}

func ExampleUint64s_Or() {
	{
		out := Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		out.Or(v1, v2)
	}
	{
		out := Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		v3 := []uint64{1, 1, 0, 4, 1, 6, 7, 2}
		out.OrV3(v1, v2, v3)
	}
}

func ExampleUint64s_Xor() {
	{
		out := Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		out.Xor(v1, v2)
	}
	{
		out := Uint64s(make([]uint64, 8))
		v1 := []uint64{1, 2, 3, 4, 5, 6, 7, 8}
		v2 := []uint64{1, 7, 3, 4, 3, 6, 7, 2}
		v3 := []uint64{1, 1, 0, 4, 1, 6, 7, 2}
		out.XorV3(v1, v2, v3)
	}
}
