package skiplist


// `FromSlice` creates a skip list where its elements are ordered by `cmp`.  The
// skip list contains the values in `ds`.
func FromSlice[Data any](cmp func(Data, Data) int, ds []Data) *Skiplist[Data] {
    r := New[Data](cmp)
    for _, d := range ds {
        r.Add(d)
    }
    return r
}

// `ToSlice` returns the values in the skip list `l` as slices.
func (l *Skiplist[Data]) ToSlice() []Data {
    ds := make([]Data, l.Len())
    for e, i := l.Front(), 0; e != nil; e = e.Next() {
        ds[i] = e.Value
        i++
    }
    return ds
}


