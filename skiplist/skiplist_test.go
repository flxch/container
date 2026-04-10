package skiplist

import (
    "math/bits"
    "fmt"
    "math/rand/v2"
    "testing"
)


func TestOptHeight(t *testing.T) {
    N := 33
    mid := prevPow2(N) / 2
    count := 0
    for i := 0; i < N; i++ {
        var height int
        if i <= mid {
            count++
            height = bits.TrailingZeros(uint(count)) + 1
        } else if i > N - mid {
            count--
            height = bits.TrailingZeros(uint(count-1)) + 1
        } else {
            height = 2 + rand.IntN(6)
        }

        t.Logf("%d: %d (count %d)", i, height, count)
    }
}


// Test whether the element height is well distributed.
// Test does not fail.  It only provides information.
func TestElementHeight(t *testing.T) {
    N := 100000
    list := New[int](func(x, y int) int { return x - y })
    list.SetSeed(34, 3)
    list.max = 16

    heights := make([]int, list.max)
    for i := 0; i < N; i++ {
        h := list.newElementHeight()
        heights[h - 1]++
    }

    for i := 0; i < list.max - 1; i++ {
        if float64(heights[i]) > 2.35 * float64(heights[i + 1]) {
            t.Logf("Warning: too many nodes of height %d (%d) compared to %d (%d)", i, heights[i], i+1, heights[i+1])
        }
        if float64(heights[i]) < 1.65 * float64(heights[i + 1]) {
            t.Logf("Warning: too few nodes of height %d (%d) compared to %d (%d)", i, heights[i], i+1, heights[i+1])
        }
    }

    s := ""
    for h, c := range heights {
        s += fmt.Sprintf("\nheight[%d] = #%d ", h, c)
    }
    t.Logf("node heights: %s", s)
}


// `ordering` checks the order of the keys at height `h` in the skip list `l`.
func (l *Skiplist[Data]) ordering(h int) bool {
    p := &l.root
    for e := p.neighbors[h].next; e != &l.root; p, e = e, e.neighbors[h].next {
        if p != &l.root && l.cmp(e.Value, p.Value) < 0 {
            return false
        }
    }
    q := &l.root
    for e := q.neighbors[h].prev; e != &l.root; q, e = e, e.neighbors[h].prev {
        if q != &l.root && l.cmp(q.Value, e.Value) < 0 {
            return false
        }
    }
    return true
}

// `linking` checks the prev and next pointers at height `h` in the skip list
// `l`.
func (l *Skiplist[Data]) linking(h int) bool {
    p := &l.root
    for e := p.neighbors[h].next; e != &l.root; p, e = e, e.neighbors[h].next {
        if e.neighbors[h].prev.neighbors[h].next != e ||
            p.neighbors[h].next.neighbors[h].prev != p {
            return false
        }
    }
    return true
}

// `heightcount` sums the number of elements of height 'h' in the skip list `l`.
func (l *Skiplist[Data]) heightcount(h int) int {
    c := 0
    for e := l.root.neighbors[0].next; e != &l.root; e = e.neighbors[0].next {
        if len(e.neighbors) - 1 == h {
            c++
        }
    }
    return c
}

// `sanity` performs a sanity check for the skip list `l`, e.g., it checks the
// element ordering and next/prev pointers at evey height.
func (l *Skiplist[Data]) sanity() error {
    max := 0
    for h := 0; h < l.Height(); h++ {
        if !l.ordering(h) {
            return fmt.Errorf("elements in skip list are not correctly ordered at level %d", h)
        }
        if !l.linking(h) {
            return fmt.Errorf("elements in skip list are not correctly linked at level %d", h)
        }

        count := l.heightcount(h)
        if count != l.heights[h] {
            return fmt.Errorf("wrong number of height count (%d!=%d) for level %d", count, l.heights[h], h)
        }
        if count > 0 {
            max = h + 1
        }
    }
    if max != l.max {
        return fmt.Errorf("wrong maximal height %d, not %d", l.max, max)
    }
    return nil
}


// Unit tests

