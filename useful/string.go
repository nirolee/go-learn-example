package main

import "fmt"

func main() {
	//修改字符串中一个字符
	str := "hello"
	//c:= [] byte(str)
	//c[0] = '1'
	//println(string(c))

	//substr := str[0:4]
	//println(substr)

	for _, value := range []byte(str) {
		fmt.Println(value)
	}
}
