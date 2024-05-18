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

type Node struct {
	sum, add int
	replaced int
}

func createNode() *Node {
	return &Node{
		sum:      0,
		add:      0,
		replaced: -1,
	}
}

func pushTree(tree *[]Node, v, treeLength int) {
	if (*tree)[v].replaced != -1 {
		if treeLength != 1 {
			(*tree)[v*2].replaced = (*tree)[v].replaced
			(*tree)[v*2+1].replaced = (*tree)[v].replaced
			(*tree)[v*2].add = 0
			(*tree)[v*2+1].add = 0
		}
		(*tree)[v].sum = ((*tree)[v].replaced) * (treeLength)
	}
	(*tree)[v].replaced = -1

	if treeLength != 1 {
		(*tree)[v*2].add += (*tree)[v].add
		(*tree)[v*2+1].add += (*tree)[v].add
	}
	(*tree)[v].sum += (*tree)[v].add * treeLength
	(*tree)[v].add = 0
}

func getSum(tree *[]Node, v, tl, tr, l, r int) int {
	pushTree(tree, v, tr-tl+1)

	if l > r {
		return 0
	}

	if l == tl && tr == r {
		return (*tree)[v].sum
	}

	tm := (tl + tr) / 2
	return getSum(tree, v*2, tl, tm, l, min(r, tm)) + getSum(tree, v*2+1, tm+1, tr, max(l, tm+1), r)
}

func updateReplace(tree *[]Node, v, tl, tr, l, r, value int) {
	pushTree(tree, v, tr-tl+1)

	if l > r {
		return
	}

	if l == tl && r == tr {
		(*tree)[v].replaced = value
		(*tree)[v].add = 0
		pushTree(tree, v, tr-tl+1)
	} else {
		tm := (tl + tr) / 2
		updateReplace(tree, v*2, tl, tm, l, min(r, tm), value)
		updateReplace(tree, v*2+1, tm+1, tr, max(l, tm+1), r, value)
		(*tree)[v].sum = (*tree)[v*2].sum + (*tree)[v*2+1].sum
	}
}

func updateAdd(tree *[]Node, v, tl, tr, l, r, value int) {
	pushTree(tree, v, tr-tl+1)

	if l > r {
		return
	}

	if l == tl && r == tr {
		(*tree)[v].add += value
		pushTree(tree, v, tr-tl+1)
	} else {
		tm := (tl + tr) / 2
		updateAdd(tree, v*2, tl, tm, l, min(r, tm), value)
		updateAdd(tree, v*2+1, tm+1, tr, max(l, tm+1), r, value)
		(*tree)[v].sum = (*tree)[v*2].sum + (*tree)[v*2+1].sum
	}
}

func main() {
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

	tree := make([]Node, n*4)

	for i := range tree {
		tree[i] = Node{
			sum:      0,
			add:      0,
			replaced: -1,
		}
	}

	var query_type string
	var l, r, v int

	for i := 0; i < m; i++ {
		fmt.Fscan(in, &query_type)
		if query_type == "1" {
			fmt.Fscan(in, &l, &r, &v)
			updateReplace(&tree, 1, 0, n-1, l, r-1, v)
		} else if query_type == "2" {
			fmt.Fscan(in, &l, &r, &v)
			updateAdd(&tree, 1, 0, n-1, l, r-1, v)
		} else {
			fmt.Fscan(in, &l, &r)
			fmt.Fprintln(out, getSum(&tree, 1, 0, n-1, l, r-1))
		}
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
