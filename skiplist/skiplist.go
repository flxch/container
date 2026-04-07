package skiplist

import (
    "math/bits"
    "math/rand/v2"
    "slices"
)


const (
    // Internally, the elements in a skip list have a height, which is randomly
    // chosen for each element.  For small skip lists, small heights works best
    // (performance and memory).  For large skip lists, a too small maximal
    // height will most likely result in performance reductions.
    // `delfaultMaxHeight` seems to work well for skip lists with up to
    // 1,000,000 elements.  If the skip list has significantly more than
    // 1,000,000 elements, a larger the maximal height might work better.  Use
    // the function `SetHeight` to change a skip list's maximal element height.
    defaultMaxHeight = 32

    // As the height of internal elements are randomly chosen, one can set the
    // seed for generating the random height values.  Use the function `SetSeed`
    // to change a skip list's seeds.
    defaultSeed1     = 1
    defaultSeed2     = 2
)


// A variable of the type `Skiplist` stores data elements of a type in a skip
// list.  The skip list's type can be any type but elements must be ordered.
// The order of elements is determined by a compare function.
type Skiplist[Data any] struct {
    // Comparison operator on Data for ordering the elements in the skip list.
    cmp     func(Data, Data) int
    // Root of the elements in the skip list (front and back).
    root    Element[Data]
    // Number of elements in the skip list.
    len     int
    // Height count of the elements in the skip list
    heights []int
    // Current maximal height of an element in the skip list.
    max     int
    // Random number generator for setting the height of an element.
    rand    *rand.Rand
}


// `New` returns an empty skip list.  List elements are ordered by the function
// `cmp`.
func New[Data any] (cmp func(Data, Data) int) *Skiplist[Data] {
    return new(Skiplist[Data]).Init(cmp)
}

// `Init` initializes or clears the skip list `l`.  List elements are ordered by
// the function `cmp`.
func (l *Skiplist[Data]) Init(cmp func(Data, Data) int) *Skiplist[Data] {
    l.cmp = cmp
    l.heights = make([]int, max(defaultMaxHeight, len(l.heights)))
    l.root.neighbors = make([]neighbors[Data], l.Height())
    for k := 0; k < l.Height(); k++ {
        l.root.neighbors[k].next = &l.root
        l.root.neighbors[k].prev = &l.root
    }
    l.len = 0
    l.max = 0
    if l.rand == nil {
        l.rand = rand.New(rand.NewPCG(defaultSeed1, defaultSeed2))
    }
    return l
}

// `Reset` removes all elements from the skip list `l`.
// NOTE: Elements still point to l. We could run trough them and set the owner
// to nil. However, this would be an expensive operation if the skip list is
// long.
func (l *Skiplist[Data]) Reset() {
    l.len = 0
    l.max = 0
    for k := 0; k < l.Height(); k++ {
        l.root.neighbors[k].next = &l.root
        l.root.neighbors[k].prev = &l.root
        l.heights[k] = 0
    }
}


// `Clone` returns a clone of the skip list `l`.
func (l *Skiplist[Data]) Clone(clone func(Data) Data) *Skiplist[Data] {
    c := &Skiplist[Data]{
        cmp:     l.cmp,
        root:    Element[Data]{
            neighbors: make([]neighbors[Data], len(l.root.neighbors), cap(l.root.neighbors)),
        },
        len:     l.Len(),
        heights: slices.Clone(l.heights),
        max:     l.max,
        rand:    l.rand, //rand.New(rand.NewPCG(defaultSeed1 + 1, defaultSeed2 + 1)),
    }
    // QUESTION: Should we use the same random generator or should we create a
    // new one.  If we create a new one, we should use different seeds,
    // otherwise we obtain the same sequence of numbers again for c.

    p := make([]*Element[Data], l.Height())
    for k := 0; k < l.Height(); k++ {
        p[k] = &c.root
        c.root.neighbors[k].next = &c.root
        c.root.neighbors[k].prev = &c.root
    }

    e := l.root.neighbors[0].next
    for i := 0; i < l.Len(); i++ {
        f := e.clone(c, clone)
        for k := 0; k < len(e.neighbors); k++ {
            f.neighbors[k].prev = p[k]
            p[k].neighbors[k].next = f
            p[k] = f
            if e.neighbors[k].next == &l.root {
                f.neighbors[k].next = &c.root
                c.root.neighbors[k].prev = f
            }
        }
        e = e.neighbors[0].next
    }

    return c
}


// `Front` returns the first element of the skip list `l` or nil if `l` is
// empty.
func (l *Skiplist[Data]) Front() *Element[Data] {
    if l.Len() == 0 {
        return nil
    }
    return l.root.neighbors[0].next
}

// `Back` returns the last element of the skip list `l` or nil if `l` is empty.
func (l *Skiplist[Data]) Back() *Element[Data] {
    if l.Len() == 0 {
        return nil
    }
    return l.root.neighbors[0].prev
}


// `Len` returns the number of elements in the skip list `l`.
func (l *Skiplist[Data]) Len() int {
    return l.len
}

// `Height` returns the maximal height of elements in the skip list `l`.
func (l *Skiplist[Data]) Height() int {
    return len(l.heights)
}

// `SetHeight` changes the maximal height of elements in the skip list `l`.
func (l *Skiplist[Data]) SetHeight(h int) {
    if h >= 1 && h <= 64 {
        d := h - len(l.heights)
        switch {
        case d < 0: // Cut elements with height above h.
            e := l.root.neighbors[h].next
            for e != &l.root {
                f := e.neighbors[h].next
                e.neighbors = e.neighbors[0:h]
                e = f
            }
            l.root.neighbors = l.root.neighbors[0:h]
            l.heights = l.heights[0:h]
            if l.max > h {
                l.max = h
            }
        case d > 0: // Extend root to height h.
            l.root.neighbors = append(l.root.neighbors, make([]neighbors[Data], d)...)
            for k := h - 1; k >= len(l.heights); k-- {
                l.root.neighbors[k].next = &l.root
                l.root.neighbors[k].prev = &l.root
            }
            l.heights = append(l.heights, make([]int, d)...)
        }
    }
}

