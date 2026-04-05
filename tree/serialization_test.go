package tree_test

import (
    "fmt"
    "testing"
    "github.com/flxch/container/tree"
)


type kvPair struct {
    Key int    `json:"key"`
    Val string `json:"val"`
}

func compareKVPair(p, q kvPair) int {
    if p.Key < q.Key {
        return -1
    } else if p.Key > q.Key {
        return 1
    }
    return 0
}

// Make the pair type an instance of the fmt.Stringer interface.  The generic
// method String() of the type Tree[kvPair] will use this instance to convert
// tree elements to strings.
func (p kvPair) String() string {
    return fmt.Sprintf(`(%d,"%s")`, p.Key, p.Val)
}

// Make the pair type an instance of the json.Marshaler interface.
func (p kvPair) MarshalJSON() ([]byte, error) {
    if p.Key == 3 {
        return nil, fmt.Errorf("fake marshaling error")
    }
    return []byte(fmt.Sprintf(`{"key":%d,"val":"%s"}`, p.Key, p.Val)), nil
}


func TestString(t *testing.T) {
    T := tree.New(compareKVPair)

    if T.String() != "[]" {
        t.Errorf("expected [] for the tree with no elements")
    }

    T.Add(kvPair{4, "foo"})
    if s, exp := T.String(), `[(4,"foo")]`; s != exp {
        t.Errorf("expected %s for the tree with a single element, not %s", exp, s)
    }

    T.Add(kvPair{6, "moo"})
    T.Add(kvPair{1, "bar"})
    T.Add(kvPair{3, "baz"})

    if s, exp := T.String(), `[(1,"bar"),(3,"baz"),(4,"foo"),(6,"moo")]`; s != exp {
        t.Errorf("expected %s, not %s", exp, s)
    }
}

func TestMarshalJSON(t *testing.T) {
    T := tree.New(compareKVPair)

    bs, err := T.MarshalJSON()
    if err != nil {
        t.Errorf("Failed to marshal empty tree: %v", err)
    } else if exp := `[]`; string(bs) != exp {
        t.Errorf("expected %s, not %s", exp, bs)
    }

    T.Add(kvPair{4, "foo"})
    bs, err = T.MarshalJSON()
    if err != nil {
        t.Errorf("Failed to marshal tree: %v", err)
    } else if exp := `[{"key":4,"val":"foo"}]`; string(bs) != exp {
        t.Errorf("expected %s, not %s", exp, bs)
    }

    T.Add(kvPair{6, "moo"})
    T.Add(kvPair{1, "bar"})

    bs, err = T.MarshalJSON()
    if err != nil {
        t.Errorf("Failed to marshal tree: %v", err)
    } else if exp := `[{"key":1,"val":"bar"},{"key":4,"val":"foo"},{"key":6,"val":"moo"}]`; string(bs) != exp {
        t.Errorf("expected %s, not %s", exp, bs)
    }

    T.Add(kvPair{3, "baz"})
    bs, err = T.MarshalJSON()
    if err == nil {
        t.Errorf("expected a marshaling error: %s", bs)
    } else if exp := `[{"key":1,"val":"bar"},null]`; string(bs) != exp {
        t.Errorf("expected %s, not %s", exp, bs)
    }
}

func TestUnmarshalJSON(t *testing.T) {
    T := tree.New(compareKVPair)
    T.Add(kvPair{4, "foo"})
    T.Add(kvPair{6, "moo"})

    err := T.UnmarshalJSON([]byte(`[]`))
    if err != nil {
        t.Errorf("failed to unmarshal tree: %v", err)
    } else if T.Len() != 0 {
        t.Errorf("expected the empty tree")
    }

    err = T.UnmarshalJSON([]byte(`[{"key":4,"val":"foo"}]`))
    if err != nil {
        t.Errorf("failed to unmarshal tree: %v", err)
    } else if T.Len() != 1 {
        t.Errorf("expected a tree with a single element")
    }

    err = T.UnmarshalJSON([]byte(`[{"key":4,"val":"foo"},{"key":1,"val":"bar"},{"key":6,"val":"moo"}]`))
    if err != nil {
        t.Errorf("failed to unmarshal tree: %v", err)
    } else if T.Len() != 3 {
        t.Errorf("expected a tree with three elements")
    }
}
