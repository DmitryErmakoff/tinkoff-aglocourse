package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
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
	for nodes, neighbors := range g.nodes {
		fmt.Fprintf(out, "%d -> %v\n", nodes, neighbors)
	}
}

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
	used := make([]int, N+1)
	used[0] = 999
	countComponents := 0
	for i := 1; i <= N; i++ {
		if used[i] == 0 {
			arr := make([]int, 0)
			dfs(g, i, used, &arr)
			countComponents++
		}
	}
	fmt.Fprintln(out, countComponents)
	for i := 1; i < len(used); i++ {
		used[i] = 0
	}
	for i := 1; i <= N; i++ {
		if used[i] == 0 {
			arr := make([]int, 0)
			arr = append(arr, i)
			dfs(g, i, used, &arr)
			fmt.Fprintln(out, len(arr))
			sort.Ints(arr)
			for j := range arr {
				fmt.Fprintf(out, "%d ", arr[j])
			}
			fmt.Fprintln(out)
		}
	}
}

func dfs(g *Graph, start int, used []int, arr *[]int) {
	used[start] = 1
	for _, neighbors := range g.nodes[start] {
		if used[neighbors] == 0 {
			*arr = append(*arr, neighbors)
			dfs(g, neighbors, used, arr) // Здесь исправленный вызов функции с передачей верного соседа
		}
	}
}
