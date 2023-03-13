package main

import "fmt"

type Employee struct {
	id int
	name string
}

func main() {
	e := Employee{}
	fmt.Printf("%v", e)
}