package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner *bufio.Scanner

func init() {
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
}

func scanInt() int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func build(tree []int, a []int, v int, tl int, tr int) {
	if tl == tr {
		tree[v] = a[tl]
	} else {
		tm := (tl + tr) / 2
		build(tree, a, v*2, tl, tm)
		build(tree, a, v*2+1, tm+1, tr)
		tree[v] = max(tree[v*2], tree[v*2+1])
	}
}

func getGex(tree []int, v int, tl int, tr int, l int, r int, x int) int {
	if l > r {
		return 2147483647
	}

	if tl == tr {
		if tree[v] >= x {
			return tl
		} else {
			return 2147483647
		}
	}

	tm := (tl + tr) / 2

	if tl < l {
		return min(getGex(tree, v*2, tl, tm, l, min(r, tm), x),
			getGex(tree, v*2+1, tm+1, tr, max(l, tm+1), r, x))
	}

	if tree[v*2] >= x {
		return getGex(tree, v*2, tl, tm, l, min(r, tm), x)
	} else {
		return getGex(tree, v*2+1, tm+1, tr, max(l, tm+1), r, x)
	}
}

func update(tree []int, v int, tl int, tr int, pos int, newVal int) {
	if tl == tr {
		tree[v] = newVal
	} else {
		tm := (tl + tr) / 2
		if pos <= tm {
			update(tree, v*2, tl, tm, pos, newVal)
		} else {
			update(tree, v*2+1, tm+1, tr, pos, newVal)
		}
		tree[v] = max(tree[v*2], tree[v*2+1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	scanner.Split(bufio.ScanWords)

	n := scanInt()
	m := scanInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	tree := make([]int, n*4)
	build(tree, a, 1, 0, n-1)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for req := 0; req < m; req++ {
		reqType := scanInt()

		if reqType == 1 {
			i := scanInt()
			v := scanInt()
			update(tree, 1, 0, n-1, i, v)
		} else {
			x := scanInt()
			l := scanInt()
			res := getGex(tree, 1, 0, n-1, l, n-1, x)
			if res == math.MaxInt32 {
				fmt.Fprintln(w, -1)
			} else {
				fmt.Fprintln(w, res)
			}
		}
	}
}
