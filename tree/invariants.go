package tree

import (
    "fmt"
    "math"
    "slices"
    "github.com/flxch/container/option"
)


// `IsLlrbTree` returns nil if the tree `t` satisfies all the invariants of a
// left-leaning red-black tree.  If one the invariants is violated, an error is
// returned.  `IsLlrbTree` should always return nil, since all exported tree
// operations (namely, `Insert` and `Remove`) from the tree package should
// maintain the invariants.  This function is useful for debugging.
func (t *Tree[Data]) IsLlrbTree() error {
    if !t.root.isBlack() {
        fmt.Errorf("tree has red node")
    }
    if err := t.noDuplicates(); err != nil {
        fmt.Errorf("tree has duplicates: %w", err)
    }
    if err := t.root.isSearchTree(t.compare, option.None[Data](), option.None[Data]()); err != nil {
        fmt.Errorf("no search tree: %w", err)
    }
    if err := t.root.noConsecutiveReds(); err != nil {
        fmt.Errorf("tree has consecutive red nodes: %w", err)
    }
    if err := t.root.isLeftLeaning(); err != nil {
        fmt.Errorf("tree is not left leaning: %w", err)
    }
    if err := t.isBlackBalanced(); err != nil {
        fmt.Errorf("tree is not balanced: %w", err)
    }
    return nil
}


func (t *Tree[Data]) noDuplicates() error {
    visited := make([]Data, 0, t.Len())
    for d := range t.Ascend {
        visited = append(visited, d)
    }
    slices.SortFunc(visited, t.compare)
    for i := 1; i < len(visited); i++ {
        if t.compare(visited[i - 1], visited[i]) == 0 {
            return fmt.Errorf("multiple nodes with value %v", visited[i])
        }
    }
    return nil
}


func (n *node[Data]) isSearchTree(cmp func(Data, Data) int, min, max option.Option[Data]) error {
    if n == nil {
        return nil
    }

    if min.IsSome() && cmp(n.data, min.Value()) == -1 {
        return fmt.Errorf("value %v in right subtree too small (%v) ", n.data, min.Value())
    }
    if max.IsSome() && cmp(n.data, max.Value()) == 1 {
        return fmt.Errorf("value %v in left subtree too big (%v) ", n.data, max.Value())
    }

    if err := n.left.isSearchTree(cmp, min, option.Some(n.data)); err != nil {
        return err
    }
    if err := n.right.isSearchTree(cmp, option.Some(n.data), max); err != nil {
        return err
    }
    return nil
}


func (n *node[Data]) noConsecutiveReds() error {
    if n == nil {
        return nil
    }

    if n.isRed() && n.left != nil && n.left.isRed() {
        return fmt.Errorf("red node %v with red left child %v", n.data, n.left.data)
    }

    if err := n.left.noConsecutiveReds(); err != nil {
        return err
    }
    if err := n.right.noConsecutiveReds(); err != nil {
        return err
    }
    return nil
}


func (n *node[Data]) isLeftLeaning() error {
    if n == nil {
        return nil
    }

    if n.right != nil && n.right.isRed() {
        return fmt.Errorf("node %v with red right child %v", n.data, n.left.data)
    }

    if err := n.left.isLeftLeaning(); err != nil {
        return err
    }
    if err := n.right.isLeftLeaning(); err != nil {
        return err
    }
    return nil
}


func (t *Tree[Data]) isBlackBalanced() error {
    paths := make(map[*node[Data]]int, t.Len())
    t.root.computePathLengths(paths, 0)
    if len(paths) == 0 {
        return nil
    }

    m := math.MaxInt
    n := math.MinInt
    for _, c := range paths {
        m = min(m, c)
        n = max(n, c)
    }
    if m != n {
        return fmt.Errorf("min path length is %d and max path length is %d", m, n)
    }

    return nil
}

func (n *node[Data]) computePathLengths(paths map[*node[Data]]int, c int) {
    if n != nil {
        if n.isBlack() {
            c++
        }

        if n.left == nil && n.right == nil {
            // Store path length from root to leaf.
            paths[n] = c
        } else {
            // Continue with children.
            n.left.computePathLengths(paths, c)
            n.right.computePathLengths(paths, c)
        }
    }
}

