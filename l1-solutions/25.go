package main

import (
	"fmt"
	"time"
)

func Sleep(d int) {
	if d < 0 {
		return
	}
	duration := time.Duration(d) * time.Second
	done := make(chan struct{})
	timer := time.NewTimer(duration)
	go func() {
		defer close(done)
		<-timer.C
	}()
	<-done
}

func main() {
	Sleep(3)
	fmt.Println("Done")
}
