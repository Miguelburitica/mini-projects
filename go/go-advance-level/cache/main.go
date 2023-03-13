package main

import (
	"fmt"
	"log"
	"time"
)

func Fibonacci(cache *Memory, n int) int {
	if n <= 1 {
		return n
	}

	fb1, _ := GetFibonacci(cache, n-1)
	fb2, _ := GetFibonacci(cache, n-2)
	
	return fb1.(int) + fb2.(int)
}

type Memory struct {
	f Function
	cache map[int]FunctionResult
}

type Function func(cache *Memory, key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f: f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]

	if !exists {
		result.value, result.err = m.f(m, key)
		m.cache[key] = result
	}
	return result.value, result.err
}

func GetFibonacci(cache *Memory, n int) (interface{}, error) {
	var value, err = cache.Get(n)

	if err != nil {
		return nil, err
	}

	return value, nil
}

func main() {
	cache := NewCache(func(cache *Memory, n int) (interface{}, error) {
		return Fibonacci(cache, n), nil
	})

	
	for {
		var fibo int = 0
		var quitFlag string
		fmt.Printf("Ingresa la posición de fibonacci a consultar: ")
		fmt.Scanf("%d\n", &fibo)
		start := time.Now()
		var value, err = GetFibonacci(cache, fibo)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("posición: %d, tardado: %s, con resultado: %d\n\n", fibo, time.Since(start), value)

		fmt.Printf("Para salir del loop ingresa 'n': ")
		fmt.Scanf("%s", &quitFlag)
		if quitFlag == "n" {
			break
		}
		fmt.Printf("\n")
		fmt.Printf("--------------------------------------------\n\n")
	}
}
