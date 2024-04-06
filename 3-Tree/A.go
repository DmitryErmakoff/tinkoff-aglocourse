package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	Depth    int
	Visited  bool
	Children []*Node
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	var n, t int
	fmt.Fscanln(in, &n)
	tree := make([]Node, n)
	depth := make([]int, n)

	for i := 1; i < n; i++ {
		fmt.Fscan(in, &t)
		tree[i].Depth = tree[t].Depth + 1
		tree[i].Children = append(tree[i].Children, &tree[t])
		tree[t].Children = append(tree[t].Children, &tree[i])
	}

	for i := 0; i < n; i++ {
		depth[i] = tree[i].Depth
	}

	// Находим самую удаленную вершину от корня
	dfs(&tree[0], 0)
	maxDepth := 0
	var farthestNode *Node

	for i := range tree {
		tree[i].Visited = false
		if tree[i].Depth > maxDepth {
			maxDepth = tree[i].Depth
			farthestNode = &tree[i]
		}
	}

	// Находим диаметр дерева
	dfs(farthestNode, 0)
	diameter := 0

	for i := range tree {
		if tree[i].Depth > diameter {
			diameter = tree[i].Depth
		}
	}

	fmt.Println(maxDepth, diameter)
	for i := 0; i < len(depth); i++ {
		fmt.Printf("%d ", depth[i])
	}
}

func dfs(node *Node, depth int) {
	node.Depth = depth
	node.Visited = true

	for _, child := range node.Children {
		if !child.Visited {
			dfs(child, depth+1)
		}
	}
}
