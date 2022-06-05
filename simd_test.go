// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package simd

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Result represents a result of a benchmark
type Result struct {
	Type    string
	Name    string
	Size    int
	Rate    float64
	Speedup float64
}

// makeVector generates a test vector
func makeVector[T Number](count int) []T {
	arr := make([]T, count)
	for i := 0; i < count; i++ {
		arr[i] = T((i % 100) + 1)
	}
	return arr
}

// runBenchmark runs a benchmark and compares it with the baseline
func runBenchmark(b *testing.B, typ, name string, size int, fn func(b *testing.B)) Result {
	rate0 := measure(b, fmt.Sprintf("%v-%v-%v-base", typ, name, size), "base", fn)
	rate1 := measure(b, fmt.Sprintf("%v-%v-%v-simd", typ, name, size), "simd", fn)
	return Result{
		Type:    typ,
		Name:    name,
		Size:    size,
		Rate:    rate1,
		Speedup: rate0 / rate1,
	}
}

// measure measures the runtime
func measure(b *testing.B, name, mode string, fn func(b *testing.B)) float64 {
	setMode(mode)

	// Run the benchmark and measure time and operations
	var start time.Time
	var ops int
	b.Run(name, func(b *testing.B) {
		start = time.Now()
		fn(b)
		ops = b.N
	})

	// Calculate nanoseconds per operation (rate)
	return float64(time.Since(start)) / float64(ops)
}

// Iterates over all modes
func rangeModes(fn func(mode string)) {
	for _, mode := range []string{"base", "simd"} {
		setMode(mode)
		fn(mode)
	}
}

// setMode changes the mode
func setMode(mode string) {
	switch mode {
	case "simd":
		avx2 = true
	case "base":
		avx2 = false
	}
}

func TestSum(t *testing.T) {
	assert.Equal(t, 3, int(Sum([]int8{1, 2})))
	assert.Equal(t, 3, int(Sum([]int16{1, 2})))
	assert.Equal(t, 3, int(Sum([]int32{1, 2})))
	assert.Equal(t, 3, int(Sum([]int64{1, 2})))
	assert.Equal(t, 3, int(Sum([]uint8{1, 2})))
	assert.Equal(t, 3, int(Sum([]uint16{1, 2})))
	assert.Equal(t, 3, int(Sum([]uint32{1, 2})))
	assert.Equal(t, 3, int(Sum([]uint64{1, 2})))
	assert.Equal(t, 3, int(Sum([]float32{1, 2})))
	assert.Equal(t, 3, int(Sum([]float64{1, 2})))
	assert.Equal(t, 3, int(Sum([]int{1, 2})))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, int(Min([]int8{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]int16{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]int32{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]int64{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]uint8{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]uint16{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]uint32{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]uint64{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]float32{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]float64{3, 1, 2})))
	assert.Equal(t, 1, int(Min([]int{3, 1, 2})))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 2, int(Max([]int8{1, 2})))
	assert.Equal(t, 2, int(Max([]int16{1, 2})))
	assert.Equal(t, 2, int(Max([]int32{1, 2})))
	assert.Equal(t, 2, int(Max([]int64{1, 2})))
	assert.Equal(t, 2, int(Max([]uint8{1, 2})))
	assert.Equal(t, 2, int(Max([]uint16{1, 2})))
	assert.Equal(t, 2, int(Max([]uint32{1, 2})))
	assert.Equal(t, 2, int(Max([]uint64{1, 2})))
	assert.Equal(t, 2, int(Max([]float32{1, 2})))
	assert.Equal(t, 2, int(Max([]float64{1, 2})))
	assert.Equal(t, 2, int(Max([]int{1, 2})))
}
