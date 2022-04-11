package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
