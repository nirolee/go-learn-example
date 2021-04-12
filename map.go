package main

import (
	"fmt"
)

func main() {
	//var j = map[int]string{
	//	1:"1j",
	//	2:"2j",
	//}
	//fmt.Println(j)

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, value := range slice {
		m[key] = &value
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}
