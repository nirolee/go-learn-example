package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Println("arr[]", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s := arr[2:6]
	s1 := s[:]
	//s = s[2:]
	//updateSlice(s)
	//fmt.Println("arr[2:6]", s1, len(s1), cap(s1))
	s2 := append(s1, 8)
	s2 = s2[2:6]
	s3 := make([]int, 10, 32)
	s3 = append(s2, s3[:2]...)
	s3 = s3[:len(s3)+1]
	printSlice(s3)
	printSlice(s2)
	//fmt.Println("arr[2:6]", s2, len(s2), cap(s2))
	//fmt.Println("arr[2:6]", s1)
}
