package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type Node struct {
	Number   int
	Segments int
	Set      int
	L        int
	R        int
	Up       bool
}

var t []Node
var color []byte
var cord []int
var delta []int
var n int
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func getMiddle(l, r int) int {
	return l + (r-l)/2
}

func build(v, tl, tr int) {
	if tl == tr {
		t[v] = Node{Number: 0, Segments: 0, Set: 0, L: tl, R: tr, Up: false}
	} else {
		tm := getMiddle(tl, tr)
		build(v*2, tl, tm)
		build(v*2+1, tm+1, tr)
		t[v] = Node{Number: 0, Segments: 0, Set: 0, L: tl, R: tr, Up: false}
	}
}

func push(v int) {
	if !t[v].Up {
		return
	}

	t[v].Number = t[v].Set * (t[v].R - t[v].L + 1)
	t[v].Segments = 1 * t[v].Set
	t[v].Up = false

	if t[v].L == t[v].R {
		return
	}

	t[v*2].Set = t[v].Set
	t[v*2+1].Set = t[v].Set

	t[v*2].Up = true
	t[v*2+1].Up = true
}

func leftIsBlack(v int) bool {
	push(v)

	if t[v].L == t[v].R {
		return t[v].Number == 1
	}

	return leftIsBlack(v * 2)
}

func rightIsBlack(v int) bool {
	push(v)

	if t[v].L == t[v].R {
		return t[v].Number == 1
	}

	return rightIsBlack(v*2 + 1)
}

func update(v, value, l, r int) {
	if t[v].R < l || t[v].L > r {
		return
	}

	if t[v].R <= r && t[v].L >= l {
		push(v)
		t[v].Set = value
		t[v].Up = true
		return
	}

	push(v)
	update(v*2, value, l, r)
	update(v*2+1, value, l, r)

	left := rightIsBlack(v * 2)
	right := leftIsBlack(v*2 + 1)

	t[v].Number = t[v*2].Number + t[v*2+1].Number
	t[v].Segments = t[v*2+1].Segments + t[v*2].Segments

	if left && right {
		t[v].Segments--
	}
}

func main() {
	defer out.Flush()
	var maxDelta, maxCord int
	maxDelta = math.MaxInt32
	maxCord = math.MinInt32

	fmt.Fscanf(in, "%d", &n)
	color = make([]byte, n)
	cord = make([]int, n)
	delta = make([]int, n)

	for i := 0; i < n; i++ {
		var c string
		fmt.Fscanf(in, "\n%s %d %d", &c, &cord[i], &delta[i])

		if c != "" {
			color[i] = c[0] // Получаем байтовое значение из символа
		}

		if delta[i] > 0 {
			delta[i]--
		} else {
			delta[i]++
		}

		del := cord[i] + delta[i]
		if del > maxCord {
			maxCord = del
		}

		if maxDelta > cord[i] {
			maxDelta = cord[i]
		}
	}

	var length int
	if maxDelta < 0 {
		length = maxCord - maxDelta + 1
	} else {
		length = maxCord + 1
	}

	t = make([]Node, 4*length)
	build(1, 0, length)

	for i := 0; i < n; i++ {
		if color[i] == 'W' {
			update(1, 0, cord[i]-maxDelta, cord[i]+delta[i]-maxDelta)
			fmt.Fprintf(out, "%d %d\n", t[1].Segments, t[1].Number)
		}

		if color[i] == 'B' {
			update(1, 1, cord[i]-maxDelta, cord[i]+delta[i]-maxDelta)
			fmt.Fprintf(out, "%d %d\n", t[1].Segments, t[1].Number)
		}
	}
}
