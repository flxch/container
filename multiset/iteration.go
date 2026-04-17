package multiset

import (
    "iter"
)


// `Elements` returns an iterator for iterating through the elements of the
// multiset `S` with their multiplicites.  There is guarantee on the iteration
// order.
func (S *Multiset[A]) Elements() iter.Seq2[A, int] {
    return func(yield func(A, int) bool) {
        for e, n := range S.elems {
            if !yield(e, n) {
                return
            }
        }
    }
}

// `Support` returns an iterator for iterating through the support of the
// multiset `S`.  There is guarantee on the iteration order.
func (S *Multiset[A]) Support() iter.Seq[A] {
    return func(yield func(A) bool) {
        for e := range S.elems {
            if !yield(e) {
                return
            }
        }
    }
}
