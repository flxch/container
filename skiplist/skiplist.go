package skiplist

import (
    "math/bits"
    "math/rand/v2"
    "slices"
)


// A variable of the type `Skiplist` stores data elements of a type in a skip
// list.  The skip list's type can be any type but elements must be ordered.
// The order of elements is determined by a compare function.
type Skiplist[Data any] struct {
    // Comparison operator on Data for ordering the elements in the skip list.
    cmp      func(Data, Data) int
    // Root of the elements in the skip list (front and back).
    root     Element[Data]
    // Number of elements in the skip list.
    len      int
    // Height count of the elements in the skip list
    heights  []int
    // Maximal height an element can have in the skip list.
    max      int
    // Random number generator for setting the height of an element.
    rand     *rand.Rand
    // Preallocation of memory for some performance gains.  Higher memory
    // consumption though.
    prealloc int
    preelems []neighbors[Data]
}


// `New` returns an empty skip list.  List elements are ordered by the function
// `cmp`.
func New[Data any](cmp func(Data, Data) int) *Skiplist[Data] {
    return new(Skiplist[Data]).Init(cmp)
}

// `Init` initializes or clears the skip list `l`.  List elements are ordered by
// the function `cmp`.
func (l *Skiplist[Data]) Init(cmp func(Data, Data) int) *Skiplist[Data] {
    l.cmp = cmp
    l.heights = make([]int, max(DefaultMaxHeight, len(l.heights)))
    l.root.neighbors = make([]neighbors[Data], l.Height())
    for k := 0; k < l.Height(); k++ {
        l.root.neighbors[k].next = &l.root
        l.root.neighbors[k].prev = &l.root
    }
    l.len = 0
    l.max = 0
    if l.rand == nil {
        l.rand = rand.New(rand.NewPCG(DefaultSeed1, DefaultSeed2))
    }
    l.prealloc = DefaultPrealloc
    return l
}

