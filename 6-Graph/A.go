package main

import (
	"bufio"
	"fmt"
	"os"
)

// Реализация графа
type Graph struct {
	nodes map[int][]int
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int][]int)}
}

func (g *Graph) AddNode(n int) {
	g.nodes[n] = make([]int, 0)
}

func (g *Graph) AddEdge(a, b int) {
	g.nodes[a] = append(g.nodes[a], b)
	g.nodes[b] = append(g.nodes[b], a)
}

func (g *Graph) Print() {
	for nodes, neghbors := range g.nodes {
		fmt.Fprintf(out, "%d -> %v\n", nodes, neghbors)
	}
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	g := NewGraph()
	var N, M int
	fmt.Fscanln(in, &N, &M)
	for i := 0; i < N; i++ {
		g.AddNode(i + 1)
	}
	for i := 0; i < M; i++ {
		var n1, n2 int
		fmt.Fscanln(in, &n1, &n2)
		g.AddEdge(n1, n2)
	}
	used := make([]int, N)
	for i := 0; i < N; i++ {
		used[i] = 0
	}
	count := 0

}

func dfs1(g *Graph, used []int, start int) {
	used[start-1] = 1
	for _, neighbor := range g.nodes[start] {
		if used[neighbor-1] == 0 {

			dfs1(g, used, neighbor)
		}
	}
}

func dfs(g *Graph, used []int, start int) {
	used[start-1] = 1
	for _, neighbor := range g.nodes[start] {
		if used[neighbor-1] == 0 {
			fmt.Printf("%d ", neighbor)
			dfs(g, used, neighbor)
		}
	}
}
