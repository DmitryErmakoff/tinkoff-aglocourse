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
	var n, m, q int

	fmt.Fscanln(in, &n, &m, &q)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &matrix[i][j])
		}
	}

	var x1, y1, x2, y2 int
	pref := prefSum(matrix, n, m)
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &x1, &y1, &x2, &y2)
		fmt.Fprintln(out, query(x1-1, y1-1, x2-1, y2-1, pref))
	}

}

func prefSum(matrix [][]int, n, m int) [][]int {
	pref := make([][]int, n)
	for i := 0; i < n; i++ {
		pref[i] = make([]int, m)
	}
	pref[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		pref[i][0] = pref[i-1][0] + matrix[i][0]
	}

	for i := 1; i < m; i++ {
		pref[0][i] = pref[0][i-1] + matrix[0][i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			pref[i][j] = pref[i-1][j] + pref[i][j-1] - pref[i-1][j-1] + matrix[i][j]
		}
	}
	return pref
}

func query(x1, y1, x2, y2 int, pref [][]int) int {
	result := pref[x2][y2]
	if x1 > 0 {
		result -= pref[x1-1][y2]
	}
	if y1 > 0 {
		result -= pref[x2][y1-1]
	}
	if x1 > 0 && y1 > 0 {
		result += pref[x1-1][y1-1]
	}
	return result
}
