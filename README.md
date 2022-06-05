<p align="center">
<img width="330" height="110" src=".github/logo.png" border="0" alt="kelindar/simd">
<br>
<img src="https://img.shields.io/github/go-mod/go-version/kelindar/simd" alt="Go Version">
<a href="https://pkg.go.dev/github.com/kelindar/simd"><img src="https://pkg.go.dev/badge/github.com/kelindar/simd" alt="PkgGoDev"></a>
<a href="https://goreportcard.com/report/github.com/kelindar/simd"><img src="https://goreportcard.com/badge/github.com/kelindar/simd" alt="Go Report Card"></a>
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
float32   sum     256     8.86 ns/op   20.54x
float32   min     256    27.57 ns/op    4.15x
float32   max     256    27.65 ns/op    4.07x
float32   add     256    15.35 ns/op    7.43x
float32   sub     256    15.31 ns/op    7.44x
float32   mul     256    15.39 ns/op    7.39x
float32   div     256    27.60 ns/op    5.90x
float32   sum    4096   107.74 ns/op   31.65x
float32   min    4096   133.08 ns/op   12.89x
float32   max    4096   140.77 ns/op   12.26x
float32   add    4096   373.34 ns/op    4.71x
float32   sub    4096   389.79 ns/op    4.40x
float32   mul    4096   369.63 ns/op    4.69x
float32   div    4096   457.12 ns/op    5.65x
float32   sum   16384   466.87 ns/op   29.47x
float32   min   16384   504.66 ns/op   13.64x
float32   max   16384   515.63 ns/op   13.33x
float32   add   16384  1422.20 ns/op    4.88x
float32   sub   16384  1414.59 ns/op    4.86x
float32   mul   16384  1426.91 ns/op    4.84x
float32   div   16384  1847.74 ns/op    5.54x

   TYPE    OP    SIZE     RATE        SPEEDUP
