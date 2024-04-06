package main

import (
	"bufio"
	"fmt"
	"os"
)

func binarySearch(n int, arr []int) bool {
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] == n {
			return true
		} else if arr[mid] > n {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscanln(in, &n, &k)

	arr := make([]int, n)
	fmt.Fscanln(in, packAddr(arr)...)

	queries := make([]int, k)
	fmt.Fscanln(in, packAddr(queries)...)

	for _, q := range queries {
		if binarySearch(q, arr) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}

func packAddr(arr []int) []any {
	a := make([]any, len(arr))
	for i := 0; i < len(a); i++ {
		a[i] = &arr[i]
	}
	return a
}
