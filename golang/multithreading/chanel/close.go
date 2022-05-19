package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	go func() {
		for i := 1; i <= 10; i++ {
			message <- fmt.Sprintf("i:=%v", i)
			time.Sleep(100 * time.Millisecond)
		}
		close(message) // если не закрывать канал, то получим ошибку
	}()
	for {
		msg, open := <-message
		if !open {
			break
		}
		fmt.Println(msg)
	}
}
