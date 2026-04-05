package option

// The type `Option` either varries a value of type `T` (i.e., some value) or no
// value (i.e., none).
type Option[T any] struct {
    value T
    some  bool
}


// `None` returns an option with no value.
func None[T any]() Option[T] {
    return Option[T]{}
}

// `Some` returns an option with the value `v`.
func Some[T any](v T) Option[T] {
    return Option[T]{value: v, some: true}
}

// `Zero` returns an option with the zero value of the type T.
func Zero[T any]() Option[T] {
    return Option[T]{some: true}
}


// `IsNone` returns true if `v` carries no value.
func (v Option[T]) IsNone() bool {
    return !v.some
}

// `IsSome` returns true if `v` carries a value.
func (v Option[T]) IsSome() bool {
    return v.some
}

// `IsZero` returns true if `v` carries a value and this value is zero of the
// type T.
//func (v Option[T]) IsZero() bool {
//    var z T
//    return v.some && v.value == z
//}


// `Wrap` returns an option of type T with value `v` if `some` is true.
// Otherwise, `Wrap` returns an option with no value.
// `Wrap` is useful for functions or expressions that return a pair (T, bool),
// which often the case in Go and is often used by Go programmers.  For example,
// a lookup in a map of a type cast returs such a pair.  `Wrap` can be used to
// convert the pair into a value fo the Option type.
func Wrap[T any](v T, some bool) Option[T] {
    return Option[T]{value: v, some: some}
}

// `Unwrap` is the counterpart of the `Wrap` function.  It returns the value and
// the boolean true if the option `v` carries a value of type T.  Otherwise, if
// `v` carries no value, te returned boolean value is false.
func Unwrap[T any](v Option[T]) (T, bool) {
    return v.value, v.some
}

// `Unwrap` is the correspondig method for the function `Unpack`.
func (v Option[T]) Unwrap() (T, bool) {
    return v.value, v.some
}

// `Value` returns the value of `v` if `v` carries a value.  Otherwise, `Value`
// panics.
func (v Option[T]) Value() T {
    if !v.some {
        panic("no value in option")
    }
    return v.value
}


// `Equal` lifts the equality of type T to Option[T].  Note that
// if `u` or `v` carries no value then `u` is not equal to `v`.
func Equal[T comparable](u, v Option[T]) bool {
    if u.IsNone() || v.IsNone() {
        return false
    }
    return u.Value() == v.Value()
}

// `LiftOp` lifts the operation `op` to the operation on Option[T].
func LiftOp[T any](op func(T, T) T) func(Option[T], Option[T]) Option[T] {
    return func(u, v Option[T]) Option[T] {
        if u.IsNone() || v.IsNone() {
            return None[T]()
        }
        return Some[T](op(u.Value(), v.Value()))
    }
}

