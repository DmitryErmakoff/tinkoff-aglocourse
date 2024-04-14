package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	cost := make([]int, n)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&cost[i])
	}

	ans[0] = cost[0]
	if n > 1 {
		ans[1] = cost[1]
	}

	for i := 2; i < n; i++ {
		ans[i] = min(ans[i-1], ans[i-2]) + cost[i]
	}

	fmt.Println(ans[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
