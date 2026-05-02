package tree_test

import (
    "cmp"
    "fmt"
    "math/rand/v2"
    "testing"
    "github.com/flxch/container/tree"
)


// Run all benchmarks as `go test -run=XXX -bench=. -cpu=1`.  Note that for the
// benchmarks, randomly generated trees are used.  Hence, the running times
// might vary from run to run.  However, the variance should be small, in
// particular, for the benchmarks with larger trees.

// Tree sizes for the benchmarks.
var sizes []int = []int{100, 1000, 10000}

var global *tree.Tree[int]

// Benchmark the lookup of an existing value in a tree.
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
                    if _, ok := global.Lookup(value); !ok {
                        global.Add(value)
                        break
                    }
                }
                b.StartTimer()
                _, _ = global.Lookup(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a tree of a fixed size.
                global.Remove(value)
            }
        })
    }
}

// Benchmark the insertion of a value into a tree.
func BenchmarkTree_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            global = random(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the tree.
                for {
                    value = rand.Int()
                    if _, ok := global.Lookup(value); !ok {
                        break
                    }
                }
                b.StartTimer()
                global.Add(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a tree of a fixed size.
                global.Remove(value)
            }
        })
    }
}

// Benchmark the insertion of a value into a tree.
// (no dublicates)
func BenchmarkTree_Insert(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            global = random(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the tree.
                for {
                    value = rand.Int()
                    if _, ok := global.Lookup(value); !ok {
                        break
                    }
                }
                b.StartTimer()
                global.Insert(value)
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a tree of a fixed size.
                global.Remove(value)
            }
        })
    }
}

// Benchmark the deletion of an existing value in a tree.
func BenchmarkTree_Remove(b *testing.B) {
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
                    if _, ok := global.Lookup(value); !ok {
                        global.Add(value)
                        break
                    }
                }
                b.StartTimer()
                global.Remove(value)
                b.StopTimer()
            }
        })
    }
}


// Benchmark the cloning of a tree.
var clone *tree.Tree[int]
func BenchmarkTree_Clone(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            for i := 0; i < b.N; i++ {
                // Generate a random tree containing `size` values.
                global = random(size)
                b.StartTimer()
                clone = global.Clone(func(v int) int { return v })
                b.StopTimer()
            }
        })
    }
}


// Benchmark the savings of the `go:nosplit` Go compiler directive.

// This benchmark hints that with `go:nosplit` some nanoseconds are saved.
// However, the other benchmarks do not support this.  A reason is that in the
// latter benchmarks not so many of the respective helper functions are called.
func BenchmarkBuild(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    perm := rand.Perm(100000)
    base := tree.New[int](cmp.Compare)
    for _, v := range perm[0:50000] {
        base.Add(v)
    }
    for i := 0; i < b.N; i++ {
        global = base.Clone(func(v int) int { return v })
        b.StartTimer()
        for _, v := range perm[50000:] {
            global.Add(v)
        }
        b.StopTimer()
    }
}

var resultIterate int

// Benchmark the iterations over a tree.

func BenchmarkIterate(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            for i := 0; i < b.N; i++ {
                // Generate a random tree containing [size] values.
                global = random(size)
                resultIterate = 0
                b.StartTimer()
                global.WalkAscend(func(_ int) { resultIterate++ })
                b.StopTimer()
                if resultIterate != size {
                    b.Errorf("wrong number of visited nodes")
                }
            }
        })
    }
}

func BenchmarkIterateGeq(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            for i := 0; i < b.N; i++ {
                // Generate a random tree containing [size] values.
                global = random(size)
                resultIterate = 0
                b.StartTimer()
                global.WalkAscendGeq(0, func(_ int) bool { resultIterate++; return true })
                b.StopTimer()
                if resultIterate != size {
                    b.Errorf("wrong number of visited nodes")
                }
            }
        })
    }
}
