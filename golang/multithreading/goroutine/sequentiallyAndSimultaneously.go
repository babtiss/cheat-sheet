package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	time.Sleep(time.Millisecond * 1000)
	for n := 0; n < 3; n++ {
		go f(n)
		// При вызове данной горутины мы получим рандомную перестановку пар чисел n и i
		// Это не заметно, но горутины выполняются последовательно

		//go f2(n)
		// При вызове данной горутины мы получим рандомные перестановки чисел, но ...
		// В этом случае они будут выводится одновременно по 3, в виде 0:i , 1:i , 2:i , где i - число от 0 до 9
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
