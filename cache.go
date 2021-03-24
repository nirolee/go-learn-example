package main

import "fmt"

func printArr(arr *[10]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
	arr[0] = 100
}

func main() {
	//cache := make(map[string]string)
	//cache["name"] = "ccmoou"
	//cache := [...]int{2,4,5,6}
	var var1 [10]int
	//var2 := [3]int{1, 2, 5}
	printArr(&var1)
	fmt.Println(var1)
	//s := 0
	//for i := range cache{
	//	fmt.Println(cache[i])
	//	s += cache[i]
	//}
	//fmt.Println(s)
}
