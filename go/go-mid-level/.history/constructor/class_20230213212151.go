package main

import "fmt"

type Employee struct {
	 id int
	 name string
	 vacation bool
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
}