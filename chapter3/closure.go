package chapter3

import "fmt"

func Closure() {
	result1 := add()
	fmt.Println(result1())
	fmt.Println(result1())
	fmt.Println(result1())

	result2 := add()
	fmt.Println(result2())
}

func add() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
