package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 41
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "John", age: 30})

	fmt.Println(person{name: "Tom"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jack"))

	// s := person{"Pony", 50}
	// s.age = 51
	// fmt.Println(s)

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s)
}
