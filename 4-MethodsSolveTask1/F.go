package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Pair struct {
	First  int
	Second int
}

type ByValue []Pair

func (p ByValue) Len() int           { return len(p) }
func (p ByValue) Less(i, j int) bool { return p[i].First < p[j].First }
func (p ByValue) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	defer out.Flush()
	var n int
	var arr ByValue
	var x, y int
	k := 0
	fmt.Fscanln(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &x, &y)
		arr = append(arr, Pair{
			First:  x,
			Second: 1,
		})
		arr = append(arr, Pair{
			First:  y,
			Second: -1,
		})
	}
	sort.Sort(ByValue(arr))

	p := arr[0].Second
	k += arr[1].First - arr[0].First
	for i := 1; i < len(arr)-1; i++ {
		if p+arr[i].Second > 0 && p+arr[i].Second != 0 {
			k += arr[i+1].First - arr[i].First
		}
		p += arr[i].Second
	}
	fmt.Fprintln(out, k)
}
