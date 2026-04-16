package multiset

import (
    "iter"
)


func (S *Multiset[A]) All() iter.Seq2[A, int] {
    return func(yield func(A, int) bool) {
        for e, n := range S.elems {
            if !yield(e, n) {
                return
            }
        }
    }
}

func (S *Multiset[A]) Elems() iter.Seq[A] {
    return func(yield func(A) bool) {
        for e := range S.elems {
            if !yield(e) {
                return
            }
        }
    }
}
