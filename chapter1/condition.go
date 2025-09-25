package chapter1

import (
	"fmt"
	"time"
)

func Conditions() {
	i := 3
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon and time is", t)
	default:
		fmt.Println("its after noon and the time is", t)
	}

}
