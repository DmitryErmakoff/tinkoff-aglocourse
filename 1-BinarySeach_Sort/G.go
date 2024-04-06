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
	fmt.Fscan(in, &n)
	solve(n)
}

func solve(n int) {
	arr := make([]int, n)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	for i := 2; i < len(arr); i++ {
		arr[i], arr[i/2] = arr[i/2], arr[i]
	}

	printArr(arr)
}

func printArr(arr []int) {
	for _, v := range arr {
		fmt.Fprintf(out, "%v ", v)
	}
}
