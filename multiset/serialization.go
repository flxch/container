package multiset

import (
    "fmt"
    "strings"
)


// `ElemSep` specifies the string that separates elements when converting a
// multiset into a string by the `String()` method.
var ElemSep string = ","


// `String` returns the string representation of the multiset `S`.
func (S *Multiset[A]) String() string {
    var b strings.Builder
    b.WriteRune('{')
    for e, n := range S.elems {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v#%d", e, n))
    }
    b.WriteRune('}')
    return b.String()
}


//func (S *Multiset[A]) MarshalJSON() ([]byte, error) {
//}

//func (S *Multiset[A]) UnmarshalJSON(data []byte) error {
//}
