package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0
	for i := 0; i < 1000; i++ {
		go func(i int) {
			count += 1
			fmt.Printf("Hello %d\n", i)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println(count)
}

// Программа выводит разом count чисел - столько, сколько успеет напечатать за отведенный timeout
