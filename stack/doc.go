package stack

// A simple, non-thread-safe, generic stack implementation that uses slices.

// Stacks are created by the New() function.  The stack operations are Push()
// and Pop().  Peek() returns the top element without removing it from the
// stack.

// Len() returns the number of the elements on the stack and Reset() empties the
// stack.  Other operations are Cap(), Shrink(), and Grow() for fine tuning
// memory management (whenever needed).

