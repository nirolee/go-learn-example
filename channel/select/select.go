package main

import "fmt"

//
//func generator() chan int {
//	out := make(chan int)
//	go func() {
//
//	}()
//}

func main() {
	var c1, c2 chan int
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
		default:
			fmt.Println(1)
		}
	}

}
