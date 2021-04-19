package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	//var j = map[int]string{
	//	1:"1j",
	//	2:"2j",
	//}
	//fmt.Println(j)

	//slice := []int{0, 1, 2, 3}
	//m := make(map[int]*int)
	//for key, value := range slice {
	//	m[key] = &value
	//}
	//for k, v := range m {
	//	fmt.Println(k, "->", *v)
	//}

	fmt.Println(cap([4]int{1, 2, 4}))
}

func hello(num ...int) {
	num[0] = 18
}
