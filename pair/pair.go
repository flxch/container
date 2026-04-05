package pair

import (
    "cmp"
)


// The pair type comprises two coordinates.
type Pair[T, U any] struct {
    fst T
    snd U
}


// `New` returns a pair.
func New[T, U any](x T, y U) Pair[T, U] {
    return Pair[T, U]{fst: x, snd: y}
}


// `Fst` returns the first coordinate of the pair `p`.
func (p Pair[T, U]) Fst() T {
    return p.fst
}

// `Snd` returns the second coordinate of the pair `p`.
func (p Pair[T, U]) Snd() U {
    return p.snd
}


// `Swap` returns the pair in which the coordinates of the pair `p` are swapped.
func (p Pair[T, U]) Swap() Pair[U, T] {
    return Pair[U, T]{fst: p.snd, snd: p.fst}
}


// `Compare` orders pairs lexicographically, i.e, `Compare` returns 0 if `p` and
// `q` are equal, 1 if `p` is greater than `q`, and -1 if `p` is smaller than
// `q`.
func Compare[T, U cmp.Ordered](p, q Pair[T, U]) int {
    if c := cmp.Compare(p.Fst(), q.Fst()); c != 0 {
        return c
    }
    return cmp.Compare(p.Snd(), q.Snd())
}
