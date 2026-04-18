# container

Go packages for various container data types.  Most data types in the
packages here use generics.  They are complementing the container
packages in Go's standard library.

Note that the packages are under development and thus not stable.
Feedback, suggestions for improvements, and contributions are always
welcome.

## Benchmarks

Run the benchmarks with `run_benchmarks.sh`.  It measures the
dictionary operations of the container packages.  Some packages have
further benchmarks.

* Go version: go version go1.26.1 linux/amd64
* Date: 2026-04-18

```
goos: linux
goarch: amd64
pkg: github.com/flxch/container/tree
cpu: Intel(R) Core(TM) i7-2600 CPU @ 3.40GHz
```

### Trees

```
BenchmarkTree_Lookup/100         	12640875	        96.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkTree_Lookup/1000        	 9144016	       134.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkTree_Lookup/10000       	 7120966	       173.6 ns/op	       0 B/op	       0 allocs/op

BenchmarkTree_Add/100         	 3386713	       353.2 ns/op	      32 B/op	       1 allocs/op
BenchmarkTree_Add/1000        	 2690682	       439.6 ns/op	      32 B/op	       1 allocs/op
BenchmarkTree_Add/10000       	 2004470	       617.8 ns/op	      32 B/op	       1 allocs/op

BenchmarkTree_Remove/100         	 4112431	       295.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkTree_Remove/1000        	 2881444	       405.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkTree_Remove/10000       	 1988073	       587.1 ns/op	       0 B/op	       0 allocs/op

BenchmarkTree_Clone/100         	  451818	      2759 ns/op	    3480 B/op	       2 allocs/op
BenchmarkTree_Clone/1000        	   52746	     23074 ns/op	   32792 B/op	       2 allocs/op
BenchmarkTree_Clone/10000       	    3345	    378581 ns/op	  327704 B/op	       2 allocs/op
```

#### Comparison with llrb Package

Source: [github.com/petar/GoLLRB/llrb](https://pkg.go.dev/github.com/petar/GoLLRB/llrb)
```
BenchmarkLlrb_Lookup/100         	 4447273	       266.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkLlrb_Lookup/1000        	 3717766	       321.4 ns/op	      16 B/op	       1 allocs/op
BenchmarkLlrb_Lookup/10000       	 2938110	       401.2 ns/op	      16 B/op	       1 allocs/op

BenchmarkLlrb_Add/100            	 2508259	       493.3 ns/op	      64 B/op	       2 allocs/op
BenchmarkLlrb_Add/1000           	 2003164	       593.6 ns/op	      64 B/op	       2 allocs/op
BenchmarkLlrb_Add/10000          	 1515982	       818.0 ns/op	      64 B/op	       2 allocs/op

BenchmarkLlrb_Remove/100         	 2508806	       456.0 ns/op	      16 B/op	       1 allocs/op
BenchmarkLlrb_Remove/1000        	 1817882	       643.2 ns/op	      16 B/op	       1 allocs/op
BenchmarkLlrb_Remove/10000       	 1275520	       961.1 ns/op	      16 B/op	       1 allocs/op
```

### Skip Lists

```
BenchmarkSkiplist_Lookup/100         	 8277027	       169.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSkiplist_Lookup/1000        	 5257615	       230.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSkiplist_Lookup/10000       	 3533071	       315.8 ns/op	       0 B/op	       0 allocs/op

BenchmarkSkiplist_Add/100         	 2997916	       394.2 ns/op	      76 B/op	       1 allocs/op
BenchmarkSkiplist_Add/1000        	 2385184	       498.0 ns/op	      79 B/op	       1 allocs/op
BenchmarkSkiplist_Add/10000       	 1780845	       683.1 ns/op	      79 B/op	       1 allocs/op

BenchmarkSkiplist_Remove/100         	 7020540	       177.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSkiplist_Remove/1000        	 4274025	       266.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSkiplist_Remove/10000       	 3230275	       367.5 ns/op	       0 B/op	       0 allocs/op

BenchmarkSkiplist_Clone/100         	   79369	     15584 ns/op	   18240 B/op	     106 allocs/op
BenchmarkSkiplist_Clone/1000        	    6622	    192717 ns/op	  221184 B/op	    1006 allocs/op
BenchmarkSkiplist_Clone/10000       	     291	   4226021 ns/op	 2807680 B/op	   10006 allocs/op
```

### Multisets

```
BenchmarkMultiset_Lookup/100         	16841822	        78.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkMultiset_Lookup/1000        	18149518	        65.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkMultiset_Lookup/10000       	15787074	        71.81 ns/op	       0 B/op	       0 allocs/op

BenchmarkMultiset_Add/100         	12578094	        96.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkMultiset_Add/1000        	14574499	        82.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkMultiset_Add/10000       	13339276	        91.38 ns/op	       0 B/op	       0 allocs/op

BenchmarkMultiset_Remove/100         	10635895	       118.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkMultiset_Remove/1000        	12661615	        98.35 ns/op	       0 B/op	       0 allocs/op
BenchmarkMultiset_Remove/10000       	11515209	       102.2 ns/op	       0 B/op	       0 allocs/op

BenchmarkMultiset_Clone/100         	  773154	      1724 ns/op	    2408 B/op	       5 allocs/op
BenchmarkMultiset_Clone/1000        	  133608	      9053 ns/op	   37008 B/op	       7 allocs/op
BenchmarkMultiset_Clone/10000       	   12285	    100183 ns/op	  295616 B/op	      35 allocs/op
```

### Sets

```
BenchmarkSet_Lookup/100         	18730075	        62.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet_Lookup/1000        	20370180	        63.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet_Lookup/10000       	19140957	        64.31 ns/op	       0 B/op	       0 allocs/op

BenchmarkSet_Add/100         	16835692	        77.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet_Add/1000        	15802293	        76.14 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet_Add/10000       	13839002	        93.74 ns/op	       0 B/op	       0 allocs/op

BenchmarkSet_Remove/100         	13805802	        81.23 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet_Remove/1000        	16370157	        74.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkSet_Remove/10000       	15275640	        77.65 ns/op	       0 B/op	       0 allocs/op

BenchmarkSet_Clone/100         	  869959	      1422 ns/op	    2392 B/op	       4 allocs/op
BenchmarkSet_Clone/1000        	  131106	      8521 ns/op	   36992 B/op	       6 allocs/op
BenchmarkSet_Clone/10000       	   13689	     89506 ns/op	  295600 B/op	      34 allocs/op
```

## TODOs

1. Implement packages for tries, heaps, ...
2. Improve documentation.
3. More thorough testing and benchmarking (also compare with other
   packages with similar container implementations).



