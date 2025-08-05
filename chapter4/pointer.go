package chapter4

import "fmt"

type Person struct {
	Name string
	Age  int
}

func Pointers() {
	x := 12
	p := &x

	fmt.Println(p)
	fmt.Println(*p)

	s := Person{Name: "skarekroe", Age: 26}
	fmt.Println(s)

	changeValue(&s)
	fmt.Println(s)

	increase(&s)
	fmt.Println(s)
}

func changeValue(p *Person) {
	p.Age = 27
	p.Name = "sanskar"
}

func increase(p *Person) {
	p.Age++
}
