package chapter2

import "fmt"

func Maps() {
	m := make(map[string]int)

	m["sanskar"] = 26
	m["sanjana"] = 33

	fmt.Println(m)
	fmt.Println("length:", len(m))

	delete(m, "sanjana")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)
}
