package set_test

import (
    "fmt"
    "github.com/flxch/container/set"
)


func ExampleSet() {
    // Create an empty set.
    S := set.New[string]()

    // Add some elements to the multiset.
    S.Add("foo")
    S.Add("bar")
    S.Add("baz")

    // Iterate through the set elements.
    for e := range S {
        // Note that there is no order on the set elements.
        if e[0] != 'b' {
            fmt.Println(e)
        }
    }

    // Output:
    // foo
}



