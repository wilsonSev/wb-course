package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	result := make([]string, 0, len(set))
	for x := range set {
		result = append(result, x)
	}
	fmt.Println(result)
}
