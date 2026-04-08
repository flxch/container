package tree


// A variable of the type `Tree` stores the data elements in a (left-leaning)
// red-black tree.  The data elements can be of `any` type.  A compare function
// orders the data elements.
type Tree[Data any] struct {
    // The number of nodes in the tree.
    count   int
    // The root of the tree.
    root    *node[Data]
    // The ordering on the elements that can be stored in the tree.
    compare func(Data, Data) int
}

// The `node` data structure represents a node in the search tree.
type node[Data any] struct {
    // The data element that the node carries.
    data  Data
    // The left child and the right child of the node.
    left  *node[Data]
    right *node[Data]
    // The color of the node.  If true, the color of the link (incoming from the
    // parent) is black, otherwise, the coler is red.  The color of new nodes is
    // always red.
    black bool
}


// `New` creates and returns a new balanced search (left-leaning red-black)
// tree.  Tree elements are ordered by the function `cmp`.
func New[Data any](cmp func(Data, Data) int) *Tree[Data] {
    return &Tree[Data]{compare: cmp}
}

// `Clone` returns a copy of the tree `t`.  The function `clone` is used to
// clone the tree's data elements.  `clone` must be order preserving.
func (t *Tree[Data]) Clone(clone func(Data) Data) *Tree[Data] {
    if t.root == nil {
        return New(t.compare)
    }
    // It is faster to the allocate memory for all nodes in the cloned tree at
    // once instead of creating the cloned nodes separately.
    r, _ := t.root.clone(clone, make([]node[Data], t.Len()))
    return &Tree[Data]{
        count:   t.count,
        root:    r,
        compare: t.compare,
    }
}

// `clone` assumes that `n` is not nil.
func (n *node[Data]) clone(clone func(Data) Data, ns []node[Data]) (*node[Data], []node[Data]) {
    // Get fresh node from the node list for cloning.
    c := &ns[0]
    ns = ns[1:]
    // Set fields for clone.
    c.data = clone(n.data)
    c.black = n.black
    // It is slightly faster to check whether the node's children exist instead
    // of always calling clone (even when they are nil).
    if n.left != nil {
        c.left, ns = n.left.clone(clone, ns)
    }
    if n.right != nil {
        c.right, ns = n.right.clone(clone, ns)
    }
    return c, ns
}

// `Reset` empties the tree `t`.
func (t *Tree[Data]) Reset() {
    t.count = 0
    t.root  = nil
}

// `Len` returns the number of nodes in the tree `t`.
func (t *Tree[Data]) Len() int {
    return t.count
}


// `Lookup` returns an element from the tree `t` whose order is the same as the
// one of `key`.  If `key` is not contained in `t` then the returned element is
//  undefined together with the Boolean value false.
func (t *Tree[Data]) Lookup(key Data) (Data, bool) {
    h := t.root
    for h != nil {
        cmp := t.compare(key, h.data)
        switch {
        case cmp < 0:
            h = h.left
        case cmp > 0:
            h = h.right
        default:
            return h.data, true
        }
    }
    // Key not in tree.
    var zero Data
    return zero, false
}


// `Remove` deletes an element with `key` from the tree `t`.  The deleted item
// is returned together with the Boolean true.  If no element is deleted,
// `Remove` returns an undefined value together with false.
func (t *Tree[Data]) Remove(key Data) (Data, bool) {
    var data Data
    var ok   bool
    t.root, data, ok = t.remove(t.root, key)
    if t.root != nil {
        t.root.black = true
    }
    if !ok {
        var zero Data
        return zero, false
    }
    t.count--
    return data, true
}

