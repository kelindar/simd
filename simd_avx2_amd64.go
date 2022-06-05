// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build amd64
// +build amd64

package simd

import (
	"unsafe"
)


//go:noescape
func _uint8_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint8_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint8_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint8_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint8_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint8_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint8_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _uint16_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint16_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint16_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint16_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint16_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint16_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint16_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _uint32_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint32_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint32_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint32_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint32_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint32_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint32_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _uint64_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint64_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint64_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _uint64_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint64_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint64_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _uint64_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _int8_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int8_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int8_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int8_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int8_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int8_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int8_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _int16_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int16_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int16_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int16_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int16_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int16_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int16_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _int32_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int32_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int32_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int32_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int32_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int32_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int32_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _int64_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int64_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int64_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _int64_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int64_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int64_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _int64_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _float32_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _float32_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _float32_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _float32_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _float32_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _float32_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _float32_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

//go:noescape
func _float64_avx2_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _float64_avx2_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _float64_avx2_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _float64_avx2_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _float64_avx2_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _float64_avx2_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _float64_avx2_div(input1, input2, output unsafe.Pointer, info uint64)

