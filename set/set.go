package set

import (
    "maps"
)


// (Finite) Set are implemented by maps.
type Set[A comparable] map[A]struct{}


// `New` returns a new empty set with elements of type `A`.
func New[A comparable]() Set[A] {
    return make(map[A]struct{})
}


// `Len` returns the size of the set `S`.
func (S Set[A]) Len() int {
    return len(S)
}

// `Reset` empties the set `S`.
func (S Set[A]) Reset() {
    clear(S)
}

// `Clone` returns a clone of the set `S`.
func (S Set[A]) Clone() Set[A] {
    return maps.Clone(S)
}


// `Lookup` returns true if the element `e` is in the set `S`.
func (S Set[A]) Lookup(e A) bool {
    _, ok := S[e]
    return ok
}

// `Add` adds the element `e` to the set `S`.
func (S Set[A]) Add(e A) {
    S[e] = struct{}{}
}

// `Remove` removes the element `e` from the set `S`.
func (S Set[A]) Remove(e A) {
    delete(S, e)
}


// `Union` makes `S` the union of the sets `S` and `T`.
func (S Set[A]) Union(T Set[A]) {
    for e := range T {
        S[e] = struct{}{}
    }
}

// `Union` returns the union of the sets `S` and `T`.
func Union[A comparable](S, T Set[A]) Set[A] {
    R := S.Clone()
    R.Union(T)
    return R
}

// `Intersection` intersects the set `T` with the set `S`.
func (S Set[A]) Intersection(T Set[A]) {
    for e := range T {
        if !S.Lookup(e) {
            delete(S, e)
        }
    }
}

// `Intersection` returns the intersection of the sets `S` and `T`.
func Intersection[A comparable](S, T Set[A]) Set[A] {
    R := S.Clone()
    R.Intersection(T)
    return R
}

// `Difference` subtracts the set `T` from the set `S`.
func (S Set[A]) Difference(T Set[A]) {
    for e := range T {
        delete(S, e)
    }
}

// `Difference` returns the difference of the set `S` from `T`.
func Difference[A comparable](S, T Set[A]) Set[A] {
    R := S.Clone()
    R.Difference(T)
    return R
}


// `Equal` returns true if the set `S` is equal to the set `T`.
func (S Set[A]) Equal(T Set[A]) bool {
    return maps.Equal(S, T)
}

// `Equal` returns true if the sets `S` and `T` are equal.
func Equal[A comparable](S, T Set[A]) bool {
    return S.Equal(T)
}

// `Subset` returns true if the Set `S` is a subset of the set `T`.
func (S Set[A]) Subset(T Set[A]) bool {
    if len(S) > len(T) {
        return false
    }

    for e := range S {
        if !T.Lookup(e) {
            return false
        }
    }
    return true
}

// `Superset` returns true if the set `S` is a superset of the set `T`.
func (S Set[A]) Superset(T Set[A]) bool {
    return T.Subset(S)
}


