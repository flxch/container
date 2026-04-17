package multiset

import (
    "iter"
)


// `Elems` returns an iterator for iterating through the elements with their
// multiplicites in the multiset `S`.  There is guarantee on the order of the
// elements in `S`.
func (S *Multiset[A]) Elems() iter.Seq2[A, int] {
    return func(yield func(A, int) bool) {
        for e, n := range S.elems {
            if !yield(e, n) {
                return
            }
        }
    }
}

// `Sup` returns an iterator for iterating through the support of the
// multiset `S`.  There is guarantee on the support order.
func (S *Multiset[A]) Sup() iter.Seq[A] {
    return func(yield func(A) bool) {
        for e := range S.elems {
            if !yield(e) {
                return
            }
        }
    }
}
