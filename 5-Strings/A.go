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

type hstring struct {
	base []int64
	pref []int64
	str  string
}

func (h *hstring) hash(l, r int) int64 {
	mod := int64(1e9 + 7)
	result := h.pref[r]
	if l > 0 {
		result -= h.pref[l-1] * h.base[r-l+1]
		result = (result%mod + mod) % mod // Adjust for negative result
	}
	return result
}

func (h *hstring) create() {
	h.base = make([]int64, len(h.str))
	h.base[0] = 1
	h.base[1] = 57
	mod := int64(1e9 + 7)
	for i := 2; i < len(h.str); i++ {
		h.base[i] = (h.base[i-1] * h.base[1]) % mod
	}

	// pref[i] = хеш первых i символов строки
	h.pref = make([]int64, len(h.str)+1) // увеличиваем длину на 1
	h.pref[0] = 0
	for i := 1; i <= len(h.str); i++ { // исправляем условие цикла
		h.pref[i] = (h.pref[i-1]*h.base[1] + num(h.str[i-1])) % mod
	}
}

func main() {
	defer out.Flush()
	var str string
	fmt.Fscanln(in, &str)
	var q, a, b, c, d int
	fmt.Fscanln(in, &q)
	h := hstring{str: str}
	h.create()
	for i := 0; i < q; i++ {
		fmt.Fscanln(in, &a, &b, &c, &d)
		if a == c && str[a:b] == str[c:d] { // проверка на равенство строк
			fmt.Fprintln(out, "Yes")
			continue
		}
		if a-b != c-d {
			fmt.Fprintln(out, "No")
			continue
		}
		hash1 := h.hash(a, b)
		hash2 := h.hash(c, d)
		if hash1 == hash2 {
			fmt.Fprintln(out, "Yes")
		} else {
			fmt.Fprintln(out, "No")
		}
	}

}

func num(ch byte) int64 {
	return int64(ch)
}
