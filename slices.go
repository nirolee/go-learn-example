package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Println("arr[]", s, len(s), cap(s))
}

func main() {
	s4 := make([]int, 0, 0)
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s1 = append(s1, s2[:]...)
	fmt.Println(s1)
	//arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//s := arr[2:6]
	//s1 := s[:]
	////s = s[2:]
	////updateSlice(s)
	////fmt.Println("arr[2:6]", s1, len(s1), cap(s1))
	//s2 := append(s1, 8)
	//s2 = s2[2:6]
	//s3 := make([]int, 10, 32)
	//s3 = append(s2, s3[:2]...)
	//s3 = s3[:len(s3)+1]
	//printSlice(s3)
	//printSlice(s2)
	printSlice(s4)
	//fmt.Println("arr[2:6]", s2, len(s2  ), cap(s2))
	//fmt.Println("arr[2:6]", s1)
}
