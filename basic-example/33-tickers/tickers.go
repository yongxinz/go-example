package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTicker(2 * time.Second)
	ch := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-ch:
				return
			case i := <-timer.C:
				fmt.Println("pro run at ", i)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	timer.Stop()
	ch <- true
}
