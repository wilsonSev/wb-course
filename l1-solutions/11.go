package main

import "fmt"

func main() {
	set := make(map[int]struct{})
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 6, 7, 8, 9, 10}
	for _, x := range a {
		set[x] = struct{}{}
	}
	for _, x := range b {
		set[x] = struct{}{}
	}

	result := make([]int, 0, len(set))
	for x := range set {
		result = append(result, x)
	}
	fmt.Println(result)
}
