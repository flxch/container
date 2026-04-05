package result

type Result[T any] struct {
    ok  T
    err error
}


func Ok[T any](val T) Result[T] {
    return Result[T]{ok: val}
}

func Err[T any](err error) Result[T] {
    return Result[T]{err: err}
}


func (r Result[T]) IsOk() bool {
    return r.err == nil
}

func (r Result[T]) IsErr() bool {
    return r.err != nil
}


func Wrap[T any](val T, err error) Result[T] {
    return Result[T]{ok: val, err: err}
}

func Unwrap[T any](r Result[T]) (T, error) {
    return r.ok, r.err
}

func (r Result[T]) Unwrap() (T, error) {
    return r.ok, r.err
}


func (r Result[T]) Value() T {
    if r.err != nil {
        panic("result contains an error")
    }
    return r.ok
}

func (r Result[T]) Error() error {
    if r.err == nil {
        panic("result contains no error")
    }
    return r.err
}



