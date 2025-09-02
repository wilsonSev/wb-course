package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

// воркер, читающий значения
func worker(c <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		v, ok := <-c
		if !ok { // канал закрыт
			return
		}
		fmt.Println(v)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Too few arguments")
		os.Exit(1)
	}
	n, _ := strconv.Atoi(os.Args[1])
	c := make(chan string, 16) // создаем буферезирванный канал
	// задаем таймаут
	timeout := time.After(time.Duration(n) * time.Second)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go worker(c, wg)
	for {
		select {
		case <-timeout:
			close(c)  // сначала закрываем канал
			wg.Wait() // потом ждем воркер
			return
		case c <- "Data":
			time.Sleep(100 * time.Millisecond)
			// данные отправлены в канал, продолжаем цикл
		}
	}
}
