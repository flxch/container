package set

import (
    "fmt"
    "encoding/json"
    "strings"
)


// `ElemSep` specifies the string that separates elements when converting a
// set into a string by the `String()` method.
var ElemSep string = ","


// `String` returns the string representation of the set `S`.
func (S Set[A]) String() string {
    var b strings.Builder
    b.WriteRune('{')
    for e := range S {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v", e))
    }
    b.WriteRune('}')
    return b.String()
}


// `MarshalJSON` returns the JSON representation of the set `S` as a list.
// There are no guarantees on the order of the set elements.
func (S Set[A]) MarshalJSON() ([]byte, error) {
    bs := []byte("[")
    var err error
    for e := range S {
        if len(bs) > 1 {
            bs = append(bs, ',')
        }

        var ds []byte
        if ds, err = json.Marshal(e); err != nil {
            break
        }
        bs = append(bs, ds...)
    }
    if err != nil {
        // Append the null JSON value in case of an error.
        bs = append(bs, []byte("null")...)
    }
    bs = append(bs, ']')
    return bs, err

}

// `UnmarshalJSON` modifies the set `S` into the set that contains the JSON
// elements of `data`.  The JSON object must be a list.
func (S Set[A]) UnmarshalJSON(data []byte) error {
    // Unmarshal data first into a slice.  Not very efficient but simple to
    // implement.
    var ds []A
    if err := json.Unmarshal(data, &ds); err != nil {
        return err
    }

    // Add slice elements to the empty set.
    S.Reset()
    for _, d := range ds {
        S.Add(d)
    }
    return nil
}
