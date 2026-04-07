package tree_test

import (
    "math/rand/v2"
    "testing"
    "github.com/flxch/container/tree"
)


// Tests with trees with data values of type int.

// `compare` compares the two integers `k` and `l`.  The function is used for
// comparing the keys in the trees.  Note that we cannot just return `k` - `l`.
// Taking the difference instead of a switch would be slightly faster.
// However, underflow can occur, e.g., very small negative integer - very large
// positive integer.
//go:nosplit
func compare(k, l int) int {
    switch {
    case k < l: // less than
        return -1
    case k > l: // greater than
        return 1
    default:    // equal to
        return 0
    }
}

// `build` returns a tree with the values 0, ..., `n`-1.  The values are added
// in a random order.
func build(n int) *tree.Tree[int] {
    t := tree.New(compare)
    for _, v := range rand.Perm(n) {
        t.Add(v)
    }
    return t
}

// `random` returns a tree with `n` random integers.  Note that the tree can
// contain duplicates.
func random(n int) *tree.Tree[int] {
    t := tree.New(compare)
    for t.Len() < n {
        t.Add(rand.Int())
    }
    return t
}


func TestSimple(t *testing.T) {
    T := tree.New(compare)

    T.Add(1)
    T.Add(1)
    T.Add(2)
    if T.Len() != 3 {
        t.Errorf("expecting that the tree has 3 elements")
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
    }
    if k, ok := T.Lookup(1); !ok {
        t.Errorf("expecting to find an element with the key 1 in the tree")
    } else if k != 1 {
        t.Errorf("expecting the value 1 (not %d)", k)
    }

    T.Remove(1)
    if T.Len() != 1 {
        t.Errorf("expecting that the tree has 1 element")
    }
    if _, ok := T.Lookup(1); ok {
        t.Errorf("expecting to find no element with the key 1 in the tree")
    }

    T.Remove(1)
    if T.Len() != 1 {
        t.Errorf("expecting that the tree has 1 element")
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
    }
    for i := 0; i < n; i++ {
        T.Remove(i)
    }
    if T.Len() != 0 {
        t.Errorf("expecting that the tree has no elements")
    }
}

func TestRemove(t *testing.T) {
    n := 100
    T := build(n)
    if _, ok := T.Remove(n + 3); ok {
        t.Errorf("deleted non-existent item")
    }
    if _, ok := T.Remove(-2); ok {
        t.Errorf("deleted non-existent item")
    }
    for i := 0; i < n; i++ {
        if u, ok := T.Remove(i); !ok || u != i {
            t.Errorf("deletion failed")
        }
    }
    if _, ok := T.Remove(n + 2); ok {
        t.Errorf("deleted non-existent item")
    }
    if _, ok := T.Remove(-3); ok {
        t.Errorf("deleted non-existent item")
    }
}

func TestPartialRemove(t *testing.T) {
    n := 100
    T := build(n)
    for i := 1; i < n-1; i++ {
        v, ok := T.Remove(i)
        if !ok {
            t.Errorf("item %d not removed", i)
        }
        if i != v {
            t.Errorf("item %d removed, not %d", v, i)
        }
    }
    j := 0
    T.AscendGeq(0, func(item int) bool {
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
    T.AscendGeq(0, func(item int) bool {
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


