package main

import "fmt"

type Employee struct {
	id int
	name string
}

func (e *Employee) setId(newId int) () {
	e.id = newId
}

func main() {
	e := Employee{}
	fmt.Printf("e: %+v\n", e)

	e.id = 1
	e.name = "Andrea"
	fmt.Printf("e: %+v\n", e)

	e.setId(5)
	fmt.Printf("e: %+v\n", e)
}