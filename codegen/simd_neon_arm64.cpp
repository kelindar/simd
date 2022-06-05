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


// ---------------------------------- Uint8 ----------------------------------

extern "C" void uint8_neon_sum(uint8 *input, uint8 *result, uint64_t size) {
    uint8 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void uint8_neon_min(uint8 *input, uint8 *result, uint64_t size) {
    uint8 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void uint8_neon_max(uint8 *input, uint8 *result, uint64_t size) {
    uint8 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void uint8_neon_add(uint8 *input1, uint8 *input2, uint8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void uint8_neon_sub(uint8 *input1, uint8 *input2, uint8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void uint8_neon_mul(uint8 *input1, uint8 *input2, uint8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void uint8_neon_div(uint8 *input1, uint8 *input2, uint8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Uint16 ----------------------------------

extern "C" void uint16_neon_sum(uint16 *input, uint16 *result, uint64_t size) {
    uint16 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void uint16_neon_min(uint16 *input, uint16 *result, uint64_t size) {
    uint16 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void uint16_neon_max(uint16 *input, uint16 *result, uint64_t size) {
    uint16 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void uint16_neon_add(uint16 *input1, uint16 *input2, uint16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void uint16_neon_sub(uint16 *input1, uint16 *input2, uint16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void uint16_neon_mul(uint16 *input1, uint16 *input2, uint16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void uint16_neon_div(uint16 *input1, uint16 *input2, uint16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Uint32 ----------------------------------

extern "C" void uint32_neon_sum(uint32 *input, uint32 *result, uint64_t size) {
    uint32 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void uint32_neon_min(uint32 *input, uint32 *result, uint64_t size) {
    uint32 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void uint32_neon_max(uint32 *input, uint32 *result, uint64_t size) {
    uint32 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void uint32_neon_add(uint32 *input1, uint32 *input2, uint32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void uint32_neon_sub(uint32 *input1, uint32 *input2, uint32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void uint32_neon_mul(uint32 *input1, uint32 *input2, uint32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void uint32_neon_div(uint32 *input1, uint32 *input2, uint32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Uint64 ----------------------------------

extern "C" void uint64_neon_sum(uint64 *input, uint64 *result, uint64_t size) {
    uint64 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void uint64_neon_min(uint64 *input, uint64 *result, uint64_t size) {
    uint64 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void uint64_neon_max(uint64 *input, uint64 *result, uint64_t size) {
    uint64 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void uint64_neon_add(uint64 *input1, uint64 *input2, uint64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void uint64_neon_sub(uint64 *input1, uint64 *input2, uint64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void uint64_neon_mul(uint64 *input1, uint64 *input2, uint64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void uint64_neon_div(uint64 *input1, uint64 *input2, uint64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Int8 ----------------------------------

extern "C" void int8_neon_sum(int8 *input, int8 *result, uint64_t size) {
    int8 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void int8_neon_min(int8 *input, int8 *result, uint64_t size) {
    int8 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void int8_neon_max(int8 *input, int8 *result, uint64_t size) {
    int8 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void int8_neon_add(int8 *input1, int8 *input2, int8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void int8_neon_sub(int8 *input1, int8 *input2, int8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void int8_neon_mul(int8 *input1, int8 *input2, int8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void int8_neon_div(int8 *input1, int8 *input2, int8 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Int16 ----------------------------------

extern "C" void int16_neon_sum(int16 *input, int16 *result, uint64_t size) {
    int16 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void int16_neon_min(int16 *input, int16 *result, uint64_t size) {
    int16 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void int16_neon_max(int16 *input, int16 *result, uint64_t size) {
    int16 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void int16_neon_add(int16 *input1, int16 *input2, int16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void int16_neon_sub(int16 *input1, int16 *input2, int16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void int16_neon_mul(int16 *input1, int16 *input2, int16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void int16_neon_div(int16 *input1, int16 *input2, int16 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Int32 ----------------------------------

extern "C" void int32_neon_sum(int32 *input, int32 *result, uint64_t size) {
    int32 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void int32_neon_min(int32 *input, int32 *result, uint64_t size) {
    int32 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void int32_neon_max(int32 *input, int32 *result, uint64_t size) {
    int32 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void int32_neon_add(int32 *input1, int32 *input2, int32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void int32_neon_sub(int32 *input1, int32 *input2, int32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void int32_neon_mul(int32 *input1, int32 *input2, int32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void int32_neon_div(int32 *input1, int32 *input2, int32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Int64 ----------------------------------

extern "C" void int64_neon_sum(int64 *input, int64 *result, uint64_t size) {
    int64 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void int64_neon_min(int64 *input, int64 *result, uint64_t size) {
    int64 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void int64_neon_max(int64 *input, int64 *result, uint64_t size) {
    int64 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void int64_neon_add(int64 *input1, int64 *input2, int64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void int64_neon_sub(int64 *input1, int64 *input2, int64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void int64_neon_mul(int64 *input1, int64 *input2, int64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void int64_neon_div(int64 *input1, int64 *input2, int64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Float32 ----------------------------------

extern "C" void float32_neon_sum(float32 *input, float32 *result, uint64_t size) {
    float32 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void float32_neon_min(float32 *input, float32 *result, uint64_t size) {
    float32 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void float32_neon_max(float32 *input, float32 *result, uint64_t size) {
    float32 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void float32_neon_add(float32 *input1, float32 *input2, float32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void float32_neon_sub(float32 *input1, float32 *input2, float32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void float32_neon_mul(float32 *input1, float32 *input2, float32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void float32_neon_div(float32 *input1, float32 *input2, float32 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}

// ---------------------------------- Float64 ----------------------------------

extern "C" void float64_neon_sum(float64 *input, float64 *result, uint64_t size) {
    float64 sum = 0.0;
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; ++i) {
        sum += input[i];
    }
    *result = sum;
}

extern "C" void float64_neon_min(float64 *input, float64 *result, uint64_t size) {
    float64 min = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] < min) {
            min = input[i];
        }
    }
    *result = min;
}

extern "C" void float64_neon_max(float64 *input, float64 *result, uint64_t size) {
    float64 max = input[0];
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        if (input[i] > max) {
            max = input[i];
        }
    }
    *result = max;
}

extern "C" void float64_neon_add(float64 *input1, float64 *input2, float64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] + input2[i];
    }
}

extern "C" void float64_neon_sub(float64 *input1, float64 *input2, float64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] - input2[i];
    }
}

extern "C" void float64_neon_mul(float64 *input1, float64 *input2, float64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] * input2[i];
    }
}

extern "C" void float64_neon_div(float64 *input1, float64 *input2, float64 *output, uint64_t size) {
    #pragma clang loop vectorize(enable) interleave(enable)
    for (int i = 0; i < (int)size; i++) {
        output[i] = input1[i] / input2[i];
    }
}
