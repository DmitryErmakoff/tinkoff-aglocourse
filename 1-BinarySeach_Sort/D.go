package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func function(x float64, n float64) float64 {
	return math.Pow(x, 2) + math.Sqrt(x+1) - n
}

func solve(n float64) float64 {
	eps := 1e-10
	low, high := 0.0, n
	for high-low > eps {
		mid := (low + high) / 2
		if function(low, n)*function(mid, n) <= 0 {
			high = mid
		} else {
			low = mid
		}
	}
	return (low + high) / 2
}

func main() {
	defer out.Flush()
	var n float64
	fmt.Fscanln(in, &n)
	result := solve(n)
	fmt.Fprintf(out, "%.6f\n", result)
}
