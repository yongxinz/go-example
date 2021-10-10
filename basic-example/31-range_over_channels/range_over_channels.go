package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "one"
	ch <- "two"
	close(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
