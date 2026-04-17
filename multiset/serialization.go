package multiset

import (
    "fmt"
    "encoding/json"
    "strings"
)


// `ElemSep` specifies the string that separates elements when converting a
// multiset into a string by the `String()` method.
var ElemSep string = ","


// `String` returns the string representation of the multiset `S`.
func (S *Multiset[A]) String() string {
    var b strings.Builder
    b.WriteRune('{')
    for e, n := range S.elems {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v#%d", e, n))
    }
    b.WriteRune('}')
    return b.String()
}


// `MarshalJSON` returns the JSON representation of the multiset `S` as a
// list. An element is listed multiple time If its multiplicity is greater than
// 1.  There are not guarantees on the order of the multiset elements.
func (S *Multiset[A]) MarshalJSON() ([]byte, error) {
    bs := []byte("[")
    var err error
    for e, n := range S.Elems() {
        var ds []byte
        if ds, err = json.Marshal(e); err != nil {
            break
        }
        // Include the element's JSON representation as often as its
        // multiplicity.
        for i := 0; i < n; i++ {
            if len(bs) > 1 {
                bs = append(bs, ',')
            }
            bs = append(bs, ds...)
        }
    }
    if err != nil {
        // Append the null JSON value in case of an error.
        if len(bs) > 1 {
            bs = append(bs, ',')
        }
        bs = append(bs, []byte("null")...)
    }
    bs = append(bs, ']')
    return bs, err

}

// `UnmarshalJSON` modifies the multiset `S` into the multiset that contains the
// JSON elements of `data`.  The JSON object must be a list.
func (S *Multiset[A]) UnmarshalJSON(data []byte) error {
    // Unmarshal data first into a slice.  Not very efficient but simple to
    // implement.
    var ds []A
    if err := json.Unmarshal(data, &ds); err != nil {
        return err
    }

    // Add slice elements to the empty multiset..
    S.Reset()
    for _, d := range ds {
        S.Add(d, 1)
    }
    return nil
}
