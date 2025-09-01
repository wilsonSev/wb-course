package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func process(ctx context.Context, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case val, ok := <-c:
			if !ok {
				return
			}
			fmt.Println("Processed ", val)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: не передан аргумент")
		os.Exit(1)
	}
	workersNumber, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	iterations := 10
	wg := &sync.WaitGroup{}
	wg.Add(workersNumber)
	for i := 0; i < workersNumber; i++ {
		go process(ctx, c, wg)
	}
	for i := 0; i < iterations; i++ {
		c <- i
		time.Sleep(time.Millisecond * 50)
	}
	close(c)
	wg.Wait()
}
