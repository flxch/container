package pair

import (
    "fmt"
)


// `ElemSep` specifies the string that separates the coordinates when converting
// a pair into a string by the `String()` method.
var ElemSep string = ","

// `String` returns the string representation of the pair `p`.
func (p Pair[T, U]) String() string {
    return fmt.Sprintf("(%v%s%v)", p.Fst(), ElemSep, p.Snd())
}
