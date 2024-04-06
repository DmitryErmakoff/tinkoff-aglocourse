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

func compare(n int) bool {
	fmt.Fprintln(out, n)
	out.Flush()
	var c string
	fmt.Fscanln(in, &c)
	return c[0] == '<'
}

func binarySearch(n int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		if compare(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func main() {
	defer out.Flush()
	var n int
	fmt.Fscanln(in, &n)
	fmt.Fprintln(out, "!", binarySearch(n))
}
