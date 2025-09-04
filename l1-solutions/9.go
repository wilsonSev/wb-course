package main

import (
	"fmt"
)

// Функция, которая пишет переданные аргументы в канал
func write(args ...int) <-chan int {
	c := make(chan int) // создаем канал, который вернем в конце
	go func() {
		defer close(c)
		for _, arg := range args {
			c <- arg
		}
	}()
	return c
}

// Функция, которая пишет обработанные аргументы в канал
func multiply(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			v, ok := <-in
			if !ok {
				return
			}
			out <- v * 2
		}
	}()
	return out
}

func main() {
	c := multiply(write(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20))
	for v := range c { // завершается, когда закрывается канал
		fmt.Println(v)
	}
}
