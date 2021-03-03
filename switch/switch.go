package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it is weekend")
	default:
		fmt.Println("it is workday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it is before none")
	default:
		fmt.Println("it is after none")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it is bool")
		case int:
			fmt.Println("it is int")
		default:
			fmt.Printf("it is unknown type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")
}
