package main

import "fmt"

func main() {
	foo(233)
	foo("666")
}

func foo(v interface{}) {
	if v1, ok1 := v.(string); ok1 {
		fmt.Println(v1)
	} else if v2, ok2 := v.(int); ok2 {
		fmt.Println(v2)
	}
}
