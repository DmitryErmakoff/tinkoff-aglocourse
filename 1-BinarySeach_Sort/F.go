package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	in             = bufio.NewReader(os.Stdin)
	out            = bufio.NewWriter(os.Stdout)
	inversionCount int
)

func main() {
	defer out.Flush()
	var n int64

	fmt.Fscanln(in, &n)
	arr := make([]float64, n)
	fmt.Fscanln(in, packAddr(arr)...)

	arr = MergeSort(arr)

	fmt.Fprintln(out, inversionCount)

	for i := range arr {
		fmt.Fprintf(out, "%v ", arr[i])
	}
}

func packAddr(arr []float64) []any {
	a := make([]any, len(arr))
	for i := 0; i < len(arr); i++ {
		a[i] = &arr[i]
	}
	return a
}

// функция сортировки слиянием
func MergeSort(a []float64) []float64 {
	N1 := len(a) / 2
	a1 := a[:N1] // деление массива на примерно равные доли
	a2 := a[N1:]

	if len(a1) > 1 {
		a1 = MergeSort(a1)
	}
	if len(a2) > 1 {
		a2 = MergeSort(a2)
	}
	return MergeArray(a1, a2)
}

func MergeArray(a, b []float64) []float64 {
	n := len(a)
	m := len(b)
	c := make([]float64, n+m)

	i := 0
	j := 0
	k := 0

	for i < n && j < m {
		if a[i] >= b[j] {
			c[k] = b[j]
			j++
			k++
			inversionCount += n - i // увеличиваем счётчик инверсий
		} else {
			c[k] = a[i]
			i++
			k++
		}

		if i == n {
			for j < m {
				c[k] = b[j]
				j++
				k++
			}
		}
		if j == m {
			for i < n {
				c[k] = a[i]
				i++
				k++
			}
		}
	}
	return c
}
