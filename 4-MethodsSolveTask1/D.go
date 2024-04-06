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
	var k int64
	fmt.Fscanln(in, &n, &k)

	var left, right, middle int64
	left, right = 0, 999999999999999

	for right-left > 1 {
		middle = (left + right) / 2

		i, j := 1, n
		var count int64

		for i <= n && j > 0 {
			if int64(i)*int64(j) < middle {
				count += int64(j)
				i++
			} else {
				j--
			}
		}

		if count < k {
			left = middle
		} else {
			right = middle
		}
	}
	fmt.Fprintln(out, left)
}