// `SetSeed` changes the seed for the pseudo-random generator for the height of
// new elements to `s`.
func (l *Skiplist[Data]) SetSeed(s, t uint64) {
    l.rand = rand.New(rand.NewPCG(s, t))
}


// `Lookup` returns the element of the skip list `l` with the value `val`.
// `Lookup` returns nil, if there is no element in `l` with `val`.
func (l *Skiplist[Data]) Lookup(val Data) *Element[Data] {
    p, q := &l.root, &l.root
    for h := l.max - 1; h >= 0; h-- {
        for q = p.neighbors[h].next; q != &l.root && l.cmp(q.Value, val) < 0; p, q = q, q.neighbors[h].next { }
    }
    if q == &l.root || l.cmp(val, q.Value) < 0 {
        return nil
    }
    return q
}


// `Add` returns the element in the skip list `l` with value `val`.  If there is
// no such element, then `Add` inserts it into `l`.  If `l` already contains an
// element of the same order as `val`, 'val` is still inserted to the skip list.
func (l *Skiplist[Data]) Add(val Data) *Element[Data] {
    e := l.find(val)
    if e.list == nil {
        // Complete new element and make it an element of the skip list.
        e.Value = val
        for h := e.height() - 1; h >= 0; h-- {
            e.neighbors[h].next = e.neighbors[h].prev.neighbors[h].next
            e.neighbors[h].next.neighbors[h].prev, e.neighbors[h].prev.neighbors[h].next = e, e
        }
        e.list = l
        l.len++
        l.heights[e.height() - 1]++
        //if e.height() > l.max && l.max < len(l.heights) {
        //    l.max = e.height()
        //}
    } else {
        // QUESTION: Does it make sense to store the additional value in e? That is, e has a list
        // elements. Note that if we would add to the skip list, we would have a sub skip list of
        // equal elements.  Elements would be a container again. Or, we count how often the
        // element was inserted.  Maybe want to also change the tree package.
        // ... TODO
        // We could add a new field to the Element type Misc of type any.

        // This new field could instantiated with a counter or with a slice if
        // we want to store vals with the same order.
        // Note that we return Element and the caller could instantiate Misc appropriately to its use case.

    }
    return e
}

// `Insert` returns the element in the skip list `l` with value `val`.  If there
// is no such element, then `Insert` inserts it into `l`.  If `l` already
// contains an element of the same order as `val`, the element is not changed.
func (l *Skiplist[Data]) Insert(val Data) *Element[Data] {
    e := l.find(val)
    if e.list == nil {
        // New value.  Make it an element of the skip list.
        e.Value = val
        for h := e.height() - 1; h >= 0; h-- {
            e.neighbors[h].next = e.neighbors[h].prev.neighbors[h].next
            e.neighbors[h].next.neighbors[h].prev, e.neighbors[h].prev.neighbors[h].next = e, e
        }
        e.list = l
        l.len++
        l.heights[e.height() - 1]++
        //if e.height() > l.max && l.max < len(l.heights) {
        //    l.max = e.height()
        //}
    }
    return e
}


// `Remove` removes the element `e` from the skip list `l`.  It returns `e`'s
// value.  `e` must not be nil.
func (l *Skiplist[Data]) Remove(e *Element[Data]) Data {
    if e.list == l {
        l.heights[e.height() - 1]--
        if e.height() == l.max && l.heights[e.height() - 1] == 0 {
            l.max--
        }
        l.len--
        for h := e.height() - 1; h >= 0; h-- {
            e.neighbors[h].prev.neighbors[h].next = e.neighbors[h].next
            e.neighbors[h].next.neighbors[h].prev = e.neighbors[h].prev
        }
        e.neighbors = nil // avoid memory leaks
        e.list = nil
    }
    return e.Value
}


// Auxiliary functions.

// `find` returns an element with value `val`.  If such an element does exist in
// the skip list `l`, `find` creates such an element with information about
// elements in `l` for inserting the element.  Otherwise, the respective element
// of `l` is returned.  `find` is used by the `Add` function.
func (l *Skiplist[Data]) find(val Data) *Element[Data] {
    before := make([]neighbors[Data], l.newElementHeight())
    p, q := &l.root, &l.root
    for h := l.max - 1; h >= len(before); h-- {
        for q = p.neighbors[h].next; q != &l.root && l.cmp(q.Value, val) < 0; p, q = q, q.neighbors[h].next { }
    }
    for h := len(before) - 1; h >= 0; h-- {
        for q = p.neighbors[h].next; q != &l.root && l.cmp(q.Value, val) < 0; p, q = q, q.neighbors[h].next { }
        before[h].prev = p
    }
    if q == &l.root || l.cmp(val, q.Value) < 0 {
        return &Element[Data]{neighbors: before, Value: val}
    }
    return q
}


// `newElementHeight` returns the height for a new element in the skip list `l`.
func (l *Skiplist[Data]) newElementHeight() int {
    if h := bits.TrailingZeros64(l.rand.Uint64() & ((1 << uint(l.Height())) - 1)); h < l.max {
        return 1 + h
    }
    if l.max == 0 || l.max < len(l.heights) && l.heights[l.max - 1] > 0 {
        l.max++
    }
    return 1
}
