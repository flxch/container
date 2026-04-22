package stack


// Generic stack type.  Stack elements are stored in a slice.  The stack
// implementation is not thread-safe.
type Stack[Data any] struct {
    // Elements that have been pushed on the stack.
    elems []Data
    // First free position (index) on the stack.
    top   int
    // Zero element (when popping stack elements) to free memory.
    zero  Data
}


// `New` creates a new, empty stack and returns it.
func New[Data any]() *Stack[Data] {
    return &Stack[Data]{
        elems: make([]Data, 64),
    }
}


// `Len` returns the number of elements on the stack `s`.
func (s *Stack[Data]) Num() int {
    return s.top
}

// `Size` returns the current size of the stack `s`.
func (s *Stack[Data]) Len() int {
    return len(s.elems)
}

// `Cap` returns the number of elements that the stack `s` can store
// without allocating new memory.
func (s *Stack[Data]) Cap() int {
    return cap(s.elems)
}


// `Reset` empties the stack `s`.
func (s *Stack[Data]) Reset() {
    for i := 0; i < s.top; i++ {
        s.elems[i] = s.zero // free memory
    }
    s.top = 0
}


// `Grow` extends the stack's capacity of `s` to guarantee that at least `n`
// elements can be pushed to `s` without another allocation.
func (s *Stack[Data]) Grow(n int) {
    if n <= 0 {
        panic("stack can only grow positively")
    }
    if cap(s.elems) - s.top < n {
        s.elems = append(s.elems, make([]Data, n)...)
    }
}

// `Shrink` reduces the stack size of `s` by `n`.  `Shrink` has no effect if the
// reduction is larger than the elements on the stack.
func (s *Stack[Data]) Shrink(n int) {
    if n <= 0 {
        panic("stack can only shrink positively")
    }
    if d := len(s.elems) - n; d > s.top {
        s.elems = s.elems[:d]
    }
}


// `Push` adds the element `d` to the stack `s`.
func (s *Stack[Data]) Push(d Data) {
    if s.top >= len(s.elems) {
        s.elems = append(s.elems, d)
    } else {
        s.elems[s.top] = d
    }
    s.top++
}

// `Pop` removes the top element from the stack `s` and returns it, provided
// that `s` is nonempty.  If `s` is empty, `Pop` returns as its second return
// argument false.
func (s *Stack[Data]) Pop() (Data, bool) {
    if s.top == 0 {
        return s.zero, false
    }
    s.top--
    r := s.elems[s.top]
    s.elems[s.top] = s.zero // free memory
    return r, true
}


// `Peek` returns the top element of the stack `s`, provided that `s` is
// nonempty.  If `s` is empty, `Peek` returns as its second return argument
// false.
func (s *Stack[Data]) Peek() (Data, bool) {
    if s.top == 0 {
        return s.zero, false
    }
    return s.elems[s.top - 1], true
}
