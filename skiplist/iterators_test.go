package skiplist_test

import (
    "slices"
    "strings"
    "testing"
    "github.com/flxch/container/skiplist"
)


func TestAscend(t *testing.T) {
    l := []string{"foo", "goo", "bar", "baz"}
    sl := skiplist.New[string](strings.Compare)
    for _, elem := range l {
        sl.Add(elem)
    }
    slices.Sort(l)

    r := []string{}
    for elem := range sl.WalkAscend {
        r = append(r, elem)
    }
    if !slices.Equal(l, r) {
        t.Errorf("expected %v, got %v", l, r)
    }
}

func TestDescend(t *testing.T) {
    l := []string{"foo", "goo", "bar", "baz"}
    sl := skiplist.New[string](strings.Compare)
    for _, elem := range l {
        sl.Add(elem)
    }
    slices.Sort(l)
    slices.Reverse(l)

    r := []string{}
    for elem := range sl.WalkDescend {
        r = append(r, elem)
    }
    if !slices.Equal(l, r) {
        t.Errorf("expected %v, got %v", l, r)
    }
}
