// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package simd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ops[T Number] struct {
	sum func([]T) T
	min func([]T) T
	max func([]T) T
	add func([]T, []T, []T) []T
	sub func([]T, []T, []T) []T
	mul func([]T, []T, []T) []T
	div func([]T, []T, []T) []T
}

type suite struct {
	name         string
	benchmark    func(*testing.B)
	testOps      func(*testing.T)
	testFallback func(*testing.T)
}

func makeSuite[T Number](name string, fn ops[T]) suite {
	return suite{
		name: name,
		benchmark: func(b *testing.B) {
			benchmarkOps(b, name, fn)
		},
		testOps: func(t *testing.T) {
			testOps(t, fn)
		},
		testFallback: func(t *testing.T) {
			withHardwareDisabled(func() {
				testOps(t, fn)
			})
		},
	}
}

func suites() []suite {
	return []suite{
		makeSuite("uint8", ops[uint8]{sum: SumUint8s, min: MinUint8s, max: MaxUint8s, add: AddUint8s, sub: SubUint8s, mul: MulUint8s, div: DivUint8s}),
		makeSuite("uint16", ops[uint16]{sum: SumUint16s, min: MinUint16s, max: MaxUint16s, add: AddUint16s, sub: SubUint16s, mul: MulUint16s, div: DivUint16s}),
		makeSuite("uint32", ops[uint32]{sum: SumUint32s, min: MinUint32s, max: MaxUint32s, add: AddUint32s, sub: SubUint32s, mul: MulUint32s, div: DivUint32s}),
		makeSuite("uint64", ops[uint64]{sum: SumUint64s, min: MinUint64s, max: MaxUint64s, add: AddUint64s, sub: SubUint64s, mul: MulUint64s, div: DivUint64s}),
		makeSuite("int8", ops[int8]{sum: SumInt8s, min: MinInt8s, max: MaxInt8s, add: AddInt8s, sub: SubInt8s, mul: MulInt8s, div: DivInt8s}),
		makeSuite("int16", ops[int16]{sum: SumInt16s, min: MinInt16s, max: MaxInt16s, add: AddInt16s, sub: SubInt16s, mul: MulInt16s, div: DivInt16s}),
		makeSuite("int32", ops[int32]{sum: SumInt32s, min: MinInt32s, max: MaxInt32s, add: AddInt32s, sub: SubInt32s, mul: MulInt32s, div: DivInt32s}),
		makeSuite("int64", ops[int64]{sum: SumInt64s, min: MinInt64s, max: MaxInt64s, add: AddInt64s, sub: SubInt64s, mul: MulInt64s, div: DivInt64s}),
		makeSuite("float32", ops[float32]{sum: SumFloat32s, min: MinFloat32s, max: MaxFloat32s, add: AddFloat32s, sub: SubFloat32s, mul: MulFloat32s, div: DivFloat32s}),
		makeSuite("float64", ops[float64]{sum: SumFloat64s, min: MinFloat64s, max: MaxFloat64s, add: AddFloat64s, sub: SubFloat64s, mul: MulFloat64s, div: DivFloat64s}),
	}
}

func BenchmarkSIMD(b *testing.B) {
	for _, tc := range suites() {
		b.Run(tc.name, tc.benchmark)
	}
}

func TestSIMDOps(t *testing.T) {
	for _, tc := range suites() {
		t.Run(tc.name, tc.testOps)
	}
}

func TestSIMDFallback(t *testing.T) {
	for _, tc := range suites() {
		t.Run(tc.name, tc.testFallback)
	}
}

func benchmarkOps[T Number](b *testing.B, typ string, fn ops[T]) {
	result := make([]Result, 0, 64)
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[T](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			var out T
			for i := 0; i < b.N; i++ {
				out = fn.sum(vector)
			}
			assert.NotZero(b, out)
		}))
		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			var out T
			for i := 0; i < b.N; i++ {
				out = fn.min(vector)
			}
			assert.NotZero(b, out)
		}))
		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			var out T
			for i := 0; i < b.N; i++ {
				out = fn.max(vector)
			}
			assert.NotZero(b, out)
		}))

		output := make([]T, count)
		input1 := makeVector[T](count)
		input2 := makeVector[T](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var out []T
			for i := 0; i < b.N; i++ {
				out = fn.add(output, input1, input2)
			}
			assert.NotEmpty(b, out)
		}))
		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var out []T
			for i := 0; i < b.N; i++ {
				out = fn.sub(output, input1, input2)
			}
			assert.NotEmpty(b, out)
		}))
		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var out []T
			for i := 0; i < b.N; i++ {
				out = fn.mul(output, input1, input2)
			}
			assert.NotEmpty(b, out)
		}))
		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var out []T
			for i := 0; i < b.N; i++ {
				out = fn.div(output, input1, input2)
			}
			assert.NotEmpty(b, out)
		}))
	}

	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n", r.Type, r.Name, r.Size, r.Rate, r.Speedup)
	}
}

func testOps[T Number](t *testing.T, fn ops[T]) {
	input := makeVector[T](70)
	assert.EqualValues(t, sum(input), fn.sum(input))
	assert.EqualValues(t, min(input), fn.min(input))
	assert.EqualValues(t, max(input), fn.max(input))

	input1 := makeVector[T](70)
	input2 := makeVector[T](70)
	assert.EqualValues(t, add(make([]T, 70), input1, input2), fn.add(make([]T, 70), input1, input2))
	assert.EqualValues(t, sub(make([]T, 70), input1, input2), fn.sub(make([]T, 70), input1, input2))
	assert.EqualValues(t, mul(make([]T, 70), input1, input2), fn.mul(make([]T, 70), input1, input2))
	assert.InDeltaSlice(t, div(make([]T, 70), input1, input2), fn.div(make([]T, 70), input1, input2), 0.01)
}

func withHardwareDisabled(fn func()) {
	prev := hardware
	hardware = false
	defer func() {
		hardware = prev
	}()
	fn()
}
