package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	in  = bufio.NewReader(os.Stdin)
	sc  = bufio.NewScanner(in)
	out = bufio.NewWriter(os.Stdout)
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	defer out.Flush()
	var n, root int
	fmt.Fscanln(in, &n, &root)
	tree := make([]Node, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		nums := strings.Fields(sc.Text())
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		tree[i].Value = i

		if num1 == -1 {
			tree[i].Left = nil
		} else {
			tree[i].Left = &tree[num1]
		}

		if num2 == -1 {
			tree[i].Right = nil
		} else {
			tree[i].Right = &tree[num2]
		}
	}

	if isAVLTree(&tree[root]) {
		fmt.Fprintln(out, 1)
	} else {
		fmt.Fprintln(out, 0)
	}
}

func isAVLTree(root *Node) bool {
	if root == nil {
		return true
	}

	return isAVLUtil(root, math.MinInt64, math.MaxInt64)
}

func isAVLUtil(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Value <= min || node.Value >= max {
		return false
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if abs(leftHeight-rightHeight) > 1 {
		return false
	}

	return isAVLUtil(node.Left, min, node.Value) && isAVLUtil(node.Right, node.Value, max)
}

func height(node *Node) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	return 1 + max(leftHeight, rightHeight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
