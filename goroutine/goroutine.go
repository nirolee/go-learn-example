package main

import (
	"fmt"
	"time"
)

func main() {
	//var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println(i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
