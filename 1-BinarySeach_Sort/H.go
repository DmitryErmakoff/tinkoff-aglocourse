package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(in, &n)

	str := make([]byte, n)
	for i := 0; i < n; i++ {
		str[i], _ = in.ReadByte()
	}
	in.ReadByte()
	defer out.Flush()
	fmt.Fprintln(out, solvePalindrome(str))
}

func solvePalindrome(arr []byte) string {
	freqMap := make(map[byte]int, len(arr))
	for i := 0; i < len(arr); i++ {
		freqMap[arr[i]] += 1
	}
	tmp := 0
	fLet := byte(0)
	for l := byte('A'); l <= 'Z'; l++ {
		if freqMap[l] == 0 {
			continue
		}
		for i := 0; i < freqMap[l]/2; i++ {
			arr[tmp] = l
			tmp++
		}
		if fLet == 0 && freqMap[l]%2 == 1 {
			fLet = l
		}

	}
	res := strings.Builder{}
	for i := 0; i < tmp; i++ {
		res.WriteByte(arr[i])
	}

	if fLet != 0 {
		res.WriteByte(fLet)
	}

	for i := tmp - 1; i >= 0; i-- {
		res.WriteByte(arr[i])
	}
	return res.String()
}
