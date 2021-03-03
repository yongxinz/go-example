package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// 向上取整
	fmt.Println("7.0 / 3.0 = ", math.Ceil(7.0/3.0))
	// 向下取整
	fmt.Println("7.0 / 3.0 = ", math.Floor(7.0/3.0))
	// 四舍五入
	fmt.Println("7.0 / 3.0 = ", math.Floor(7.0/3.0+0.5))

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
