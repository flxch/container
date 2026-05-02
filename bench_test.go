package container_test

import (
    "cmp"
    "math/rand/v2"
    "testing"
    "unsafe"
)


// Continer sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}


// Benchmarking various comparison functions.

var compareFn     func(int, int) int
var lessFn        func(int, int) bool
var lessgreaterFn func(int, int) (bool, bool)
var k int
var u bool
var v bool

func compare(n, m int) int {
    if n < m {
        return -1
    }
    if m > n {
        return 1
    }
    return 0
    //    switch {
    //    case n < m:
    //        return -1
    //    case n > m:
    //        return 1
    //    default:
    //        return 0
    //    }
}


// Overflows and underflows are possibe for negative integers.  Save if only
// using the positive integers of the int type.
func comparesub(n, m int) int {
    return n - m
}

func bool2int(b bool) int {
    return *(*int)(unsafe.Pointer(&b))
}

func comparebool(n, m int) int {
    return bool2int(n < m) - bool2int(m < n)
}

// Two calls are necessary when checking for equality.
func less(n, m int) bool {
    return n < m
}

func lessgreater(n, m int) (bool, bool) {
    return n < m, n > m
}


func BenchmarkCompare(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    compareFn = compare
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = compareFn(x, y)
    }
}

func BenchmarkCompare_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = compare(x, y)
    }
}

func BenchmarkCompareGeneric(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    compareFn = cmp.Compare
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = compareFn(x, y)
    }
}

func BenchmarkCompareGeneric_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = cmp.Compare(x, y)
    }
}

func BenchmarkCompareSub(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    compareFn = comparesub
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = compareFn(x, y)
    }
}

func BenchmarkCompareSub_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = comparesub(x, y)
    }
}

func BenchmarkCompareBool(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    compareFn = comparebool
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = compareFn(x, y)
    }
}

func BenchmarkCompareBool_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = comparebool(x, y)
    }
}


func BenchmarkCompareLess(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    lessFn = less
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        u = lessFn(x, y)
    }
}

func BenchmarkCompareLess_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        u = less(x, y)
    }
}

func BenchmarkCompareLessgreater(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    lessgreaterFn = lessgreater
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        u, v = lessgreaterFn(x, y)
    }
}

func BenchmarkCompareLessGreater_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        u, v = lessgreater(x, y)
    }
}

