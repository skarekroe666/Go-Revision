package chapter4

import "fmt"

func Strings() {
	s := "hello"

	fmt.Println("length of the string is:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
