// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package simd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

{{ range .Types }}
// ---------------------------------- {{.Name}} ----------------------------------

func Benchmark{{.Name}}(b *testing.B) {
	result := make([]Result, 0, 64)

	typ := "{{.Type}}"
	for _, count := range []int{256, 4096, 16384} {
		vector := makeVector[{{.Type}}](count)
		result = append(result, runBenchmark(b, typ, "sum", count, func(b *testing.B) {
			result := {{.Type}}(0)
			for i := 0; i < b.N; i++ {
				result = Sum{{.Name}}s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "min", count, func(b *testing.B) {
			result := {{.Type}}(0)
			for i := 0; i < b.N; i++ {
				result = Min{{.Name}}s(vector)
			}
			assert.NotZero(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "max", count, func(b *testing.B) {
			result := {{.Type}}(0)
			for i := 0; i < b.N; i++ {
				result = Max{{.Name}}s(vector)
			}
			assert.NotZero(b, result)
		}))

		output := make([]{{.Type}}, count)
		input1 := makeVector[{{.Type}}](count)
		input2 := makeVector[{{.Type}}](count)
		result = append(result, runBenchmark(b, typ, "add", count, func(b *testing.B) {
			var result []{{.Type}}
			for i := 0; i < b.N; i++ {
				result = Add{{.Name}}s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "sub", count, func(b *testing.B) {
			var result []{{.Type}}
			for i := 0; i < b.N; i++ {
				result = Sub{{.Name}}s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "mul", count, func(b *testing.B) {
			var result []{{.Type}}
			for i := 0; i < b.N; i++ {
				result = Mul{{.Name}}s(output, input1, input2)
			}
			assert.NotEmpty(b, result)
		}))

		result = append(result, runBenchmark(b, typ, "div", count, func(b *testing.B) {
			var result []{{.Type}}
			for i := 0; i < b.N; i++ {
				result = Div{{.Name}}s(output, input1, input2)
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

func Test{{.Name}}_Sum(t *testing.T) {
	input := makeVector[{{.Type}}](70)
	expect := sum(input)
	result := Sum{{.Name}}s(input)
	assert.EqualValues(t, expect, result)
}

func Test{{.Name}}_Min(t *testing.T) {
	input := makeVector[{{.Type}}](70)
	expect := min(input)
	result := Min{{.Name}}s(input)
	assert.EqualValues(t, expect, result)
}

func Test{{.Name}}_Max(t *testing.T) {
	input := makeVector[{{.Type}}](70)
	expect := max(input)
	result := Max{{.Name}}s(input)
	assert.EqualValues(t, expect, result)
}

func Test{{.Name}}_Add(t *testing.T) {
	input1 := makeVector[{{.Type}}](70)
	input2 := makeVector[{{.Type}}](70)
	expect := add(make([]{{.Type}}, 70), input1, input2)
	result := Add{{.Name}}s(make([]{{.Type}}, 70), input1, input2)
	assert.EqualValues(t, expect, result)
}

func Test{{.Name}}_Sub(t *testing.T) {
	input1 := makeVector[{{.Type}}](70)
	input2 := makeVector[{{.Type}}](70)
	expect := sub(make([]{{.Type}}, 70), input1, input2)
	result := Sub{{.Name}}s(make([]{{.Type}}, 70), input1, input2)
	assert.EqualValues(t, expect, result)
}

func Test{{.Name}}_Mul(t *testing.T) {
	input1 := makeVector[{{.Type}}](70)
	input2 := makeVector[{{.Type}}](70)
	expect := mul(make([]{{.Type}}, 70), input1, input2)
	result := Mul{{.Name}}s(make([]{{.Type}}, 70), input1, input2)
	assert.EqualValues(t, expect, result)
}

func Test{{.Name}}_Div(t *testing.T) {
	input1 := makeVector[{{.Type}}](70)
	input2 := makeVector[{{.Type}}](70)
	expect := div(make([]{{.Type}}, 70), input1, input2)
	result := Div{{.Name}}s(make([]{{.Type}}, 70), input1, input2)
	assert.InDeltaSlice(t, expect, result, 0.01)
}
{{ end }}
