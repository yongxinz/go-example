package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "tow"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case msg := <-c2:
		fmt.Println(msg)
	}
}
