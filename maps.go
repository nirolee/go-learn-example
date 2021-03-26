package main

import "fmt"

func main() {
	m := map[string]string{
		"m": "mm",
		"b": "bb",
	}
	for _, v := range m {
		fmt.Println("key,value,", v)
	}
	delete(m, "m")
	fmt.Println(m)
}
