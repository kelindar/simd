<p align="center">
<img width="300" height="100" src=".github/logo.png" border="0" alt="kelindar/simd">
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

Below are the results for the Apple M3 Pro (Apple Silicon) machine.

```go
oos: darwin
goarch: arm64
pkg: github.com/kelindar/simd
cpu: Apple M3 Pro

   TYPE    OP    SIZE     RATE        SPEEDUP
  uint8   sum     256     4.91 ns/op   14.77x
  uint8   min     256     4.85 ns/op   37.35x
  uint8   max     256     5.08 ns/op   36.97x
  uint8   add     256     8.29 ns/op   13.05x
  uint8   sub     256     8.31 ns/op   12.96x
  uint8   mul     256     8.71 ns/op   12.48x
  uint8   div     256   130.96 ns/op    1.12x
  uint8   sum    4096    48.78 ns/op   21.38x
  uint8   min    4096    51.58 ns/op   61.09x
  uint8   max    4096    48.59 ns/op   65.16x
  uint8   add    4096    59.04 ns/op   27.71x
  uint8   sub    4096    59.42 ns/op   27.47x
  uint8   mul    4096    59.57 ns/op   27.45x
  uint8   div    4096  2074.64 ns/op    1.13x
  uint8   sum   16384   235.86 ns/op   17.73x
  uint8   min   16384   234.78 ns/op   53.59x
  uint8   max   16384   238.65 ns/op   53.78x
  uint8   add   16384   277.92 ns/op   23.88x
  uint8   sub   16384   275.56 ns/op   24.04x
  uint8   mul   16384   280.81 ns/op   23.59x
  uint8   div   16384  8163.29 ns/op    1.15x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint16   sum     256     6.80 ns/op   10.69x
 uint16   min     256     6.91 ns/op   26.50x
 uint16   max     256     6.94 ns/op   26.21x
 uint16   add     256    11.80 ns/op    9.16x
 uint16   sub     256    11.87 ns/op    9.20x
 uint16   mul     256    11.80 ns/op    9.26x
 uint16   div     256   129.94 ns/op    1.12x
 uint16   sum    4096   108.85 ns/op   10.18x
 uint16   min    4096   105.79 ns/op   29.66x
 uint16   max    4096   106.19 ns/op   29.43x
 uint16   add    4096   112.26 ns/op   14.59x
 uint16   sub    4096   118.79 ns/op   13.80x
 uint16   mul    4096   116.68 ns/op   14.21x
 uint16   div    4096  2056.41 ns/op    1.13x
 uint16   sum   16384   529.53 ns/op    7.93x
 uint16   min   16384   497.89 ns/op   25.35x
 uint16   max   16384   512.68 ns/op   24.40x
 uint16   add   16384   548.44 ns/op   11.99x
 uint16   sub   16384   579.93 ns/op   11.40x
 uint16   mul   16384   526.33 ns/op   12.69x
 uint16   div   16384  8454.70 ns/op    1.11x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint32   sum     256    11.29 ns/op    6.73x
 uint32   min     256    11.25 ns/op   11.40x
 uint32   max     256    11.30 ns/op   11.11x
 uint32   add     256    17.86 ns/op    6.09x
 uint32   sub     256    18.64 ns/op    5.79x
 uint32   mul     256    17.84 ns/op    6.10x
 uint32   div     256   132.50 ns/op    1.11x
 uint32   sum    4096   240.50 ns/op    4.46x
 uint32   min    4096   246.75 ns/op    8.63x
 uint32   max    4096   242.99 ns/op    8.72x
 uint32   add    4096   254.21 ns/op    6.49x
 uint32   sub    4096   258.73 ns/op    6.33x
 uint32   mul    4096   280.35 ns/op    5.87x
 uint32   div    4096  2187.58 ns/op    1.12x
 uint32   sum   16384  1039.29 ns/op    4.20x
 uint32   min   16384  1067.80 ns/op    7.97x
 uint32   max   16384  1023.83 ns/op    8.29x
 uint32   add   16384   887.07 ns/op    7.41x
 uint32   sub   16384   889.97 ns/op    7.66x
 uint32   mul   16384   886.21 ns/op    7.45x
 uint32   div   16384  9012.67 ns/op    1.04x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint64   sum     256    21.81 ns/op    3.38x
 uint64   min     256    42.46 ns/op    2.95x
 uint64   max     256    41.39 ns/op    3.08x
 uint64   add     256    30.89 ns/op    3.52x
 uint64   sub     256    30.91 ns/op    3.49x
 uint64   mul     256    74.32 ns/op    1.45x
 uint64   div     256   134.35 ns/op    1.10x
 uint64   sum    4096   491.83 ns/op    2.12x
 uint64   min    4096   981.65 ns/op    2.17x
 uint64   max    4096   992.13 ns/op    2.11x
 uint64   add    4096   549.37 ns/op    2.97x
 uint64   sub    4096   484.83 ns/op    3.50x
 uint64   mul    4096  1091.51 ns/op    1.50x
 uint64   div    4096  2136.43 ns/op    1.09x
 uint64   sum   16384  2091.84 ns/op    2.13x
 uint64   min   16384  4061.30 ns/op    2.07x
 uint64   max   16384  4356.20 ns/op    1.97x
 uint64   add   16384  3391.09 ns/op    1.95x
 uint64   sub   16384  3518.09 ns/op    1.88x
 uint64   mul   16384  4433.94 ns/op    1.48x
 uint64   div   16384  8670.50 ns/op    1.09x

   TYPE    OP    SIZE     RATE        SPEEDUP
   int8   sum     256     4.80 ns/op   15.42x
   int8   min     256     4.86 ns/op   38.25x
   int8   max     256     4.86 ns/op   37.66x
   int8   add     256     8.38 ns/op   13.32x
   int8   sub     256     8.24 ns/op   13.54x
   int8   mul     256     8.71 ns/op   12.38x
   int8   div     256   129.52 ns/op    1.12x
   int8   sum    4096    49.14 ns/op   21.24x
   int8   min    4096    50.77 ns/op   60.68x
   int8   max    4096    48.70 ns/op   63.33x
   int8   add    4096    62.65 ns/op   26.14x
   int8   sub    4096    62.24 ns/op   26.46x
   int8   mul    4096    59.96 ns/op   27.77x
   int8   div    4096  2073.18 ns/op    1.17x
   int8   sum   16384   247.78 ns/op   16.55x
   int8   min   16384   257.64 ns/op   52.10x
   int8   max   16384   236.66 ns/op   53.97x
   int8   add   16384   262.95 ns/op   26.29x
   int8   sub   16384   254.03 ns/op   27.76x
   int8   mul   16384   272.69 ns/op   26.59x
   int8   div   16384  8479.32 ns/op    1.15x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int16   sum     256     7.05 ns/op   10.97x
  int16   min     256     7.19 ns/op   26.90x
  int16   max     256     6.90 ns/op   26.61x
  int16   add     256    13.51 ns/op    8.06x
  int16   sub     256    12.27 ns/op    9.59x
  int16   mul     256    12.21 ns/op    8.90x
  int16   div     256   130.96 ns/op    1.13x
  int16   sum    4096   112.66 ns/op    9.51x
  int16   min    4096   108.11 ns/op   28.60x
  int16   max    4096   108.40 ns/op   29.06x
  int16   add    4096   125.41 ns/op   13.54x
  int16   sub    4096   119.49 ns/op   13.86x
  int16   mul    4096   123.22 ns/op   13.78x
  int16   div    4096  2074.25 ns/op    1.11x
  int16   sum   16384   494.65 ns/op    8.57x
  int16   min   16384   489.44 ns/op   25.05x
  int16   max   16384   493.22 ns/op   25.27x
  int16   add   16384   522.42 ns/op   12.49x
  int16   sub   16384   535.28 ns/op   12.33x
  int16   mul   16384   559.30 ns/op   11.72x
  int16   div   16384  8296.67 ns/op    1.12x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int32   sum     256    11.30 ns/op    6.41x
  int32   min     256    11.29 ns/op   11.20x
  int32   max     256    11.28 ns/op   11.24x
  int32   add     256    17.78 ns/op    6.06x
  int32   sub     256    17.78 ns/op    6.09x
  int32   mul     256    17.78 ns/op    6.07x
  int32   div     256   129.66 ns/op    1.13x
  int32   sum    4096   236.73 ns/op    4.40x
  int32   min    4096   237.77 ns/op    8.79x
  int32   max    4096   235.05 ns/op    8.97x
  int32   add    4096   225.48 ns/op    7.24x
  int32   sub    4096   240.99 ns/op    6.79x
  int32   mul    4096   258.20 ns/op    6.33x
  int32   div    4096  2075.19 ns/op    1.11x
  int32   sum   16384  1011.10 ns/op    4.18x
  int32   min   16384  1011.42 ns/op    8.41x
  int32   max   16384  1002.50 ns/op    8.39x
  int32   add   16384   881.46 ns/op    7.42x
  int32   sub   16384   884.55 ns/op    7.38x
  int32   mul   16384   887.31 ns/op    7.40x
  int32   div   16384  8352.29 ns/op    1.12x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int64   sum     256    35.13 ns/op    2.07x
  int64   min     256    41.41 ns/op    3.08x
  int64   max     256    41.26 ns/op    3.02x
  int64   add     256    30.90 ns/op    3.49x
  int64   sub     256    30.88 ns/op    3.49x
  int64   mul     256    71.46 ns/op    1.51x
  int64   div     256   134.15 ns/op    1.09x
  int64   sum    4096   527.85 ns/op    1.98x
  int64   min    4096   981.92 ns/op    2.15x
  int64   max    4096   985.04 ns/op    2.15x
  int64   add    4096   486.18 ns/op    3.36x
  int64   sub    4096   476.42 ns/op    3.43x
  int64   mul    4096  1094.60 ns/op    1.50x
  int64   div    4096  2141.80 ns/op    1.09x
  int64   sum   16384  2094.27 ns/op    2.13x
  int64   min   16384  4036.02 ns/op    2.07x
  int64   max   16384  4101.59 ns/op    2.07x
  int64   add   16384  3500.60 ns/op    1.92x
  int64   sub   16384  3485.66 ns/op    1.88x
  int64   mul   16384  4372.74 ns/op    1.50x
  int64   div   16384  9099.17 ns/op    1.05x

   TYPE    OP    SIZE     RATE        SPEEDUP
float32   sum     256    11.76 ns/op   10.34x
float32   min     256    11.25 ns/op   19.61x
float32   max     256    11.25 ns/op   15.58x
float32   add     256    18.06 ns/op    6.11x
float32   sub     256    17.85 ns/op    6.11x
float32   mul     256    17.81 ns/op    6.05x
float32   div     256    21.08 ns/op    5.11x
float32   sum    4096   320.75 ns/op    8.54x
float32   min    4096   232.22 ns/op   17.46x
float32   max    4096   231.89 ns/op   17.14x
float32   add    4096   277.66 ns/op    5.87x
float32   sub    4096   248.42 ns/op    6.56x
float32   mul    4096   240.00 ns/op    6.79x
float32   div    4096   288.28 ns/op    5.65x
float32   sum   16384  1384.83 ns/op    7.98x
float32   min   16384  1009.17 ns/op   16.04x
float32   max   16384  1006.63 ns/op   16.19x
float32   add   16384   884.13 ns/op    7.39x
float32   sub   16384   882.45 ns/op    7.42x
float32   mul   16384   882.46 ns/op    7.43x
float32   div   16384  1100.18 ns/op    5.95x

   TYPE    OP    SIZE     RATE        SPEEDUP
float64   sum     256    27.91 ns/op    4.33x
float64   min     256    21.68 ns/op   10.27x
float64   max     256    21.79 ns/op    8.05x
float64   add     256    30.51 ns/op    3.53x
float64   sub     256    30.42 ns/op    3.56x
float64   mul     256    30.48 ns/op    3.52x
float64   div     256    37.69 ns/op    2.86x
float64   sum    4096   669.96 ns/op    4.08x
float64   min    4096   489.15 ns/op    8.23x
float64   max    4096   499.26 ns/op    7.96x
float64   add    4096   485.25 ns/op    3.37x
float64   sub    4096   485.85 ns/op    3.37x
float64   mul    4096   476.16 ns/op    3.42x
float64   div    4096   574.07 ns/op    2.84x
float64   sum   16384  2805.05 ns/op    3.90x
float64   min   16384  2052.30 ns/op    7.90x
float64   max   16384  2070.18 ns/op    7.79x
float64   add   16384  3488.30 ns/op    1.87x
float64   sub   16384  3492.81 ns/op    1.87x
float64   mul   16384  3501.81 ns/op    1.86x
float64   div   16384  3490.82 ns/op    1.87x
```

# Acknowledgements

This library was originally inspired by the work of Valery Carey & Adrian Witas in [viant/vec](https://github.com/viant/vec) package, but instead of hand-rolled assembly and intrinsics I opted for using auto-vectorization for maintainability reasons.
