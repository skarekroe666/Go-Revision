package chapter3

import "fmt"

func Vardiac() {
	sum(1, 2, 3, 4, 5)

	nums := []int{3, 34, 6, 5, 3}
	sum(nums...)
}

func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
