package skiplist


// `neighbors` stores the pointers to the next and the previous elements in the
// skip list for an element for each level.
type neighbors[Data any] struct {
    next *Element[Data]
    prev *Element[Data]
}

// `Element` is a wrapper around the elements in a skip list.
type Element[Data any] struct {
    // The list to which the element belongs.
    list      *Skiplist[Data]
    // The next and previous pointers in the doubly-linked skip list for each
    // height.
    neighbors []neighbors[Data]
    // The value of the element.
    Value     Data
}


// `Next` returns `e`'s next element.  The return value is nil, if `e` is nil,
// not an element of a skip list, or the last element of a skip list.
func (e *Element[Data]) Next() *Element[Data] {
    if e == nil || e.list == nil || &e.list.root == e.neighbors[0].next {
        return nil
    }
    return e.neighbors[0].next
}

// `Prev` returns `e`'s previous element.  The return value is nil, if `e` is
// nil, not an element of a skip list, or the first element of a skip list.
func (e *Element[Data]) Prev() *Element[Data] {
    if e == nil || e.list == nil || &e.list.root == e.neighbors[0].prev {
        return nil
    }
    return e.neighbors[0].prev
}


// `height` returns the height of the element `e`.
func (e *Element[Data]) height() int {
    return len(e.neighbors)
}


// `clone` returns a clone of the element `e` and links it to the skip list `l`.
// For clining the element's value, it uses the function `clone`.
// NOTE: The next and previous pointers are not set, since this cannot be done
// properly locally here.
// (Helper function used by the Clone method for skip lists.)
func (e *Element[Data]) clone(l *Skiplist[Data], clone func(Data) Data, elems []Element[Data], ns []neighbors[Data]) (*Element[Data], []Element[Data], []neighbors[Data]) {
    if len(elems) == 0 || len(ns) == 0 {
        return &Element[Data]{
            list:      l,
            neighbors: make([]neighbors[Data], e.height(), e.height()),
            Value:     clone(e.Value),
        }, nil, nil
    }
    f := elems[0]
    f.list = l
    f.neighbors = ns[:e.height():e.height()]
    f.Value = clone(e.Value)
    return &f, elems[1:], ns[e.height():]
}
