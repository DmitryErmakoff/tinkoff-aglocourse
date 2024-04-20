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

func main() {
	defer out.Flush()

	var n int
	fmt.Fscanln(in, &n)

	cost := make([]int, n)
	for i := 0; i < len(cost); i++ {
		fmt.Fscan(in, &cost[i])
	}

	dp := make([]int, n)
	dp[0] = cost[0]

	if n > 1 {
		dp[1] = cost[1]
	}

	for i := 2; i < len(dp); i++ {
		dp[i] = min(dp[i-1], dp[i-2]) + cost[i]
	}
	fmt.Fprintln(out, dp[n-1])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
