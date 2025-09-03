package main

import (
	"math/rand"
	"sync"
)

var data map[int]string = make(map[int]string)
var mutex sync.Mutex
var wg = sync.WaitGroup{}

// Функция, которая пишет в мапу
func writeData(id int) {
	defer wg.Done()
	mutex.Lock() // мьютекс для предотвращения гонки
	data[id] = "someData"
	mutex.Unlock()
}

func main() {
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go writeData(rand.Intn(100)) // запускаем каждого воркера в отдельной горутине
	}
	wg.Wait() // ждем пока горутины закончат выполнение
}
