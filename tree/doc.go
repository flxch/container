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


// Benchmarks
// ----------

// The following benchmarks provide some rough estimate about the performance
// inmprovements of this package of the llrb package.  Please note that the
// exact numbers are most likely not up to date.  Some changes to the package
// might have been done in the meanwhile.  Furthermore, using a more recent Go
// version might give different measurements.  The differences should be small
// though.

// go version go1.25.7 linux/amd64

// > go test -run=XXX -bench=Tree -cpu=1
// goos: linux
// goarch: amd64
// pkg: github.com/flxch/container/tree
// cpu: Intel(R) Core(TM) i7-3667U CPU @ 2.00GHz
// BenchmarkTree_Add/100         	 3327306	       356.4 ns/op	      32 B/op	       1 allocs/op
// BenchmarkTree_Add/1000        	 2468887	       494.3 ns/op	      32 B/op	       1 allocs/op
// BenchmarkTree_Add/10000       	 1971072	       615.8 ns/op	      32 B/op	       1 allocs/op
// BenchmarkTree_Remove/100      	 4494888	       275.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTree_Remove/1000     	 2888192	       411.5 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTree_Remove/10000    	 1985682	       596.6 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTree_Lookup/100      	10766779	       114.9 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTree_Lookup/1000     	 7410687	       158.0 ns/op	       0 B/op	       0 allocs/op
// BenchmarkTree_Lookup/10000    	 6030325	       200.3 ns/op	       0 B/op	       0 allocs/op

// Some slight performance gains can be achieved when separating the data items
// at the nodes into keys and values, where a key is a generic type that
// satisfies the interface cmp.Ordered and value is any generic type.  Instead
// of a compare function (provided when calling the New function for creating a
// tree) for ordering data items, keys are directly compared by the operators <,
// >, and ==.  This is slighly faster (but also less general).  I made a quick
// test and the performance improvements are fairly small.  It seems that
// trading generality for a slight performance improvement is not worth the
// effort.

// The corresponding methods for the dictonary operations `Add`, `Remove`, and
// `Lookup` of the llrb package are slower.  Furthermore, the bst package avoids
// memory allocation by the use of generics.  In contrast, the llrb package uses
// interfaces for determining the order of tree elements.

// > go test -run=XXX -bench=Llrb -cpu=1
// oos: linux
// goarch: amd64
// pkg: github.com/flxch/container/tree
// cpu: Intel(R) Core(TM) i7-3667U CPU @ 2.00GHz
// BenchmarkLlrb_Add/100         	 2405500	       489.4 ns/op	      64 B/op	       2 allocs/op
// BenchmarkLlrb_Add/1000        	 2005843	       615.1 ns/op	      64 B/op	       2 allocs/op
// BenchmarkLlrb_Add/10000       	 1568086	       782.6 ns/op	      64 B/op	       2 allocs/op
// BenchmarkLlrb_Remove/100      	 2520348	       484.5 ns/op	      16 B/op	       1 allocs/op
// BenchmarkLlrb_Remove/1000     	 1875494	       627.6 ns/op	      16 B/op	       1 allocs/op
// BenchmarkLlrb_Remove/10000    	 1333600	       885.0 ns/op	      16 B/op	       1 allocs/op
// BenchmarkLlrb_Lookup/100      	 4326193	       282.0 ns/op	      16 B/op	       1 allocs/op
// BenchmarkLlrb_Lookup/1000     	 3517029	       351.1 ns/op	      16 B/op	       1 allocs/op
// BenchmarkLlrb_Lookup/10000    	 2805577	       426.6 ns/op	      16 B/op	       1 allocs/op

package tree

