package main

import (
	"fmt"
	"time"
)

func main() {
	nums := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		nums <- i
	}
	close(nums)

	limiter := time.Tick(time.Second)
	for n := range nums {
		<-limiter
		fmt.Println("request ", n, time.Now())
	}

	burstyLimiters := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiters <- time.Now()
	}

	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimiters <- t
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println("request ", i, <-burstyLimiters)
	}
}
