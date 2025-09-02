package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// функция-воркер
func process(ctx context.Context, c chan int, wg *sync.WaitGroup) {
	defer wg.Done() // в самом конце декрементим wg
	for {
		val, ok := <-c
		if !ok { // при ok = false канал закрыт
			return
		}
		fmt.Println("Processed ", val) // если канал открыт, выводим значение в консоль
	}
}

func main() {
	// проверка, что количество переданных аргументов верное
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Ошибка: не передан аргумент")
		os.Exit(1)
	}
	workersNumber, err := strconv.Atoi(os.Args[1]) // конвертируем string в int
	if err != nil {                                // обрабатываем ошибку приведения типов
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	c := make(chan int, 16) // создаем конал, с которым мы работаем

	iterations := 100 // количество итераций записи в канал
	// создаем waitgrouop и добавляем туда кол-во воркеров
	wg := &sync.WaitGroup{}
	wg.Add(workersNumber)
	// добавляем контекст с обработкой SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	// запускаем воркеры
	for i := 0; i < workersNumber; i++ {
		go process(ctx, c, wg)
	}
	// пишем в канал
produce: // создаем метку, чтобы выйти из цикла при завершении контекста
	for i := 0; i < iterations; i++ {
		select {
		case <-ctx.Done():
			break produce
		default:
			c <- i
			time.Sleep(time.Millisecond * 50)
		}
	}
	close(c)  // закрываем канал, чтобы остановить горутины
	wg.Wait() // ждем, пока все горутины завершат свою работу
}
