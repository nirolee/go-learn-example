package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred", err)
		} else {
			panic(fmt.Sprintf("i don't know what to do:%v", r))
		}
	}()
	//panic(errors.New("this is an error"))
	b := 1
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}
