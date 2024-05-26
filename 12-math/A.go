package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	var N, K int
	fmt.Scanf("%d %d", &N, &K)

	lcmValue := lcm(N, K)

	fmt.Println(lcmValue)
}
