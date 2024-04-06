package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	value int
	count int
}

func main() {
	var n int
	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numbers := strings.Fields(input)

	stack := []Pair{{0, 0}}
	count := 0

	numbers = append(numbers, "0") // Для завершения обработки последнего элемента

	for _, numStr := range numbers {
		num, _ := strconv.Atoi(numStr)
		curr := Pair{num, 1}

		if stack[len(stack)-1].value == curr.value {
			stack = append(stack, Pair{curr.value, stack[len(stack)-1].count + 1})
		} else {
			if stack[len(stack)-1].count > 2 {
				for len(stack) > 1 && stack[len(stack)-1].value == stack[len(stack)-2].value {
					stack = stack[:len(stack)-1]
					count++
				}
				stack = stack[:len(stack)-1]
				count++
			}

			if stack[len(stack)-1].value == curr.value {
				stack = append(stack, Pair{curr.value, stack[len(stack)-1].count + 1})
			} else {
				stack = append(stack, curr)
			}
		}
	}

	fmt.Println(count)
}
