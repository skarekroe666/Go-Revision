package chapter6

import (
	"fmt"
	"time"
)

func Routines() {
	go printNums()
	time.Sleep(3 * time.Second)
}

func printNums() {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}
