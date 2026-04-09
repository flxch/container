// Go package for generic skiplists.

// Under development.

package skiplist

// TODOs
// * Change API (avoid exporting the element type, unify with tree package).
// * Allow skiplists to contain equal data values (insert versus add).
// * Add examples (for various skiplist methods and the iterators).
// * Add more tests.
// * Add JSON marshaling (serialization.go).
// * Add pool for elements (optional when creating skip list).
//   However, we should first do some profiling to see we can gain a speedup
//   with a pool.  Note that we only allocate memory when adding a new element
//   to a skip list.

