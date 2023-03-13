package main

import (
	"fmt"
	"time"
)

func main() {
	// x := 5

	// y := func() int {
	// 	return x * x
	// }

	// fmt.Printf("y: %v\n", y)
	
	c := make(chan int)
	go func() {
		fmt.Println("Starting process")
		time.Sleep(5 * time.Second)
		fmt.Println("Ending process")
		c <- 1
	}()

	<-c
}