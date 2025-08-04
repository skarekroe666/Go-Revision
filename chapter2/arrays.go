package chapter2

import (
	"fmt"
)

func Arrays() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 34
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	c := [...]int{3, 5: 43, 3}
	fmt.Println("idx:", c)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d array:", twoD)
}
