package option_test

import (
    "fmt"
    "testing"
    "github.com/flxch/container/option"
)


// Examples.

func ExampleOption() {
    u := option.Some("foo")
    v := option.Some(1)
    w := option.None[int]()

    fmt.Printf("%s, %s, %s\n", u, v, w)

    // Return values from functions.
    kvstore := map[string]int{"foo": 0, "moo": 1}

    // Unfortunately, the assignment `u := Wrap[int](kvstore["bar"])` is ill
    // typed.  The return value of the lookup kvpairs["bar"] is an int and not
    // and int and a bool.  (Maybe this will change in later Go versions and
    // both types are possible.)

    // With the auxiliary function lookup below, the return type is the pair of
    // an int and a bool that can wrapped into an Option.  (Similarly, a type
    // cast like `u := Wrap(v.(int))` is ill typed since v.(int) is not a pair
    // of an int and a bool.)
    lookup := func(s string) (int, bool) { n, ok := kvstore[s]; return n, ok }
    // Go does not allow type parameters here.  A generic lookup function could
    // be defined as follows:
    // func lookup[K comparable, V any](m map[K]V, k K) (V, bool) {
    //     v, ok := m[k]
    //     return v, ok
    // }

    // Wrap return value.
    x := option.Wrap(lookup("bar"))
    y := option.Wrap(lookup("foo"))

    fmt.Printf("bar -> %s\n", x)
    fmt.Printf("foo -> %s\n", y)

    // Output:
    // Some(foo), Some(1), None
    // bar -> None
    // foo -> Some(0)
}


// Tests.

func TestOptional_String(t *testing.T) {
    if u := option.Some(1); u.String() != "Some(1)" {
        t.Errorf("expected Some(1), not %s", u)
    }
    if u := option.Some("foo"); u.String() != "Some(foo)" {
        t.Errorf("expected Some(foo), not %s", u)
    }
    if u := option.None[string](); u.String() != "None" {
        t.Errorf("expected None, not %s", u)
    }
}

func TestOptional_Some(t *testing.T) {
    u := option.Some(1)

    if !u.IsSome() {
        t.Errorf("expected some value")
    } else if u.IsNone() {
        t.Errorf("did not expect no value")
    }

    if u.Value() != 1 {
        t.Errorf("expected the value 1, not %d", u.Value())
    }

    if val, ok := u.Unwrap(); !ok || val != 1 {
        t.Errorf("expected the valid value 1, not %d, %t", val, ok)
    }
}

func TestOptional_None(t *testing.T) {
    u := option.None[int]()

    if !u.IsNone() {
        t.Errorf("expected no value")
    } else if u.IsSome() {
        t.Errorf("did not expect some valued")
    }

    if _, ok := u.Unwrap(); ok {
        t.Errorf("did not expect a valid value")
    }

    // Must be the last check in this test.
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("exptected a panic")
        }
    }()
    // Panic and check in the defer function whether the panic happens.
    u.Value()
}

func TestOptional_Zero(t *testing.T) {
    u := option.Zero[int]()

    var n int
    if u.IsNone() {
        t.Errorf("expected some value, not none")
    } else if !u.IsSome() {
        t.Errorf("expected some value")
    } else if u.Value() != n {
        t.Errorf("expected the zero value, not %d", u.Value())
    }
}
