package main

// printer считывает сообщение из канала 'c', выводит его и засыпает на секунду
// pinger записывает в 'c' канал i-ые числа, пока printer спит
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

	// Запуск двух горутин pinger и printer
	go pinger(c)
	go printer(c)

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
