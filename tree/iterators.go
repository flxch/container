package tree


// `Ascend` calls the function `iterator` in ascending order for all the
// elements' data items in the tree `t`.
func (t *Tree[Data]) Ascend(iterator func(Data)) {
    if t.root != nil {
        ascend(t.root, iterator)
    }
}

func ascend[Data any](h *node[Data], iterator func(Data)) {
    // ascend assumes that h is not nil.
    if h.left != nil {
        ascend(h.left, iterator)
    }
    iterator(h.data)
    if h.right != nil {
        ascend(h.right, iterator)
    }
}


// `Descend` is dual to `Ascend`.
func (t *Tree[Data]) Descend(iterator func(Data)) {
    if t.root != nil {
        descend(t.root, iterator)
    }
}

func descend[Data any](h *node[Data], iterator func(Data)) {
    if h.right != nil {
        descend(h.right, iterator)
    }
    iterator(h.data)
    if h.left != nil {
        descend(h.left, iterator)
    }
}


// `AscendGeq` calls the function `iterator` in ascending order for the
// elements' data items in the tree `t` that are greater than or equal to the
// data item `pivot`.  It stops whenever `iterator` returns false.
func (t *Tree[Data]) AscendGeq(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.ascendGeq(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) ascendGeq(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    // ascendGeq assumes that h is not nil.
    if t.compare(h.data, pivot) >= 0 {
        if h.left != nil && !t.ascendGeq(h.left, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.right == nil || t.ascendGeq(h.right, pivot, iterator)
}


// `DescendLeq` is dual to `AscendGeq`.
func (t *Tree[Data]) DescendLeq(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.descendLeq(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) descendLeq(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    if t.compare(pivot, h.data) >= 0 {
        if h.right != nil && !t.descendLeq(h.right, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.left == nil || t.descendLeq(h.left, pivot, iterator)
}


// `AscendGeater` is similar to `AscendGeq`.
func (t *Tree[Data]) AscendGreater(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.ascendGreater(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) ascendGreater(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    // ascendGeq assumes that h is not nil.
    if t.compare(h.data, pivot) > 0 {
        if h.left != nil && !t.ascendGreater(h.left, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.right == nil || t.ascendGreater(h.right, pivot, iterator)
}

// `DescendLess` is similiar to `AscendGeq`.
func (t *Tree[Data]) DescendLess(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.descendLess(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) descendLess(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    if t.compare(pivot, h.data) > 0 {
        if h.right != nil && !t.descendLess(h.right, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.left == nil || t.descendLess(h.left, pivot, iterator)
}


// Range iterators are supported since Go1.23.

// The `WalkAscend` function is an alternative to the `Ascend` function.
// It can be used as range in for loops.
func (t *Tree[Data]) WalkAscend(yield func(data Data) bool) {
    t.root.walkAscend(yield)
}

func (n *node[Data]) walkAscend(yield func(data Data) bool) bool {
    return n == nil ||
        n.left.walkAscend(yield) && yield(n.data) && n.right.walkAscend(yield)
}

// The `WalkDescend` function is an alternative to the `Descend` function.
// It can be used as range in for loops.
func (t *Tree[Data]) WalkDescend(yield func(data Data) bool) {
    t.root.walkDescend(yield)
}

func (n *node[Data]) walkDescend(yield func(data Data) bool) bool {
    return n == nil ||
        n.right.walkDescend(yield) && yield(n.data) && n.left.walkDescend(yield)
}

