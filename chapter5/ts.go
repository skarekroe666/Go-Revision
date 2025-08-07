package chapter5

import "fmt"

func TypeAssertion() {
	// var a any = 42
	var a any = "hello"
	b, ok := a.(int)
	if ok {
		fmt.Println("type assertion done:", b)
	} else {
		fmt.Println("type assertion failed")
	}

	// b := a.(int)

	checkType(443)
	checkType("bitch")
	checkType(43.54)
	checkType('c')
}

func checkType(a any) {
	switch v := a.(type) {
	case int:
		fmt.Println("type is int:", v)
	case string:
		fmt.Println("type is string:", v)
	case float64:
		fmt.Println("type is float:", v)
	default:
		fmt.Println("unknown type")
	}
}
