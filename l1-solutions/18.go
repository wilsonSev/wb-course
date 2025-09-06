package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var mu = sync.Mutex{}

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go increment(wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
