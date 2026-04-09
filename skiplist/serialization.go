package skiplist

import (
    "fmt"
    "encoding/json"
    "strings"
)


// `String` returns the string representation of the skip list `l`.
func (l *Skiplist[Data]) String() string {
    var b strings.Builder
    b.WriteRune('[')
    for e := l.Front(); e != nil; e = e.Next() {
        if b.Len() > 1 {
            b.WriteString(ElemSep)
        }
        b.WriteString(fmt.Sprintf("%v", e.Value))
    }
    b.WriteRune(']')
    return b.String()
}


// `MarshalJSON` returns the JSON representation of the skiplist `l` as a list in
// which the list's elements are ordered ascendingly.
func (l *Skiplist[Data]) MarshalJSON() ([]byte, error) {
    bs := []byte("[")
    var err error
    for d := range l.Ascend {
        if len(bs) > 1 {
            bs = append(bs, ',')
        }
        var ds []byte
        if ds, err = json.Marshal(d); err != nil {
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

// `UnmarshalJSON` modifies the skip list `l` into the skip list that contains
// the JSON elements of `data`.  The JSON object must be a list.  Its elements
// do not need to be ordered.
func (l *Skiplist[Data]) UnmarshalJSON(data []byte) error {
    // Unmarshal data first into a Data slice.  Not very efficient but simple to
    // implement.
    var ds []Data
    if err := json.Unmarshal(data, &ds); err != nil {
        return err
    }

    // Add slice elements to the empty skip list.
    l.Reset()
    for _, d := range ds {
        l.Add(d)
    }
    return nil
}
