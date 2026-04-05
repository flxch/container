package tree


// `Minimum` returns the smallest data value in the tree `t`.  If `t` has no
// smallest element (i.e., `t` is empty`), then the second return value is false.
func (t *Tree[Data]) Minimum() (Data, bool) {
    h := t.root
    if h == nil {
        // The empty tree has no minimum.
        var z Data
        return z, false
    }
    for h.left != nil {
        // Go left, since the left node (if it exists) has a smaller element.
        h = h.left
    }
    return h.data, true
}

// `Maximum` returns the largest data value in the tree `t`.  If `t` has no
// largest element (i.e., `t` is empty`), then the second return value is false.
func (t *Tree[Data]) Maximum() (Data, bool) {
    h := t.root
    if h == nil {
        var z Data
        return z, false
    }
    for h.right != nil {
        h = h.right
    }
    return h.data, true
}
