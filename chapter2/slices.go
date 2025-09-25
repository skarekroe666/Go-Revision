package chapter2

import "fmt"

func Slices() {
	s := make([]string, 3)
	fmt.Println(s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[:4]
	fmt.Println(l)

	t := []int{1, 3, 5, 9}
	fmt.Println(t)

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d slice:", twoD)
}
