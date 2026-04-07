package skiplist_test

import (
    "fmt"
    "math/rand"
    "testing"
    "github.com/flxch/container/skiplist"
)


//go:nosplit
func compare(k, l int) int {
    switch {
    case k < l: // less than
        return -1
    case k > l: // greater than
        return 1
    default:    // equal to
        return 0
    }
}

// `build` returns a skiplist with the values 0, ..., `n`-1.  The values are
// added in a random order.
func build(n int) *skiplist.Skiplist[int] {
    sl := skiplist.New(true, compare)
    for _, v := range rand.Perm(n) {
        sl.Add(v)
    }
    return sl
}

// `random` returns a skiplist with `n` random integers.
func random(n int) *skiplist.Skiplist[int] {
    sl := skiplist.New(true, compare)
    //sl.SetHeight(8)
    for sl.Len() < n {
        sl.Add(rand.Int())
    }
    return sl
}


// Skiplist sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}

var global *skiplist.Skiplist[int]

// Benchmark the insertion of a value into a skiplist.
func BenchmarkSkiplist_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random skiplist containing `size` values.
            global = random(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the skiplist.
                for {
                    value = rand.Int()
                    if global.Lookup(value) == nil {
                        break
                    }
                }
                b.StartTimer()
                global.Add(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a skiplist of a fixed size.
                global.Remove(global.Lookup(value))
            }
        })
    }
}

// Benchmark the deletion of an existing value in a skiplist.
func BenchmarkSkiplist_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            global = random(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the
                // deletion of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if global.Lookup(value) == nil {
                        global.Add(value)
                        break
                    }
                }
                b.StartTimer()
                global.Remove(global.Lookup(value))
                b.StopTimer()
            }
        })
    }
}

// Benchmark the lookup of an existing value in a skiplist.
func BenchmarkTree_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            global = random(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if global.Lookup(value) == nil {
                        global.Add(value)
                        break
                    }
                }
                b.StartTimer()
                elem := global.Lookup(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a tree of a fixed size.
                global.Remove(elem)
            }
        })
    }
}


func BenchmarkBuild(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            perm := rand.Perm(size)
            base := skiplist.New(true, compare)
            for _, v := range perm[0:size/2] {
                base.Add(v)
            }
            for i := 0; i < b.N; i++ {
                global = base.Clone(func(v int) int { return v })
                b.StartTimer()
                for _, v := range perm[size/2:] {
                    global.Add(v)
                }
                b.StopTimer()
            }
        })
    }
}
