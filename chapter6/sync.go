package chapter6

import (
	"fmt"
	"sync"
	"time"
)

func Sync() {
	var wg sync.WaitGroup
	// wg.Add(1)

	// go sayHello(&wg)

	// wg.Wait()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("all workers finished")
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	fmt.Println("hello from go routine")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Second)
	fmt.Println("Worker", id, "done")
}
