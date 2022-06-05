// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build amd64
// +build amd64

package simd

import (
	"unsafe"
)


// ---------------------------------- Uint8 ----------------------------------

// SumUint8s sums up all of the elements of the slice and returns the value
func SumUint8s(input []uint8) (out uint8) {
	switch {
	case avx2:
		_uint8_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinUint8s returns the smallest element value in the slice
func MinUint8s(input []uint8) (out uint8) {
	switch {
	case avx2:
		_uint8_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxUint8s returns the largest element value in the slice
func MaxUint8s(input []uint8) (out uint8) {
	switch {
	case avx2:
		_uint8_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddUint8s adds input1 to input2 and writes back the result into dst slice
func AddUint8s(dst, input1, input2 []uint8) []uint8 {
	if avx2 {
		_uint8_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubUint8s subtracts input2 from input1 and writes back the result into dst slice
func SubUint8s(dst, input1, input2 []uint8) []uint8 {
	if avx2 {
		_uint8_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulUint8s multiplies input1 by input2 and writes back the result into dst slice
func MulUint8s(dst, input1, input2 []uint8) []uint8 {
	if avx2 {
		_uint8_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivUint8s divides input1 by input2 and writes back the result into dst slice
func DivUint8s(dst, input1, input2 []uint8) []uint8 {
	if avx2 {
		_uint8_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Uint16 ----------------------------------

// SumUint16s sums up all of the elements of the slice and returns the value
func SumUint16s(input []uint16) (out uint16) {
	switch {
	case avx2:
		_uint16_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinUint16s returns the smallest element value in the slice
func MinUint16s(input []uint16) (out uint16) {
	switch {
	case avx2:
		_uint16_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxUint16s returns the largest element value in the slice
func MaxUint16s(input []uint16) (out uint16) {
	switch {
	case avx2:
		_uint16_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddUint16s adds input1 to input2 and writes back the result into dst slice
func AddUint16s(dst, input1, input2 []uint16) []uint16 {
	if avx2 {
		_uint16_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubUint16s subtracts input2 from input1 and writes back the result into dst slice
func SubUint16s(dst, input1, input2 []uint16) []uint16 {
	if avx2 {
		_uint16_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulUint16s multiplies input1 by input2 and writes back the result into dst slice
func MulUint16s(dst, input1, input2 []uint16) []uint16 {
	if avx2 {
		_uint16_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivUint16s divides input1 by input2 and writes back the result into dst slice
func DivUint16s(dst, input1, input2 []uint16) []uint16 {
	if avx2 {
		_uint16_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Uint32 ----------------------------------

// SumUint32s sums up all of the elements of the slice and returns the value
func SumUint32s(input []uint32) (out uint32) {
	switch {
	case avx2:
		_uint32_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinUint32s returns the smallest element value in the slice
func MinUint32s(input []uint32) (out uint32) {
	switch {
	case avx2:
		_uint32_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxUint32s returns the largest element value in the slice
func MaxUint32s(input []uint32) (out uint32) {
	switch {
	case avx2:
		_uint32_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddUint32s adds input1 to input2 and writes back the result into dst slice
func AddUint32s(dst, input1, input2 []uint32) []uint32 {
	if avx2 {
		_uint32_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubUint32s subtracts input2 from input1 and writes back the result into dst slice
func SubUint32s(dst, input1, input2 []uint32) []uint32 {
	if avx2 {
		_uint32_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulUint32s multiplies input1 by input2 and writes back the result into dst slice
func MulUint32s(dst, input1, input2 []uint32) []uint32 {
	if avx2 {
		_uint32_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivUint32s divides input1 by input2 and writes back the result into dst slice
func DivUint32s(dst, input1, input2 []uint32) []uint32 {
	if avx2 {
		_uint32_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Uint64 ----------------------------------

// SumUint64s sums up all of the elements of the slice and returns the value
func SumUint64s(input []uint64) (out uint64) {
	switch {
	case avx2:
		_uint64_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinUint64s returns the smallest element value in the slice
func MinUint64s(input []uint64) (out uint64) {
	switch {
	case avx2:
		_uint64_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxUint64s returns the largest element value in the slice
func MaxUint64s(input []uint64) (out uint64) {
	switch {
	case avx2:
		_uint64_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddUint64s adds input1 to input2 and writes back the result into dst slice
func AddUint64s(dst, input1, input2 []uint64) []uint64 {
	if avx2 {
		_uint64_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubUint64s subtracts input2 from input1 and writes back the result into dst slice
func SubUint64s(dst, input1, input2 []uint64) []uint64 {
	if avx2 {
		_uint64_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulUint64s multiplies input1 by input2 and writes back the result into dst slice
func MulUint64s(dst, input1, input2 []uint64) []uint64 {
	if avx2 {
		_uint64_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivUint64s divides input1 by input2 and writes back the result into dst slice
func DivUint64s(dst, input1, input2 []uint64) []uint64 {
	if avx2 {
		_uint64_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Int8 ----------------------------------

// SumInt8s sums up all of the elements of the slice and returns the value
func SumInt8s(input []int8) (out int8) {
	switch {
	case avx2:
		_int8_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinInt8s returns the smallest element value in the slice
func MinInt8s(input []int8) (out int8) {
	switch {
	case avx2:
		_int8_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxInt8s returns the largest element value in the slice
func MaxInt8s(input []int8) (out int8) {
	switch {
	case avx2:
		_int8_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddInt8s adds input1 to input2 and writes back the result into dst slice
func AddInt8s(dst, input1, input2 []int8) []int8 {
	if avx2 {
		_int8_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubInt8s subtracts input2 from input1 and writes back the result into dst slice
func SubInt8s(dst, input1, input2 []int8) []int8 {
	if avx2 {
		_int8_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulInt8s multiplies input1 by input2 and writes back the result into dst slice
func MulInt8s(dst, input1, input2 []int8) []int8 {
	if avx2 {
		_int8_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivInt8s divides input1 by input2 and writes back the result into dst slice
func DivInt8s(dst, input1, input2 []int8) []int8 {
	if avx2 {
		_int8_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Int16 ----------------------------------

// SumInt16s sums up all of the elements of the slice and returns the value
func SumInt16s(input []int16) (out int16) {
	switch {
	case avx2:
		_int16_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinInt16s returns the smallest element value in the slice
func MinInt16s(input []int16) (out int16) {
	switch {
	case avx2:
		_int16_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxInt16s returns the largest element value in the slice
func MaxInt16s(input []int16) (out int16) {
	switch {
	case avx2:
		_int16_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddInt16s adds input1 to input2 and writes back the result into dst slice
func AddInt16s(dst, input1, input2 []int16) []int16 {
	if avx2 {
		_int16_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubInt16s subtracts input2 from input1 and writes back the result into dst slice
func SubInt16s(dst, input1, input2 []int16) []int16 {
	if avx2 {
		_int16_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulInt16s multiplies input1 by input2 and writes back the result into dst slice
func MulInt16s(dst, input1, input2 []int16) []int16 {
	if avx2 {
		_int16_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivInt16s divides input1 by input2 and writes back the result into dst slice
func DivInt16s(dst, input1, input2 []int16) []int16 {
	if avx2 {
		_int16_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Int32 ----------------------------------

// SumInt32s sums up all of the elements of the slice and returns the value
func SumInt32s(input []int32) (out int32) {
	switch {
	case avx2:
		_int32_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinInt32s returns the smallest element value in the slice
func MinInt32s(input []int32) (out int32) {
	switch {
	case avx2:
		_int32_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxInt32s returns the largest element value in the slice
func MaxInt32s(input []int32) (out int32) {
	switch {
	case avx2:
		_int32_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddInt32s adds input1 to input2 and writes back the result into dst slice
func AddInt32s(dst, input1, input2 []int32) []int32 {
	if avx2 {
		_int32_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubInt32s subtracts input2 from input1 and writes back the result into dst slice
func SubInt32s(dst, input1, input2 []int32) []int32 {
	if avx2 {
		_int32_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulInt32s multiplies input1 by input2 and writes back the result into dst slice
func MulInt32s(dst, input1, input2 []int32) []int32 {
	if avx2 {
		_int32_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivInt32s divides input1 by input2 and writes back the result into dst slice
func DivInt32s(dst, input1, input2 []int32) []int32 {
	if avx2 {
		_int32_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Int64 ----------------------------------

// SumInt64s sums up all of the elements of the slice and returns the value
func SumInt64s(input []int64) (out int64) {
	switch {
	case avx2:
		_int64_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinInt64s returns the smallest element value in the slice
func MinInt64s(input []int64) (out int64) {
	switch {
	case avx2:
		_int64_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxInt64s returns the largest element value in the slice
func MaxInt64s(input []int64) (out int64) {
	switch {
	case avx2:
		_int64_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddInt64s adds input1 to input2 and writes back the result into dst slice
func AddInt64s(dst, input1, input2 []int64) []int64 {
	if avx2 {
		_int64_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubInt64s subtracts input2 from input1 and writes back the result into dst slice
func SubInt64s(dst, input1, input2 []int64) []int64 {
	if avx2 {
		_int64_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulInt64s multiplies input1 by input2 and writes back the result into dst slice
func MulInt64s(dst, input1, input2 []int64) []int64 {
	if avx2 {
		_int64_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivInt64s divides input1 by input2 and writes back the result into dst slice
func DivInt64s(dst, input1, input2 []int64) []int64 {
	if avx2 {
		_int64_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Float32 ----------------------------------

// SumFloat32s sums up all of the elements of the slice and returns the value
func SumFloat32s(input []float32) (out float32) {
	switch {
	case avx2:
		_float32_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinFloat32s returns the smallest element value in the slice
func MinFloat32s(input []float32) (out float32) {
	switch {
	case avx2:
		_float32_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxFloat32s returns the largest element value in the slice
func MaxFloat32s(input []float32) (out float32) {
	switch {
	case avx2:
		_float32_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddFloat32s adds input1 to input2 and writes back the result into dst slice
func AddFloat32s(dst, input1, input2 []float32) []float32 {
	if avx2 {
		_float32_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubFloat32s subtracts input2 from input1 and writes back the result into dst slice
func SubFloat32s(dst, input1, input2 []float32) []float32 {
	if avx2 {
		_float32_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulFloat32s multiplies input1 by input2 and writes back the result into dst slice
func MulFloat32s(dst, input1, input2 []float32) []float32 {
	if avx2 {
		_float32_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivFloat32s divides input1 by input2 and writes back the result into dst slice
func DivFloat32s(dst, input1, input2 []float32) []float32 {
	if avx2 {
		_float32_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

// ---------------------------------- Float64 ----------------------------------

// SumFloat64s sums up all of the elements of the slice and returns the value
func SumFloat64s(input []float64) (out float64) {
	switch {
	case avx2:
		_float64_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// MinFloat64s returns the smallest element value in the slice
func MinFloat64s(input []float64) (out float64) {
	switch {
	case avx2:
		_float64_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// MaxFloat64s returns the largest element value in the slice
func MaxFloat64s(input []float64) (out float64) {
	switch {
	case avx2:
		_float64_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// AddFloat64s adds input1 to input2 and writes back the result into dst slice
func AddFloat64s(dst, input1, input2 []float64) []float64 {
	if avx2 {
		_float64_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// SubFloat64s subtracts input2 from input1 and writes back the result into dst slice
func SubFloat64s(dst, input1, input2 []float64) []float64 {
	if avx2 {
		_float64_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// MulFloat64s multiplies input1 by input2 and writes back the result into dst slice
func MulFloat64s(dst, input1, input2 []float64) []float64 {
	if avx2 {
		_float64_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// DivFloat64s divides input1 by input2 and writes back the result into dst slice
func DivFloat64s(dst, input1, input2 []float64) []float64 {
	if avx2 {
		_float64_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}

