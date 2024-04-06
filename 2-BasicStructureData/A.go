package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	sc  = bufio.NewScanner(in)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	stack := list.New()
	var n int

	fmt.Fscanln(in, &n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums := strings.Fields(sc.Text())
		if len(nums) == 2 {
			num2, _ := strconv.ParseInt(nums[1], 10, 64)
			inputArr := [2]int64{num2, num2}
			if stack.Len() == 0 {
				stack.PushBack(inputArr)
			} else {
				el := stack.Back().Value.([2]int64)[1]
				if el < inputArr[1] {
					stack.PushBack([2]int64{inputArr[0], el})
				} else {
					stack.PushBack(inputArr)
				}
			}
		} else {
			num1, _ := strconv.ParseInt(nums[0], 10, 64)
			if num1 == 2 {
				stack.Remove(stack.Back())
			} else if num1 == 3 {
				el := stack.Back().Value.([2]int64)[1]
				fmt.Fprintln(out, el)
				out.Flush()
			}
		}

	}
}

func printList(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
