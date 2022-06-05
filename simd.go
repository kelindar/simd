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
	~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func sum[T number](input []T) (sum T) {
	for _, v := range input {
		sum += v
	}
	return
}

func max[T number](input []T) T {
	max := input[0]
	for _, v := range input[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func min[T number](input []T) T {
	min := input[0]
	for _, v := range input[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func add[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v + input2[i]
	}
	return dst
}

func sub[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v - input2[i]
	}
	return dst
}

func mul[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v * input2[i]
	}
	return dst
}

func div[T number](dst, input1, input2 []T) []T {
	for i, v := range input1 {
		dst[i] = v / input2[i]
	}
	return dst
}
