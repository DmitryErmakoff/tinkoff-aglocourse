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

func minShift(s string) string {
	sLen := len(s)
	doubled := s + s
	minIdx := 0

	for i := 1; i < sLen; i++ {
		if shifted := doubled[i : i+sLen]; shifted < doubled[minIdx:minIdx+sLen] {
			minIdx = i
		}
	}

	return doubled[minIdx : minIdx+sLen]
}

func main() {
	var str string
	fmt.Fscanln(in, &str)
	result := minShift(str)
	fmt.Println(result)
}
