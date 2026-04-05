package tree_test

import (
    "fmt"
    "testing"
    "math/rand"

    "github.com/petar/GoLLRB/llrb"
)


// Benchmarks for some functions of the llrb package for comparing them with the
// corresponding functions of the tree package.  No functions from the tree package
// are used in this file.


// Elements of the llrb trees.
type llrbElem struct {
    key int
}

// Comparison operator for the elements by implementing the interface llrb.Item.
func (e llrbElem) Less(item llrb.Item) bool {
    f, ok := item.(llrbElem)
    if !ok {
        panic("no tree element")
    }
    return e.key < f.key
}


func llrbRandom(n int) *llrb.LLRB {
    t := llrb.New()
    for t.Len() < n {
        t.InsertNoReplace(llrbElem{rand.Int()})
    }
    return t
}


var llrbTree  *llrb.LLRB

func BenchmarkLlrb_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            llrbTree = llrbRandom(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the tree.
                for {
                    value = rand.Int()
                    if !llrbTree.Has(llrbElem{value}) {
                        break
                    }
                }
                b.StartTimer()
                llrbTree.InsertNoReplace(llrbElem{value})
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a tree of a fixed size.
                llrbTree.Delete(llrbElem{value})
            }
        })
    }
}

func BenchmarkLlrb_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            llrbTree = llrbRandom(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the
                // deletion of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if !llrbTree.Has(llrbElem{value}) {
                        llrbTree.InsertNoReplace(llrbElem{value})
                        break
                    }
                }
                b.StartTimer()
                llrbTree.Delete(llrbElem{value})
                b.StopTimer()
            }
        })
    }
}

func BenchmarkLlrb_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            llrbTree = llrbRandom(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if !llrbTree.Has(llrbElem{value}) {
                        llrbTree.InsertNoReplace(llrbElem{value})
                        break
                    }
                }
                b.StartTimer()
                item := llrbTree.Get(llrbElem{value})
                e, ok := item.(llrbElem)
                b.StopTimer()
                if !ok || e.key != value {
                    panic("wrong tree element")
                }
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a tree of a fixed size.
                llrbTree.Delete(llrbElem{value})
            }
        })
    }
}
