package skiplist_test

import (
    "slices"
    "strings"
    "testing"
    "github.com/flxch/container/skiplist"
)


type tcIterator[Data any] struct {
    elements []Data
    pivot    Data
}

var tcsIterators []tcIterator[string] = []tcIterator[string]{
    tcIterator[string]{},
    tcIterator[string]{
        elements: []string{"foo", "goo", "bar", "baz"},
        pivot:    "aaaa",
    },
    tcIterator[string]{
        elements: []string{"foo", "goo", "bar", "baz"},
        pivot:    "zzzz",
    },
    tcIterator[string]{
        elements: []string{"foo", "goo", "bar", "baz"},
        pivot:    "baz",
    },
    tcIterator[string]{
        elements: []string{"foo", "goo", "bar", "baz"},
        pivot:    "foo",
    },
    tcIterator[string]{
        elements: []string{"foo", "goo", "bar", "baz"},
        pivot:    "hoo",
    },
}


func TestAscend(t *testing.T) {
    for i, tc := range tcsIterators {
        sl := skiplist.New[string](strings.Compare)
        for _, elem := range tc.elements {
            sl.Add(elem)
        }
        r := []string{}
        for elem := range sl.Ascend {
            r = append(r, elem)
        }
        // Check.
        expected := slices.Clone(tc.elements)
        slices.Sort(expected)
        if !slices.Equal(expected, r) {
            t.Errorf("#%d: expected %v, got %v", i, expected, r)
        }
    }
}

func TestDescend(t *testing.T) {
    for i, tc := range tcsIterators {
        sl := skiplist.New[string](strings.Compare)
        for _, elem := range tc.elements {
            sl.Add(elem)
        }
        r := []string{}
        for elem := range sl.Descend {
            r = append(r, elem)
        }
        // Check.
        expected := slices.Clone(tc.elements)
        slices.Sort(expected)
        slices.Reverse(expected)
        if !slices.Equal(expected, r) {
            t.Errorf("#%d: expected %v, got %v", i, expected, r)
        }
    }
}


func TestAscendGeq(t *testing.T) {
    for i, tc := range tcsIterators {
        sl := skiplist.New[string](strings.Compare)
        for _, elem := range tc.elements {
            sl.Add(elem)
        }
        r := []string{}
        for elem := range sl.AscendGeq(tc.pivot) {
            r = append(r, elem)
        }
        // Check.
        expected := slices.Clone(tc.elements)
        slices.Sort(expected)
        p, _ := slices.BinarySearch(expected, tc.pivot)
        expected = expected[p:]
        if !slices.Equal(expected, r) {
            t.Errorf("#%d: expected %v, got %v", i, expected, r)
        }
    }
}

func TestDescendLeq(t *testing.T) {
    for i, tc := range tcsIterators {
        sl := skiplist.New[string](strings.Compare)
        for _, elem := range tc.elements {
            sl.Add(elem)
        }
        r := []string{}
        for elem := range sl.DescendLeq(tc.pivot) {
            r = append(r, elem)
        }
        // Check.
        expected := slices.Clone(tc.elements)
        slices.Sort(expected)
        p, ok := slices.BinarySearch(expected, tc.pivot)
        if ok {
            p++
        }
        p = len(expected) - p
        slices.Reverse(expected)
        expected = expected[p:]
        if !slices.Equal(expected, r) {
            t.Errorf("#%d: expected %v, got %v", i, expected, r)
        }
    }
}

func TestAscendGreater(t *testing.T) {
    for i, tc := range tcsIterators {
        sl := skiplist.New[string](strings.Compare)
        for _, elem := range tc.elements {
            sl.Add(elem)
        }
        r := []string{}
        for elem := range sl.AscendGreater(tc.pivot) {
            r = append(r, elem)
        }
        // Check.
        expected := slices.Clone(tc.elements)
        slices.Sort(expected)
        p, ok := slices.BinarySearch(expected, tc.pivot)
        if ok {
            p++
        }
        expected = expected[p:]
        if !slices.Equal(expected, r) {
            t.Errorf("#%d: expected %v, got %v", i, expected, r)
        }
    }
}

func TestDescendLess(t *testing.T) {
    for i, tc := range tcsIterators {
        sl := skiplist.New[string](strings.Compare)
        for _, elem := range tc.elements {
            sl.Add(elem)
        }
        r := []string{}
        for elem := range sl.DescendLess(tc.pivot) {
            r = append(r, elem)
        }
        // Check.
        expected := slices.Clone(tc.elements)
        slices.Sort(expected)
        p, _ := slices.BinarySearch(expected, tc.pivot)
        p = len(expected) - p
        slices.Reverse(expected)
        expected = expected[p:]
        if !slices.Equal(expected, r) {
            t.Errorf("#%d: expected %v, got %v", i, expected, r)
        }
    }
}