func TestString(t *testing.T) {
    l := New[int](func(x, y int) int { return x - y})
    if s := l.String(); s != "[]" {
        t.Errorf("Wrong string representation: %s, epxected []", s)
    }
    l = FromSlice[int](func(x, y int) int { return x - y}, []int{1, 2, 3, 4, 5})
    if s := l.String(); s != "[1,2,3,4,5]" {
        t.Errorf("Wrong string representation: %s, epxected []", s)
    }
}

func TestClone(t *testing.T) {
    l := New[int](func(x, y int) int { return x - y})
    c := l.Clone(func(x int) int { return x })
    if l.Len() != c.Len() {
        t.Errorf("Lengths are not the same.")
    }
    for e, f := l.Front(), c.Front(); e != nil; e, f = e.Next(), f.Next() {
        if e.Value != f.Value {
            t.Errorf("Values are not equal: %d != %d.", e.Value, f.Value)
        }
    }

    l = FromSlice[int](func(x, y int) int { return x - y}, []int{1, 2, 3, 4, 5})
    c = l.Clone(func(x int) int { return x })
    if l.Len() != c.Len() {
        t.Errorf("Lengths are not the same.")
    }
    for e, f := l.Front(), c.Front(); e != nil; e, f = e.Next(), f.Next() {
        if e.Value != f.Value {
            t.Errorf("Values are not equal: %d != %d.", e.Value, f.Value)
        }
    }
}


// For the following tests, the skip lists store key-value pairs, where the keys
// are ordered.
type kvpair struct {
    key int
    val string
}

// `cmp` orders the key-value pairs.
func cmp(x, y kvpair) int {
    switch {
    case x.key < y.key:
        return -1
    case x.key > y.key:
        return 1
    default:
        return 0
    }
}

func TestNew(t *testing.T) {
    l := New[kvpair](cmp)
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 0 {
        t.Errorf("wrong number (%d) of elements in skip list: %s", l.Len(), l)
    }
    if l.Front() != nil {
        t.Errorf("expected nil as front of skip list: %s", l)
    }
    if l.Back() != nil {
        t.Errorf("expected nil as back of skip list: %s", l)
    }
}

func TestAdd1(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 1 {
        t.Errorf("wrong number (%d) of elements in skip list: %s", l.Len(), l)
    }
    if e := l.Front(); e == nil {
        t.Errorf("front is nil in skip list: %s", l)
    } else if e.Value.val != "foo" {
        t.Errorf("front has wrong value (%v)", e.Value)
    }
    if e := l.Back(); e == nil {
        t.Errorf("back is nil in skip list: %s", l)
    } else if e.Value.val != "foo" {
        t.Errorf("back has wrong value (%v)", e.Value)
    }
}

func TestAdd(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    l.Add(kvpair{2, "goo"})
    l.Add(kvpair{1, "moo"})
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 3 {
        t.Errorf("wrong number (%d) of elements in skip list: %s", l.Len(), l)
    }
    if e := l.Front(); e == nil {
        t.Errorf("front is nil in skip list: %s", l)
    } else if e.Value.val != "foo" {
        t.Errorf("front has wrong value (%v)", e.Value)
    }
    if e := l.Back(); e == nil {
        t.Errorf("back is nil in skip list: %s", l)
    } else if e.Value.val != "goo" {
        t.Errorf("back has wrong value (%v)", e.Value)
    }
}

func TestAddEqual(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    l.Add(kvpair{2, "goo"})
    l.Add(kvpair{4, "baz"})
    l.Add(kvpair{1, "moo"})
    l.Add(kvpair{5, "hoo"})
    l.Add(kvpair{6, "bat"})
    l.Add(kvpair{3, "bar"})
    l.Add(kvpair{2, "XYZ"})
    t.Logf("%s", l)

    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 7 {
        t.Errorf("wrong number (%d) of elements in skip list: %s", l.Len(), l)
    }
}

func TestLookup1(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    l.Add(kvpair{1, "moo"})
    l.Add(kvpair{3, "goo"})
    if e := l.Lookup(kvpair{key: 1}); e == nil {
        t.Errorf("expected to find the key 1 in skip list: %s", l)
    } else if e.Value.val  != "moo" {
        t.Errorf("wrong value (%v) of key 1 in skip list: %s", e.Value, l)
    }
    if e := l.Lookup(kvpair{key: 2}); e != nil {
        t.Errorf("expected to not find the key 2 in skip list: %s", l)
    }
}

