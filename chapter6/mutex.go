package chapter6

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	count int
	mu    sync.Mutex
}

func Mutex() {
	counter := SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("final count value:", counter.Value())
}

func (s *SafeCounter) Increment() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *SafeCounter) Value() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}
