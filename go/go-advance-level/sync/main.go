package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 100
)

func Deposit(amount int, wg *sync.WaitGroup, mux *sync.RWMutex) {
	defer wg.Done()
	mux.Lock()
	b := balance
	balance = b + amount
	mux.Unlock()
}

func Balance(mux *sync.RWMutex) int {
	mux.RLock()
	b := balance
	mux.RUnlock()
	return b
}

func main() {
	var wg sync.WaitGroup
	var mux sync.RWMutex

	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go Deposit(i * 1, &wg, &mux)
	}
	wg.Wait()
	fmt.Println(Balance(&mux))
}