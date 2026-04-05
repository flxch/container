package result

import (
    "fmt"
)


func (r Result[T]) String() string {
    if r.IsErr() {
        return fmt.Sprintf("Err(%v)", r.err)
    }
    return fmt.Sprintf("%v", r.ok)
}
