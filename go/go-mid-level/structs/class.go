package main

import "fmt"

type Employee struct {
	id int
	name string
}

func (e *Employee) SetId(newId int) () {
	e.id = newId
}

func (e *Employee) SetName(newName string) {
	e.name = newName
}

func (e Employee) GetId() int {
	return e.id
}

func (e Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}
	fmt.Printf("e: %+v\n", e)

	e.id = 1
	e.name = "Andrea"
	fmt.Printf("e: %+v\n", e)

	e.SetId(5)
	fmt.Printf("e: %+v\n", e)

	e.SetName("Miguel")
	fmt.Printf("e: %+v\n", e)
}