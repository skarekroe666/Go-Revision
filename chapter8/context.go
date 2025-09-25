package chapter8

import (
	"fmt"
	"sync"
	"time"
)

type Order struct {
	TableNumber int
	PrepTime    time.Duration
}

func ContextEx() {
	orders := []Order{
		{TableNumber: 1, PrepTime: 2 * time.Second},
		{TableNumber: 2, PrepTime: 3 * time.Second},
		{TableNumber: 3, PrepTime: 1 * time.Second},
		{TableNumber: 4, PrepTime: 5 * time.Second},
		{TableNumber: 5, PrepTime: 4 * time.Second},
	}

	wg := sync.WaitGroup{}

	for _, order := range orders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processOrder(order)
		}()
	}

	wg.Wait()
}

func processOrder(order Order) {
	fmt.Printf("Preparing order for table %d...\n", order.TableNumber)

	time.Sleep(order.PrepTime)

	fmt.Printf("Order ready for table %d...\n", order.TableNumber)
}
