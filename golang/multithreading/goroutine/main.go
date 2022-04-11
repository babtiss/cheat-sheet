package main

import (
	"fmt"
	"time"
)

func f(n int) {
	fmt.Println(n)
}

func f2(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	for i := 0; i < 3; i++ {
		//go f(i)
		// При вызове данной горутины мы получим рандомную перестановку чисел от 0 до 2
		// Это не заметно, но горутины выполняются последовательно

		go f2(i)
		// При вызове данной горутины мы получим рандомные перестановки чисел от 0 до 2
		// В этом случае они будут выводится одновременно, т.к.
	}
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		return
	}
}
