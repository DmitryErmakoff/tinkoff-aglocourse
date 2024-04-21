package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	n    int
	a    []int
	tree []Pair
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
)

type Pair struct {
	First  int
	Second int
}

func makePair(first, second int) Pair {
	return Pair{First: first, Second: second}
}

func combine(a, b Pair) Pair {
	if a.First < b.First {
		return a
	}
	if b.First < a.First {
		return b
	}
	return makePair(a.First, a.Second+b.Second)
}

func buildTree(v, tl, tr int) {
	if tl == tr {
		tree[v] = makePair(a[tl], 1)
	} else {
		tm := (tr + tl) / 2
		buildTree(v*2, tl, tm)
		buildTree(v*2+1, tm+1, tr)
		tree[v] = combine(tree[v*2], tree[v*2+1])
	}
}

func getMin(v, tl, tr, l, r int) Pair {
	if l > r {
		return makePair(math.MaxInt, 0)
	}
	if l == tl && r == tr {
		return tree[v]
	}
	tm := (tl + tr) / 2
	return combine(
		getMin(v*2, tl, tm, l, min(r, tm)),
		getMin(v*2+1, tm+1, tr, max(l, tm+1), r))
}

func update(v, tl, tr, pos, newVal int) {
	if tl == tr {
		tree[v] = makePair(newVal, 1)
	} else {
		tm := (tr + tl) / 2
		if pos <= tm {
			update(v*2, tl, tm, pos, newVal)
		} else {
			update(v*2+1, tm+1, tr, pos, newVal)
		}
		tree[v] = combine(tree[v*2], tree[v*2+1])
	}
}

func main() {
	defer out.Flush()

	var m int
	fmt.Fscanln(in, &n, &m)
	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
	}
	fmt.Fscanln(in)

	tree = make([]Pair, 4*n)
	buildTree(1, 0, n-1)

	for i := 0; i < m; i++ {
		var typeQuery, tmp1, tmp2 int
		fmt.Fscanf(in, "%v %v %v", &typeQuery, &tmp1, &tmp2)
		if typeQuery == 2 {
			p := getMin(1, 0, n-1, tmp1, tmp2-1)
			fmt.Fprintln(out, p.First, p.Second)
		} else {
			update(1, 0, n-1, tmp1, tmp2)
		}
		fmt.Fscanln(in)
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
