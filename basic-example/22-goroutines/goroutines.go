package main

import (
	"fmt"
	"time"
)

func f1(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, " : ", i)
	}
}

func main() {
	f1("direct")

	go f1("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
