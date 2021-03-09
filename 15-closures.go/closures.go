package main

import "fmt"

func intSeq() func() int {
	i := 1
	return func() int {
		i++
		return i
	}
}

func main() {
	nxtInt := intSeq()

	fmt.Println(nxtInt())
	fmt.Println(nxtInt())
	fmt.Println(nxtInt())

	nxtInts := intSeq()
	fmt.Println(nxtInts())
}
