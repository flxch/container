package stack

import (
    "fmt"
    "strings"
)


// `ElemSep` specifies the string that separates stack elements when converting a
// stack into a string by the `String()` method.
var ElemSep string = ","

// `String` returns a string representation of the stack `s`.  This function can
// be useful for debugging.
func (s *Stack[Data]) String() string {
    var b strings.Builder
    b.WriteByte('[')

    for i := 0; i < s.top; i++ {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v", s.elems[i]))
    }

    b.WriteByte(']')
    return b.String()
}
