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


// `MarshalJSON` returns the JSON representation of the multiset `S` as a list.
// There are not guarantees on the order of the multiset elements.
func (S *Multiset[A]) MarshalJSON() ([]byte, error) {
    bs := []byte("[")
    var err error
    for e, n := range S.Elems() {
        if len(bs) > 1 {
            bs = append(bs, ',')
        }
        var ds []byte
        if ds, err = json.Marshal(e); err != nil {
            break
        }
        bs = append(bs, []byte(fmt.Sprintf(`{"element":%s,"multiplicity":%d}`, ds, n))...)
    }
    if err != nil {
        // Append the null JSON value in case of an error.
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
    var ds []struct{
        Element      A   `json:"element"`
        Multiplicity int `json:"multiplicity"`
    }
    if err := json.Unmarshal(data, &ds); err != nil {
        return err
    }

    // Add slice elements to the empty skip list.
    S.Reset()
    for _, d := range ds {
        for i := 0; i < d.Multiplicity; i++ {
            S.Add(d.Element)
        }
    }
    return nil
}
