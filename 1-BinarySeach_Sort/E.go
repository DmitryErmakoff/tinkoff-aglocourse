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

func main() {
	defer out.Flush()
	var a, b, c, d float64
	fmt.Fscanln(in, &a, &b, &c, &d)

	f := func(x float64) float64 {
		return a*math.Pow(x, 3) + b*math.Pow(x, 2) + c*x + d
	}
	l := -1.0
	r := 1.0
	exp := 1e-6

	for f(l)*f(r) >= 0 {
		l *= 2
		r *= 2
	}

	for r-l > exp {
		mid := (l + r) / 2
		if f(mid)*f(r) <= 0 {
			l = mid
		} else {
			r = mid
		}
	}
	fmt.Fprintf(out, "%.7f\n", (l+r)/2)
}
