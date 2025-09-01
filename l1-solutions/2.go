package main

import (
	"fmt"
	"sync"
)

func printSquare(x int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(x * x)
}

func main() {
	slice := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(slice))
	for i := 0; i < len(slice); i++ {
		go printSquare(slice[i], &wg)
	}
	wg.Wait()
}

