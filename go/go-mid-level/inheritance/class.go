package main

import "fmt"

type Person struct {
	name string
	age uint8
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Miguel"
	ftEmployee.age = 12
	ftEmployee.id = 5245124

	fmt.Printf("ftEmployee: %+v\n", ftEmployee)
	GetMessage(ftEmployee.Person)
}