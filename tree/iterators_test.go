package tree_test

import (
    "slices"
    "testing"
    "github.com/flxch/container/tree"
)


func TestAscendGeq(t *testing.T) {
    expected := []int{1, 3, 4, 6}
    T := tree.New(compare)
    T.Add(4)
    T.Add(6)
    T.Add(1)
    T.Add(3)
    var actual []int
    op := func(elem int) bool { actual = append(actual, elem); return true }

    T.AscendGeq(-1, op)
    if !slices.Equal(expected, actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected)
    }
    actual = nil
    T.AscendGeq(3, op)
    if !slices.Equal(expected[1:], actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected[1:])
    }
    actual = nil
    T.AscendGeq(2, op)
    if !slices.Equal(expected[1:], actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected[1:])
    }
    actual = nil
    T.AscendGeq(2, func(elem int) bool {
        actual = append(actual, elem)
        return elem != 4
    })
    if !slices.Equal(expected[1:3], actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected[1:3])
    }
}

func TestDescendLeq(t *testing.T) {
    expected := []int{6, 4, 3, 1}
    T := tree.New(compare)
    T.Add(4)
    T.Add(6)
    T.Add(1)
    T.Add(3)
    var actual []int
    op := func(elem int) bool { actual = append(actual, elem); return true }

    T.DescendLeq(10, op)
    if !slices.Equal(expected, actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected)
    }
    actual = nil
    T.DescendLeq(4, op)
    if !slices.Equal(expected[1:], actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected[1:])
    }
    actual = nil
    T.DescendLeq(5, op)
    if !slices.Equal(expected[1:], actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected[1:])
    }
    actual = nil
    T.DescendLeq(5, func(elem int) bool {
        actual = append(actual, elem)
        return elem != 3
    })
    if !slices.Equal(expected[1:3], actual) {
        t.Errorf("wrong result: %v (should be %v)", actual, expected[1:3])
    }
}


func TestOrder(t *testing.T) {
    T := build(1000)
    j := 0
    T.AscendGeq(0, func(elem int) bool {
        if elem != j {
            t.Fatalf("wrong order")
        }
        j++
        return true
    })
}

func TestReverseOrder(t *testing.T) {
    T := tree.New(compare)
    n := 100
    for i := 0; i < n; i++ {
        T.Add(n - i)
    }
    i := 0
    T.AscendGeq(0, func(elem int) bool {
        i++
        if elem != i {
            t.Errorf("wrong order: got %d, expect %d", elem, i)
        }
        return true
    })
}

func TestEmptyWalkAscend(t *testing.T) {
    T := tree.New(compare)
    ds := make([]int, 0, 10)
    for d := range T.WalkAscend {
        t.Logf("%v ", d)
        ds = append(ds, d)
    }
    if len(ds) != 0 {
        t.Errorf("expected no data from the empty tree")
    }
}

func TestNonemptyWalkAscend(t *testing.T) {
    T := build(10)
    ds := make([]int, 0, 10)
    for d := range T.WalkAscend {
        t.Logf("%v ", d)
        ds = append(ds, d)
    }
    if !slices.Equal(ds, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
        t.Errorf("data values do not match: %v", ds)
    }
}

func TestEmptyWalkDescend(t *testing.T) {
    T := tree.New(compare)
    ds := make([]int, 0, 10)
    for d := range T.WalkDescend {
        t.Logf("%v ", d)
        ds = append(ds, d)
    }
    if len(ds) != 0 {
        t.Errorf("expected no data from the empty tree")
    }
}

func TestNonemptyWalkDescend(t *testing.T) {
    T := build(10)
    ds := make([]int, 0, 10)
    for d := range T.WalkDescend {
        t.Logf("%v ", d)
        ds = append(ds, d)
    }
    if !slices.Equal(ds, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}) {
        t.Errorf("data values do not match: %v", ds)
    }
}
