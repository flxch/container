package option

import (
    "fmt"
    "encoding/json"
)


// `String` returns the string version of the option `v`.
func (v Option[T]) String() string {
    if v.IsNone() {
        return "None"
    }
    return fmt.Sprintf("Some(%v)", v.Value())
}


// `MarshalJSON` marshals the option `v` into a JSON object.
func (v Option[T]) MarshalJSON() ([]byte, error) {
    if !v.IsNone() {
        return []byte("null"), nil
    }
    return json.Marshal(v.Value())
}

// `UnmarshalJSON` unmarshals the JSON object `data` and save the value into
// `v`.
func (v *Option[T]) UnmarshalJSON(data []byte) error {
    // TODO: Compare if data is null.  If this is the case v should be None.
    if err := json.Unmarshal(data, &v.value); err != nil {
        return err
    }
    v.some = true
    return nil
}

