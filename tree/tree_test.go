package tree_test

import (
    "cmp"
    "math/rand/v2"
    "testing"
    "github.com/flxch/container/tree"
)


// Tests with trees with data values of type int.

// `compare` compares the two integers `k` and `l`.  Note that the result is
// only valid for non-negative values of the type int.  For large negative
// integers, `k` - `l` might overflow or underflow.  Currently `compare` is not
// used in the tests and benchmarks.  Instead, `cmp.Compare` from the standard
// library is used.  However, according to benchmarksm `compare` is
// significantly faster than `cmp.Compare`. `compare` takes roughly the same
// time as directly comparing integers with <, >, or == when inlined `compare`.
// If not inlined, the time is roughly doubled.
//go:nosplit
func compare(k, l int) int {
    return k - l
}

// `build` returns a tree with the values 0, ..., `n`-1.  The values are added
// in a random order.
func build(n int) *tree.Tree[int] {
    t := tree.New[int](cmp.Compare)
    for _, v := range rand.Perm(n) {
        t.Add(v)
    }
    return t
}

// `random` returns a tree with `n` random, non-negative integers.  Note that
// the tree does not contain duplicates.
func random(n int) *tree.Tree[int] {
    t := tree.New[int](cmp.Compare)
    for t.Len() < n {
        t.Insert(rand.Int())
    }
    return t
}

// `worst` returns a tree with the worst-case height for the elements, 0, ...,
// `n`-1.  Insert alternating smallest and largest element.
func worst(n int) *tree.Tree[int] {
    t := tree.New[int](cmp.Compare)
    for i := 0; i < n / 2; i++ {
        t.Add(i)
        t.Add(n - i)
    }
    return t
}

// `best` returns a tree with the best-case height for the elements 0, ...,
// `n`-1.  Insert medians of remaining elements.
func best(n int) *tree.Tree[int] {
    t := tree.New[int](cmp.Compare)
    for s := n / 2; s > 0; s /= 2 {
        for i := s; i < n; i += s {
            t.Insert(i)
        }
    }
    return t
}



