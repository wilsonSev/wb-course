package main

import (
	"strings"
)

func createHugeString(n int) string {
	return strings.Repeat("a", n) // строка из n букв 'a'
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]byte(v[:100])) // создаем копию
}

func main() {
	someFunc()
}
