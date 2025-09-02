package main

import (
	"fmt"
	"sync"
)

// функция, возвращающая квадрат числа
func printSquare(x int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик на 1
	fmt.Println(x * x)
}

func main() {
	// заданный срез
	slice := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	wg.Add(len(slice)) // увеличиваем счетчик на кол-во элементов массива
	// вызываем по горутине на каждый элемент
	for i := 0; i < len(slice); i++ {
		go printSquare(slice[i], &wg)
	}
	wg.Wait() // ждем пока каждая горутина завершит выполнение
}

