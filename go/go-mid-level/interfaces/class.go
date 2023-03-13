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
	endDate string
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}

	ftEmployee.name = "Miguel"
	ftEmployee.age = 12
	ftEmployee.id = 5245124

	// fmt.Printf("ftEmployee: %+v\n", ftEmployee)
	tEmployee := TemporaryEmployee{}

	getMessage(tEmployee)
	getMessage(ftEmployee)
	
	// GetMessage(ftEmployee.Person)
}