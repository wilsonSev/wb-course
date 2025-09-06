package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	return string(runes)
}

func main() {
	s := "hello world"
	fmt.Println(reverse(s))
}
