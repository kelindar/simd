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

//go:generate go run main.go

//go:embed templates/*
var templates embed.FS

type Type struct {
	Name string
	Type string
}

var types = []Type{
	{Name: "Uint8", Type: "uint8"},
	{Name: "Uint16", Type: "uint16"},
	{Name: "Uint32", Type: "uint32"},
	{Name: "Uint64", Type: "uint64"},
	{Name: "Int8", Type: "int8"},
	{Name: "Int16", Type: "int16"},
	{Name: "Int32", Type: "int32"},
	{Name: "Int64", Type: "int64"},
	{Name: "Float32", Type: "float32"},
	{Name: "Float64", Type: "float64"},
}

func main() {
	genCode("amd64", "avx2")
	//genCode("amd64", "avx512")
	//genCode("arm64", "neon")
	genFuncs("amd64")

	if err := execute("../simd_generic.go", "generic.go", "---", "---"); err != nil {
		panic(err)
	}

	if err := execute("../bench_test.go", "bench.go", "---", "---"); err != nil {
		panic(err)
	}
}

// Generates API functions and their tests
func genFuncs(arch string) {
	if err := execute(fmt.Sprintf("../simd_%s.go", arch), "funcs.go", arch, "---"); err != nil {
		panic(err)
	}
}

// Generates the underling code for vectorization and companions
func genCode(arch, mode string) {
	if err := execute(fmt.Sprintf("simd_%s_%s.cpp", mode, arch), "source.cpp", arch, mode); err != nil {
		panic(err)
	}

	if err := execute(fmt.Sprintf("../simd_%s_%s.go", mode, arch), "binding.go", arch, mode); err != nil {
		panic(err)
	}
}

// Executes the template
func execute(dst, src, arch, mode string) error {
	body, err := templates.ReadFile("templates/" + src + ".tt")
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
