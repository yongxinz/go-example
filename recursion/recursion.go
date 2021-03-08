package main

import "fmt"

func fact(a int) int {
	if a == 0 {
		return 1
	}
	return a * fact(a-1)
}

func main() {
	n := fact(7)
	fmt.Println(n)
}
