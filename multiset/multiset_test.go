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

    if S.Len() != 4 {
        t.Errorf("wrong size")
    }

    t.Logf("%s", S)
}
