package container_test

import (
    "fmt"
    "math/rand/v2"
    "testing"
    "github.com/MauriceGit/skiplist"
)


type mauriceElement int

func (e mauriceElement) ExtractKey() float64 {
    return float64(e)
}

func (e mauriceElement) String() string {
    return fmt.Sprintf("%03d", e)
}


func mauriceRandom(n int) skiplist.SkipList {
    sl := skiplist.New()
    i := 0
    for i < n {
        e := mauriceElement(rand.Int())
        if _, ok := sl.Find(e); !ok {
            sl.Insert(e)
            i++
        }
    }
    return sl
}


var mauriceSkiplist skiplist.SkipList

func BenchmarkMauriceSkiplist_Lookup(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            // Generate a random skip list containing `size` values.
            mauriceSkiplist = mauriceRandom(size)
            for i := 0; i < b.N; i++ {
                // Insert the value first, since we want to benchmark the lookup
                // of an existing value.
                var value int
                for {
                    value = rand.Int()
                    if _, ok := mauriceSkiplist.Find(mauriceElement(value)); !ok {
                        mauriceSkiplist.Insert(mauriceElement(value))
                        break
                    }
                }
                b.StartTimer()
                _, ok := mauriceSkiplist.Find(mauriceElement(value))
                b.StopTimer()
                if !ok {
                    panic("wrong skip list element")
                }
                // Remove the value again, since we want to benchmark the lookup
                // of a value for a skip list of a fixed size.
                mauriceSkiplist.Delete(mauriceElement(value))
            }
        })
    }
}


func BenchmarkMauriceSkiplist_Add(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            mauriceSkiplist = mauriceRandom(size)
            for i := 0; i < b.N; i++ {
                var value int
                for {
                    value = rand.Int()
                    if _, ok := mauriceSkiplist.Find(mauriceElement(value)); !ok {
                        break
                    }
                }
                b.StartTimer()
                mauriceSkiplist.Insert(mauriceElement(value))
                b.StopTimer()
                mauriceSkiplist.Delete(mauriceElement(value))
            }
        })
    }
}

func BenchmarkMauriceSkiplist_Remove(b *testing.B) {
    for _, size := range sizes {
        b.Run(fmt.Sprintf("%d", size), func(b *testing.B) {
            b.StopTimer()
            b.ReportAllocs()
            mauriceSkiplist = mauriceRandom(size)
            for i := 0; i < b.N; i++ {
                var value int
                for {
                    value = rand.Int()
                    if _, ok := mauriceSkiplist.Find(mauriceElement(value)); !ok {
                        mauriceSkiplist.Insert(mauriceElement(value))
                        break
                    }
                }
                b.StartTimer()
                mauriceSkiplist.Delete(mauriceElement(value))
                b.StopTimer()
            }
        })
    }
}
