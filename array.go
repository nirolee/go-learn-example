package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 4}
	s1 = append(s1, s2[:]...)
	fmt.Println(s1)
}
