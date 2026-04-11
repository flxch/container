// Go package for generic skip lists.

// Under development.

package skiplist

// TODOs
// * Change API (avoid exporting the element type, unify with tree package). (?)
// * Change names (also tree package) (?)
//   - Add -> Insert
//   - Remove -> Delete
// * Add pool for elements (optional when creating skip list).  However, we
//   should first do some profiling to see we gain a speedup with a pool.  Note
//   that we only allocate memory when adding a new element to a skip list.
// * Compare with other skip list packages, in particular, compare benchmarks.
