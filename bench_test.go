package container_test

import (
    "math/rand/v2"
    "testing"
)


// Continer sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}


// Benchmarking various comparison functions.

var cmpFn  func(int, int) int
var lessFn func(int, int) bool
var k int
var u bool

func cmp(n, m int) int {
    switch {
    case n < m:
        return -1
    case n > m:
        return 1
    default:
        return 0
    }
}

func cmpsub(n, m int) int { return n - m }

func less(n, m int) bool { return n < m }

func BenchmarkCompare(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    cmpFn = cmp
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = cmpFn(x, y)
    }
}

func BenchmarkCompare_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = cmp(x, y)
    }
}

func BenchmarkCompareSub(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    cmpFn = cmpsub
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = cmpFn(x, y)
    }
}

func BenchmarkCompareSub_Inlined(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    x, y := rand.Int(), rand.Int()
    b.StartTimer()
    for i := 0; i < b.N; i++ {
        k = cmpsub(x, y)
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

