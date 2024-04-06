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

func findApproximateMatches(p, t string) []int {
	var result []int
	n, m := len(p), len(t)

	for i := 0; i <= m-n; i++ {
		countMismatches := 0
		for j := 0; j < n; j++ {
			if p[j] != t[i+j] {
				countMismatches++
				if countMismatches > 1 {
					break
				}
			}
		}
		if countMismatches <= 1 {
			result = append(result, i+1)
		}
	}

	return result
}

func main() {
	var p, t string
	fmt.Fscanln(in, &p)
	fmt.Fscanln(in, &t)

	positions := findApproximateMatches(p, t)

	fmt.Println(len(positions))
	for _, pos := range positions {
		fmt.Printf("%v ", pos)
	}
}
