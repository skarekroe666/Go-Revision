package chapter5

import "fmt"

func Generics() {
	fmt.Println(add(54, 343))
}

func add[T Number](a, b T) T {
	return a + b
}

type Number interface {
	int | float64
}
