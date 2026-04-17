package multiset_test

import (
    "testing"
    "github.com/flxch/container/multiset"
)


func TestString(t *testing.T) {
    S := multiset.New[string]()

    S.Add("foo")
    S.Add("foo")
    S.Add("bar")
    S.Add("baz")

    if S.Card() != 4 {
        t.Errorf("wrong cardinality")
    }
    if S.Len() != 3 {
        t.Errorf("wrong support size")
    }

    t.Logf("%s", S)
}


func TestAddRemove(t *testing.T) {
    S := multiset.New[string]()

    S.Add("foo")
    S.Add("foo")
    S.Add("bar")
    S.Add("baz")
    if m := S.Lookup("foo"); m != 2 {
        t.Errorf("expected that foo's multiplicity is 2, not %d", m)
    }
    if m := S.Lookup("goo"); m != 0 {
        t.Errorf("expected that goo's multiplicity is 0, not %d", m)
    }
    if m := S.Lookup("bar"); m != 1 {
        t.Errorf("expected that bar's multiplicity is 1, not %d", m)
    }

    S.Remove("foo")
    S.Remove("baz")
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
    if m := S.Lookup("foo"); m != 0 {
        t.Errorf("expected that foo's multiplicity is 0, not %d", m)
    }
    if S.Card() != 0 || S.String() != "{}" {
        t.Errorf("exptected the empty multiset")
    }
}

func TestUnion(t *testing.T) {
    S := multiset.New[string]()
    S.Add("foo")
    S.Add("foo")
    S.Add("bar")
    S.Add("baz")

    T := multiset.New[string]()
    T.Add("foo")
    T.Add("bar")
    T.Add("bar")
    T.Add("goo")

    S.Union(T)

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


func TestIterate(t *testing.T) {
    S := multiset.New[string]()
    S.Add("foo")
    S.Add("foo")
    S.Add("bar")
    S.Add("baz")

    c := 0
    for _, n := range S.Elems() {
        c += n
    }

    if c != S.Card() {
        t.Errorf("expected size %d, not %d", c, S.Len())
    }
}