func TestLookup2(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    l.Add(kvpair{2, "goo"})
    l.Add(kvpair{5, "baz"})
    l.Add(kvpair{1, "moo"})
    l.Add(kvpair{8, "hoo"})
    l.Add(kvpair{7, "bat"})
    l.Add(kvpair{3, "bar"})
    if l.Lookup(kvpair{key: 3}) == nil {
        t.Errorf("expected to find the key 3 in skip list: %s", l)
    }
    if l.Lookup(kvpair{key: -1}) != nil {
        t.Errorf("expected to not find the key -1 in skip list: %s", l)
    }
    if l.Lookup(kvpair{key: 4}) != nil {
        t.Errorf("expected to not find the key 4 in skip list: %s", l)
    }
    if l.Lookup(kvpair{key: 10}) != nil {
        t.Errorf("expected to not find the key 10 in skip list: %s", l)
    }
}

func TestRemove(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    l.Add(kvpair{2, "goo"})
    l.Add(kvpair{5, "baz"})
    l.Add(kvpair{1, "moo"})
    l.Add(kvpair{8, "hoo"})
    l.Add(kvpair{7, "bat"})
    l.Add(kvpair{3, "bar"})
    e := l.Lookup(kvpair{key: 5})
    if d := l.Remove(e); d.key != 5 && d.val != "baz" {
        t.Errorf("wrong key (%d) or value (%s) of removed element from skip list: %s", d.key, d.val, l)
    } else if l.Len() != 6 {
        t.Errorf("wrong number (%d) of elements in skip list: %s", l.Len(), l)
    } else if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    l.Remove(e)
    if l.Len() != 6 {
        t.Errorf("wrong number (%d) of elements in skip list: %s", l.Len(), l)
    } else if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
}

func TestReset(t *testing.T) {
    l := New[kvpair](cmp)
    l.Add(kvpair{0, "foo"})
    l.Add(kvpair{2, "goo"})
    l.Add(kvpair{5, "baz"})
    l.Add(kvpair{1, "moo"})
    l.Add(kvpair{8, "hoo"})
    l.Add(kvpair{7, "bat"})
    l.Add(kvpair{3, "bar"})

    l.Reset()
    if l.Len() != 0 {
        t.Errorf("skip list should be empty (%d): %s", l.Len(), l)
    }
    if e := l.Lookup(kvpair{5, "baz"}); e != nil {
        t.Errorf("expected not find element 5 in skip list")
    }
}

func TestSetHeight(t *testing.T) {
    l := New[int](func(x, y int) int { return x - y })
    l.SetHeight(2)

    elems := rand.Perm(300)
    for _, v := range elems[:100] {
        l.Add(v)
    }
    t.Logf("%v", l)
    t.Logf("max height: %d, height count: %v", l.max, l.heights)
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 100 {
        t.Errorf("wrong length %d of skip list", l.Len())
    }

    l.SetHeight(4)
    for _, v := range elems[100:200] {
        l.Add(v)
    }
    t.Logf("%v", l)
    t.Logf("max height: %d, height count: %v", l.max, l.heights)
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 200 {
        t.Errorf("wrong length %d of skip list", l.Len())
    }

    l.SetHeight(3)
    for _, v := range elems[200:300] {
        l.Add(v)
    }
    t.Logf("%v", l)
    t.Logf("max height: %d, height count: %v", l.max, l.heights)
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
    if l.Len() != 300 {
        t.Errorf("wrong length %d of skip list", l.Len())
    }
    u := 0
    for v := range l.AscendGeq(1) {
        if u + 1 != v {
            t.Errorf("expected the successor of %d, not %d", u, v)
        }
        u = v
    }
}

func TestResetHeights(t *testing.T) {
    l := New[int](func(x, y int) int { return x - y })
    elems := rand.Perm(200)
    for _, v := range elems[:100] {
        l.Add(v)
    }
    t.Logf("max height: %d, height count: %v", l.max, l.heights)
    l.SetHeight(8)
    for _, v := range elems[100:200] {
        l.Add(v)
    }
    l.ResetHeights()
    t.Logf("max height: %d, height count: %v", l.max, l.heights)
    if err := l.sanity(); err != nil {
        t.Errorf("%v: %s", err, l)
    }
}
