<p align="center">
<img width="330" height="110" src=".github/logo.png" border="0" alt="kelindar/simd">
<br>
<img src="https://img.shields.io/github/go-mod/go-version/kelindar/simd" alt="Go Version">
<a href="https://pkg.go.dev/github.com/kelindar/simd"><img src="https://pkg.go.dev/badge/github.com/kelindar/simd" alt="PkgGoDev"></a>
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg" alt="License"></a>
<a href="https://coveralls.io/github/kelindar/simd"><img src="https://coveralls.io/repos/github/kelindar/simd/badge.svg" alt="Coverage"></a>
</p>

## Vectorized Math Functions

This library contains a set of vectorized mathematical functions which were [auto-vectorized](https://llvm.org/docs/Vectorizers.html) using clang compiler and translated into PLAN9 assembly code for Go. Generic version is also provided for CPUs where vectorization is not available, or for which this library doesn't have a generated code.

It currently supports only `AVX2`, but `AVX512` and `SVE` (for `ARM`) should be easy enough to generate. Most of the code in this library is auto-generated, which helps with maintenance.

## Usage

Usage of this library is very straightforward, it simply exposes a set of non-opinionated functions. For example, if you want to sum up a slice of floating point numbers, you can use `SumFloat32s([]float32) float32` function to do so.

```go
sum := simd.SumFloat32s([]float32{1, 2, 3, 4, 5})
```

## Benchmarks

```go
goos: windows
goarch: amd64
pkg: github.com/kelindar/simd
cpu: Intel(R) Core(TM) i7-9700K CPU @ 3.60GHz

   TYPE    OP    SIZE     RATE        SPEEDUP
float32   sum     256     8.96 ns/op   20.51x
float32   min     256    27.78 ns/op    4.15x
float32   max     256    25.94 ns/op    4.43x
float32   add     256    15.31 ns/op    7.46x
float32   sub     256    15.48 ns/op    7.34x
float32   mul     256    15.36 ns/op    7.54x
float32   div     256    27.88 ns/op    5.79x
float32   sum    4096   107.92 ns/op   31.60x
float32   min    4096   134.06 ns/op   13.01x
float32   max    4096   133.90 ns/op   12.94x
float32   add    4096   353.01 ns/op    4.90x
float32   sub    4096   353.62 ns/op    4.92x
float32   mul    4096   353.64 ns/op    4.87x
float32   div    4096   455.60 ns/op    5.68x
float32   sum   16384   469.18 ns/op   29.44x
float32   min   16384   497.91 ns/op   13.94x
float32   max   16384   500.43 ns/op   13.84x
float32   add   16384  1442.28 ns/op    4.82x
float32   sub   16384  1366.98 ns/op    5.07x
float32   mul   16384  1382.20 ns/op    4.97x
float32   div   16384  1821.82 ns/op    5.71x

   TYPE    OP    SIZE     RATE        SPEEDUP
float64   sum     256    14.56 ns/op   12.61x
float64   min     256    24.51 ns/op    4.73x
float64   max     256    24.61 ns/op    3.71x
float64   add     256    25.81 ns/op    4.42x
float64   sub     256    25.84 ns/op    4.41x
float64   mul     256    25.66 ns/op    4.45x
float64   div     256   109.60 ns/op    1.97x
float64   sum    4096   220.98 ns/op   15.51x
float64   min    4096   229.76 ns/op    7.62x
float64   max    4096   227.68 ns/op    5.86x
float64   add    4096   707.07 ns/op    2.44x
float64   sub    4096   716.65 ns/op    2.45x
float64   mul    4096   699.45 ns/op    2.50x
float64   div    4096  1726.40 ns/op    2.01x
float64   sum   16384   930.43 ns/op   14.77x
float64   min   16384   930.08 ns/op    7.45x
float64   max   16384   938.31 ns/op    5.81x
float64   add   16384  5522.65 ns/op    1.45x
float64   sub   16384  5433.96 ns/op    1.55x
float64   mul   16384  5564.75 ns/op    1.51x
float64   div   16384  6913.10 ns/op    2.01x

   TYPE    OP    SIZE     RATE        SPEEDUP
  uint8   sum     100     5.94 ns/op    6.34x
  uint8   min     100     7.09 ns/op    7.31x
  uint8   max     100     7.82 ns/op    6.25x
  uint8   add     100     9.13 ns/op    5.76x
  uint8   sub     100     9.11 ns/op    5.73x
  uint8   mul     100    11.15 ns/op    6.37x
  uint8   div     100   186.89 ns/op    1.17x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint16   sum     256     6.13 ns/op   13.19x
 uint16   min     256    10.00 ns/op   15.94x
 uint16   max     256    11.11 ns/op    9.58x
 uint16   add     256    10.04 ns/op   16.83x
 uint16   sub     256    10.07 ns/op   16.77x
 uint16   mul     256    10.45 ns/op   16.26x
 uint16   div     256   438.34 ns/op    1.26x
 uint16   sum    4096    38.15 ns/op   32.42x
 uint16   min    4096    57.84 ns/op   44.50x
 uint16   max    4096    58.61 ns/op   29.17x
 uint16   add    4096    86.01 ns/op   30.29x
 uint16   sub    4096    91.34 ns/op   30.54x
 uint16   mul    4096    89.25 ns/op   29.65x
 uint16   div    4096  7203.14 ns/op    1.26x
 uint16   sum   16384   148.55 ns/op   34.66x
 uint16   min   16384   231.87 ns/op   44.76x
 uint16   max   16384   227.01 ns/op   31.32x
 uint16   add   16384   722.47 ns/op   14.60x
 uint16   sub   16384   737.93 ns/op   14.21x
 uint16   mul   16384   711.59 ns/op   14.84x
 uint16   div   16384 28084.66 ns/op    1.24x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint32   sum     256     8.06 ns/op   10.39x
 uint32   min     256    22.96 ns/op    6.96x
 uint32   max     256    26.21 ns/op    4.12x
 uint32   add     256    15.91 ns/op    8.18x
 uint32   sub     256    15.49 ns/op    7.98x
 uint32   mul     256    16.81 ns/op    7.40x
 uint32   div     256   436.08 ns/op    1.29x
 uint32   sum    4096    75.90 ns/op   16.47x
 uint32   min    4096   120.89 ns/op   21.54x
 uint32   max    4096   125.33 ns/op   13.76x
 uint32   add    4096   380.00 ns/op    5.16x
 uint32   sub    4096   345.19 ns/op    5.69x
 uint32   mul    4096   367.35 ns/op    4.91x
 uint32   div    4096  6695.04 ns/op    1.29x
 uint32   sum   16384   462.89 ns/op   10.85x
 uint32   min   16384   475.29 ns/op   21.98x
 uint32   max   16384   481.76 ns/op   14.43x
 uint32   add   16384  1437.51 ns/op    5.49x
 uint32   sub   16384  1427.22 ns/op    5.57x
 uint32   mul   16384  1367.51 ns/op    5.71x
 uint32   div   16384 26786.72 ns/op    1.30x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint64   sum     256    12.35 ns/op    9.18x
 uint64   min     256    41.42 ns/op    3.84x
 uint64   max     256    44.52 ns/op    2.55x
 uint64   add     256    24.86 ns/op    4.92x
 uint64   sub     256    25.35 ns/op    4.82x
 uint64   mul     256    50.52 ns/op    2.37x
 uint64   div     256  1250.57 ns/op    1.18x
 uint64   sum    4096   145.98 ns/op   11.78x
 uint64   min    4096   507.57 ns/op    5.11x
 uint64   max    4096   508.87 ns/op    3.42x
 uint64   add    4096   701.61 ns/op    2.79x
 uint64   sub    4096   719.04 ns/op    2.74x
 uint64   mul    4096   881.98 ns/op    2.18x
 uint64   div    4096 20114.02 ns/op    1.18x
 uint64   sum   16384   930.65 ns/op    7.44x
 uint64   min   16384  2059.33 ns/op    5.08x
 uint64   max   16384  2071.32 ns/op    3.34x
 uint64   add   16384  6001.34 ns/op    1.45x
 uint64   sub   16384  5713.91 ns/op    1.52x
 uint64   mul   16384  6147.43 ns/op    1.44x
 uint64   div   16384 80476.33 ns/op    1.17x

   TYPE    OP    SIZE     RATE        SPEEDUP
   int8   sum     100     5.89 ns/op    6.18x
   int8   min     100     7.51 ns/op    6.81x
   int8   max     100     7.46 ns/op    6.33x
   int8   add     100     9.10 ns/op    5.74x
   int8   sub     100     9.13 ns/op    5.75x
   int8   mul     100    10.97 ns/op    6.46x
   int8   div     100   480.25 ns/op    0.63x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int16   sum     256     6.20 ns/op   13.04x
  int16   min     256    10.80 ns/op    9.77x
  int16   max     256    10.13 ns/op   10.53x
  int16   add     256     9.99 ns/op   16.97x
  int16   sub     256    10.00 ns/op   17.03x
  int16   mul     256    10.47 ns/op   16.24x
  int16   div     256   499.02 ns/op    1.53x
  int16   sum    4096    37.97 ns/op   32.87x
  int16   min    4096    59.55 ns/op   29.15x
  int16   max    4096    58.42 ns/op   29.54x
  int16   add    4096    85.73 ns/op   30.27x
  int16   sub    4096    86.22 ns/op   30.07x
  int16   mul    4096    86.61 ns/op   30.00x
  int16   div    4096  7936.80 ns/op    1.52x
  int16   sum   16384   144.40 ns/op   34.61x
  int16   min   16384   226.17 ns/op   30.59x
  int16   max   16384   222.39 ns/op   30.90x
  int16   add   16384   708.78 ns/op   14.55x
  int16   sub   16384   722.46 ns/op   14.33x
  int16   mul   16384   715.68 ns/op   14.49x
  int16   div   16384 31537.05 ns/op    1.52x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int32   sum     256     8.04 ns/op   10.35x
  int32   min     256    22.58 ns/op    4.72x
  int32   max     256    22.80 ns/op    4.67x
  int32   add     256    14.89 ns/op    8.29x
  int32   sub     256    15.08 ns/op    8.20x
  int32   mul     256    15.86 ns/op    7.61x
  int32   div     256   490.90 ns/op    1.37x
  int32   sum    4096    76.31 ns/op   16.23x
  int32   min    4096   121.09 ns/op   14.37x
  int32   max    4096   120.26 ns/op   14.35x
  int32   add    4096   347.35 ns/op    5.64x
  int32   sub    4096   351.04 ns/op    5.58x
  int32   mul    4096   341.38 ns/op    5.34x
  int32   div    4096  7849.51 ns/op    1.37x
  int32   sum   16384   467.90 ns/op   10.59x
  int32   min   16384   486.13 ns/op   14.18x
  int32   max   16384   480.85 ns/op   14.48x
  int32   add   16384  1426.40 ns/op    5.53x
  int32   sub   16384  1421.44 ns/op    5.53x
  int32   mul   16384  1448.81 ns/op    5.41x
  int32   div   16384 31078.60 ns/op    1.38x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int64   sum     256    12.13 ns/op    6.71x
  int64   min     256    30.63 ns/op    3.47x
  int64   max     256    31.69 ns/op    3.84x
  int64   add     256    25.72 ns/op    5.17x
  int64   sub     256    25.66 ns/op    5.45x
  int64   mul     256    51.14 ns/op    2.38x
  int64   div     256  1524.32 ns/op    1.18x
  int64   sum    4096   145.87 ns/op    8.59x
  int64   min    4096   327.28 ns/op    5.26x
  int64   max    4096   331.75 ns/op    5.24x
  int64   add    4096   696.99 ns/op    2.83x
  int64   sub    4096   711.28 ns/op    2.76x
  int64   mul    4096   890.62 ns/op    2.13x
  int64   div    4096 24268.84 ns/op    1.18x
  int64   sum   16384   933.71 ns/op    5.59x
  int64   min   16384  1415.62 ns/op    4.90x
  int64   max   16384  1419.97 ns/op    4.88x
  int64   add   16384  5541.14 ns/op    1.53x
  int64   sub   16384  5183.70 ns/op    1.62x
  int64   mul   16384  6164.29 ns/op    1.36x
  int64   div   16384 96899.48 ns/op    1.18x
```

# Acknowledgements

This library was originally inspired by the work of Valery Carey & Adrian Witas in [viant/vec](https://github.com/viant/vec) package, but instead of hand-rolled assembly and intrinsics I opted for using auto-vectorization for maintainability reasons.
