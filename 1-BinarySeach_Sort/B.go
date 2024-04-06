package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func binarySearch(arr []int64, numb int64) int64 {
	start := 0
	end := len(arr)
	for end-start > 1 {
		mid := (start + end) / 2
		if arr[mid] == numb {
			return arr[mid]
		} else if arr[mid] > numb {
			end = mid
		} else if arr[mid] < numb {
			start = mid
		}

		if end-start == 1 {
			if end == len(arr) {
				return arr[start]
			} else if math.Abs(float64(numb-arr[start])) > math.Abs(float64(numb-arr[end])) {
				return arr[end]
			} else {
				return arr[start]
			}
		}
	}
	return -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int64
	fmt.Fscanln(in, &n, &k)

	arr := make([]int64, n)
	fmt.Fscanln(in, packAddr(arr)...)

	queries := make([]int64, k)
	fmt.Fscanln(in, packAddr(queries)...)

	for _, q := range queries {
		fmt.Println(binarySearch(arr, q))
	}
}

func packAddr(arr []int64) []any {
	a := make([]any, len(arr))
	for i := 0; i < len(arr); i++ {
		a[i] = &arr[i]
	}
	return a
}
