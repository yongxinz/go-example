package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j <= 5; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for m := 0; m <= 10; m++ {
		if m%2 == 0 {
			continue
		}

		fmt.Println(m)
	}
}
