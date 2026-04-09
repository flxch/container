package tree


// `WalkAscend` calls the function `iterator` in ascending order for all the
// elements' data items in the tree `t`.
func (t *Tree[Data]) WalkAscend(iterator func(Data)) {
    if t.root != nil {
        walkAscend(t.root, iterator)
    }
}

func walkAscend[Data any](h *node[Data], iterator func(Data)) {
    // ascend assumes that h is not nil.
    if h.left != nil {
        walkAscend(h.left, iterator)
    }
    iterator(h.data)
    if h.right != nil {
        walkAscend(h.right, iterator)
    }
}


// `WalkDescend` is dual to `WalkAscend`.
func (t *Tree[Data]) WalkDescend(iterator func(Data)) {
    if t.root != nil {
        walkDescend(t.root, iterator)
    }
}

func walkDescend[Data any](h *node[Data], iterator func(Data)) {
    if h.right != nil {
       walkDescend(h.right, iterator)
    }
    iterator(h.data)
    if h.left != nil {
        walkDescend(h.left, iterator)
    }
}


// `WalkAscendGeq` calls the function `iterator` in ascending order for the
// elements' data items in the tree `t` that are greater than or equal to the
// data item `pivot`.  It stops whenever `iterator` returns false.
func (t *Tree[Data]) WalkAscendGeq(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.walkAscendGeq(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) walkAscendGeq(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    // ascendGeq assumes that h is not nil.
    if t.compare(h.data, pivot) >= 0 {
        if h.left != nil && !t.walkAscendGeq(h.left, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.right == nil || t.walkAscendGeq(h.right, pivot, iterator)
}


// `WalkDescendLeq` is dual to `WalkAscendGeq`.
func (t *Tree[Data]) WalkDescendLeq(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.walkDescendLeq(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) walkDescendLeq(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    if t.compare(pivot, h.data) >= 0 {
        if h.right != nil && !t.walkDescendLeq(h.right, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.left == nil || t.walkDescendLeq(h.left, pivot, iterator)
}


// `WalkAscendGeater` is similar to `WalkAscendGeq`.
func (t *Tree[Data]) WalkAscendGreater(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.walkAscendGreater(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) walkAscendGreater(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    // ascendGeq assumes that h is not nil.
    if t.compare(h.data, pivot) > 0 {
        if h.left != nil && !t.walkAscendGreater(h.left, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.right == nil || t.walkAscendGreater(h.right, pivot, iterator)
}

// `WalkDescendLess` is similiar to `WalkAscendGeq`.
func (t *Tree[Data]) WalkDescendLess(pivot Data, iterator func(Data) bool) {
    if t.root != nil {
        t.walkDescendLess(t.root, pivot, iterator)
    }
}

func (t *Tree[Data]) walkDescendLess(h *node[Data], pivot Data, iterator func(Data) bool) bool {
    if t.compare(pivot, h.data) > 0 {
        if h.right != nil && !t.walkDescendLess(h.right, pivot, iterator) {
            return false
        }
        if !iterator(h.data) {
            return false
        }
    }
    return h.left == nil || t.walkDescendLess(h.left, pivot, iterator)
}


// Range iterators (supported since Go 1.23).

// The `Ascend` function is an alternative to the `WalkAscend` function.
// It can be used as range in for loops.
func (t *Tree[Data]) Ascend(yield func(data Data) bool) {
    t.root.ascend(yield)
}

func (n *node[Data]) ascend(yield func(data Data) bool) bool {
    return n == nil ||
        n.left.ascend(yield) && yield(n.data) && n.right.ascend(yield)
}

// The `Descend` function is an alternative to the `WalkDescend` function.
// It can be used as range in for loops.
func (t *Tree[Data]) Descend(yield func(data Data) bool) {
    t.root.descend(yield)
}

func (n *node[Data]) descend(yield func(data Data) bool) bool {
    return n == nil ||
        n.right.descend(yield) && yield(n.data) && n.left.descend(yield)
}

