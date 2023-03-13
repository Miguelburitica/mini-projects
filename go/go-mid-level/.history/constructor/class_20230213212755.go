package main

import "fmt"

type Employee struct {
	 id int
	 name string
	 vacation bool
}

func NewEmployee (id int, name string, vacation bool) *Employee {
	return &Employee{
		id,
		name,
		vacation,
	}
}

func main() {
	e := Employee{}
	fmt.Printf("e: %+v\n", e)

	e2 := Employee{
		id: 1,
		name: "Andrea",
		vacation: true,
	}
	fmt.Printf("e2: %+v\n", e2)

	e3 := new(Employee)
	fmt.Printf("e3: %+v\n", *e3)

	e4 := NewEmployee(1, "Miguel", true)
}