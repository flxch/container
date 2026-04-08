package skiplist

import (
    "math/rand/v2"
)


var (
    // `ElemSep` specifies the string that separates data items when converting a
    // skip list into a string by the `String()` method.
    ElemSep          string = ","

    // Internally, the elements in a skip list have a height, which is randomly
    // chosen for each element.  `DefaultMaxHeight` is the default maximal
    // height.  It is globally used when initializing a skip list.
    // The value 32 seems to work well for skip lists with up to 1,000,000
    // elements.  If the skip list has significantly more than 1,000,000
    // elements, a larger maximal height might work better, either by setting
    // `DefaultMaxHeight` to another value or using the methoid `SetHeight` to
    // change a skip list's maximal element height.
    DefaultMaxHeight int = 32

    // As the height of internal elements are randomly chosen, one can set the
    // seed for generating the random height values.  The default values are
    // used when initializing a skip list.  Alternatively, use the method
    // `SetSeed` to change a skip list's seeds.
    DefaultSeed1     uint64 = 1
    DefaultSeed2     uint64 = 2

    // Preallocation of large memory junks for elements that might be added to a
    // skiplist in the future can result in some performance gains.  The reason
    // is that allocating a big junk of memory at once is less expensive than
    // allocating many times small memory junks.  However, the running times of
    // some operations become less predicable.  In the rare case, when they
    // allocate a big memory junk, they take more time.  `DefaultPrealloc`
    // provides a "factor" of preallocating memory for the elements of a skip
    // list.  A value smaller than or equal to 0 means that no memory is
    // preallocated.  Use the method `SetPrealloc` to the preallocation factor
    // for a skip list.
    DefaultPrealloc  int = -1
)


// `SetSeed` changes the seed for the pseudo-random generator for the height of
// new elements to `s`.
func (l *Skiplist[Data]) SetSeed(s, t uint64) {
    l.rand = rand.New(rand.NewPCG(s, t))
}

// `SetPreall` preallocates a large junk of memory for skip list elements, if n
// is positive.
func (l *Skiplist[Data]) SetPrealloc(n int) {
    if l.prealloc = n; l.prealloc > 0 {
        l.preelems = make([]neighbors[Data], max(l.Height() << 2, l.Height() << l.prealloc))
    }
}
