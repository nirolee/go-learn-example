package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	s1 := s[:]
	//s = s[2:]
	//updateSlice(s)
	fmt.Println("arr[2:6]", s1, len(s1), cap(s1))
	s2 := append(s1, 8)
	s2 = s2[2:6]
	fmt.Println("arr[2:6]", s2, len(s2), cap(s2))
	//fmt.Println("arr[2:6]", s1)
}
