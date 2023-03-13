package main

import "fmt"

type Employee struct {
	 id int
	 name string
}

func main() {
	e := Employee{}
	fmt.Printf("e: %+v\n", e)
}