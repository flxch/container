package multiset

import (
    "maps"
)


// A (finite) multiset is similar to a (finite) set but it additionally counts
// how often an element is contained in the multiset, i.e., the element's
// multiplicity.
type Multiset[A comparable] struct {
    // Elements with their multiplicites.
    elems map[A]int
    // Total number of elements in the multiset, i.e., sum of all multiplicites.
    count int
}


// `New` returns a new empty multiset with elements of type `A`.  Since the
// implementation of multisets is based on maps, elements of `A` must be
// comparable.
func New[A comparable]() *Multiset[A] {
    return &Multiset[A]{
        elems: make(map[A]int),
        count: 0,
    }
}


// `Card` returns the cardinality of the multiset `S`, i.e., the sum of the
// elements' multiplicities.
func (S *Multiset[A]) Card() int {
    return S.count
}

// `Len` returns the size of the support of the multiset `S`.
func (S *Multiset[A]) Len() int {
    return len(S.elems)
}

// `Reset` empties the multiset `S`.
func (S *Multiset[A]) Reset() {
    clear(S.elems)
    S.count = 0
}

// `Clone` returns a clone of the multiset `S`.
func (S *Multiset[A]) Clone() *Multiset[A] {
    return &Multiset[A]{
        elems: maps.Clone(S.elems),
        count: S.count,
    }
}


// `Support` returns the support (as a slice) of the multiset `S`.
func (S *Multiset[A]) Support() []A {
    s := make([]A, 0, len(S.elems))
    for e := range S.elems {
        s = append(s, e)
    }
    return s
}


// `Lookup` returns the multiplicity of the element `e` in the multiset `S`.
// Note that if the multiplicity is 0, `S` does not contain `e`.
func (S *Multiset[A]) Lookup(e A) int {
    n, ok := S.elems[e]
    if !ok {
        return 0
    }
    return n
}

// `Add` adds the element `e` to the multiset `S` with multiplicity `m`.  Note
// that if `S` already contains `e`, `e`'s multiplicity in `S` increases by `m`.
func (S *Multiset[A]) Add(e A, m int) {
    S.elems[e] = m + S.Lookup(e)
    S.count += m
}

// `Remove` removes the element `e` from the multiset `S` with multiplicity `m`.
// That is, `e`'s multiplicity in `S` decreses by `m`, provided that `n > m`.
// Note that if `m >= n`, `e` is completely removed from `S`.
func (S Multiset[A]) Remove(e A, m int) {
    if n := S.Lookup(e); n <= m {
        delete(S.elems, e)
        S.count -= n
    } else if n > m {
        S.elems[e] = n - m
        S.count -= m
    }
}


// `Union` makes `S` the union of the multisets `S` and `T`..
func (S *Multiset[A]) Union(T *Multiset[A]) {
    for e, m := range T.elems {
        n := S.Lookup(e)
        S.elems[e] = max(m, n)
        if n > m {
            S.count += n - m
        }
    }
}

// `Union` returns the union of the multisets `S` and `T`.
func Union[A comparable](S, T *Multiset[A]) *Multiset[A] {
    R := S.Clone()
    R.Union(T)
    return R
}

// `Sum` adds the elements of the multiset `T` to the multiset `S`.
func (S *Multiset[A]) Sum(T *Multiset[A]) {
    for e, m := range T.elems {
        n := S.Lookup(e)
        S.elems[e] = m + n
        S.count += n
    }
}

// `Sum` returns the sum of the multisets `S` and `T`.
func Sum[A comparable](S, T *Multiset[A]) *Multiset[A] {
    R := S.Clone()
    R.Sum(T)
    return R
}

// `Intersection` intersects the multiset `T` with the multiset `S`.
func (S *Multiset[A]) Intersection(T *Multiset[A]) {
    for e, n := range S.elems {
        if m := T.Lookup(e); m == 0 {
            delete(S.elems, e)
            S.count -= n
        } else {
            S.elems[e] = min(n, m)
            if n > m {
                S.count -= m - n
            }
        }
    }
}

// `Intersection` returns the intersection of the multisets `S` and `T`.
func Intersection[A comparable](S, T *Multiset[A]) *Multiset[A] {
    R := S.Clone()
    R.Intersection(T)
    return R
}

// `Difference` subtracts the multiset `T` from the multiset `S`.
func (S *Multiset[A]) Difference(T *Multiset[A]) {
    for e, n := range S.elems {
        if m := T.Lookup(e); m > n {
            delete(S.elems, e)
            S.count -= n
        } else {
            S.elems[e] = n - m
            S.count -= n - m
        }
    }
}

// `Difference` returns the difference of the muliset `S` from `T`.
func Difference[A comparable](S, T *Multiset[A]) *Multiset[A] {
    R := S.Clone()
    R.Difference(T)
    return R
}


// `Equal` returns true if the multiset `S` is equal to the multiset `T`.
func (S *Multiset[A]) Equal(T *Multiset[A]) bool {
    return S.count == T.count && maps.Equal(S.elems, T.elems)
}

// `Equal` returns true if the mulitsets `S` and `T` are equal.
func Equal[A comparable](S, T *Multiset[A]) bool {
    return S.Equal(T)
}

// `Subset` returns true if the multiset `S` is a subset of the multiset `T`.
func (S *Multiset[A]) Subset(T *Multiset[A]) bool {
    if S.count > T.count || len(S.elems) > len(T.elems) {
        return false
    }

    for e, n := range S.elems {
        if m := T.Lookup(e); m > n {
            return false
        }
    }
    return true
}

// `Superset` returns true if the multiset `S` is a superset of the multiset `T`.
func (S *Multiset[A]) Superset(T *Multiset[A]) bool {
    return T.Subset(S)
}


