package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 2
	m["k2"] = 300

	fmt.Println("m: ", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("map len: ", len(m))

	m["k1"] = 3
	fmt.Println("m: ", m)

	delete(m, "k1")
	v1 = m["k1"]
	fmt.Println("v1: ", v1)

	_, prs := m["k1"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2, "baz": 3}
	fmt.Println("n: ", n)
}
