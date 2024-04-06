package main

import (
	"bufio"
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

type Heap []int

func (h *Heap) Insert(n int) {
	*h = binarySearchAndInsert(*h, n)
}

func (h *Heap) Extract() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	defer out.Flush()
	h := Heap{}
	var n, num2 int
	fmt.Fscanln(in, &n)
	for i := 0; i < n; i++ {
		sc.Scan()
		nums := strings.Fields(sc.Text())
		if len(nums) == 1 {
			fmt.Fprintln(out, h.Extract())
		} else {
			num2, _ = strconv.Atoi(nums[1])
			h.Insert(num2)
		}
	}
}

func binarySearchAndInsert(arr []int, target int) []int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			// Вставка элемента в середину массива без изменения порядка сортировки
			return append(arr[:mid+1], append([]int{target}, arr[mid+1:]...)...)
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// Если элемент не найден, вставляем его на правильную позицию
	insertIndex := left

	return append(arr[:insertIndex], append([]int{target}, arr[insertIndex:]...)...)
}
