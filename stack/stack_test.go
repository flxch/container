package stack_test

import (
    "testing"
    "github.com/flxch/container/stack"
)


func TestStack(t *testing.T) {
    s := stack.New[int]()
    if s.Num() != 0 {
        t.Errorf("stack should be empty")
    }

    s.Push(1)
    if s.Num() != 1 {
        t.Errorf("length should be 1")
    }

    if v, ok := s.Peek(); !ok {
        t.Errorf("failed to peek element")
    } else if v != 1 {
        t.Errorf("top element should be 1")
    }
    if v, ok := s.Pop(); !ok {
        t.Errorf("failed to pop element")
    } else if v != 1 {
        t.Errorf("top element should be 1")
    }
    if s.Num() != 0 {
        t.Errorf("stack should be empty")
    }

    s.Push(1)
    s.Push(2)
    if s.Num() != 2 {
        t.Errorf("length should be 2")
    }
    if v, ok := s.Peek(); !ok || v != 2 {
        t.Errorf("top element should be 2")
    }

    s.Reset()
    if s.Num() != 0 {
        t.Errorf("stack should be 0")
    }
    s.Push(1)
    if s.Num() != 1 {
        t.Errorf("length should be 1")
    }
}

func TestPush(t *testing.T) {
    s := stack.New[int]()
    for i := 0; i < 100; i++ {
        s.Push(i)
        if s.Num() != i + 1 {
            t.Errorf("length should be %d", i + 1)
        }
        if v, ok := s.Peek(); !ok || v != i {
            t.Errorf("top element should be %d", i)
        }
    }
}

func TestPop(t *testing.T) {
    s := stack.New[int]()
    for i := 0; i < 100; i++ {
        s.Push(i)
    }
    for i := 99; i >= 0; i-- {
        v, ok := s.Pop()
        if !ok {
            t.Errorf("failed to pop element")
        } else if v != i {
            t.Errorf("wrong value %d", v)
        }
        if s.Num() != i {
            t.Errorf("length should be %d", i)
        }
    }
    if _, ok := s.Pop(); ok {
        t.Errorf("pop should fail")
    }
}
