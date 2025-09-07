package main

import (
	"fmt"
	"strings"
)

func isUniq(s string) bool {
	strings.ToLower(s)
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if m[s[i]] > 1 {
			return false
		}
	}
	return true
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isUniq(s))
}
