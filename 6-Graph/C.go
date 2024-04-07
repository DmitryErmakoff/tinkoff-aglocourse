package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var g [][]int
var top []int
var used []int
var n, m, flag int

func dfs(v int) {
	used[v] = 1
	for i := 0; i < len(g[v]); i++ {
		to := g[v][i]
		if used[to] == 1 {
			flag = 1
			return
		}
		if used[to] == 0 {
			dfs(to)
		}
	}
	used[v] = 2
	top = append(top, v)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &m)

	g = make([][]int, n+1)
	used = make([]int, n+1)

	for i := range g {
		g[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(reader, &a, &b)
		g[a] = append(g[a], b)
		fmt.Fscanln(reader)
	}

	for i := 1; i <= n; i++ {
		if used[i] == 0 {
			dfs(i)
		}
	}

	var str string
	var inputTopArr []string
	byteStr, _ := reader.ReadString('\n')
	str = strings.TrimSpace(byteStr)
	inputTopArr = strings.Split(str, " ")

	match := true
	if len(inputTopArr) != len(top) {
		match = false
	} else {
		for i := len(inputTopArr) - 1; i >= 0; i-- {
			node, _ := strconv.Atoi(inputTopArr[i])
			if top[len(inputTopArr)-1-i] != node {
				match = false
				break
			}
		}
	}

	if match {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
