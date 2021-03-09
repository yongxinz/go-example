package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, i := range num {
		sum += i
	}
	fmt.Println("sum = ", sum)

	for i, n := range num {
		if n == 3 {
			fmt.Println("i = ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("k -> v: ", k, v)
	}

	for k := range kvs {
		fmt.Println("key: ", k)
	}

	for k, v := range "abc" {
		fmt.Println("k, v: ", k, v)
	}
}
