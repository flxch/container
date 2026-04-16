package multiset_test

import (
    "fmt"
    "github.com/flxch/container/multiset"
)


func ExampleMultiset() {
    // Create an empty multiset.
    S := multiset.New[string]()

    // Add elements to the multiset.
    S.Add("foo")
    S.Add("foo")
    S.Add("bar")
    S.Add("baz")

    // Iterate through multiset elements.
    for e, n := range S.All() {
        // Note that the order is not guaranteed.  However, since only one
        // element has a multiplicity greater than 1, the output is always the
        // string "foo".
        if n > 1 {
            fmt.Println(e)
        }
    }

    // Output:
    // foo
}



