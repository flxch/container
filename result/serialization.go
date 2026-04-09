package result

import (
    "fmt"
)


// `String` returns the string representation of the result `r`.
func (r Result[T]) String() string {
    if r.IsErr() {
        return fmt.Sprintf("Err(%v)", r.err)
    }
    return fmt.Sprintf("%v", r.ok)
}
