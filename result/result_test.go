package result_test

import (
    "fmt"
    "testing"
    "github.com/flxch/container/result"
)


// Examples.

func ExampleResult() {
    div := func(n, m int) (int, error) {
        if m == 0 {
            return 0, fmt.Errorf("division by zero")
        }
        return n / m, nil
    }

    r := result.Wrap(div(8, 2))
    s := result.Wrap(div(8, 0))

    fmt.Printf("%v\n", r)
    fmt.Printf("%v\n", s)

    // Output:
    // 4
    // Err(division by zero)
}


// Tests.

func TestResult(t *testing.T) {
    r := result.Ok(4)
    s := result.Ok("foo")
    u := result.Err[int](fmt.Errorf("division by zero"))

    if !r.IsOk() || !s.IsOk() {
        t.Errorf("expected a value")
    }
    if r.IsErr() || s.IsErr() {
        t.Errorf("did not expect an error")
    }

    if !u.IsErr() {
        t.Errorf("expected an error")
    }
    if u.IsOk() {
        t.Errorf("did not expect a value")
    }

    if r.Value() != 4 || s.Value() != "foo" {
        t.Errorf("values do not match")
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

func TestWrap(t *testing.T) {
    div := func(n, m int) (int, error) {
        if m == 0 {
            return 0, fmt.Errorf("division by 0")
        }
        return n / m, nil
    }

    r := result.Wrap(div(8, 2))
    if k, err := r.Unwrap(); err != nil || k != 4 {
        t.Errorf("failed to unwrap result %s", r)
    }
    s := result.Wrap(div(8, 0))
    if _, err := s.Unwrap(); err == nil {
        t.Errorf("failed to unwrap result %s", r)
    }
}
