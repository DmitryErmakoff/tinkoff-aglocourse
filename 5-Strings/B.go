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

func zFunction(s string) []int {
	n := len(s)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l, r = i, i+z[i]-1
		}
	}
	return z
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findAllOccurrences(S, T string) {

	newString := S + "$" + T
	z := zFunction(newString)

	occurrences := []int{}
	for i := len(S) + 1; i < len(newString); i++ {
		if z[i] == len(S) {
			occurrences = append(occurrences, i-len(S)-1)
		}
	}

	fmt.Fprintf(out, "%d ", len(occurrences))
	for _, index := range occurrences {
		fmt.Fprint(out, index, " ")
	}
	fmt.Fprintln(out)
}

func main() {
	defer out.Flush()
	var q int
	var S, T string
	fmt.Fscanln(in, &T)
	fmt.Fscanln(in, &q)

	for i := 0; i < q; i++ {
		fmt.Fscanln(in, &S)
		findAllOccurrences(S, T)

	}
}
