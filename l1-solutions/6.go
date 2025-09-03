package main

// В этой программе будут представлены разные способы выйти из горутины (4 основных способа, упомянутых в условии)
import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Выходим из горутины с помощью return, когда выполнено необходимое условие
func goroutine1(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		res := rand.Intn(3)
		fmt.Println("Random number is ", res)
		// Выходим из гортины, если достигается условие
		if res == 1 {
			fmt.Println("goroutine1 ok")
			return
		}
	}
}

// Выходим из горутины с помощью специального канала, созаданного нами
func goroutine2(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("goroutine2 ok")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
	}
}

// Горутина будет завершена через контекст спустя одну секунду
func goroutine3(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("goroutine3 ok")
			return
		default:
			fmt.Println("Working")
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
	}
}

// Завершение через runtime.Goexit()
func goroutine4(wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Println("goroutine4 ok")
	fmt.Println("before Goexit")
	runtime.Goexit()
	fmt.Println("this will never be printed")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("---Starting goroutine1---")
	go goroutine1(&wg)
	wg.Wait()

	stop := make(chan struct{}) // специальный канал, который сигнализирует об остановке
	wg.Add(1)
	fmt.Println("---Starting goroutine2---")
	go goroutine2(stop, &wg)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	close(stop)
	wg.Wait()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	wg.Add(1)
	fmt.Println("---Starting goroutine3---")
	go goroutine3(ctx, &wg)
	wg.Wait()

	fmt.Println("---Starting goroutine4---")
	wg.Add(1)
	go goroutine4(&wg)
	wg.Wait()
}
