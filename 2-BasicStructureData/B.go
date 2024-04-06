package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	deq := list.New()

	var N, K int
	fmt.Fscanln(in, &N, &K)
	numbers := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(in, &numbers[i])
	}

	for i := 0; i < N; i++ {
		for deq.Len() > 0 && numbers[i] <= deq.Back().Value.([2]int)[0] {
			deq.Remove(deq.Back())
		}
		deq.PushBack([2]int{numbers[i], i})

		for deq.Front().Value.([2]int)[1] < i-K+1 {
			deq.Remove(deq.Front())
		}

		if i >= K-1 {
			fmt.Fprintf(out, "%d ", deq.Front().Value.([2]int)[0])
		}
	}
}
