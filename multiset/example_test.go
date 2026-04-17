package multiset_test

import (
    "fmt"
    "github.com/flxch/container/multiset"
)


func ExampleMultiset() {
    // Create an empty multiset.
    S := multiset.New[string]()

    // Add some elements to the multiset.
    S.Add("foo", 2)
    S.Add("bar", 1)
    S.Add("baz", 1)

    // Iterate through the multiset's elements.
    for e, m := range S.Elements() {
        // Note that there is no order on the multiset elements.  However, since
        // only one element has a multiplicity greater than 1 in S, the output
        // is always the string "foo" here.
        if m > 1 {
            fmt.Println(e)
        }
    }

    // Output:
    // foo
}



