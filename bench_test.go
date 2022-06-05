// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package simd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ---------------------------------- Benchmark Uint8 ----------------------------------

func BenchmarkUint8(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "uint8"
	for _, count := range []int{100} {
		vector := makeVector[uint8](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := uint8(0)
			for i := 0; i < b.N; i++ {
				result = SumUint8s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := uint8(0)
			for i := 0; i < b.N; i++ {
				result = MinUint8s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := uint8(0)
			for i := 0; i < b.N; i++ {
				result = MaxUint8s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]uint8, count)
		input1 := makeVector[uint8](count)
		input2 := makeVector[uint8](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []uint8
			for i := 0; i < b.N; i++ {
				result = AddUint8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []uint8
			for i := 0; i < b.N; i++ {
				result = SubUint8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []uint8
			for i := 0; i < b.N; i++ {
				result = MulUint8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []uint8
			for i := 0; i < b.N; i++ {
				result = DivUint8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Uint8 ----------------------------------

func TestUint8_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[uint8](70)
		expect := sum(input)
		result := SumUint8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint8](70)
		expect := min(input)
		result := MinUint8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint8](70)
		expect := max(input)
		result := MaxUint8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := add(make([]uint8, 70), input1, input2)
		result := AddUint8s(make([]uint8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := sub(make([]uint8, 70), input1, input2)
		result := SubUint8s(make([]uint8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := mul(make([]uint8, 70), input1, input2)
		result := MulUint8s(make([]uint8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := div(make([]uint8, 70), input1, input2)
		result := DivUint8s(make([]uint8, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Uint8 ----------------------------------

func TestUint8_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[uint8](70)
		expect := sum(input)
		result := SumUint8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint8](70)
		expect := min(input)
		result := MinUint8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint8](70)
		expect := max(input)
		result := MaxUint8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := add(make([]uint8, 70), input1, input2)
		result := AddUint8s(make([]uint8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := sub(make([]uint8, 70), input1, input2)
		result := SubUint8s(make([]uint8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := mul(make([]uint8, 70), input1, input2)
		result := MulUint8s(make([]uint8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint8](70)
		input2 := makeVector[uint8](70)
		expect := div(make([]uint8, 70), input1, input2)
		result := DivUint8s(make([]uint8, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Uint16 ----------------------------------

func BenchmarkUint16(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "uint16"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[uint16](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := uint16(0)
			for i := 0; i < b.N; i++ {
				result = SumUint16s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := uint16(0)
			for i := 0; i < b.N; i++ {
				result = MinUint16s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := uint16(0)
			for i := 0; i < b.N; i++ {
				result = MaxUint16s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]uint16, count)
		input1 := makeVector[uint16](count)
		input2 := makeVector[uint16](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []uint16
			for i := 0; i < b.N; i++ {
				result = AddUint16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []uint16
			for i := 0; i < b.N; i++ {
				result = SubUint16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []uint16
			for i := 0; i < b.N; i++ {
				result = MulUint16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []uint16
			for i := 0; i < b.N; i++ {
				result = DivUint16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Uint16 ----------------------------------

func TestUint16_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[uint16](70)
		expect := sum(input)
		result := SumUint16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint16](70)
		expect := min(input)
		result := MinUint16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint16](70)
		expect := max(input)
		result := MaxUint16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := add(make([]uint16, 70), input1, input2)
		result := AddUint16s(make([]uint16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := sub(make([]uint16, 70), input1, input2)
		result := SubUint16s(make([]uint16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := mul(make([]uint16, 70), input1, input2)
		result := MulUint16s(make([]uint16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := div(make([]uint16, 70), input1, input2)
		result := DivUint16s(make([]uint16, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Uint16 ----------------------------------

func TestUint16_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[uint16](70)
		expect := sum(input)
		result := SumUint16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint16](70)
		expect := min(input)
		result := MinUint16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint16](70)
		expect := max(input)
		result := MaxUint16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := add(make([]uint16, 70), input1, input2)
		result := AddUint16s(make([]uint16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := sub(make([]uint16, 70), input1, input2)
		result := SubUint16s(make([]uint16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := mul(make([]uint16, 70), input1, input2)
		result := MulUint16s(make([]uint16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint16](70)
		input2 := makeVector[uint16](70)
		expect := div(make([]uint16, 70), input1, input2)
		result := DivUint16s(make([]uint16, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Uint32 ----------------------------------

func BenchmarkUint32(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "uint32"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[uint32](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := uint32(0)
			for i := 0; i < b.N; i++ {
				result = SumUint32s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := uint32(0)
			for i := 0; i < b.N; i++ {
				result = MinUint32s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := uint32(0)
			for i := 0; i < b.N; i++ {
				result = MaxUint32s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]uint32, count)
		input1 := makeVector[uint32](count)
		input2 := makeVector[uint32](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []uint32
			for i := 0; i < b.N; i++ {
				result = AddUint32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []uint32
			for i := 0; i < b.N; i++ {
				result = SubUint32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []uint32
			for i := 0; i < b.N; i++ {
				result = MulUint32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []uint32
			for i := 0; i < b.N; i++ {
				result = DivUint32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Uint32 ----------------------------------

func TestUint32_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[uint32](70)
		expect := sum(input)
		result := SumUint32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint32](70)
		expect := min(input)
		result := MinUint32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint32](70)
		expect := max(input)
		result := MaxUint32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := add(make([]uint32, 70), input1, input2)
		result := AddUint32s(make([]uint32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := sub(make([]uint32, 70), input1, input2)
		result := SubUint32s(make([]uint32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := mul(make([]uint32, 70), input1, input2)
		result := MulUint32s(make([]uint32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := div(make([]uint32, 70), input1, input2)
		result := DivUint32s(make([]uint32, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Uint32 ----------------------------------

func TestUint32_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[uint32](70)
		expect := sum(input)
		result := SumUint32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint32](70)
		expect := min(input)
		result := MinUint32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint32](70)
		expect := max(input)
		result := MaxUint32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := add(make([]uint32, 70), input1, input2)
		result := AddUint32s(make([]uint32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := sub(make([]uint32, 70), input1, input2)
		result := SubUint32s(make([]uint32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := mul(make([]uint32, 70), input1, input2)
		result := MulUint32s(make([]uint32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint32](70)
		input2 := makeVector[uint32](70)
		expect := div(make([]uint32, 70), input1, input2)
		result := DivUint32s(make([]uint32, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Uint64 ----------------------------------

func BenchmarkUint64(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "uint64"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[uint64](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := uint64(0)
			for i := 0; i < b.N; i++ {
				result = SumUint64s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := uint64(0)
			for i := 0; i < b.N; i++ {
				result = MinUint64s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := uint64(0)
			for i := 0; i < b.N; i++ {
				result = MaxUint64s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]uint64, count)
		input1 := makeVector[uint64](count)
		input2 := makeVector[uint64](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []uint64
			for i := 0; i < b.N; i++ {
				result = AddUint64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []uint64
			for i := 0; i < b.N; i++ {
				result = SubUint64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []uint64
			for i := 0; i < b.N; i++ {
				result = MulUint64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []uint64
			for i := 0; i < b.N; i++ {
				result = DivUint64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Uint64 ----------------------------------

func TestUint64_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[uint64](70)
		expect := sum(input)
		result := SumUint64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint64](70)
		expect := min(input)
		result := MinUint64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint64](70)
		expect := max(input)
		result := MaxUint64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := add(make([]uint64, 70), input1, input2)
		result := AddUint64s(make([]uint64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := sub(make([]uint64, 70), input1, input2)
		result := SubUint64s(make([]uint64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := mul(make([]uint64, 70), input1, input2)
		result := MulUint64s(make([]uint64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := div(make([]uint64, 70), input1, input2)
		result := DivUint64s(make([]uint64, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Uint64 ----------------------------------

func TestUint64_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[uint64](70)
		expect := sum(input)
		result := SumUint64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[uint64](70)
		expect := min(input)
		result := MinUint64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[uint64](70)
		expect := max(input)
		result := MaxUint64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := add(make([]uint64, 70), input1, input2)
		result := AddUint64s(make([]uint64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := sub(make([]uint64, 70), input1, input2)
		result := SubUint64s(make([]uint64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := mul(make([]uint64, 70), input1, input2)
		result := MulUint64s(make([]uint64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[uint64](70)
		input2 := makeVector[uint64](70)
		expect := div(make([]uint64, 70), input1, input2)
		result := DivUint64s(make([]uint64, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Int8 ----------------------------------

func BenchmarkInt8(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "int8"
	for _, count := range []int{100} {
		vector := makeVector[int8](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := int8(0)
			for i := 0; i < b.N; i++ {
				result = SumInt8s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := int8(0)
			for i := 0; i < b.N; i++ {
				result = MinInt8s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := int8(0)
			for i := 0; i < b.N; i++ {
				result = MaxInt8s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]int8, count)
		input1 := makeVector[int8](count)
		input2 := makeVector[int8](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []int8
			for i := 0; i < b.N; i++ {
				result = AddInt8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []int8
			for i := 0; i < b.N; i++ {
				result = SubInt8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []int8
			for i := 0; i < b.N; i++ {
				result = MulInt8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []int8
			for i := 0; i < b.N; i++ {
				result = DivInt8s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Int8 ----------------------------------

func TestInt8_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[int8](70)
		expect := sum(input)
		result := SumInt8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int8](70)
		expect := min(input)
		result := MinInt8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int8](70)
		expect := max(input)
		result := MaxInt8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := add(make([]int8, 70), input1, input2)
		result := AddInt8s(make([]int8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := sub(make([]int8, 70), input1, input2)
		result := SubInt8s(make([]int8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := mul(make([]int8, 70), input1, input2)
		result := MulInt8s(make([]int8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := div(make([]int8, 70), input1, input2)
		result := DivInt8s(make([]int8, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Int8 ----------------------------------

func TestInt8_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[int8](70)
		expect := sum(input)
		result := SumInt8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int8](70)
		expect := min(input)
		result := MinInt8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int8](70)
		expect := max(input)
		result := MaxInt8s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := add(make([]int8, 70), input1, input2)
		result := AddInt8s(make([]int8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := sub(make([]int8, 70), input1, input2)
		result := SubInt8s(make([]int8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := mul(make([]int8, 70), input1, input2)
		result := MulInt8s(make([]int8, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int8](70)
		input2 := makeVector[int8](70)
		expect := div(make([]int8, 70), input1, input2)
		result := DivInt8s(make([]int8, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Int16 ----------------------------------

func BenchmarkInt16(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "int16"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[int16](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := int16(0)
			for i := 0; i < b.N; i++ {
				result = SumInt16s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := int16(0)
			for i := 0; i < b.N; i++ {
				result = MinInt16s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := int16(0)
			for i := 0; i < b.N; i++ {
				result = MaxInt16s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]int16, count)
		input1 := makeVector[int16](count)
		input2 := makeVector[int16](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []int16
			for i := 0; i < b.N; i++ {
				result = AddInt16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []int16
			for i := 0; i < b.N; i++ {
				result = SubInt16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []int16
			for i := 0; i < b.N; i++ {
				result = MulInt16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []int16
			for i := 0; i < b.N; i++ {
				result = DivInt16s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Int16 ----------------------------------

func TestInt16_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[int16](70)
		expect := sum(input)
		result := SumInt16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int16](70)
		expect := min(input)
		result := MinInt16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int16](70)
		expect := max(input)
		result := MaxInt16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := add(make([]int16, 70), input1, input2)
		result := AddInt16s(make([]int16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := sub(make([]int16, 70), input1, input2)
		result := SubInt16s(make([]int16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := mul(make([]int16, 70), input1, input2)
		result := MulInt16s(make([]int16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := div(make([]int16, 70), input1, input2)
		result := DivInt16s(make([]int16, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Int16 ----------------------------------

func TestInt16_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[int16](70)
		expect := sum(input)
		result := SumInt16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int16](70)
		expect := min(input)
		result := MinInt16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int16](70)
		expect := max(input)
		result := MaxInt16s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := add(make([]int16, 70), input1, input2)
		result := AddInt16s(make([]int16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := sub(make([]int16, 70), input1, input2)
		result := SubInt16s(make([]int16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := mul(make([]int16, 70), input1, input2)
		result := MulInt16s(make([]int16, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int16](70)
		input2 := makeVector[int16](70)
		expect := div(make([]int16, 70), input1, input2)
		result := DivInt16s(make([]int16, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Int32 ----------------------------------

func BenchmarkInt32(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "int32"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[int32](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := int32(0)
			for i := 0; i < b.N; i++ {
				result = SumInt32s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := int32(0)
			for i := 0; i < b.N; i++ {
				result = MinInt32s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := int32(0)
			for i := 0; i < b.N; i++ {
				result = MaxInt32s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]int32, count)
		input1 := makeVector[int32](count)
		input2 := makeVector[int32](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []int32
			for i := 0; i < b.N; i++ {
				result = AddInt32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []int32
			for i := 0; i < b.N; i++ {
				result = SubInt32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []int32
			for i := 0; i < b.N; i++ {
				result = MulInt32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []int32
			for i := 0; i < b.N; i++ {
				result = DivInt32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Int32 ----------------------------------

func TestInt32_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[int32](70)
		expect := sum(input)
		result := SumInt32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int32](70)
		expect := min(input)
		result := MinInt32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int32](70)
		expect := max(input)
		result := MaxInt32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := add(make([]int32, 70), input1, input2)
		result := AddInt32s(make([]int32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := sub(make([]int32, 70), input1, input2)
		result := SubInt32s(make([]int32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := mul(make([]int32, 70), input1, input2)
		result := MulInt32s(make([]int32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := div(make([]int32, 70), input1, input2)
		result := DivInt32s(make([]int32, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Int32 ----------------------------------

func TestInt32_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[int32](70)
		expect := sum(input)
		result := SumInt32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int32](70)
		expect := min(input)
		result := MinInt32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int32](70)
		expect := max(input)
		result := MaxInt32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := add(make([]int32, 70), input1, input2)
		result := AddInt32s(make([]int32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := sub(make([]int32, 70), input1, input2)
		result := SubInt32s(make([]int32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := mul(make([]int32, 70), input1, input2)
		result := MulInt32s(make([]int32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int32](70)
		input2 := makeVector[int32](70)
		expect := div(make([]int32, 70), input1, input2)
		result := DivInt32s(make([]int32, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Int64 ----------------------------------

func BenchmarkInt64(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "int64"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[int64](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := int64(0)
			for i := 0; i < b.N; i++ {
				result = SumInt64s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := int64(0)
			for i := 0; i < b.N; i++ {
				result = MinInt64s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := int64(0)
			for i := 0; i < b.N; i++ {
				result = MaxInt64s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]int64, count)
		input1 := makeVector[int64](count)
		input2 := makeVector[int64](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []int64
			for i := 0; i < b.N; i++ {
				result = AddInt64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []int64
			for i := 0; i < b.N; i++ {
				result = SubInt64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []int64
			for i := 0; i < b.N; i++ {
				result = MulInt64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []int64
			for i := 0; i < b.N; i++ {
				result = DivInt64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Int64 ----------------------------------

func TestInt64_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[int64](70)
		expect := sum(input)
		result := SumInt64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int64](70)
		expect := min(input)
		result := MinInt64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int64](70)
		expect := max(input)
		result := MaxInt64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := add(make([]int64, 70), input1, input2)
		result := AddInt64s(make([]int64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := sub(make([]int64, 70), input1, input2)
		result := SubInt64s(make([]int64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := mul(make([]int64, 70), input1, input2)
		result := MulInt64s(make([]int64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := div(make([]int64, 70), input1, input2)
		result := DivInt64s(make([]int64, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Int64 ----------------------------------

func TestInt64_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[int64](70)
		expect := sum(input)
		result := SumInt64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[int64](70)
		expect := min(input)
		result := MinInt64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[int64](70)
		expect := max(input)
		result := MaxInt64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := add(make([]int64, 70), input1, input2)
		result := AddInt64s(make([]int64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := sub(make([]int64, 70), input1, input2)
		result := SubInt64s(make([]int64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := mul(make([]int64, 70), input1, input2)
		result := MulInt64s(make([]int64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[int64](70)
		input2 := makeVector[int64](70)
		expect := div(make([]int64, 70), input1, input2)
		result := DivInt64s(make([]int64, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Float32 ----------------------------------

func BenchmarkFloat32(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "float32"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[float32](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := float32(0)
			for i := 0; i < b.N; i++ {
				result = SumFloat32s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := float32(0)
			for i := 0; i < b.N; i++ {
				result = MinFloat32s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := float32(0)
			for i := 0; i < b.N; i++ {
				result = MaxFloat32s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]float32, count)
		input1 := makeVector[float32](count)
		input2 := makeVector[float32](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []float32
			for i := 0; i < b.N; i++ {
				result = AddFloat32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []float32
			for i := 0; i < b.N; i++ {
				result = SubFloat32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []float32
			for i := 0; i < b.N; i++ {
				result = MulFloat32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []float32
			for i := 0; i < b.N; i++ {
				result = DivFloat32s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Float32 ----------------------------------

func TestFloat32_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[float32](70)
		expect := sum(input)
		result := SumFloat32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[float32](70)
		expect := min(input)
		result := MinFloat32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[float32](70)
		expect := max(input)
		result := MaxFloat32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := add(make([]float32, 70), input1, input2)
		result := AddFloat32s(make([]float32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := sub(make([]float32, 70), input1, input2)
		result := SubFloat32s(make([]float32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := mul(make([]float32, 70), input1, input2)
		result := MulFloat32s(make([]float32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := div(make([]float32, 70), input1, input2)
		result := DivFloat32s(make([]float32, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Float32 ----------------------------------

func TestFloat32_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[float32](70)
		expect := sum(input)
		result := SumFloat32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[float32](70)
		expect := min(input)
		result := MinFloat32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[float32](70)
		expect := max(input)
		result := MaxFloat32s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := add(make([]float32, 70), input1, input2)
		result := AddFloat32s(make([]float32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := sub(make([]float32, 70), input1, input2)
		result := SubFloat32s(make([]float32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := mul(make([]float32, 70), input1, input2)
		result := MulFloat32s(make([]float32, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[float32](70)
		input2 := makeVector[float32](70)
		expect := div(make([]float32, 70), input1, input2)
		result := DivFloat32s(make([]float32, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Benchmark Float64 ----------------------------------

func BenchmarkFloat64(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "float64"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[float64](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := float64(0)
			for i := 0; i < b.N; i++ {
				result = SumFloat64s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := float64(0)
			for i := 0; i < b.N; i++ {
				result = MinFloat64s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := float64(0)
			for i := 0; i < b.N; i++ {
				result = MaxFloat64s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]float64, count)
		input1 := makeVector[float64](count)
		input2 := makeVector[float64](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []float64
			for i := 0; i < b.N; i++ {
				result = AddFloat64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []float64
			for i := 0; i < b.N; i++ {
				result = SubFloat64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []float64
			for i := 0; i < b.N; i++ {
				result = MulFloat64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []float64
			for i := 0; i < b.N; i++ {
				result = DivFloat64s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))
	}

	// Print out the result and respective speed-up
	fmt.Println()
	fmt.Println("   TYPE    OP    SIZE     RATE        SPEEDUP")
	for _, r := range result {
		fmt.Printf("%7s %5s %7d %8.2f ns/op %7.2fx\n",
			r.Type, r.Name, r.Size, r.Rate, r.Speedup,
		)
	}
}

// ---------------------------------- Test Float64 ----------------------------------

func TestFloat64_Ops(t *testing.T) {
	{ // Sum
		input := makeVector[float64](70)
		expect := sum(input)
		result := SumFloat64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[float64](70)
		expect := min(input)
		result := MinFloat64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[float64](70)
		expect := max(input)
		result := MaxFloat64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := add(make([]float64, 70), input1, input2)
		result := AddFloat64s(make([]float64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := sub(make([]float64, 70), input1, input2)
		result := SubFloat64s(make([]float64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := mul(make([]float64, 70), input1, input2)
		result := MulFloat64s(make([]float64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := div(make([]float64, 70), input1, input2)
		result := DivFloat64s(make([]float64, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}

// ---------------------------------- Test Fallback Float64 ----------------------------------

func TestFloat64_Fallback_Sum(t *testing.T) {
	defer func(v bool) {
		avx2 = v
	}(avx2)
	avx2 = false

	{ // Sum
		input := makeVector[float64](70)
		expect := sum(input)
		result := SumFloat64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Min
		input := makeVector[float64](70)
		expect := min(input)
		result := MinFloat64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Max
		input := makeVector[float64](70)
		expect := max(input)
		result := MaxFloat64s(input)
		assert.EqualValues(t, expect, result)
	}

	{ // Add
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := add(make([]float64, 70), input1, input2)
		result := AddFloat64s(make([]float64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Sub
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := sub(make([]float64, 70), input1, input2)
		result := SubFloat64s(make([]float64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Mul
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := mul(make([]float64, 70), input1, input2)
		result := MulFloat64s(make([]float64, 70), input1, input2)
		assert.EqualValues(t, expect, result)
	}

	{ // Div
		input1 := makeVector[float64](70)
		input2 := makeVector[float64](70)
		expect := div(make([]float64, 70), input1, input2)
		result := DivFloat64s(make([]float64, 70), input1, input2)
		assert.InDeltaSlice(t, expect, result, 0.01)
	}
}
