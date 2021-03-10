package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	// message := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("receive msg: ", msg)
	default:
		fmt.Println("no receive message")
	}

	msg := "hello"
	select {
	case message <- msg:
		fmt.Println("send msg: ", msg)
	default:
		fmt.Println("no send message")
	}

	select {
	case msg := <-message:
		fmt.Println("receive message: ", msg)
	case sig := <-signals:
		fmt.Println("receive sig: ", sig)
	default:
		fmt.Println("no active")
	}
}
