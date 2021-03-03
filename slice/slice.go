package main

import "fmt"

func main() {
	i := make([]int, 3)
	fmt.Println("i: ", i)

	s := make([]string, 3)
	fmt.Println("s: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append: ", s)
	fmt.Println("len: ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("c: ", c)

	fmt.Println("c[2:5]: ", c[2:5])

	d := make([][]int, 3)
	for m := 0; m < 3; m++ {
		innerLen := m + 1
		d[m] = make([]int, innerLen)
		for n := 0; n < innerLen; n++ {
			d[m][n] = m + n
		}
	}
	fmt.Println("d: ", d)
}
