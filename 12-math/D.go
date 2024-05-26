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

	var n, num int64
	num = 1
	fmt.Fscan(in, &n)
	for z := int64(1); z <= n; z++ {
		num *= z
		for num%10 == 0 {
			num /= 10
		}
		num %= 1000000
	}
	fmt.Fprintln(out, num%10)
}
