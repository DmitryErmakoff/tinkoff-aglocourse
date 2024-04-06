package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack []int

func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	expression, _ := reader.ReadString('\n')
	expression = strings.TrimSpace(expression)
	tokens := strings.Fields(expression)
	stack := Stack{}

	for _, token := range tokens {
		switch token {
		case "+":
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			stack.Push(operand1 + operand2)
		case "-":
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			stack.Push(operand2 - operand1)
		case "*":
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			stack.Push(operand1 * operand2)
		default:
			num, err := strconv.Atoi(token)
			if err != nil {
				fmt.Printf("Некорректный ввод: %s\n", token)
				return
			}
			stack.Push(num)
		}
	}

	if len(stack) == 1 {
		result := stack.Pop()
		fmt.Println(result)
	} else {
		fmt.Println("Некорректное входное выражение")
	}
}
