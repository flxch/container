package skiplist

import (
    "fmt"
    "strings"
)


// `String` returns the string representation of the skip list `l`.
func (l *Skiplist[Data]) String() string {
    var b strings.Builder
    b.WriteRune('[')
    for e := l.Front(); e != nil; e = e.Next() {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v", e.Value))
    }
    b.WriteRune(']')
    return b.String()
}
