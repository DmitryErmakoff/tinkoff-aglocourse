package main

import (
	"fmt"
)

func main() {
	var n, answ int
	fmt.Scanln(&n)

	a := make([][3]int, n)
	a[0] = [3]int{1, 1, 1}

	for i := 1; i < n; i++ {
		a[i][0] = a[i-1][1] + a[i-1][2]
		a[i][1] = a[i-1][0] + a[i-1][1] + a[i-1][2]
		a[i][2] = a[i][1]
	}

	for i := 0; i < 3; i++ {
		answ += a[n-1][i]
	}

	fmt.Println(answ)
}
