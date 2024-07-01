package main

import "fmt"

type Person struct {
	name   string
	age    int8
	salary float32
	status bool
}

func main() {
	var p Person
	p.name = "Matheus"
	p.age = 21
	p.salary = 2000
	p.status = true

	p2 := Person{
		name:   "Tiago",
		age:    26,
		salary: 3000,
		status: true,
	}

	fmt.Println(p)
	fmt.Println(p2)

}
