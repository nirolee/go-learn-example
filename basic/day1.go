package basic

import (
	"bufio"
	"fmt"
)

func first() {
	// hello world
	fmt.Println("Hello, World!")
	//计算两数之和
	a := 1
	b := 2
	c := a + b
	fmt.Println(c)
	//接受输入用户姓名和年龄，打印出来
}

func second(scanner *bufio.Reader) string {
	fmt.Println("请输入：")

	name, _ := scanner.ReadString('\n')
	return fmt.Sprintf(name)
}

func third(n int) {
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
