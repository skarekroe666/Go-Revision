package chapter13

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Conc2() {
	done := make(chan any, 5)
	defer close(done)

	cows := make(chan any, 100)
	pigs := make(chan any, 100)

	go func() {
		for {
			select {
			case <-done:
				return
			case cows <- "moo":
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case pigs <- "oink":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows)
	wg.Add(1)
	go consumePigs(done, pigs)

	wg.Wait()
}

func orDone(done, c <-chan any) <-chan any {
	relayStream := make(chan any)

	go func() {
		defer close(relayStream)

		for {
			select {
			case <-done:
				return
			case v, ok := <-c:
				if !ok {
					return
				}
				select {
				case relayStream <- v:
				case <-done:
					return
				}
			}
		}
	}()

	return relayStream
}

func consumePigs(done <-chan any, pigs <-chan any) {
	defer wg.Done()
	for v := range orDone(done, pigs) {
		fmt.Println(v)
	}
}

func consumeCows(done <-chan any, cows <-chan any) {
	defer wg.Done()
	for v := range orDone(done, cows) {
		fmt.Println(v)
	}
}
