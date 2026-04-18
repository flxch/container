// Go package for balanced search trees based on left-leaning red-black trees
// (2-3 balanced search trees).

// This package is based on Petar Maymounkov's llrb Go package
// (github.com/petar/GoLLRB/llrb).  The main change to the llrb package is that
// this package uses generics, which is supported since Go 1.18.  This package
// also provides range iteraators, which were introduced in Go 1.23.  Names for
// methods and the like have also been changed.  Overall, this package provides
// a simpler and cleaner API than the llrb package.  Finally, this package is
// optimized for speed and memory usage.  Benchmarks show that this package
// performs better than the llrb package.

// Trees are created by the New() function.  Len() returns the number of the
// stored data values in a tree and Reset() empties the tree.  The Clone()
// operation makes a copy of a tree.  The dictionary operations on a tree are
// Add(), Remove(), and Lookup().  Finally, iterators like AscendGeq() and
// DescendLeq() iterate over the data values in a tree in ascending or
// descending order, respectively.

// Some slight performance gains can be achieved when separating the data items
// at the nodes into keys and values, where a key is a generic type that
// satisfies the interface cmp.Ordered and value is any generic type.  Instead
// of a compare function (provided when calling the New function for creating a
// tree) for ordering data items, keys are directly compared by the operators <,
// >, and ==.  This is slighly faster (but also less general).  I made a quick
// test and the performance improvements are fairly small.  It seems that
// trading generality for a slight performance improvement is not worth the
// effort.

package tree

