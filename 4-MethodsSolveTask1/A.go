package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func main() {
	defer out.Flush()
	var n, t int64
	fmt.Fscanln(in, &n)
	arr := make([]int64, n)
	for i := int64(0); i < n; i++ {
		fmt.Fscan(in, &t)
		arr[i] = t
	}

	fmt.Fscanln(in)

	var q, op, l, r int64
	fmt.Fscanln(in, &q)
	arrPrefSum := prefSum(arr)
	arrPrefXOR := prefXOR(arr)
	for i := int64(0); i < q; i++ {
		fmt.Fscan(in, &op, &l, &r)
		if op == 1 {
			fmt.Fprintln(out, queryPrefSum(arrPrefSum, l-1, r-1))
		} else {
			fmt.Fprintln(out, queryPrefXOR(arrPrefXOR, l-1, r-1))
		}
	}
}

func prefSum(arr []int64) []int64 {
	pref := make([]int64, len(arr))
	pref[0] = arr[0]
	for i := int64(1); i < int64(len(arr)); i++ {
		pref[i] = pref[i-1] + arr[i]
	}
	return pref
}

func queryPrefSum(pref []int64, l, r int64) int64 {
	if l > 0 {
		return pref[r] - pref[l-1]
	} else {
		return pref[r]
	}
}

func prefXOR(arr []int64) []int64 {
	pref := make([]int64, len(arr))
	pref[0] = arr[0]
	for i := int64(1); i < int64(len(arr)); i++ {
		pref[i] = pref[i-1] ^ arr[i]
	}
	return pref
}

func queryPrefXOR(pref []int64, l, r int64) int64 {
	if l > 0 {
		return pref[r] ^ pref[l-1]
	} else {
		return pref[r]
	}
}
