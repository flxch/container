package multiset_test

import (
    "fmt"
    "testing"
    "github.com/flxch/container/multiset"
)


func valid[A comparable](S *multiset.Multiset[A]) error {
    n := 0
    s := 0
    for e, m := range S.Elements() {
        n++
        if m <= 0 {
            return fmt.Errorf("element %v with invalid multiplicity: %d", e, m)
        }
        s += m
    }
    if s != S.Card() {
        return fmt.Errorf("wrong cardinality %d != %d", s, S.Card())
    }
    if n != S.Len() {
        return fmt.Errorf("wrong support %d != %d", n, S.Len())
    }
    return nil
}


func TestString(t *testing.T) {
    S := multiset.New[string]()

    S.Add("foo", 1)
    S.Add("foo", 1)
    S.Add("bar", 1)
    S.Add("baz", 1)

    t.Logf("%s", S)

    if err := valid(S); err != nil {
        t.Errorf("invalid multiset: %v", err)
    }
}


func TestAddRemove(t *testing.T) {
    S := multiset.New[string]()

    S.Add("foo", 2)
    S.Add("bar", 1)
    S.Add("baz", 1)
    if err := valid(S); err != nil {
        t.Errorf("invalid multiset: %v", err)
    }
    if m := S.Lookup("foo"); m != 2 {
        t.Errorf("expected that foo's multiplicity is 2, not %d", m)
    }
    if m := S.Lookup("goo"); m != 0 {
        t.Errorf("expected that goo's multiplicity is 0, not %d", m)
    }
    if m := S.Lookup("bar"); m != 1 {
        t.Errorf("expected that bar's multiplicity is 1, not %d", m)
    }

    S.Remove("foo", 1)
    S.Remove("baz", 1)
    if err := valid(S); err != nil {
        t.Errorf("invalid multiset: %v", err)
    }
    if m := S.Lookup("foo"); m != 1 {
        t.Errorf("expected that foo's multiplicity is 1, not %d", m)
    }
    if m := S.Lookup("baz"); m != 0 {
        t.Errorf("expected that baz's multiplicity is 0, not %d", m)
    }
    if m := S.Lookup("bar"); m != 1 {
        t.Errorf("expected that bar's multiplicity is 1, not %d", m)
    }

    S.Reset()
    if err := valid(S); err != nil {
        t.Errorf("invalid multiset: %v", err)
    }
    if m := S.Lookup("foo"); m != 0 {
        t.Errorf("expected that foo's multiplicity is 0, not %d", m)
    }
    if S.Card() != 0 || S.String() != "{}" {
        t.Errorf("exptected the empty multiset")
    }
}

func TestUnion(t *testing.T) {
    S := multiset.New[string]()
    S.Add("foo", 2)
    S.Add("bar", 1)
    S.Add("baz", 1)
    T := multiset.New[string]()
    T.Add("foo", 1)
    T.Add("bar", 2)
    T.Add("goo", 1)

    S.Union(T)

    if err := valid(S); err != nil {
        t.Errorf("invalid multiset: %v", err)
    }
    if m := S.Lookup("foo"); m != 2 {
        t.Errorf("expected that foo's multiplicity is 2, not %d", m)
    }
    if m := S.Lookup("goo"); m != 1 {
        t.Errorf("expected that goo's multiplicity is 1, not %d", m)
    }
    if m := S.Lookup("bar"); m != 2 {
        t.Errorf("expected that bar's multiplicity is 2, not %d", m)
    }
}

func TestIntersection(t *testing.T) {
    S := multiset.New[string]()
    S.Add("foo", 2)
    S.Add("bar", 1)
    S.Add("baz", 1)
    T := multiset.New[string]()
    T.Add("foo", 1)
    T.Add("bar", 2)
    T.Add("goo", 3)

    S.Intersection(T)

    if err := valid(S); err != nil {
        t.Errorf("invalid multiset: %v", err)
    }
    if m := S.Lookup("foo"); m != 1 {
        t.Errorf("expected that foo's multiplicity is 1, not %d", m)
    }
    if m := S.Lookup("bar"); m != 1 {
        t.Errorf("expected that bar's multiplicity is 1, not %d", m)
    }
    if m := S.Lookup("goo"); m != 0 {
        t.Errorf("expected that gos'ss multiplicity is 0, not %d", m)
    }
    if m := S.Lookup("baz"); m != 0 {
        t.Errorf("expected that baz's multiplicity is 0, not %d", m)
    }
}


func TestIterate(t *testing.T) {
    S := multiset.New[string]()
    S.Add("foo", 2)
    S.Add("bar", 1)
    S.Add("baz", 1)

    c := 0
    for _, n := range S.Elements() {
        c += n
    }

    if c != S.Card() {
        t.Errorf("expected cardinality %d, not %d", c, S.Len())
    }
}
