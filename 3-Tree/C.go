package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Index    int
	Visited  bool
	Parents  *Node
	Children []*Node
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func dfs(node *Node, u, v int) *Node {
	if node == nil {
		return nil
	}

	if node.Index == u || node.Index == v {
		return node
	}

	var lca *Node = nil
	for _, child := range node.Children {
		found := dfs(child, u, v)
		if found != nil {
			if lca != nil {
				return node // если уже нашли LCA
			}
			lca = found
		}
	}

	return lca
}

func main() {
	defer out.Flush()
	var n int
	fmt.Fscanln(in, &n)
	tree := make([]Node, n)
	for i := 1; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		tree[i].Index = i
		tree[t].Children = append(tree[t].Children, &tree[i])
		tree[i].Parents = &tree[t]
	}
	fmt.Fscanln(in)
	var q int
	fmt.Fscanln(in, &q)
	for i := 0; i < q; i++ {
		var u, v int
		fmt.Fscanln(in, &u, &v)
		res := dfs(&tree[0], u, v)
		fmt.Fprintln(out, res.Index)
	}
}
