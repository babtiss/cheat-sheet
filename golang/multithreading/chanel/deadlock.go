package main

import "fmt"

func main() {
	message := make(chan string)
	message <- "start"
	fmt.Println(<-message)
}
