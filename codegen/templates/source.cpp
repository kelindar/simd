// Copyright (c) Roman Atachiants and contributors. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

#include <stdint.h>
#include <immintrin.h>

typedef float float32;
typedef double float64;
typedef int8_t int8;
typedef int16_t int16;
typedef int32_t int32;
typedef int64_t int64;
typedef uint8_t uint8;
typedef uint16_t uint16;
typedef uint32_t uint32;
typedef uint64_t uint64;
{{ $Mode := .Mode }}
{{ range .Types }}
// ---------------------------------- {{.Name}} ----------------------------------

extern "C" void {{.Type}}_{{$Mode}}_sum({{.Type}} *input, {{.Type}} *result, uint64_t size) {
    {{.Type}} sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void {{.Type}}_{{$Mode}}_min({{.Type}} *input, {{.Type}} *result, uint64_t size) {
    {{.Type}} min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void {{.Type}}_{{$Mode}}_max({{.Type}} *input, {{.Type}} *result, uint64_t size) {
    {{.Type}} max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void {{.Type}}_{{$Mode}}_add({{.Type}} *input1, {{.Type}} *input2, {{.Type}} *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void {{.Type}}_{{$Mode}}_sub({{.Type}} *input1, {{.Type}} *input2, {{.Type}} *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void {{.Type}}_{{$Mode}}_mul({{.Type}} *input1, {{.Type}} *input2, {{.Type}} *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void {{.Type}}_{{$Mode}}_div({{.Type}} *input1, {{.Type}} *input2, {{.Type}} *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}
{{ end }}