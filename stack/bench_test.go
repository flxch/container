package stack_test

import (
    "testing"
    "github.com/flxch/container/stack"
)


const size = 1024

func BenchmarkPush(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    elems := make([]int, size)
    for i := range elems {
        elems[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := stack.New[int]()
        //s.Grow(size)
        b.StartTimer()
        for _, v := range elems {
            s.Push(v)
        }
        b.StopTimer()
        s.Reset()
    }
}

func BenchmarkPop(b *testing.B) {
    b.StopTimer()
    b.ReportAllocs()
    elems := make([]int, size)
    for i := range elems {
        elems[i] = i
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := stack.New[int]()
        //s.Grow(size)
        for _, v := range elems {
            s.Push(v)
        }
        b.StartTimer()
        for n := 0; n < size; n++ {
            s.Pop()
        }
        b.StopTimer()
        s.Reset()
    }
}
