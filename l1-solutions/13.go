package main

import "fmt"

func main() {
	a := 7
	b := 10
	a = a + b
	b = a - b
	a = a - b
	fmt.Println(a)
	fmt.Println(b)
}
