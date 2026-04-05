package pair

import (
    "fmt"
)


// `String` provides the string representation of the pair `p`.
func (p Pair[T, U]) String() string {
    return fmt.Sprintf("(%v,%v)", p.Fst(), p.Snd())
}
