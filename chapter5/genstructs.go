package chapter5

import "fmt"

type Box[T any] struct {
	value T
}

type Container[T any] struct {
	data []T
}

func GenericStruct() {
	intBox := Box[int]{value: 324}
	fmt.Println(intBox.get())

	strBox := Box[string]{value: "bitch"}
	fmt.Println(strBox.get())

	c := Container[int]{}
	c.add(4)
	c.add(66)
	fmt.Println(c.data)
}

func (b Box[T]) get() T {
	return b.value
}

func (c *Container[T]) add(item T) {
	c.data = append(c.data, item)
}
