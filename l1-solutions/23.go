package main

import "fmt"

func deleteEl(slice []int, ind int) []int {
	copy(slice[ind:], slice[ind+1:])
	return slice[:len(slice)-1]
}

func main() {
	var slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(deleteEl(slice, 4))
}
