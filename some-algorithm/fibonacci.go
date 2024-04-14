package main

import "fmt"

func main() {
	fmt.Println(Fibonacci(20))
}

// Фибоначи через динамику
func Fibonacci(n int) int {
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[len(arr)-1]
}
