package container_test

import (
    "fmt"
    "testing"
    "math/rand/v2"
    "github.com/mikenye/gotrees/rbtree"
)


// Benchmarks for some functions of the llrb package for comparing them with the
// corresponding functions of the tree package.


// Random tree with n elements.
func rbtreeRandom(n int) *rbtree.Tree[int, struct{}] {
    t := rbtree.New[int, struct{}](func(a, b int) bool { return a < b })
    for t.Size() < n {
        t.Insert(rand.Int(), struct{}{})
    }
    return t
}


var rbtreeTree  *rbtree.Tree[int, struct{}]

func BenchmarkRbtree_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            rbtreeTree = rbtreeRandom(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if _, ok := rbtreeTree.Search(value); !ok {
                        rbtreeTree.Insert(value, struct{}{})
                        break
                    }
                }
                b.StartTimer()
                node, ok := rbtreeTree.Search(value)
                b.StopTimer()
                if !ok || node == nil {
                    panic("wrong tree element")
                }
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a tree of a fixed size.
                rbtreeTree.Delete(node)
            }
        })
    }
}

func BenchmarkRbtree_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            rbtreeTree = rbtreeRandom(size)
            for i := 0; i < b.N; i++ {
                var value int
                // Find a value that is not in the tree.
                for {
                    value = rand.Int()
                    if _, ok := rbtreeTree.Search(value); !ok {
                        break
                    }
                }
                b.StartTimer()
                rbtreeTree.Insert(value, struct{}{})
                b.StopTimer()
                // Remove the value again, since we want to benchmark the
                // insertion of a value into a tree of a fixed size.
                node, ok := rbtreeTree.Search(value)
                if !ok || node == nil {
                    panic("wrong tree element")
                }
                rbtreeTree.Delete(node)
            }
        })
    }
}

func BenchmarkRbtree_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random tree containing `size` values.
            rbtreeTree = rbtreeRandom(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the
                // deletion of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if _, ok := rbtreeTree.Search(value); !ok {
                        rbtreeTree.Insert(value, struct{}{})
                        break
                    }
                }
                b.StartTimer()
                if node, ok := rbtreeTree.Search(value); ok && node != nil {
                    rbtreeTree.Delete(node)
                } else {
                    panic("wrong tree element")
                }
                b.StopTimer()
            }
        })
    }
}