float64   sum     256    14.28 ns/op   13.01x
float64   min     256    24.58 ns/op    4.73x
float64   max     256    24.28 ns/op    3.77x
float64   add     256    25.44 ns/op    4.50x
float64   sub     256    25.40 ns/op    4.44x
float64   mul     256    25.66 ns/op    4.47x
float64   div     256   107.58 ns/op    2.01x
float64   sum    4096   217.11 ns/op   15.64x
float64   min    4096   237.68 ns/op    7.21x
float64   max    4096   238.10 ns/op    5.68x
float64   add    4096   779.90 ns/op    2.33x
float64   sub    4096   752.78 ns/op    2.31x
float64   mul    4096   737.30 ns/op    2.43x
float64   div    4096  1735.34 ns/op    1.98x
float64   sum   16384   922.49 ns/op   15.07x
float64   min   16384   989.75 ns/op    6.98x
float64   max   16384  1017.51 ns/op    5.42x
float64   add   16384  5987.32 ns/op    1.42x
float64   sub   16384  5383.73 ns/op    1.50x
float64   mul   16384  5674.55 ns/op    1.44x
float64   div   16384  7063.43 ns/op    2.00x

   TYPE    OP    SIZE     RATE        SPEEDUP
  uint8   sum     100    35.63 ns/op    1.02x
  uint8   min     100    23.77 ns/op    2.17x
  uint8   max     100    29.69 ns/op    1.59x
  uint8   add     100    37.46 ns/op    1.41x
  uint8   sub     100    36.63 ns/op    1.45x
  uint8   mul     100     9.42 ns/op    8.03x
  uint8   div     100   193.95 ns/op    1.14x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint16   sum     256     6.24 ns/op   13.95x
 uint16   min     256    29.83 ns/op    5.49x
 uint16   max     256    43.57 ns/op    2.54x
 uint16   add     256    10.30 ns/op   17.19x
 uint16   sub     256    10.29 ns/op   16.14x
 uint16   mul     256    10.36 ns/op   16.57x
 uint16   div     256   440.84 ns/op    1.27x
 uint16   sum    4096    42.39 ns/op   30.41x
 uint16   min    4096    79.48 ns/op   32.70x
 uint16   max    4096    95.83 ns/op   18.66x
 uint16   add    4096    90.78 ns/op   29.94x
 uint16   sub    4096    87.41 ns/op   29.95x
 uint16   mul    4096    86.62 ns/op   29.73x
 uint16   div    4096  6867.75 ns/op    1.28x
 uint16   sum   16384   155.46 ns/op   32.00x
 uint16   min   16384   246.67 ns/op   42.19x
 uint16   max   16384   262.69 ns/op   26.06x
 uint16   add   16384   739.30 ns/op   14.01x
 uint16   sub   16384   721.78 ns/op   14.33x
 uint16   mul   16384   787.84 ns/op   12.89x
 uint16   div   16384 26838.80 ns/op    1.28x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint32   sum     256     7.94 ns/op   10.47x
 uint32   min     256    22.62 ns/op    6.96x
 uint32   max     256    25.75 ns/op    4.20x
 uint32   add     256    14.72 ns/op    8.31x
 uint32   sub     256    14.80 ns/op    8.32x
 uint32   mul     256    15.93 ns/op    7.52x
 uint32   div     256   419.09 ns/op    1.30x
 uint32   sum    4096    74.25 ns/op   16.59x
 uint32   min    4096   119.60 ns/op   21.61x
 uint32   max    4096   125.50 ns/op   13.67x
 uint32   add    4096   382.51 ns/op    5.09x
 uint32   sub    4096   369.66 ns/op    5.27x
 uint32   mul    4096   376.69 ns/op    4.82x
 uint32   div    4096  6611.42 ns/op    1.30x
 uint32   sum   16384   470.78 ns/op   10.45x
 uint32   min   16384   509.34 ns/op   20.05x
 uint32   max   16384   509.63 ns/op   14.20x
 uint32   add   16384  1430.07 ns/op    5.62x
 uint32   sub   16384  1470.22 ns/op    5.37x
 uint32   mul   16384  1469.16 ns/op    5.45x
 uint32   div   16384 26749.18 ns/op    1.31x

   TYPE    OP    SIZE     RATE        SPEEDUP
 uint64   sum     256    12.16 ns/op    9.29x
 uint64   min     256    41.61 ns/op    3.80x
 uint64   max     256    44.33 ns/op    2.59x
 uint64   add     256    24.87 ns/op    4.95x
 uint64   sub     256    24.85 ns/op    4.95x
 uint64   mul     256    50.84 ns/op    2.35x
 uint64   div     256  1241.17 ns/op    1.21x
 uint64   sum    4096   142.52 ns/op   11.94x
 uint64   min    4096   502.09 ns/op    5.15x
 uint64   max    4096   501.85 ns/op    3.43x
 uint64   add    4096   736.55 ns/op    2.66x
 uint64   sub    4096   768.94 ns/op    2.59x
 uint64   mul    4096   899.77 ns/op    2.24x
 uint64   div    4096 19940.91 ns/op    1.21x
 uint64   sum   16384   976.09 ns/op    7.30x
 uint64   min   16384  2107.63 ns/op    5.02x
 uint64   max   16384  2123.41 ns/op    3.33x
 uint64   add   16384  5605.09 ns/op    1.53x
 uint64   sub   16384  5382.83 ns/op    1.57x
 uint64   mul   16384  5934.57 ns/op    1.43x
 uint64   div   16384 79974.54 ns/op    1.18x

   TYPE    OP    SIZE     RATE        SPEEDUP
   int8   sum     100    35.60 ns/op    1.01x
   int8   min     100    24.12 ns/op    2.10x
   int8   max     100    24.06 ns/op    1.96x
   int8   add     100    35.74 ns/op    1.45x
   int8   sub     100    36.45 ns/op    1.43x
   int8   mul     100     9.42 ns/op    7.55x
   int8   div     100   203.37 ns/op    1.46x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int16   sum     256     6.11 ns/op   13.56x
  int16   min     256    30.70 ns/op    3.43x
  int16   max     256    29.39 ns/op    3.66x
  int16   add     256     9.81 ns/op   17.06x
  int16   sub     256     9.84 ns/op   17.05x
  int16   mul     256    10.21 ns/op   16.58x
  int16   div     256   489.38 ns/op    1.55x
  int16   sum    4096    37.85 ns/op   32.80x
  int16   min    4096    79.82 ns/op   21.40x
  int16   max    4096    78.70 ns/op   21.86x
  int16   add    4096    85.27 ns/op   30.57x
  int16   sub    4096    84.76 ns/op   30.43x
  int16   mul    4096    88.35 ns/op   29.32x
  int16   div    4096  7896.65 ns/op    1.51x
  int16   sum   16384   152.95 ns/op   32.75x
  int16   min   16384   243.98 ns/op   28.40x
  int16   max   16384   245.85 ns/op   27.93x
  int16   add   16384   783.82 ns/op   13.01x
  int16   sub   16384   740.70 ns/op   14.02x
  int16   mul   16384   736.79 ns/op   14.00x
  int16   div   16384 31574.74 ns/op    1.53x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int32   sum     256     7.92 ns/op   10.50x
  int32   min     256    21.82 ns/op    4.82x
  int32   max     256    22.71 ns/op    4.65x
  int32   add     256    14.98 ns/op    8.25x
  int32   sub     256    14.91 ns/op    8.22x
  int32   mul     256    15.95 ns/op    7.54x
  int32   div     256   488.63 ns/op    1.38x
  int32   sum    4096    74.84 ns/op   16.85x
  int32   min    4096   119.84 ns/op   14.27x
  int32   max    4096   119.55 ns/op   14.33x
  int32   add    4096   383.44 ns/op    5.07x
  int32   sub    4096   373.52 ns/op    5.20x
  int32   mul    4096   373.64 ns/op    4.76x
  int32   div    4096  7791.16 ns/op    1.37x
  int32   sum   16384   468.25 ns/op   10.68x
  int32   min   16384   483.94 ns/op   14.19x
  int32   max   16384   485.43 ns/op   14.13x
  int32   add   16384  1411.90 ns/op    5.57x
  int32   sub   16384  1401.61 ns/op    5.55x
  int32   mul   16384  1422.65 ns/op    5.40x
  int32   div   16384 31080.74 ns/op    1.37x

   TYPE    OP    SIZE     RATE        SPEEDUP
  int64   sum     256    12.01 ns/op    6.79x
  int64   min     256    29.99 ns/op    3.47x
  int64   max     256    29.96 ns/op    3.53x
  int64   add     256    24.87 ns/op    4.96x
  int64   sub     256    24.83 ns/op    4.94x
  int64   mul     256    50.67 ns/op    2.34x
  int64   div     256  1546.44 ns/op    1.15x
  int64   sum    4096   148.93 ns/op    8.48x
  int64   min    4096   323.19 ns/op    5.27x
  int64   max    4096   322.25 ns/op    5.35x
  int64   add    4096   721.78 ns/op    2.72x
  int64   sub    4096   725.84 ns/op    2.71x
  int64   mul    4096   901.09 ns/op    2.10x
  int64   div    4096 24406.86 ns/op    1.17x
  int64   sum   16384   930.98 ns/op    5.64x
  int64   min   16384  1427.03 ns/op    4.78x
  int64   max   16384  1429.45 ns/op    4.78x
  int64   add   16384  5703.49 ns/op    1.48x
  int64   sub   16384  5103.69 ns/op    1.64x
  int64   mul   16384  6019.09 ns/op    1.42x
  int64   div   16384 98294.20 ns/op    1.15x

```

# Acknowledgements

This library was originally inspired by the work of Valery Carey & Adrian Witas in [viant/vec](https://github.com/viant/vec) package, but instead of hand-rolled assembly and intrinsics I opted for using auto-vectorization for maintainability reasons.
