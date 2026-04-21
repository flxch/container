package multiset


// `Elements` is an iterator for iterating through the elements of the multiset
// `S` with their multiplicites.  There is no guarantee on the iteration order.
func (S *Multiset[A]) Elements(yield func(A, int) bool) {
    for e, n := range S.elems {
        if !yield(e, n) {
            return
        }
    }
}

// `Support` is an iterator for iterating through the support of the multiset
// `S`.  There is no guarantee on the iteration order.
func (S *Multiset[A]) Support(yield func(A) bool) {
    for e := range S.elems {
        if !yield(e) {
            return
        }
    }
}
