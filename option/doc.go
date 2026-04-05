package option

// Go package with the generic type Option[T].  It is inspired by the Haskell
// abstract data type Maybe a.  Option[T] either carries some value of type T or
// none.  It can be used to wrap return values (T, bool), which are often used
// in Go (e.g., map lookups and type casts), where the boolean signals whether
// the T value is valid.


