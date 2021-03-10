package main

import "fmt"

func main() {
	message := make(chan string)

	go func(msg string) { message <- msg }("hello channel")

	fmt.Println(<-message)
}
