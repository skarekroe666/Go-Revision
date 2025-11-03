package chapter13

import (
	"fmt"
	"sync"
	"time"
)

func Conc() {
	start := time.Now()
	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5, 6}
	result := make([]int, len(input))

	for i, v := range input {
		wg.Add(1)
		go processData(&wg, &result[i], v)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
	fmt.Println(result)
}

func processData(wg *sync.WaitGroup, resultDest *int, data int) {
	defer wg.Done()

	*resultDest = process(data)
}

func process(data int) int {
	time.Sleep(500 * time.Millisecond)
	return data * 2
}
