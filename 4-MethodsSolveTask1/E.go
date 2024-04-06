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

func canSplit(nums []int, k int, maxSum int) bool {
	cnt := 1
	sum := 0

	for _, num := range nums {
		sum += num
		if sum > maxSum {
			cnt++
			sum = num
		}
	}

	return cnt <= k
}

func splitArray(nums []int, k int) int {
	left := 0
	right := 0

	for _, num := range nums {
		if num > left {
			left = num
		}
		right += num
	}

	for left < right {
		mid := left + (right-left)/2
		if canSplit(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	defer out.Flush()
	var n, k, t int
	fmt.Fscanln(in, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &t)
		arr[i] = t
	}
	result := splitArray(arr, k)
	fmt.Fprintln(out, result)
}
