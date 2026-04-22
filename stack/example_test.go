package stack_test

import (
    "fmt"
    "github.com/flxch/container/stack"
)


func ExampleStack() {
    s := stack.New[int]()

    s.Push(1)
    s.Push(2)

    n, ok := s.Pop()
    if !ok {
        panic("tried to pop an element from the empty stack")
    }
    fmt.Println(n)

    s.Push(3)

    n, ok = s.Pop()
    if !ok {
        panic("tried to pop an element from the empty stack")
    }
    fmt.Println(n)

    n, ok = s.Pop()
    if !ok {
        panic("tried to pop an element from the empty stack")
    }
    fmt.Println(n)

    // Output:
    // 2
    // 3
    // 1
}
