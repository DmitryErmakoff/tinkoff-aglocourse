package main

import (
	"fmt"
	"sort"
)

func norm(a []int, k int, s int) bool {
	prev := 0
	for i := 0; i < k-1; i++ {
		cur := sort.Search(len(a), func(j int) bool { return a[j] >= a[prev]+s })
		prev = cur
		if cur == len(a) {
			return false
		}
	}
	return true
}

func cow(a []int, k int) int {
	l := 0
	r := a[len(a)-1] - a[0] + 1
	for r-l > 1 {
		c := (l + r) / 2
		if norm(a, k, c) {
			l = c
		} else {
			r = c
		}
	}
	return l
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)
	fmt.Println(cow(a, k))
}
