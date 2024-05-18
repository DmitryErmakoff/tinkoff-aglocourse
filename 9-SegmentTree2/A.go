package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Pair struct {
	First  int64
	Second int64
}

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func get_min(tree *[]Pair, v, tl, tr, l, r int64) int64 {
	if l > r {
		return math.MaxInt64
	}

	additional := (*tree)[v].Second

	if l == tl && tr == r {
		return (*tree)[v].First + additional
	}

	tm := (tl + tr) / 2

	return min(
		get_min(tree, v*2, tl, tm, l, min(r, tm)),
		get_min(tree, v*2+1, tm+1, tr, max(l, tm+1), r),
	) + additional
}

func update(tree *[]Pair, v, tl, tr, l, r, value int64) {
	if l > r {
		return
	}

	if l == tl && tr == r {
		(*tree)[v].Second += value
	} else {
		tm := (tl + tr) / 2
		update(tree, v*2, tl, tm, l, min(r, tm), value)
		update(tree, v*2+1, tm+1, tr, max(l, tm+1), r, value)
		(*tree)[v].First = min(
			(*tree)[v*2].First+(*tree)[v*2].Second,
			(*tree)[v*2+1].First+(*tree)[v*2+1].Second,
		)
	}
}

func main() {
	defer out.Flush()

	var n, m int64
	fmt.Fscan(in, &n, &m) // Use fmt.Fscan instead of fmt.Fscanln

	tree := make([]Pair, n*4)
	for i := range tree { // Simplified syntax for iterating over slice indices
		tree[i].First = 0
		tree[i].Second = 0
	}

	var query_type string
	var l, r, v int64

	for i := int64(0); i < m; i++ {
		fmt.Fscan(in, &query_type)
		if query_type == "1" {
			fmt.Fscan(in, &l, &r, &v)
			update(&tree, 1, 0, n-1, l, r-1, v)
		} else if query_type == "2" {
			fmt.Fscan(in, &l, &r)
			fmt.Fprintln(out, get_min(&tree, 1, 0, n-1, l, r-1))
		}
	}
}

func min(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}
