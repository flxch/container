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

// `Hop` returns the element that is `d` hops away from the element `e`.  That
// is, if `d` is negative, it is before `e` and if `d` is positive, it is after
// `e`.  `Hop` returns nil if the beginning or the end of the skip list has been
// reached.
func (e *Element[Data]) Hop(d int) *Element[Data] {
    if d < 0 {
        for i := d; i < 0 && e != nil; i, e = i + 1, e.Prev() { }
    } else if d > 0 {
        for i := 0; i < d && e != nil; i, e = i + 1, e.Next() { }
    }
    return e
}


// `InsertBefore` inserts the value `val` as a new element right before the
// element `e` in the skip list of `e`.  It is not checked whether the values
// are correctly ordered.
func (e *Element[Data]) InsertBefore(val Data) {
    e = e.list.root.neighbors[0].prev
    e.insert(&Element[Data]{
        list:      e.list,
        neighbors: make([]neighbors[Data], 1),
        Value:     val,
    })
}

// `InsertBefore` inserts the value `val` as a new element right after the
// element `e` in the skip list of `e`.  It is not checked whether the values
// are correctly ordered.
func (e *Element[Data]) InsertAfter(val Data) {
    e.insert(&Element[Data]{
        list:      e.list,
        neighbors: make([]neighbors[Data], 1),
        Value:     val,
    })
}

func (e *Element[Data]) insert(f *Element[Data]) {
    // Update skip list data.
    e.list.len++
    e.list.heights[0]++
    e.list.max = max(e.list.max, 1)
    // Set links.
    t := e.neighbors[0].next
    e.neighbors[0].next = f
    f.neighbors[0].prev = e
    t.neighbors[0].prev = f
    f.neighbors[0].next = t
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
