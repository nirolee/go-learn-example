package main

import (
	"fmt"
	"time"
)

func main() {
	i:= 2
	fmt.Println("i等于:", i)
	switch i {
		case 1:
			fmt.Println(time.Now().Unix())
		case 2:
			fmt.Println("i等于:", time.Now().Format("2006-01-02 15:04:05"))
			fallthrough
		case 3:
			fmt.Println("i等于:", i)
		case 4:
			fmt.Println("i等于:", i)
	}
}
