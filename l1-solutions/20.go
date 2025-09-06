package main

import "fmt"

func reverse(runes []rune) {
	i := 0
	j := len(runes) - 1
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	reverse(runes)
	start := 0
	end := 0
	for end < len(runes) {
		for end < len(runes) && runes[end] != ' ' {
			end++
		}
		reverse(runes[start:end])
		end++
		start = end
	}
	return string(runes)
}

func main() {
	s := "snow dog sun"
	fmt.Println(reverseWords(s))
}