func TestSimple(t *testing.T) {
    T := tree.New[int](cmp.Compare)

    T.Add(1)
    T.Add(1)
    T.Add(2)
    if T.Len() != 3 {
        t.Errorf("expecting that the tree has 3 elements")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    if k, ok := T.Lookup(1); !ok {
        t.Errorf("expecting to find an element with the key 1 in the tree")
    } else if k != 1 {
        t.Errorf("expecting the value 1 (not %d)", k)
    }

    if m, ok := T.Maximum(); !ok {
        t.Errorf("expecting to find a max element in the tree")
    } else if m != 2 {
        t.Errorf("expecting the max value 2 (not %d)", m)
    }
    if m, ok := T.Minimum(); !ok {
        t.Errorf("expecting to find a min element in the tree")
    } else if m != 1 {
        t.Errorf("expecting the min value 1 (not %d)", m)
    }

    T.Remove(1)
    if T.Len() != 2 {
        t.Errorf("expecting that the tree has 2 elements")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    if k, ok := T.Lookup(1); !ok {
        t.Errorf("expecting to find an element with the key 1 in the tree")
    } else if k != 1 {
        t.Errorf("expecting the value 1 (not %d)", k)
    }

    T.Remove(1)
    if T.Len() != 1 {
        t.Errorf("expecting that the tree has 1 element")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    if _, ok := T.Lookup(1); ok {
        t.Errorf("expecting to find no element with the key 1 in the tree")
    }

    T.Remove(1)
    if T.Len() != 1 {
        t.Errorf("expecting that the tree has 1 element")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    if k, ok := T.Lookup(2); !ok {
        t.Errorf("expecting to find an element with the key 2 in the tree")
    } else if k != 2 {
        t.Errorf("expecting the value 2 (not %d)", k)
    }
}


func TestAddRemove(t *testing.T) {
    n := 1000
    T := build(n)
    if T.Len() != n {
        t.Errorf("expecting that the tree has %d elements", n)
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    for i := 0; i < n; i++ {
        T.Remove(i)
        if err := T.IsLlrbTree(); err != nil {
            t.Errorf("no left-leaning red black tree: %v", err)
        }
    }
    if T.Len() != 0 {
        t.Errorf("expecting that the tree has no elements")
    }
}

func TestRemove(t *testing.T) {
    n := 100
    T := build(n)
    if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }

    if _, ok := T.Remove(n + 3); ok {
        t.Errorf("deleted non-existent item")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    if _, ok := T.Remove(-2); ok {
        t.Errorf("deleted non-existent item")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }

    for i := 0; i < n; i++ {
        if u, ok := T.Remove(i); !ok || u != i {
            t.Errorf("deletion failed")
        } else if err := T.IsLlrbTree(); err != nil {
            t.Errorf("no left-leaning red black tree: %v", err)
        }
    }

    if _, ok := T.Remove(n + 2); ok {
        t.Errorf("deleted non-existent item")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
    if _, ok := T.Remove(-3); ok {
        t.Errorf("deleted non-existent item")
    } else if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }
}

func TestPartialRemove(t *testing.T) {
    n := 100
    T := build(n)
    if err := T.IsLlrbTree(); err != nil {
        t.Errorf("no left-leaning red black tree: %v", err)
    }

    for i := 1; i < n-1; i++ {
        v, ok := T.Remove(i)
        if !ok {
            t.Errorf("item %d not removed", i)
        } else if err := T.IsLlrbTree(); err != nil {
            t.Errorf("no left-leaning red black tree: %v", err)
        }
        if i != v {
            t.Errorf("item %d removed, not %d", v, i)
        }
    }
    j := 0
    T.WalkAscendGeq(0, func(item int) bool {
        switch j {
        case 0:
            if item != 0 {
                t.Errorf("expecting 0 but got %d", item)
            }
        case 1:
            if item != n-1 {
                t.Errorf("expecting %d but got %d", n-1, item)
            }
        }
        j++
        return true
    })
}

func TestDoubleAdd(t *testing.T) {
    n := 1000
    T := build(n)
    perm := rand.Perm(n)
    for i := 0; i < n; i++ {
        T.Add(perm[i])
    }
    j := 0
    T.WalkAscendGeq(0, func(item int) bool {
        if item != j/2 {
            t.Fatalf("incorrect order of keys")
        }
        j++
        return true
    })
}


// Tests with key-value pairs.

type element struct {
    key   int
    value interface{}
}

func TestKeyValue(t *testing.T) {
    T := tree.New(func(k, l element) int {
        switch {
        case k.key < l.key:
            return -1
        case k.key > l.key:
            return 1
        default:
            return 0
        }
    })
    T.Add(element{key: 10, value: "10"})
    T.Add(element{key: 5,  value: "5"})
    T.Add(element{key: 7,  value: "7"})

    var u, v element
    var ok   bool
    u, ok = T.Lookup(element{key: 10})
    if !ok {
        t.Errorf("key not found")
    }
    if u.key != 10 {
        t.Errorf("invalid key")
    }
    if u.value != "10" {
        t.Errorf("invalid value")
    }

    u.value = "10modified"
    v, ok = T.Lookup(element{key: 10})
    if !ok {
        t.Errorf("key not found")
    }
    if v.key != 10 {
        t.Errorf("invalid key")
    }
    if v.value != "10" {
        t.Errorf("invalid value")
    }

    t.Logf("%v (%p), %v (%p) ", u.value, &u.value, v.value, &v.value)
}

func TestKeyValuePtr(t *testing.T) {
    T := tree.New(func(k, l *element) int {
        switch {
        case k.key < l.key:
            return -1
        case k.key > l.key:
            return 1
        default:
            return 0
        }
    })
    T.Add(&element{key: 10, value: "10"})
    T.Add(&element{key: 5,  value: "5"})
    T.Add(&element{key: 7,  value: "7"})

    var u, v *element
    var ok   bool
    u, ok = T.Lookup(&element{key: 10})
    if !ok {
        t.Errorf("key not found")
    }
    if u.key != 10 {
        t.Errorf("invalid key")
    }
    if u.value != "10" {
        t.Errorf("invalid value")
    }

    u.value = "10modified"
    v, ok = T.Lookup(&element{key: 10})
    if v.key != 10 {
        t.Errorf("invalid key")
    }
    if v.value != "10modified" {
        t.Errorf("invalid value")
    }

    t.Logf("%v (%p), %v (%p) ", u.value, &u.value, v.value, &v.value)
}


