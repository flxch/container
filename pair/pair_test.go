package pair_test

import (
    "fmt"
    "testing"
    "github.com/flxch/container/pair"
)


// Examples.

func ExamplePair() {
    m := make(map[pair.Pair[string, int]] string)

    m[pair.New("foo", 0)] = "bar"
    m[pair.New("goo", 1)] = "baz"

    fmt.Printf("%v", m)

    // Output:
    // map[(foo,0):bar (goo,1):baz]
}


// Tests.

func TestPair(t *testing.T) {
    p := pair.New("foo", 0)

    if p.String() != "(foo,0)" {
        t.Errorf("wrong string conversion of pair")
    }

    if p == pair.New("foo", 1) {
        t.Errorf("Expected that pairs are not equal.")
    }
    if p != pair.New("foo", 0) {
        t.Errorf("Expected that pairs are equal.")
    }

    q := p.Swap()
    if q.Fst() != p.Snd() {
        t.Errorf("Wrong first coordinate in swapped pair.")
    }
    if q.Snd() != p.Fst() {
        t.Errorf("Wrong second coordinate in swapped pair.")
    }

    if p != p.Swap().Swap() {
        t.Errorf("Swapping twice should result in an equal pair..")
    }
}

func TestCompare(t *testing.T) {
    z := pair.New[int, int](0, 0)
    x := pair.New[int, int](1, 0)
    y := pair.New[int, int](0, 1)

    if pair.Compare(z, z) != 0 {
        t.Errorf("expected that %s and %s are equal", z, z)
    }

    if pair.Compare(z, x) != -1 {
        t.Errorf("expected that %s is smaller than %s", z, x)
    }
    if pair.Compare(z, y) != -1 {
        t.Errorf("expected that %s is smaller than %s", z, y)
    }
    if pair.Compare(x, z) != 1 {
        t.Errorf("expected that %s is greater than %s", z, x)
    }
    if pair.Compare(y, z) != 1 {
        t.Errorf("expected that %s is greater than %s", z, y)
    }
}

