package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var (
	in  = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

type Node struct {
	Value    int
	Priority int
	Left     *Node
	Right    *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value, Priority: rand.Intn(100000000)}
}

func Merge(v *Node, u *Node) *Node {
	if v == nil {
		return u
	} else if u == nil {
		return v
	} else {
		if v.Priority > u.Priority {
			v.Right = Merge(v.Right, u)
			return v
		} else {
			u.Left = Merge(v, u.Left)
			return u
		}
	}
}

func split(v *Node, value int) [2]*Node {
	if v == nil {
		return [2]*Node{}
	} else {
		if v.Value <= value {
			arr := split(v.Right, value)
			v.Right = arr[0]
			return [2]*Node{v, arr[1]}
		} else {
			arr := split(v.Left, value)
			v.Left = arr[1]
			return [2]*Node{arr[0], v}
		}
	}
}

func insert(v *Node, value int) *Node {
	arr := split(v, value)
	return Merge(Merge(arr[0], NewNode(value)), arr[1])
}

func exists(v *Node, value int) bool {
	if v == nil {
		return false
	} else if v.Value == value {
		return true
	} else if v.Value > value {
		exists(v.Left, value)
	} else if v.Value < value {
		exists(v.Right, value)
	}
	return false
}

func Print(root *Node) {
	if root == nil {
		return
	}
	Print(root.Left)
	fmt.Fprintf(out, "%d ", root.Value)
	Print(root.Right)
}

func main() {
	defer out.Flush()
	var root *Node
	root = nil

	var n, t int
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n)
		root = insert(root, n)
	}
	Print(root)
}
