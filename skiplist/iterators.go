package skiplist


// `Ascend` calls the function `op` in ascending order for all the
// elements' data items in the skiplist `l`.
func (l *Skiplist[Data]) Ascend(op func(Data)) {
    for e := l.Front(); e != nil; e = e.Next() {
        op(e.Value)
    }
}

// `Descend` calls the function `op` in descending order for all the elements'
// data items in the skiplist `l`.
func (l *Skiplist[Data]) Descend(op func(Data)) {
    for elem := l.Back(); elem != nil; elem = elem.Prev() {
        op(elem.Value)
    }
}


// Helper function that finds the element to start the iteration in ascending order.
func (l *Skiplist[Data]) findAscendGeq(pivot Data) *Element[Data] {
    // Find element that is greater than or equal to pivot.
    p, q := &l.root, &l.root
    for h := l.max - 1; h >= 0; h-- {
        for q = p.neighbors[h].next; q != &l.root && l.cmp(q.Value, pivot) < 0; p, q = q, q.neighbors[h].next { }
    }
    if q == &l.root {
        // We are at the end of the list.
        return nil
    }
    return q
}

// Helper function that finds the element to start the iteration in descending order.
func (l *Skiplist[Data]) findDescendLeq(pivot Data) *Element[Data] {
    // Find element that is less than or equal to pivot.
    p, q := &l.root, &l.root
    for h := l.max - 1; h >= 0; h-- {
        for q = p.neighbors[h].prev; q != &l.root && l.cmp(q.Value, pivot) > 0; p, q = q, q.neighbors[h].prev { }
    }
    if q == &l.root {
        // We are at the beginning of the list.
        return nil
    }
    return q
}


// `AscendGeq` calls the function `op` in ascending order for the elements' data
// items in the skiplist `t` that are greater than or equal to the data item
// `pivot`.  It stops whenever `op` returns false.
func (l *Skiplist[Data]) AscendGeq(pivot Data, op func(Data) bool) {
    for elem := l.findAscendGeq(pivot); elem != nil && op(elem.Value); elem = elem.Next() { }
}

// `DescendLeq` calls the function `op` in descending order for the elements'
// data items in the skiplist `t` that are less than or equal to the data item
// `pivot`.  It stops whenever `op` returns false.
func (l *Skiplist[Data]) DescendGeq(pivot Data, op func(Data) bool) {
    for elem := l.findDescendLeq(pivot); elem != nil && op(elem.Value); elem = elem.Prev() { }
}


// Range iterators (supported since Go 1.23).

// `WalkAscend` calls the function `yield` in ascending order for all elements'
// data values in the skiplist `l` until `yield` returns false.
func (l *Skiplist[Data]) WalkAscend(yield func(Data) bool) {
    for elem := l.Front(); elem != nil && yield(elem.Value); elem = elem.Next() { }
}

// `WalkDescend` calls the function `yield` in descending order for all elements'
// data values in the skiplist `l` until `yield` returns false.
func (l *Skiplist[Data]) WalkDescend(yield func(Data) bool) {
    for elem := l.Back(); elem != nil && yield(elem.Value); elem = elem.Prev() { }
}

// `WalkAscendGeq` is similar to `WalkAscend`, except that the iteration starts
// at the smallest element that is greater than or equal to `pivot`.
func (l *Skiplist[Data]) WalkAscendGeq(pivot Data) func(func(Data) bool) {
    elem := l.findAscendGeq(pivot)
    return func(yield func(Data) bool) {
        // Iterate over elements, starting from elem.
        for ; elem != nil && yield(elem.Value); elem = elem.Next() { }
    }
}

// `WalkAscendGreater` is similar to `WalkAscend`, except that the iteration
// starts at the smallest element that is greater than `pivot`.
func (l *Skiplist[Data]) WalkAscendGreater(pivot Data) func(func(Data) bool) {
    elem := l.findAscendGeq(pivot)
    if elem != nil && l.cmp(elem.Value, pivot) == 0 {
        // Equal; go to next element.
        elem = elem.Next()
    }
    return func(yield func(Data) bool) {
        for ; elem != nil && yield(elem.Value); elem = elem.Next() { }
    }
}


// `WalkDescendLeq` is similar to `WalkDescend`, except that the iteration
// starts at the smallest element that is smaller than or equal to `pivot`.
func (l *Skiplist[Data]) WalkDescendLeq(pivot Data) func(func(Data) bool) {
    elem := l.findDescendLeq(pivot)
    return func(yield func(Data) bool) {
        for ; elem != nil && yield(elem.Value); elem = elem.Prev() { }
    }
}

// `WalkDescendLess` is similar to `WalkDescend`, except that the iteration
// starts at the smallest element that is smaller than `pivot`.
func (l *Skiplist[Data]) WalkDescendLess(pivot Data) func(func(Data) bool) {
    elem := l.findDescendLeq(pivot)
    if elem != nil && l.cmp(elem.Value, pivot) == 0 {
        elem = elem.Prev()
    }
    return func(yield func(Data) bool) {
        for ; elem != nil && yield(elem.Value); elem = elem.Prev() { }
    }
}