func (t *Tree[Data]) remove(h *node[Data], elem Data) (*node[Data], Data, bool) {
    if h == nil {
        var zero Data
        return nil, zero, false
    }
    var del Data
    var ok  bool
    cmp := t.compare(elem, h.data)
    if cmp < 0 {
        // The data value that should be removed is smaller than the node's data
        // value.  That is, the value must be in the left subtree.
        if h.left == nil {
            // Nothing to delete, data value not present.
            var zero Data
            return h, zero, false
        }
        if h.left.isBlack() && h.left.left.isBlack() {
            h = h.moveRedLeft()
        }
        // Continue search in left subtree.
        h.left, del, ok = t.remove(h.left, elem)
        return h.fixUp(), del, ok
    }
    // The data value that should be removed is equal to or greater than the
    // node's data value, i.e., cmp >= 0.
    if h.left.isRed() {
        h = h.rotateRight()
        // h has changed.  Data values may be equal now.
        cmp = t.compare(elem, h.data)
    }
    if cmp == 0 && h.right == nil {
        // The data value that should be removed equals h.data; h is the deleted
        // node.
        return nil, h.data, true
    }
    if h.right != nil && h.left != nil && h.right.isBlack() && h.right.left.isBlack() {
        // The checks h.right != nil and h.left != nil are both necessary.
        // moveRedRight requires that both children are not nil.
        h = h.moveRedRight()
        // h may has changed.  Data values may be equal now.
        cmp = t.compare(elem, h.data)
    }
    if cmp == 0 {
        // The data value that should be removed equals h.data.  Note that
        // h.right != nil holds because otherwise we would have returned in the
        // second to last if statement.
        var sub *node[Data]
        h.right, sub = h.right.removeMin()
        if sub == nil {
            panic("nil node in BST that should carry a data value")
        }
        h.data, del = sub.data, h.data
        return h.fixUp(), del, true
    }
    // The data value that should be removed is greater than h.data, i.e., cmp >
    // 0.  Continue search in right subtree.
    h.right, del, ok = t.remove(h.right, elem)
    return h.fixUp(), del, ok
}

// The second return argument of returnMin is the node that carries the data
// item.
func (h *node[Data]) removeMin() (*node[Data], *node[Data]) {
    if h == nil {
        panic("nil node in BST that should carry a data value")
        return nil, nil
    }
    if h.left == nil {
        return nil, h
    }
    if h.left.isBlack() && h.left.left.isBlack() {
        h = h.moveRedLeft()
    }
    var sub *node[Data]
    h.left, sub = h.left.removeMin()
    return h.fixUp(), sub
}


// `Add` inserts the element `elem` into the tree `t`.  If `t` already contains
// an element of the same order, `elem` is nevertheless inserted into `t`.
func (t *Tree[Data]) Add(elem Data) {
    if t.root == nil {
        t.root = &node[Data]{data: elem}
    } else {
        t.root = t.add(t.root, elem)
    }
    t.root.black = true
    t.count++
}

func (t *Tree[Data]) add(h *node[Data], elem Data) *node[Data] {
    if t.compare(elem, h.data) < 0 {
        if h.left == nil {
            h.left = &node[Data]{data: elem}
        } else {
            h.left = t.add(h.left, elem)
        }
        return h.walkUp()
    }
    if h.right == nil {
        h.right = &node[Data]{data: elem}
    } else {
        h.right = t.add(h.right, elem)
    }
    return h.walkUp()
}


// `Insert` inserts the element `elem` into the tree `t`.  In contrast to `Add`,
// `Insert` replaces an existing element of the same order with `elem`.
// `Insert` returns the replaced element (the returned value is undefined, if
// `elem` does not replace an element in `t``).  The second return value is true
// if an element was replaced (i.e., the returned element existed in the tree).
func (t *Tree[Data]) Insert(elem Data) (Data, bool) {
    var repl Data
    var ok   bool
    t.root, repl, ok = t.insert(t.root, elem)
    t.root.black = true
    if !ok {
        // New element, i.e., the element did not replace any element.
        t.count++
    }
    return repl, ok
}

