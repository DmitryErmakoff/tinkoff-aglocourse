package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner *bufio.Scanner
var writer *bufio.Writer

func build(tree, src []int, v, tl, tr int) {
	if tl == tr {
		tree[v] = src[tl]
	} else {
		tm := (tl + tr) / 2
		build(tree, src, v*2, tl, tm)
		build(tree, src, v*2+1, tm+1, tr)
		tree[v] = tree[v*2] + tree[v*2+1]
	}
}

func getKth(tree []int, v, tl, tr, k int) int {
	if tl == tr {
		return tl
	}

	tm := (tl + tr) / 2
	if k <= tree[v*2] {
		return getKth(tree, v*2, tl, tm, k)
	}
	return getKth(tree, v*2+1, tm+1, tr, k-tree[v*2])
}

func update(tree []int, v, tl, tr, pos int) {
	if tl == tr {
		tree[v] = 1 - tree[v]
	} else {
		tm := (tl + tr) / 2
		if pos <= tm {
			update(tree, v*2, tl, tm, pos)
		} else {
			update(tree, v*2+1, tm+1, tr, pos)
		}
		tree[v] = tree[v*2] + tree[v*2+1]
	}
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	src := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		src[i] = val
	}

	tree := make([]int, n*4)
	build(tree, src, 1, 0, n-1)

	for req := 0; req < m; req++ {
		scanner.Scan()
		reqType := scanner.Text()
		if reqType == "1" {
			scanner.Scan()
			idx, _ := strconv.Atoi(scanner.Text())
			update(tree, 1, 0, n-1, idx)
		} else {
			scanner.Scan()
			k, _ := strconv.Atoi(scanner.Text())
			fmt.Fprintf(writer, "%d\n", getKth(tree, 1, 0, n-1, k+1))
		}
	}
}
