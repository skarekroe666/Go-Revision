package chapter1

import "fmt"

func Loops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	for i := range 8 {
		fmt.Println("range:", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
