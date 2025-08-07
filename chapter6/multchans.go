package chapter6

import "fmt"

func MultipleChannels() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received data from chan1:", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("received data from chan2:", chan2Val)
		}
	}
}
