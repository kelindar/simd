// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

//go:generate go run ./main.go

//go:embed *.tt
var templates embed.FS

type Type struct {
	Name  string
	Type  string
	CType string
}

var types = []Type{
	{Name: "Uint8", Type: "uint8", CType: "uint8_t"},
	{Name: "Uint16", Type: "uint16", CType: "uint16_t"},
	{Name: "Uint32", Type: "uint32", CType: "uint32_t"},
	{Name: "Uint64", Type: "uint64", CType: "uint64_t"},
	{Name: "Int8", Type: "int8", CType: "int8_t"},
	{Name: "Int16", Type: "int16", CType: "int16_t"},
	{Name: "Int32", Type: "int32", CType: "int32_t"},
	{Name: "Int64", Type: "int64", CType: "int64_t"},
	{Name: "Float32", Type: "float32", CType: "float"},
	{Name: "Float64", Type: "float64", CType: "double"},
}

func main() {
	genCode("amd64", "avx2")
	genCode("arm64", "neon")
	genCode("arm64", "apple")
	genFuncs("simd", "../simd_funcs.go")

	if err := execute("../simd_stub.go", "stub.go", "---", "---"); err != nil {
		panic(err)
	}
}

// Generates API functions and their tests
func genFuncs(arch, dst string) {
	if err := execute(dst, "funcs.go", arch, "---"); err != nil {
		panic(err)
	}
}

// Generates the underling code for vectorization and companions
func genCode(arch, mode string) {
	if err := execute(fmt.Sprintf("simd_%s_%s.c", mode, arch), "source.c", arch, mode); err != nil {
		panic(err)
	}
}

// Executes the template
func execute(dst, src, arch, mode string) error {
	body, err := templates.ReadFile(src + ".tt")
	if err != nil {
		return err
	}

	cgen, err := template.New(src).Parse(string(body))
	if err != nil {
		return err
	}

	out, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}

	return cgen.Execute(out, struct {
		Arch  string
		Mode  string
		Types []Type
	}{
		Arch:  arch,
		Mode:  mode,
		Types: types,
	})
}
