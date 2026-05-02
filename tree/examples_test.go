package tree_test

import (
    "fmt"
    "github.com/flxch/container/tree"
)


// `ExampleTree_Add`` demonstrates the use of a binary search tree with integer
// values with the standard ordering on integers.  The example creates a tree,
// adds two elements to it, and prints the elements in ascending order.
func ExampleTree_Add() {
    // Create an integer tree.
    t := tree.New(func(k, l int) int {
        switch {
        case k < l: // less than
            return -1
        case k > l: // greater than
            return 1
        default:    // equal to
            return 0
        }
    })

    // Add two elements to the tree.
    t.Add(10)
    t.Add(20)

    // Print the elements in ascending order.
    t.WalkAscend(func(elem int) { fmt.Println(elem) })

    // Output:
    // 10
    // 20
}

// `ExampleTree_WalkAscend` demonstrates the range iterators for trees.
func ExampleTree_Ascend() {
    // Create an integer tree.
    t := tree.New(func(k, l int) int {
        switch {
        case k < l: // less than
            return -1
        case k > l: // greater than
            return 1
        default:    // equal to
            return 0
        }
    })

    // Add two elements to the tree.
    t.Add(10)
    t.Add(20)
    t.Add(0)

    // Walk through the tree in ascending order.
    for elem := range t.Ascend {
        fmt.Println(elem)
    }

    // Output:
    // 0
    // 10
    // 20
}