func (t *Tree[Data]) insert(h *node[Data], elem Data) (*node[Data], Data, bool) {
    if h == nil {
        // Add the new data item.
        var zero Data
        return &node[Data]{data: elem}, zero, false
    }
    var repl Data
    var ok   bool
    cmp := t.compare(elem, h.data)
    switch {
    case cmp < 0:
        h.left, repl, ok = t.insert(h.left, elem)
    case cmp > 0:
        h.right, repl, ok = t.insert(h.right, elem)
    default:
        // Replace the node's data item.
        h.data, repl, ok = elem, h.data, true
    }
    return h.walkUp(), repl, ok
}


// Rotations for the 2-3 LLRB algorithm.

// Unfortunately, some of the helper functions below are not inlined by the Go
// compiler.  Check compiler output with the `-gcflags "-m"`.  For these helper
// functions, we use the Go compiler directive `go:nosplit` for optimizing a
// function call to these functions.  A few nanoseconds should be saved.  This
// is a low-level optimization, which also depends on the Go compiler version.
// We remark that Go compiler directives are not guaranteed to be backward
// compatible.  However, overall it is currently unclear whether this
// optimization is worthwhile to have.

// Furthermore, note that although a tree's compare function is typically very
// small and can be inlined in the tree functions like `Add` and `Remove`, it
// will not.  The reason is that we call the compare function via a function
// pointer.  Thus, a few additional nanoseconds are saved when using the Go
// compiler directive `go:nosplit` also for the compare function.

//go:nosplit
func (h *node[Data]) walkUp() *node[Data] {
    if h.right.isRed() && h.left.isBlack() {
        h = h.rotateLeft()
    }
    if h.left.isRed() && h.left.left.isRed() {
        h = h.rotateRight()
    }
    if h.left.isRed() && h.right.isRed() {
        h.flip()
    }
    return h
}


// Color of a node.

// `isRed` returns true if the color of the node `h` is red.
// Note that the color of the nil node is not red.
func (h *node[Data]) isRed() bool {
    //return h != nil && !h.black
    if h == nil {
        return false
    }
    return !h.black
}

// `black` returns true if the color of the node `h` is black.
// Note that the color of the nil node is black.
func (h *node[Data]) isBlack() bool {
    //return h == nil || h.black
    if h == nil {
        return true
    }
    return h.black
}

// Internal node manipulation routines.

// `rotateLeft` makes h.right the root.
func (h *node[Data]) rotateLeft() *node[Data] {
    x := h.right
    if x.black {
        // This should never be the case.
        panic("rotating a black node in BST (rotateLeft)")
    }
    h.right = x.left
    x.left  = h
    x.black = h.black
    h.black = false
    return x
}

// `rotateRight` makes h.left the root.
func (h *node[Data]) rotateRight() *node[Data] {
    x := h.left
    if x.black {
        // This should never be the case.
        panic("rotating a black node in BST (rotateRight)")
    }
    h.left  = x.right
    x.right = h
    x.black = h.black
    h.black = false
    return x
}

// `flip` flips the colors of the node `h` and its children.
// Left child and right child must not be nil.
func (h *node[Data]) flip() {
    h.black       = !h.black
    h.left.black  = !h.left.black
    h.right.black = !h.right.black
}

// Left child and right child must not be nil.
//go:nosplit
func (h *node[Data]) moveRedLeft() *node[Data] {
    h.flip()
    if h.right.left.isRed() {
        h.right = h.right.rotateRight()
        h = h.rotateLeft()
        h.flip()
    }
    return h
}

// Left child and right child must not be nil.
//go:nosplit
func (h *node[Data]) moveRedRight() *node[Data] {
    h.flip()
    if h.left.left.isRed() {
        h = h.rotateRight()
        h.flip()
    }
    return h
}

//go:nosplit
func (h *node[Data]) fixUp() *node[Data] {
    if h.right.isRed() {
        h = h.rotateLeft()
    }
    if h.left.isRed() && h.left.left.isRed() {
        h = h.rotateRight()
    }
    if h.left.isRed() && h.right.isRed() {
        h.flip()
    }
    return h
}
