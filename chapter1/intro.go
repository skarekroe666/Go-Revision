package chapter1

import (
	"fmt"
	"math"
)

const a string = "something"

func BasicExample() {
	fmt.Println(a)

	const n = 54899729

	const d = 3e10 / n

	fmt.Println(d)

	fmt.Println(math.Sqrt(n))
}


























