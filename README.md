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

It currently supports `AVX2` on `amd64` and `NEON` (Advanced SIMD) on `arm64` (including Apple Silicon). Most of the code in this library is auto-generated, which helps with maintenance.

## Usage

The API is intentionally simple and non-opinionated:

- Reduction ops: `Sum*`, `Min*`, `Max*`
- Element-wise ops: `Add*`, `Sub*`, `Mul*`, `Div*`
- Typed fast paths for `int*`, `uint*`, `float*` slices
- Generic fallback when SIMD is unavailable

### Examples

Compute a sum:

```go
sum := simd.SumFloat32s([]float32{1, 2, 3, 4, 5})
```

Element-wise add into a destination buffer:

```go
a := []float32{1, 2, 3, 4}
b := []float32{10, 20, 30, 40}
dst := make([]float32, len(a))

simd.AddFloat32s(dst, a, b)
// dst => []float32{11, 22, 33, 44}
```

Generic API (works across numeric slice types):

```go
values := []int16{7, 2, 9, 4}
min := simd.Min(values) // 2
max := simd.Max(values) // 9
sum := simd.Sum(values) // 22
```

## Benchmarks

```go
goos: windows
goarch: amd64
pkg: github.com/kelindar/simd
cpu: 13th Gen Intel(R) Core(TM) i7-13700K

   TYPE    OP    SIZE     RATE        SPEEDUP
  uint8   sum     256     3.43 ns/op   13.68x
  uint8   min     256     3.49 ns/op   19.12x
  uint8   max     256     3.48 ns/op   21.08x
  uint8   add     256     4.23 ns/op   19.37x
  uint8   sub     256     4.25 ns/op   19.44x
  uint8   mul     256     6.06 ns/op   13.52x
  uint8   div     256   289.94 ns/op    1.20x
  uint8   sum    4096    16.72 ns/op   46.02x
  uint8   min    4096    16.59 ns/op   55.02x
  uint8   max    4096    16.90 ns/op   54.28x
  uint8   add    4096    30.64 ns/op   38.13x
  uint8   sub    4096    29.15 ns/op   39.93x
  uint8   mul    4096    79.53 ns/op   14.69x
  uint8   div    4096  4646.76 ns/op    1.19x
  uint8   sum   16384    58.88 ns/op   52.57x
  uint8   min   16384    59.61 ns/op   61.48x
  uint8   max   16384    59.34 ns/op   61.39x
  uint8   add   16384   159.51 ns/op   29.68x
  uint8   sub   16384   147.56 ns/op   31.72x
  uint8   mul   16384   313.20 ns/op   14.88x
  uint8   div   16384 18551.34 ns/op    1.18x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint16   sum     256     4.58 ns/op    9.93x
 uint16   min     256     4.27 ns/op   32.99x
 uint16   max     256     4.45 ns/op   19.57x
 uint16   add     256     5.53 ns/op   17.34x
 uint16   sub     256     5.53 ns/op   18.43x
 uint16   mul     256     5.53 ns/op   18.62x
 uint16   div     256   342.14 ns/op    1.01x
 uint16   sum    4096    32.29 ns/op   23.92x
 uint16   min    4096    30.82 ns/op   75.31x
 uint16   max    4096    31.26 ns/op   49.09x
 uint16   add    4096    57.24 ns/op   25.48x
 uint16   sub    4096    57.14 ns/op   26.93x
 uint16   mul    4096    46.25 ns/op   34.19x
 uint16   div    4096  5439.07 ns/op    1.00x
 uint16   sum   16384   114.01 ns/op   27.21x
 uint16   min   16384   110.94 ns/op   84.12x
 uint16   max   16384   112.55 ns/op   54.88x
 uint16   add   16384   535.12 ns/op   10.39x
 uint16   sub   16384   970.32 ns/op    6.34x
 uint16   mul   16384   986.37 ns/op   12.71x
 uint16   div   16384 31041.77 ns/op    0.79x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint32   sum     256    10.71 ns/op    9.66x
 uint32   min     256    11.45 ns/op   16.55x
 uint32   max     256    11.10 ns/op   16.65x
 uint32   add     256    20.22 ns/op    7.82x
 uint32   sub     256    20.59 ns/op    7.67x
 uint32   mul     256    30.25 ns/op    5.88x
 uint32   div     256   370.89 ns/op    1.03x
 uint32   sum    4096   125.98 ns/op   12.72x
 uint32   min    4096   133.07 ns/op   22.64x
 uint32   max    4096   132.61 ns/op   22.74x
 uint32   add    4096   408.41 ns/op    6.22x
 uint32   sub    4096   412.60 ns/op    6.16x
 uint32   mul    4096   507.06 ns/op    5.55x
 uint32   div    4096  6041.48 ns/op    1.00x
 uint32   sum   16384   649.25 ns/op   10.07x
 uint32   min   16384   637.87 ns/op   18.85x
 uint32   max   16384   645.31 ns/op   18.70x
 uint32   add   16384  1975.68 ns/op    5.06x
 uint32   sub   16384  1991.51 ns/op    4.94x
 uint32   mul   16384  2033.82 ns/op    5.41x
 uint32   div   16384 24277.79 ns/op    0.99x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint64   sum     256    18.53 ns/op    5.58x
 uint64   min     256    95.22 ns/op    1.98x
 uint64   max     256    98.92 ns/op    1.91x
 uint64   add     256    36.25 ns/op    4.35x
 uint64   sub     256    35.94 ns/op    4.44x
 uint64   mul     256   101.35 ns/op    1.75x
 uint64   div     256   383.12 ns/op    1.00x
 uint64   sum    4096   296.89 ns/op    5.39x
 uint64   min    4096  1593.28 ns/op    1.90x
 uint64   max    4096  1572.11 ns/op    1.92x
 uint64   add    4096   976.07 ns/op    2.55x
 uint64   sub    4096   984.38 ns/op    2.56x
 uint64   mul    4096  1709.65 ns/op    1.63x
 uint64   div    4096  6072.12 ns/op    1.01x
 uint64   sum   16384  1280.59 ns/op    5.17x
 uint64   min   16384  6189.62 ns/op    1.96x
 uint64   max   16384  6194.25 ns/op    1.93x
 uint64   add   16384  4021.97 ns/op    2.55x
 uint64   sub   16384  3982.78 ns/op    2.57x
 uint64   mul   16384  6725.83 ns/op    1.64x
 uint64   div   16384 24463.65 ns/op    1.02x

   TYPE    OP    SIZE     RATE        SPEEDUP
   int8   sum     256     6.60 ns/op   16.73x
   int8   min     256     7.26 ns/op   19.91x
   int8   max     256     7.26 ns/op   20.19x
   int8   add     256     9.21 ns/op   18.82x
   int8   sub     256     9.39 ns/op   18.27x
   int8   mul     256    17.92 ns/op    9.83x
   int8   div     256   818.03 ns/op    0.71x
   int8   sum    4096    38.16 ns/op   42.02x
   int8   min    4096    38.91 ns/op   57.41x
   int8   max    4096    38.70 ns/op   56.66x
   int8   add    4096    75.48 ns/op   36.97x
   int8   sub    4096    74.46 ns/op   37.00x
   int8   mul    4096   226.79 ns/op   11.85x
   int8   div    4096 13120.54 ns/op    0.69x
   int8   sum   16384   131.28 ns/op   49.58x
   int8   min   16384   131.36 ns/op   68.74x
   int8   max   16384   132.08 ns/op   68.57x
   int8   add   16384   417.09 ns/op   26.35x
   int8   sub   16384   411.26 ns/op   26.84x
   int8   mul   16384   900.74 ns/op   12.24x
   int8   div   16384 52317.05 ns/op    0.69x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int16   sum     256     8.17 ns/op   13.64x
  int16   min     256     8.50 ns/op   22.13x
  int16   max     256     8.49 ns/op   21.84x
  int16   add     256    12.55 ns/op   14.16x
  int16   sub     256    12.90 ns/op   13.65x
  int16   mul     256    12.81 ns/op   15.47x
  int16   div     256   523.61 ns/op    1.10x
  int16   sum    4096    66.69 ns/op   23.65x
  int16   min    4096    66.74 ns/op   45.50x
  int16   max    4096    66.81 ns/op   44.62x
  int16   add    4096   130.95 ns/op   21.24x
  int16   sub    4096   130.76 ns/op   21.18x
  int16   mul    4096   130.26 ns/op   23.53x
  int16   div    4096  8162.28 ns/op    1.12x
  int16   sum   16384   290.53 ns/op   23.18x
  int16   min   16384   303.05 ns/op   39.85x
  int16   max   16384   306.08 ns/op   38.30x
  int16   add   16384  1000.12 ns/op   11.27x
  int16   sub   16384   996.05 ns/op   11.19x
  int16   mul   16384  1009.52 ns/op   12.17x
  int16   div   16384 32518.88 ns/op    1.13x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int32   sum     256    10.79 ns/op    9.84x
  int32   min     256    11.42 ns/op   16.40x
  int32   max     256    10.96 ns/op   17.12x
  int32   add     256    20.39 ns/op    7.45x
  int32   sub     256    19.83 ns/op    7.16x
  int32   mul     256    30.90 ns/op    5.61x
  int32   div     256   379.47 ns/op    1.01x
  int32   sum    4096   130.67 ns/op   12.45x
  int32   min    4096   134.71 ns/op   22.40x
  int32   max    4096   125.29 ns/op   24.25x
  int32   add    4096   412.80 ns/op    6.18x
  int32   sub    4096   417.97 ns/op    6.11x
  int32   mul    4096   505.23 ns/op    5.21x
  int32   div    4096  6085.25 ns/op    1.00x
  int32   sum   16384   667.40 ns/op    9.69x
  int32   min   16384   655.49 ns/op   18.18x
  int32   max   16384   648.01 ns/op   18.86x
  int32   add   16384  1995.43 ns/op    5.04x
  int32   sub   16384  1961.25 ns/op    5.03x
  int32   mul   16384  2040.80 ns/op    5.19x
  int32   div   16384 24338.73 ns/op    1.00x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int64   sum     256     9.26 ns/op   11.19x
  int64   min     256    25.42 ns/op    3.39x
  int64   max     256    80.10 ns/op    1.58x
  int64   add     256    36.52 ns/op    4.34x
  int64   sub     256    36.71 ns/op    4.36x
  int64   mul     256   106.45 ns/op    1.63x
  int64   div     256   380.91 ns/op    1.03x
  int64   sum    4096   295.11 ns/op    5.59x
  int64   min    4096  1132.89 ns/op    2.68x
  int64   max    4096  1165.34 ns/op    2.61x
  int64   add    4096   997.97 ns/op    2.53x
  int64   sub    4096   976.80 ns/op    2.58x
  int64   mul    4096  1721.22 ns/op    1.63x
  int64   div    4096  6124.12 ns/op    1.01x
  int64   sum   16384  1279.70 ns/op    5.17x
  int64   min   16384  4355.66 ns/op    2.79x
  int64   max   16384  4553.27 ns/op    2.61x
  int64   add   16384  4003.71 ns/op    2.55x
  int64   sub   16384  4150.76 ns/op    2.45x
  int64   mul   16384  6037.59 ns/op    2.43x
  int64   div   16384 24871.54 ns/op    0.99x

   TYPE    OP    SIZE     RATE        SPEEDUP
float32   sum     256    12.07 ns/op   12.44x
float32   min     256    12.33 ns/op   11.50x
float32   max     256    11.36 ns/op   13.25x
float32   add     256    19.95 ns/op    8.08x
float32   sub     256    19.54 ns/op    7.94x
float32   mul     256    19.58 ns/op    8.24x
float32   div     256    59.00 ns/op    5.31x
float32   sum    4096   132.69 ns/op   22.35x
float32   min    4096   131.46 ns/op   17.27x
float32   max    4096   131.15 ns/op   16.92x
float32   add    4096   370.71 ns/op    6.69x
float32   sub    4096   415.25 ns/op    6.06x
float32   mul    4096   412.00 ns/op    5.93x
float32   div    4096   946.05 ns/op    5.12x
float32   sum   16384   623.06 ns/op   19.16x
float32   min   16384   650.93 ns/op   13.67x
float32   max   16384   640.29 ns/op   14.44x
float32   add   16384  2056.12 ns/op    4.95x
float32   sub   16384  2002.50 ns/op    4.99x
float32   mul   16384  2048.68 ns/op    4.79x
float32   div   16384  4053.14 ns/op    5.01x

   TYPE    OP    SIZE     RATE        SPEEDUP
float64   sum     256    19.07 ns/op    8.82x
float64   min     256    19.35 ns/op    7.59x
float64   max     256    19.11 ns/op    7.89x
float64   add     256    37.08 ns/op    4.17x
float64   sub     256    32.91 ns/op    5.00x
float64   mul     256    36.15 ns/op    4.44x
float64   div     256   505.90 ns/op    1.01x
float64   sum    4096   268.04 ns/op   11.24x
float64   min    4096   284.08 ns/op    7.97x
float64   max    4096   301.93 ns/op    7.79x
float64   add    4096  1013.73 ns/op    2.53x
float64   sub    4096   992.44 ns/op    2.54x
float64   mul    4096   967.94 ns/op    2.62x
float64   div    4096  7182.29 ns/op    1.10x
float64   sum   16384  1242.69 ns/op    9.33x
float64   min   16384  1268.02 ns/op    7.10x
float64   max   16384  1273.06 ns/op    7.15x
float64   add   16384  4086.44 ns/op    2.47x
float64   sub   16384  4026.68 ns/op    2.53x
float64   mul   16384  4163.52 ns/op    2.41x
float64   div   16384 32172.72 ns/op    0.98x
PASS
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
