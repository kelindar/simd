// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

//go:build {{.Arch}}
// +build {{.Arch}}

package simd

import (
	"unsafe"
)
{{ $Mode := .Mode }}
{{ range .Types }}
//go:noescape
func _{{.Type}}_{{$Mode}}_sum(input, result unsafe.Pointer, info uint64)
//go:noescape
func _{{.Type}}_{{$Mode}}_min(input, result unsafe.Pointer, info uint64)
//go:noescape
func _{{.Type}}_{{$Mode}}_max(input, result unsafe.Pointer, info uint64)
//go:noescape
func _{{.Type}}_{{$Mode}}_add(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _{{.Type}}_{{$Mode}}_sub(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _{{.Type}}_{{$Mode}}_mul(input1, input2, output unsafe.Pointer, info uint64)
//go:noescape
func _{{.Type}}_{{$Mode}}_div(input1, input2, output unsafe.Pointer, info uint64)
{{ end }}
