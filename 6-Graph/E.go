package main

import (
	"fmt"
)

const Maxx = 100005

var Graph = make([][]pair, Maxx)

type pair struct {
	first  int
	second int
}

func Dijkstra(source int, N int) {
	PQ := make([][]int, 0)

	Distance := make([]int, N+2)
	for i := range Distance {
		Distance[i] = 1e9
	}

	// Append source in Priority Queue
	PQ = append(PQ, []int{0, source})
	src := source
	Distance[src] = 0

	for len(PQ) != 0 {
		current := PQ[0][1]
		PQ = PQ[1:]
		for _, neighbours := range Graph[current] {
			v := neighbours.first
			weight := neighbours.second
			if Distance[v] > Distance[current]+weight {
				Distance[v] = Distance[current] + weight
				PQ = append(PQ, []int{Distance[v], v})
			}
		}
	}

	fmt.Println(1 + Distance[0])
}

func minSumDigits(N int) {
	for i := 1; i <= N; i++ {
		From := (i) % N
		To := (i + 1) % N
		Wt := 1
		Graph[From] = append(Graph[From], pair{To, Wt})
	}

	for i := 1; i <= N; i++ {
		From := (i) % N
		To := (10 * i) % N
		Wt := 0
		Graph[From] = append(Graph[From], pair{To, Wt})
	}

	Dijkstra(1, N)
}

func main() {
	var N int
	fmt.Scanln(&N)

	minSumDigits(N)
}
