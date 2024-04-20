package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 1, 1, 4, 2}
	treeSegments := make([]int, len(arr)*2)

	for i := 8; i < len(treeSegments); i++ {
		treeSegments[i] = arr[i-8]
	}
	// left = i * 2 (всегда четные)
	// right = i * 2 + 1 (всегда нечетные)
	// parent = i / 2
	for i := 7; i != 0; i-- {
		treeSegments[i] = treeSegments[i*2] + treeSegments[i*2+1]
	}
	fmt.Println(querySegmentTree(treeSegments, 1, 1, len(treeSegments)/2, 0, 8))
}

// Считает от [l, r)
func querySegmentTree(treeSegment []int, node, start, end, l, r int) int {
	// Если отрезок узла полностью находится внутри интересующего нас отрезка
	if l <= start && r >= end {
		return treeSegment[node]
	}

	// Если отрезок узла полностью за пределами интересующего нас отрезка
	if l > end || r < start {
		return 0
	}
	// Отрезок узла пересц	екается с интересующим нас отрезком, поэтому спускаемся ниже
	mid := (start + end) / 2

	return querySegmentTree(treeSegment, node*2, start, mid, l, r) + querySegmentTree(treeSegment, node*2+1, mid+1, end, l, r)
}

func createPrefSum(arr []int) []int {
	pref := make([]int, len(arr))
	pref[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		pref[i] = pref[i-1] + arr[i]
	}

	return pref
}

func rangeSum(pref []int, l, r int) int {
	if l == 0 {
		return pref[r]
	}

	return pref[r] - pref[l-1]
}
