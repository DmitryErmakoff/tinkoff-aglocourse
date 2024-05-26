package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	answer := make([]string, 0)
	i := 2
	for i*i <= n {
		count := 0
		for n%i == 0 {
			n /= i
			count++
		}
		if count == 1 {
			answer = append(answer, strconv.Itoa(i))
		} else if count > 1 {
			answer = append(answer, fmt.Sprintf("%d^%d", i, count))
		}
		i++
	}
	if n > 1 {
		answer = append(answer, strconv.Itoa(n))
	}
	for i := range answer {
		if i == len(answer)-1 {
			fmt.Print(answer[i])
			return
		}
		fmt.Print(answer[i] + "*")

	}
}