// `Reset` removes all elements from the skip list `l`.  NOTE: Elements still
// point to l. We could run trough them and set the owner to nil.  However, this
// would be an expensive operation if the skip list is long.  Furthermore, the
// elements are linked to each other.  This might lead to a memory leak, as all
// elements are not garbage collected.
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
        rand:    l.rand, //rand.New(rand.NewPCG(DefaultSeed1 + 1, DefaultSeed2 + 1)),
    }
    // QUESTION: Should we use the same random generator or should we create a
    // new one.  If we create a new one, we should use different seeds,
    // otherwise we obtain the same sequence of numbers again for c.

    if l.prealloc > 0 {
        c.SetPrealloc(c.prealloc)
    }

    p := make([]*Element[Data], l.Height())
    for k := 0; k < l.Height(); k++ {
        p[k] = &c.root
        c.root.neighbors[k].next = &c.root
        c.root.neighbors[k].prev = &c.root
    }

    elems := make([]Element[Data], l.Len())
    neighbors := make([]neighbors[Data], l.max * l.Len())

    e := l.root.neighbors[0].next
    for i := 0; i < l.Len(); i++ {
        var f *Element[Data]
        f, elems, neighbors = e.clone(c, clone, elems, neighbors)

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

// `SetHeight` changes the maximal height `h` of elements in the skip list `l`.
// If `h == 1`, the skip list degenerates to a doubly-linked list.
func (l *Skiplist[Data]) SetHeight(h int) {
    if h >= 1 && h <= 64 {
        d := h - len(l.heights)
        switch {
        case d < 0: // Cut elements with height above h to h.
            e := l.root.neighbors[h].next
            for e != &l.root {
                // Update height count.
                for k := h; k < len(e.neighbors); k++ {
                    l.heights[h - 1] += l.heights[k]
                    l.heights[k] = 0
                }
                // Cut height.
                f := e.neighbors[h].next
                e.neighbors = e.neighbors[:h]
                e = f
            }
            l.root.neighbors = l.root.neighbors[:h]
            l.heights = l.heights[:h]
            l.max = min(l.max, h)
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


// `resize` returns a slice of length `n`.  The first `min(n, len(s))` elements
// of slice s are the first `min(n, len(s))` of the returned slice.
func resize[S ~[]E, E any](s []E, n int) []E {
    if n > cap(s) {
        // Make a big enough copy of s.
        t := make([]E, n)
        copy(t, s)
        return t
    }
    // Capacity of s big enough, just reset the length of s.
    t := s[:n]
    for i := len(s);  i < n; i++ {
        // Zero values to avoid memory leaks.
        var z E
        t[i] = z
    }
    return t
}

// `nextPow2` returns the smallest power of 2 that is larger than `n`.
func nextPow2(n int) int {
    if n <= 1 {
        return 1
    }
    return 1 << bits.Len(uint(n - 1))
}

// `prevPow2` returns the largest power of 2 that is smaller than `n`.
func prevPow2(n int) int {
    if n <= 1 {
        return 0 // no power of 2 < n
    }
    return 1 << (bits.Len(uint(n - 1)) - 1)
}

// `ResetHeights` sets the heights of the elements optimally in the skip list
// `l` new.  Note that the optimality is not preserved when adding or removing
// elements from the skip list.  However, when there are many lookups and few
// removals or insertions, optimizing the heights might make sense.
// The benchmarks do not show a significant performance improcements.
func (l *Skiplist[Data]) ResetHeights() {
    if l.Len() == 0 {
        // Nothing to do for the empty skip list.
        return
    }

    p := make([]*Element[Data], l.Height())
    for k := 0; k < l.Height(); k++ {
        p[k] = &l.root
        l.heights[k] = 0
    }
    l.max = 1

    mid := prevPow2(l.Len()) / 2
    count, n := 0, 0
    for elem := l.root.neighbors[0].next; elem != &l.root; elem = elem.neighbors[0].next {
        var h int
        if count <= mid {
            n++
            h = bits.TrailingZeros(uint(n)) + 1
        } else if count > l.Len() - mid {
            n--
            h = bits.TrailingZeros(uint(n - 1)) + 1
        } else {
            h = l.newElementHeight()
        }
        count++
        elem.neighbors = resize[[]neighbors[Data], neighbors[Data]](elem.neighbors, h)

        // Link element neighbors (previous).
        for k := 1; k < h; k++ {
            p[k].neighbors[k].next = elem
            elem.neighbors[k].prev = p[k]
            p[k] = elem
        }

        // Update skip list data.
        l.heights[h - 1]++
        l.max = max(l.max, h)
    }

    // Complete links of last element.
    for k := 1; k < l.Height(); k++ {
        p[k].neighbors[k].next = &l.root
        l.root.neighbors[k].prev = p[k]
    }
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


// `Add` returns the element in the skip list `l` with value `val`.  If there
// is no such element, then `Insert` inserts it into `l`.  If `l` already
// contains an element of the same order as `val`, the element is not changed.
//
// Note that by assigning a value to the element's `Value` field, one overwrites
// the value stored in the skip list.  Furthermore, note that not allowing that
// a skip list stores two or more elements with the same order is without loss
// of generality.  For instance, one could extend the `Data` type with a second
// component that stores the elements of the same order.  For some uses a
// counter as a second component would already be sufficient.  The counter
// stores the number of occurences.  For such uses, one could write a wrapper
// package that uses the skiplist package.
func (l *Skiplist[Data]) Add(val Data) *Element[Data] {
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
        l.max = max(l.max, e.height())
    }
    return e
}


// `Remove` removes the element `e` from the skip list `l`.  It returns `e`'s
// value.  `e` must not be nil.
func (l *Skiplist[Data]) Remove(e *Element[Data]) Data {
    if e.list == l {
        l.len--
        l.heights[e.height() - 1]--
        if h := e.height(); h == l.max {
            for ; l.heights[h - 1] == 0 && h >= 0; h-- { }
            l.max = h
        }
        for h := e.height() - 1; h >= 0; h-- {
            e.neighbors[h].prev.neighbors[h].next = e.neighbors[h].next
            e.neighbors[h].next.neighbors[h].prev = e.neighbors[h].prev
        }
        e.neighbors = nil // avoid memory leaks
        e.list = nil
    }
    return e.Value
}


// `IsSorted` returns true if the elements in the skip list `l` are ordered
// ascendingly.
func (l *Skiplist[Data]) IsSorted() bool {
    var p, q *Element[Data]
    for p, q = nil, l.Front(); q != nil; p, q = q, p.Next() {
        if p != nil && l.cmp(p.Value, q.Value) >= 0 {
            return false
        }
    }
    return true
}


// Auxiliary functions.

// `find` returns an element with value `val`.  If such an element does exist in
// the skip list `l`, `find` creates such an element with information about
// elements in `l` for inserting the element.  Otherwise, the respective element
// of `l` is returned.  `find` is used by the `Add` function.
func (l *Skiplist[Data]) find(val Data) *Element[Data] {
    before := l.newNeighbors()
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

// `newNeighbors` returns the neighbors part for an new element of the skip list
// `l`.
func (l *Skiplist[Data]) newNeighbors() []neighbors[Data] {
    h := l.newElementHeight()
    if l.prealloc < 1 {
        // No preallocation of memory for skip list elements.
        return make([]neighbors[Data], h)
    }
    if h > len(l.preelems) {
        // No enough memory preallocated.  New pre-allocation of a memory junk
        // for new skip list elements.
        l.SetPrealloc(l.prealloc)
    }
    r := l.preelems[:h:h]
    l.preelems = l.preelems[h:]
    return r
}

// `newElementHeight` returns the height for a new element in the skip list `l`.
func (l *Skiplist[Data]) newElementHeight() int {
    if h := bits.TrailingZeros64(l.rand.Uint64() & ((1 << uint(l.Height())) - 1)); h < l.max {
        // TODO: The probabiliy of choosing a too large height might be too
        // large. Check!
        return 1 + h
    }
    if l.max == 0 || l.max < len(l.heights) && l.heights[l.max - 1] > 0 {
        // Increase maximal height of an element if we already have
        l.max++
     }
    return 1
}
