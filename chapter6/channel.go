package chapter6

import (
	"fmt"
	"time"
)

func Channels() {
	// msg := make(chan string)

	// go func() {
	// 	msg <- "hello from goroutines"
	// }()

	// ans := <-msg
	// fmt.Println(ans)

	// done := make(chan bool)
	// go task(done)

	// <-done

	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 1; i <= 20; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}
	fmt.Println("done sending.")
	close(emailChan)
	<-done
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to:", email)
		time.Sleep(50 * time.Millisecond)
	}
}
