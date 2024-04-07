package main

import (
	"bufio"
	"fmt"
	"os"
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
	if hasCycleDirectedGraph(g) {
		fmt.Fprintln(out, 1)
	} else {
		fmt.Fprintln(out, 0)
	}

}

func hasCycleDirectedGraph(g *Graph) bool {
	visited := make(map[int]int) // 0 - белый, 1 - серый, 2 - черный
	var hasCycleUtil func(u int) bool

	hasCycleUtil = func(u int) bool {
		visited[u] = 1 // Помечаем вершину как серую

		for _, v := range g.nodes[u] {
			if visited[v] == 1 { // Обнаружен цикл
				return true
			}
			if visited[v] == 0 && hasCycleUtil(v) {
				return true
			}
		}

		visited[u] = 2 // Помечаем вершину как черную
		return false
	}

	for u := range g.nodes {
		if visited[u] == 0 {
			if hasCycleUtil(u) {
				return true
			}
		}
	}

	return false
}
