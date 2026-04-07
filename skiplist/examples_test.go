package skiplist_test

import (
    "fmt"
    "strings"
    "github.com/flxch/container/skiplist"
)


func ExampleSkiplist() {
    sl := skiplist.New[string](false, strings.Compare)

    sl.Add("foo")
    sl.Add("goo")
    fmt.Println(sl)

    sl.Remove(sl.Lookup("goo"))
    sl.Add("bar")
    sl.Add("baz")
    fmt.Println(sl)

    // Output:
    // [foo,goo]
    // [bar,baz,foo]
}
