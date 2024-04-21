//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//var n int
//var a []int
//var tree []int
//var in = bufio.NewReader(os.Stdin)
//var out = bufio.NewWriter(os.Stdout)
//
//func buildTree(v, tl, tr int) {
//	if tl == tr {
//		if tl < len(a) {
//			tree[v] = a[tl]
//		}
//	} else {
//		tm := (tr + tl) / 2
//		buildTree(v*2, tl, tm)
//		buildTree(v*2+1, tm+1, tr)
//		tree[v] = tree[v*2] + tree[v*2+1]
//	}
//}
//
//func getSum(l, r, v, tl, tr int) int {
//	if l <= tl && tr <= r {
//		return tree[v]
//	}
//
//	if tr < l || r < tl {
//		return 0
//	}
//
//	tm := (tl + tr) / 2
//	return getSum(l, r, v*2, tl, tm) + getSum(l, r, v*2+1, tm+1, tr)
//}
//
//func update(idx, val, v, tl, tr int) {
//	if idx <= tl && tr <= idx {
//		a[idx] = val
//		tree[v] = val
//		return
//	}
//
//	if tr < idx || idx < tl {
//		return
//	}
//
//	tm := (tl + tr) / 2
//	update(idx, val, v*2, tl, tm)
//	update(idx, val, v*2+1, tm+1, tr)
//	tree[v] = tree[v*2] + tree[v*2+1]
//}
//
//func main() {
//	var m int
//	fmt.Fscanln(in, &n, &m)
//
//	for i := 0; i < n; i++ {
//		var t int
//		fmt.Fscan(in, &t)
//	}
//}

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	n    int
	a    []int
	tree []int
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
)

func buildTree(v, tl, tr int) {
	if tl == tr {
		if tl < len(a) {
			tree[v] = a[tl]
		}
	} else {
		tm := (tr + tl) / 2
		buildTree(v*2, tl, tm)
		buildTree(v*2+1, tm+1, tr)
		tree[v] = tree[v*2] + tree[v*2+1]
	}
}

func getSum(l, r, v, tl, tr int) int {
	if l <= tl && r >= tr {
		return tree[v]
	}

	if tr < l || r < tl {
		return 0
	}

	tm := (tl + tr) / 2
	return getSum(l, r, v*2, tl, tm) + getSum(l, r, v*2+1, tm+1, tr)
}

func update(idx, val, v, tl, tr int) {
	if idx <= tl && tr <= idx { // Тоже самое что и idx == tr == tl
		a[idx] = val
		tree[v] = val
		return
	}

	if tr < idx || idx < tl {
		return
	}

	tm := (tl + tr) / 2
	update(idx, val, v*2, tl, tm)
	update(idx, val, v*2+1, tm+1, tr)
	tree[v] = tree[v*2] + tree[v*2+1]
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

	tree = make([]int, 4*n)
	buildTree(1, 0, n-1) // Создаем дерево

	for i := 0; i < m; i++ {
		var typeQuery, tmp1, tmp2 int
		fmt.Fscanf(in, "%v %v %v", &typeQuery, &tmp1, &tmp2)
		if typeQuery == 2 {
			fmt.Fprintln(out, getSum(tmp1, tmp2-1, 1, 0, n-1))
		} else {
			update(tmp1, tmp2, 1, 0, n-1)
		}
		fmt.Fscanln(in)
	}

}
