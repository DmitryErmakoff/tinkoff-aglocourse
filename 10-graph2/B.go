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

type Edge struct {
	u, v int
	w    int64
}

func main() {
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	graph := make([]Edge, m)
	parent := make([]int, n)
	size := make([]int, n)

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &graph[i].u, &graph[i].v, &graph[i].w)
		graph[i].u--
		graph[i].v--
	}

	sort.Slice(graph, func(i, j int) bool {
		return graph[i].w < graph[j].w
	})

	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = i
	}

	var res int64
	for i := 0; i < m; i++ {
		v := graph[i].v
		u := graph[i].u
		w := graph[i].w
		if find(v, parent) != find(u, parent) {
			res += w
			unionSet(v, u, parent, size)
		}
	}

	fmt.Fprintln(out, res)
}

func find(x int, parent []int) int {
	if x != parent[x] {
		parent[x] = find(parent[x], parent)
	}
	return parent[x]
}

func unionSet(x, y int, parent, size []int) {
	x = find(x, parent)
	y = find(y, parent)

	if size[x] < size[y] {
		x, y = y, x
	}

	parent[y] = x
	size[x] += size[y]
}
