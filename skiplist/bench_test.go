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

// `build` returns a skip list with the values 0, ..., `n`-1.  The values are
// added in a random order.
func build(n int) *skiplist.Skiplist[int] {
    sl := skiplist.New(compare)
    sl.SetPrealloc(6)
    for _, v := range rand.Perm(n) {
        sl.Add(v)
    }
    return sl
}

// `random` returns a skip list with `n` random integers.
func random(n int) *skiplist.Skiplist[int] {
    sl := skiplist.New(compare)
    sl.SetPrealloc(6)
    for sl.Len() < n {
        sl.Add(rand.Int())
    }
    return sl
}


// Skip list sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}

var global *skiplist.Skiplist[int]

// Benchmark the insertion of a value into a skip list.
func BenchmarkSkiplist_Add(b *testing.B) {
    //skiplist.DefaultSeed1 = 143
    //skiplist.DefaultSeed2 = 34
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random skip list containing `size` values.
            global = random(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the skip list.
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
                // insertion of a value into a skip list of a fixed size.
                global.Remove(global.Lookup(value))
            }
        })
    }
}

// Benchmark the deletion of an existing value in a skip list.
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

// Benchmark the lookup of an existing value in a skip list.
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


// Benchmark iterating through the elements of a skip list.
func BenchmarkIterators(b *testing.B) {
    var sum int

    // The old style before the introduction of range iterators.
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d elems", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            perm := rand.Perm(size)
            sl := skiplist.New(compare)
            for _, v := range perm {
                sl.Add(v)
            }
            b.StartTimer()
            for i := 0; i < b.N; i++ {
                for elem := sl.Front(); elem != nil; elem = elem.Next() {
                    sum += elem.Value
                }
            }
        })
    }

    // The new style using range iterators (slower!)
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d range", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            perm := rand.Perm(size)
            sl := skiplist.New(compare)
            for _, v := range perm {
                sl.Add(v)
            }
            b.StartTimer()
            for i := 0; i < b.N; i++ {
                for val := range sl.Ascend {
                    sum += val
                }
            }
        })
    }

    // Using functions to iterate through the list (slower, similar to range
    // iterators),
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d func", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            perm := rand.Perm(size)
            sl := skiplist.New(compare)
            for _, v := range perm {
                sl.Add(v)
            }
            b.StartTimer()
            for i := 0; i < b.N; i++ {
                sl.WalkAscend(func(val int) { sum += val })
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
            base := skiplist.New(compare)
            base.SetPrealloc(6)
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
