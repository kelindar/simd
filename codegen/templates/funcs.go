// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build {{.Arch}}
// +build {{.Arch}}

package simd

import (
	"unsafe"
)

{{ range .Types }}
// ---------------------------------- {{.Name}} ----------------------------------

// Sum{{.Name}}s sums up all of the elements of the slice and returns the value
func Sum{{.Name}}s(input []{{.Type}}) (out {{.Type}}) {
	switch {
	case avx2:
		_{{.Type}}_avx2_sum(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return sum(input)
	}
}

// Min{{.Name}}s returns the smallest element value in the slice
func Min{{.Name}}s(input []{{.Type}}) (out {{.Type}}) {
	switch {
	case avx2:
		_{{.Type}}_avx2_min(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return min(input)
	}
}

// Max{{.Name}}s returns the largest element value in the slice
func Max{{.Name}}s(input []{{.Type}}) (out {{.Type}}) {
	switch {
	case avx2:
		_{{.Type}}_avx2_max(unsafe.Pointer(&input[0]), unsafe.Pointer(&out), uint64(len(input)))
		return
	default:
		return max(input)
	}
}

// Add{{.Name}}s adds input1 to input2 and writes back the result into dst slice
func Add{{.Name}}s(dst, input1, input2 []{{.Type}}) []{{.Type}} {
	if avx2 {
		_{{.Type}}_avx2_add(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return add(dst, input1, input2)
}

// Sub{{.Name}}s subtracts input2 from input1 and writes back the result into dst slice
func Sub{{.Name}}s(dst, input1, input2 []{{.Type}}) []{{.Type}} {
	if avx2 {
		_{{.Type}}_avx2_sub(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return sub(dst, input1, input2)
}

// Mul{{.Name}}s multiplies input1 by input2 and writes back the result into dst slice
func Mul{{.Name}}s(dst, input1, input2 []{{.Type}}) []{{.Type}} {
	if avx2 {
		_{{.Type}}_avx2_mul(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return mul(dst, input1, input2)
}

// Div{{.Name}}s divides input1 by input2 and writes back the result into dst slice
func Div{{.Name}}s(dst, input1, input2 []{{.Type}}) []{{.Type}} {
	if avx2 {
		_{{.Type}}_avx2_div(unsafe.Pointer(&input1[0]), unsafe.Pointer(&input2[0]), unsafe.Pointer(&(dst)[0]), uint64(len(dst)))
		return dst
	}
	return div(dst, input1, input2)
}
{{ end }}
