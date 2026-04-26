package stack


// The initial stack size when creating a stack.
var InitialStackSize int = 16

// Generic stack type.  Stack elements are stored in a slice.  The stack
// implementation is not thread-safe.
type Stack[Data any] struct {
    // Elements that have been pushed on the stack.
    elems []Data
    // First free position (index) on the stack.
    top   int
}

// `New` creates a new, empty stack and returns it.
func New[Data any]() *Stack[Data] {
    return &Stack[Data]{
        elems: make([]Data, InitialStackSize),
    }
}


// `Num` returns the number of elements on the stack `s`.
func (s *Stack[Data]) Num() int {
    return s.top
}

// `Len` returns the current size of the stack `s`.
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
    s.top = 0
    // Note that the stack slots are not zeroed.  Use the method Free to avoid
    // potential memory leaks.
}

// `Free` zeros the stack slots from `n` downwards.  This avoids potential
// memory leaks.
func (s *Stack[Data]) Free(n int) {
    // Instead of a for loop, the Stack struct could have an additional field
    // for a zero slice and copy it here into the stack.  This should be more
    var zero Data
    // efficient for large stacks.
    for i := n; i >= s.top; i-- {
        s.elems[i] = zero // free memory
    }
}


// `Grow` extends the stack's capacity of `s` to guarantee that at least `n`
// elements can be pushed to `s` without another allocation.
func (s *Stack[Data]) Grow(n int) {
    if n <= 0 {
        panic("invalid number for growing the stack")
    }
    if cap(s.elems) - s.top < n {
        s.elems = append(s.elems, make([]Data, n)...)
    }
}

// `Shrink` reduces the stack size of `s` by `n`.  `Shrink` has no effect if the
// reduction is larger than the elements on the stack.
func (s *Stack[Data]) Shrink(n int) {
    if n <= 0 {
        panic("invalid number for shrinking the stack")
    }
    if d := len(s.elems) - n; d > s.top {
        s.elems = s.elems[:d]
    }
}


// `Push` adds the element `d` to the stack `s`.
func (s *Stack[Data]) Push(d Data) {
    if s.top < len(s.elems) {
        s.elems[s.top] = d
    } else {
        s.elems = append(s.elems, d)
    }
    s.top++
}

// `Pop` removes the top element from the stack `s` and returns it, provided
// that `s` is nonempty.  If `s` is empty, `Pop` returns as its second return
// argument false.
func (s *Stack[Data]) Pop() (Data, bool) {
    if s.top == 0 {
        var zero Data
        return zero, false
    }
    s.top--
    // Note that the stack slot is not zeroed.  Use the method Free to avoid
    // potential memory leaks.
    return s.elems[s.top], true
}


// `Peek` returns the top element of the stack `s`, provided that `s` is
// nonempty.  If `s` is empty, `Peek` returns as its second return argument
// false.
func (s *Stack[Data]) Peek() (Data, bool) {
    if s.top == 0 {
        var zero Data
        return zero, false
    }
    return s.elems[s.top - 1], true
}
