// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package simd

//go:generate go run ./codegen/main.go
import (
	"github.com/klauspost/cpuid/v2"
)

var (
	avx2 = cpuid.CPU.Supports(cpuid.AVX2)
	sve  = cpuid.CPU.Supports(cpuid.SVE)
)

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Sum sums up all of the elements of the slice and returns the value
func Sum[T number](input []T) T {
	switch v := any(input).(type) {
	case []int8:
		return T(SumInt8s(v))
	case []int16:
		return T(SumInt16s(v))
	case []int32:
		return T(SumInt32s(v))
	case []int64:
		return T(SumInt64s(v))
	case []uint8:
		return T(SumUint8s(v))
	case []uint16:
		return T(SumUint16s(v))
	case []uint32:
		return T(SumUint32s(v))
	case []uint64:
		return T(SumUint64s(v))
	case []float32:
		return T(SumFloat32s(v))
	case []float64:
		return T(SumFloat64s(v))
	default:
		return sum(input)
	}
}

// Sum sums up all of the elements of the slice and returns the value
func sum[T number](input []T) (sum T) {
	for _, v := range input {
		sum += v
	}
	return
}

// Min returns the smallest element value in the slice
func Min[T number](input []T) T {
	switch v := any(input).(type) {
	case []int8:
		return T(MinInt8s(v))
	case []int16:
		return T(MinInt16s(v))
	case []int32:
		return T(MinInt32s(v))
	case []int64:
		return T(MinInt64s(v))
	case []uint8:
		return T(MinUint8s(v))
	case []uint16:
		return T(MinUint16s(v))
	case []uint32:
		return T(MinUint32s(v))
	case []uint64:
		return T(MinUint64s(v))
	case []float32:
		return T(MinFloat32s(v))
	case []float64:
		return T(MinFloat64s(v))
	default:
		return min(input)
	}
}

// Min returns the smallest element value in the slice
func min[T number](input []T) T {
	min := input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// Max returns the largest element value in the slice
func Max[T number](input []T) T {
	switch v := any(input).(type) {
	case []int8:
		return T(MaxInt8s(v))
	case []int16:
		return T(MaxInt16s(v))
	case []int32:
		return T(MaxInt32s(v))
	case []int64:
		return T(MaxInt64s(v))
	case []uint8:
		return T(MaxUint8s(v))
	case []uint16:
		return T(MaxUint16s(v))
	case []uint32:
		return T(MaxUint32s(v))
	case []uint64:
		return T(MaxUint64s(v))
	case []float32:
		return T(MaxFloat32s(v))
	case []float64:
		return T(MaxFloat64s(v))
	default:
		return max(input)
	}
}

// Max returns the largest element value in the slice
func max[T number](input []T) T {
	max := input[0]
	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Add adds input1 to input2 and writes back the result into dst slice
func add[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v + input2[i]
	}
	return dst
}

// Sub subtracts input2 from input1 and writes back the result into dst slice
func sub[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v - input2[i]
	}
	return dst
}

// Mul multiplies input1 by input2 and writes back the result into dst slice
func mul[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v * input2[i]
	}
	return dst
}

// Div divides input1 by input2 and writes back the result into dst slice
func div[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v / input2[i]
	}
	return dst
}
