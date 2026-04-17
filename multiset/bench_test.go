package multiset_test

import (
    "fmt"
    "math/rand/v2"
    "testing"
    "github.com/flxch/container/multiset"
)


// `random` returns a multiset with `n` random integers as support.  The
// elements' multiplicity is also randomly chosen between 0 and 9.
func random(n int) *multiset.Multiset[int] {
    S := multiset.New[int]()
    for S.Len() < n {
        S.Add(rand.Int(), rand.IntN(10))
    }
    return S
}


// Multiset support sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}

var global *multiset.Multiset[int]

// Benchmark the insertion of a new element in a multiset.
func BenchmarkMultiset_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random multiset with a support of size `size`.
            global = random(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the multiset.
                for {
                    value = rand.Int()
                    if global.Lookup(value) == 0 {
                        break
                    }
                }
                b.StartTimer()
                global.Add(value, 1)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a multiset of a fixed size.
                global.Remove(value, 1)
            }
        })
    }
}

// Benchmark the deletion of an existing value in a multiset.
func BenchmarkMultiset_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random multiset with a support of size `size`.
            global = random(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the
                // deletion of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if global.Lookup(value) == 0 {
                        global.Add(value, 1)
                        break
                    }
                }
                b.StartTimer()
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a multiset of a fixed size.
                global.Remove(value, 1)
                b.StopTimer()
            }
        })
    }
}

// Benchmark the lookup of an existing value in a multiset.
func BenchmarkMultiset_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random multiset with a support of size `size`.
            global = random(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if global.Lookup(value) == 0 {
                        global.Add(value, 1)
                        break
                    }
                }
                b.StartTimer()
                _ = global.Lookup(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a multiset of a fixed size.
                global.Remove(value, 1)
            }
        })
    }
}


// Benchmark the cloning of a multiset.
var clone *multiset.Multiset[int]
func BenchmarkMultiset_Clone(b *testing.B) {
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
