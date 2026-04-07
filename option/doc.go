// Go package providing the generic type Option[T], which is inspired by the
// Haskell abstract data type Maybe a (or in most other functional programming
// language where it might have a different name).  Option[T] either carries
// some value of type T or none.  It can, for example, be used to wrap return
// values (T, bool), which are often used in Go (e.g., map lookups and type
// casts), where the boolean signals whether the T value is valid.

package option



