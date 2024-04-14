package main

import (
	"container/ring"
	"fmt"
)

func main() {
	splash := ring.New(2)
	splash.Value = 2
	splash = splash.Next()
	splash.Value = 3

	var n int
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		nextSplash := &ring.Ring{Value: splash.Prev().Value.(int) + splash.Value.(int)}
		splash.Link(nextSplash)
		splash = nextSplash
	}

	fmt.Println(splash.Value, "")
}
