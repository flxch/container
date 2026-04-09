package tree

import (
    "fmt"
    "encoding/json"
    "strings"
)


// `ElemSep` specifies the string that separates data items when converting a
// tree into a string by the `String()` method.
var ElemSep string = ","

// `String` returns the string representation of the tree `t`.  The tree
// elements are ordered ascendingly.
func (t *Tree[Data]) String() string {
    var b strings.Builder
    b.WriteByte('[')

    t.WalkAscend(func(d Data) {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v", d))
    })

    b.WriteByte(']')
    return b.String()
}

// `MarshalJSON` returns the JSON representation of the tree `t` as a list in
// which the tree elements are ordered ascendingly.
func (t *Tree[Data]) MarshalJSON() ([]byte, error) {
    min, ok := t.Minimum()
    if !ok {
        return []byte("[]"), nil
    }

    bs := []byte("[")

    var err error
    t.WalkAscendGeq(min, func(d Data) bool {
        if len(bs) > 1 {
            bs = append(bs, ',')
        }
        var ds []byte
        ds, err = json.Marshal(d)
        if err != nil {
            // Stop tree iteration.
            return false
        }
        bs = append(bs, ds...)
        return true
    })

    if err != nil {
        // Append the null JSON value in case of an error.
        bs = append(bs, []byte("null")...)
    }
    bs = append(bs, ']')
    return bs, err
}

// `UnmarshalJSON` modifies the tree `t` into the tree that contains the JSON
// elements of `data`.  The JSON object must be a list.  Its elements do not
// need to be ordered.
func (t *Tree[Data]) UnmarshalJSON(data []byte) error {
    // Unmarshal data first into a Data slice.  Not very efficient but simple to
    // implement.
    var ds []Data
    if err := json.Unmarshal(data, &ds); err != nil {
        return err
    }

    // Add slice elements to the empty tree.
    t.Reset()
    for _, d := range ds {
        t.Add(d)
    }
    return nil
}
