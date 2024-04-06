package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	q := list.New()
	qt := list.New()
	k1 := 0
	k2 := 0

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		t := strings.Fields(scanner.Text())

		if strings.Contains(t[0], "-") {
			val := q.Front().Value
			fmt.Println(val)
			q.Remove(q.Front())
			k1--
		} else if strings.Contains(t[0], "+") {
			qt.PushBack(t[len(t)-1])
			k2++
		} else {
			qt.PushFront(t[len(t)-1])
			k2++
		}

		if k1 < k2 {
			q.PushBack(qt.Front().Value)
			qt.Remove(qt.Front())
			k2--
			k1++
		}
	}
}
