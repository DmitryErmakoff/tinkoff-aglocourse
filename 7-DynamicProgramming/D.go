package main

import "fmt"

var n, m int
var a [][]int

func createArray(n, m int) {
	a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
		for j := 0; j < m; j++ {
			a[i][j] = -1
		}
	}
	a[0][0] = 1
}

func solveProbelm(i, j int) int {
	if i >= 0 && j >= 0 && i < n && j < m {
		if a[i][j] == -1 {
			a[i][j] = solveProbelm(i-2, j-1) + solveProbelm(i-2, j+1) + solveProbelm(i-1, j-2) + solveProbelm(i+1, j-2)
		}
	} else {
		return 0
	}
	return a[i][j]
}

func main() {
	fmt.Scanf("%d %d", &n, &m)
	createArray(n, m)
	result := solveProbelm(n-1, m-1)
	fmt.Println(result)
}
