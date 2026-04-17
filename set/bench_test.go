package set_test

import (
    "fmt"
    "math/rand/v2"
    "testing"
    "github.com/flxch/container/set"
)


// `random` returns a set containing `n` random integers.
func random(n int) set.Set[int] {
    S := set.New[int]()
    for S.Len() < n {
        S.Add(rand.Int())
    }
    return S
}


// Set sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}

var global set.Set[int]

// Benchmark the insertion of a new element in a set.
func BenchmarkSet_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random set of size `size`.
            global = random(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the set.
                for {
                    value = rand.Int()
                    if !global.Lookup(value) {
                        break
                    }
                }
                b.StartTimer()
                global.Add(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a set of a fixed size.
                global.Remove(value)
            }
        })
    }
}

// Benchmark the deletion of an existing value in a set.
func BenchmarkSet_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random set of size `size`.
            global = random(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the
                // deletion of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if !global.Lookup(value) {
                        global.Add(value)
                        break
                    }
                }
                b.StartTimer()
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a set of a fixed size.
                global.Remove(value)
                b.StopTimer()
            }
        })
    }
}

// Benchmark the lookup of an existing value in a set.
func BenchmarkSet_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random set of size `size`.
            global = random(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if !global.Lookup(value) {
                        global.Add(value)
                        break
                    }
                }
                b.StartTimer()
                _ = global.Lookup(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a set of a fixed size.
                global.Remove(value)
            }
        })
    }
}


// Benchmark the cloning of a set.
var clone set.Set[int]
func BenchmarkSet_Clone(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            for i := 0; i < b.N; i++ {
                global = random(size)
                b.StartTimer()
                clone = global.Clone()
                b.StopTimer()
            }
        })
    }
}
