#!/bin/bash

PATH="$PATH:./bin"
CLANG_OPTS="-mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti -ffast-math -Ofast"

function build_avx2_amd64 {
    SRC="simd_avx2_amd64.cpp"
    ASM="simd_avx2_amd64.s"
    clang-10 -S -mavx2 -masm=intel $CLANG_OPTS -o $ASM $SRC
    c2goasm -a -f $ASM ../$ASM
    rm $ASM
}

function build_avx512_amd64 {
    SRC="simd_avx512_amd64.cpp"
    ASM="simd_avx512_amd64.s"
    clang-10 -S -mavx512vl -masm=intel $CLANG_OPTS -o $ASM $SRC
    c2goasm -a -f $ASM ../$ASM
    rm $ASM
}

function build_neon_arm64 {
    SRC="simd_neon_arm64.cpp"
    ASM="simd_neon_arm64.s"
    clang-10 -S -arch arm64 $CLANG_OPTS -o $ASM $SRC
    c2goasm -a -f $ASM ../$ASM
    rm $ASM
}

build_avx2_amd64
#build_avx512_amd64
#build_neon_arm64
