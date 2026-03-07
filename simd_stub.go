// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build noasm || !(amd64 || arm64)
// +build noasm !amd64,!arm64

package simd

import "unsafe"

func init() {
	hardware = false
}

func _uint8_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint8_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint8_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint8_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint8_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint8_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint8_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint16_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint32_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _uint64_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int8_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int16_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int32_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _int64_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float32_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_sum(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_min(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_max(input unsafe.Pointer, result unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_add(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_sub(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_mul(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

func _float64_div(input1 unsafe.Pointer, input2 unsafe.Pointer, output unsafe.Pointer, size uint64) {
	panic("simd: assembly not available")
}

