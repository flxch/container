package container_test

import (
    "fmt"
    "math/rand/v2"
    "testing"
    "github.com/fogfish/skiplist"
)


func fogfishRandom(n int) *skiplist.Set[int] {
    sl := skiplist.NewSet[int]()
    for sl.Length() < n {
        e := rand.Int()
        if ok, _ := sl.Has(e); !ok {
            sl.Add(e)
        }
    }
    return sl
}


var fogfishSkiplist *skiplist.Set[int]

func BenchmarkFogfishSkiplist_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random skip list containing `size` values.
            fogfishSkiplist = fogfishRandom(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if ok, _ := fogfishSkiplist.Has(value); !ok {
                        fogfishSkiplist.Add(value)
                        break
                    }
                }
                b.StartTimer()
                ok, _ := fogfishSkiplist.Has(value)
                b.StopTimer()
                if !ok {
                    panic("wrong skip list element")
                }
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a skip list of a fixed size.
                fogfishSkiplist.Cut(value)
            }
        })
    }
}

func BenchmarkFogfishSkiplist_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            fogfishSkiplist = fogfishRandom(size)
            for i := 0; i < b.N; i++ {
                var value int
                for {
                    value = rand.Int()
                    if ok, _ := fogfishSkiplist.Has(value); !ok {
                        break
                    }
                }
                b.StartTimer()
                fogfishSkiplist.Add(value)
                b.StopTimer()
                fogfishSkiplist.Cut(value)
            }
        })
    }
}

func BenchmarkFogfishSkiplist_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            fogfishSkiplist = fogfishRandom(size)
            for i := 0; i < b.N; i++ {
                var value int
                for {
                    value = rand.Int()
                    if ok, _ := fogfishSkiplist.Has(value); !ok {
                        fogfishSkiplist.Add(value)
                        break
                    }
                }
                b.StartTimer()
                fogfishSkiplist.Cut(value)
                b.StopTimer()
            }
        })
    }
}
